package services

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/scinna/server/config"
	"github.com/scinna/server/dal"
	"golang.org/x/text/language"
	"html/template"
	"net/smtp"
)

var mailTemplates *template.Template

type Provider struct {
	EmbeddedAssets *Assets

	DB          *sqlx.DB
	Dal         *dal.Dal
	ArgonParams *ArgonParams
	MailClient  smtp.Auth
	Config      *config.Config

	languageMatcher language.Matcher
}

func NewProvider(cfg *config.Config, embededAssets *Assets) (*Provider, error) {
	db := cfg.Database
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", db.Username, db.Password, db.Hostname, db.Port, db.Database)

	var err error
	mailTemplates, err = template.ParseFS(*embededAssets.Templates, "assets/templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	sqlxDb, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	dalObject := dal.NewDal(sqlxDb)

	return &Provider{
		EmbeddedAssets: embededAssets,
		DB:          sqlxDb,
		Dal:         &dalObject,
		ArgonParams: defaultArgonParams(),
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