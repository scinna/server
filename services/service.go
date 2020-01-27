// Package services with the Provider struct
package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"fmt"

	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/scinna/server/configuration"
	"github.com/scinna/server/utils"

	"golang.org/x/crypto/argon2"
)

// Provider is the struct that carry all the parameters / connections for the software
type Provider struct {
	Db          *sqlx.DB
	Mail        utils.MailClient
	Templates   *template.Template
	ArgonParams *ArgonParams
	Config      configuration.Configuration
}

// GenerateUID function generates an ID for the pictures
func (prv *Provider) GenerateUID() (string, error) {
	alphabet, exists := os.LookupEnv("ID_ALPHABET")
	if !exists {
		alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ"
	}

	lengthStr, exists := os.LookupEnv("ID_LENGTH")
	length := 10
	if exists {
		length, _ = strconv.Atoi(lengthStr)
	}

	return gonanoid.Generate(alphabet, length)
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
		// @TODO Log error in DB
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func parseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}

// New function initializes the the Provider structure
func New(cfg configuration.Configuration) *Provider {
	t, err := parseTemplateDir("templates")
	if err != nil {
		fmt.Println(err)
		panic("Can't load templates!")
	}
	fmt.Println("- Templates loaded")

	db := utils.LoadDatabase(cfg.PostgresDSN)
	fmt.Println("- Connected to database")

	argonParams := &ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return &Provider{
		Config:      cfg,
		Db:          db,
		Mail:        utils.LoadMail(cfg),
		Templates:   t,
		ArgonParams: argonParams,
	}
}

// Render renders a template to the writer
func Render(prv *Provider, w http.ResponseWriter, data interface{}) error {
	return prv.Templates.ExecuteTemplate(w, "index.html", &data)
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
