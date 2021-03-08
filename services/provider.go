package services

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"crypto/tls"
	"embed"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/scinna/server/dal"
	"golang.org/x/text/language"
	"html/template"
	"io/fs"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/pariz/gountries"
	"github.com/scinna/server/config"
	"golang.org/x/crypto/argon2"
)

//go:embed templates
var templatesFS embed.FS
var mailTemplates *template.Template

var countries = gountries.New()

type Provider struct {
	Webapp      *fs.FS
	DB          *sqlx.DB
	Dal         *dal.Dal
	ArgonParams *ArgonParams
	MailClient  smtp.Auth
	Config      *config.Config

	languageMatcher language.Matcher
}

func NewProvider(cfg *config.Config, webapp *embed.FS) (*Provider, error) {
	correctedFS, err := fs.Sub(webapp, "frontend/dist")
	if err != nil {
		return nil, err
	}

	db := cfg.Database
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", db.Username, db.Password, db.Hostname, db.Port, db.Database)

	mailTemplates, err = template.ParseFS(templatesFS, "templates/*")
	if err != nil {
		return nil, err
	}

	sqlxDb, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	argonParams := &ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	dalObject := dal.NewDal(sqlxDb)

	return &Provider{
		Webapp:      &correctedFS,
		DB:          sqlxDb,
		Dal:         &dalObject,
		ArgonParams: argonParams,
		Config:      cfg,

		languageMatcher: language.NewMatcher([]language.Tag{
			language.English,
			language.French,
		}),
	}, nil
}

func (prv *Provider) GenerateUID() (string, error) {
	return gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 10)
}

func (prv *Provider) Shutdown() {
	prv.DB.Close()
}

// HashPassword will generate a hash of a password ready to be stored in the database
func (prv *Provider) HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(prv.ArgonParams.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, prv.ArgonParams.Iterations, prv.ArgonParams.Memory, prv.ArgonParams.Parallelism, prv.ArgonParams.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, prv.ArgonParams.Memory, prv.ArgonParams.Iterations, prv.ArgonParams.Parallelism, b64Salt, b64Hash), nil
}

// VerifyPassword will check for the validity of the password in the database
func (prv *Provider) VerifyPassword(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		// @TODO Log error in DB for "login attempt" section in the user profile
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

// ArgonParams represents all the parameters needed to hash the passwords
type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

/** Thanks to https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go **/
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

/** Errors specific to Scinna **/
var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func decodeHash(encodedHash string) (p *ArgonParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &ArgonParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}

// SendMail sends a mail
func (prv *Provider) SendMail(dest, subject, lang string, data interface{}) (bool, error) {
	body := bytes.Buffer{}
	headers := []byte(fmt.Sprintf("Subject: %v\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", subject))
	body.Write(headers)
	err := mailTemplates.ExecuteTemplate(&body, "validation_email." + lang + ".gohtml", data)
	if err != nil {
		// Unfortunately we'll have to go the ugly way since
		// golang's templating thing is hardcoded with a fmt.Errorf
		if err.Error()[len(err.Error())-12:] == "is undefined" {
			fmt.Println("Unknown lang: " + lang)
			err := mailTemplates.ExecuteTemplate(&body, "validation_email.en.gohtml", data)
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	// @TODO: Add a way to stay connected to the SMTP server for X amt of time (If you have a huge userbase on your server you might not want to open a connection each email sent)
	smtpCfg := prv.Config.Mail
	prv.MailClient = smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Hostname)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpCfg.Hostname + ":" + strconv.Itoa(smtpCfg.Port),
	}

	c, err := smtp.Dial(smtpCfg.Hostname + ":" + strconv.Itoa(smtpCfg.Port))
	if err != nil {
		return false, err
	}

	if prv.Config.Mail.ConnectionType == "starttls" {
		err = c.StartTLS(tlsConfig)
		if err != nil {
			return false, err
		}
	}

	if err := smtp.SendMail(smtpCfg.Hostname+":"+strconv.Itoa(smtpCfg.Port), prv.MailClient, smtpCfg.Sender, []string{dest}, body.Bytes()); err != nil {
		return false, err
	}
	return true, nil
}

// SendValidationMail sends the validation mail for a user
func (prv *Provider) SendValidationMail(r *http.Request, dest, validationCode string) (bool, error) {
	url := prv.Config.WebURL
	if url[len(url)-1:] != "/" {
		url = url + "/"
	}

	url += "app/validate/" + validationCode

	acceptLanguage, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		return false, nil
	}
	tag, _, _ := prv.languageMatcher.Match(acceptLanguage...)

	// It looks like golang can't get the 2 character code for the country
	// so yeah we'll fallback for now. >:(

	//return prv.SendMail(dest, translations.TLang(tag.String(), "registration.validation_email.subject"), tag.String(), struct {
	return prv.SendMail(dest, "Scinna: Activate your account", tag.String(), struct {
		Url string
	} {
		Url: url,
	})
}
