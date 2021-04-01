package services

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/pariz/gountries"
	"golang.org/x/text/language"
	"net/http"
	"net/smtp"
	"strconv"
)

var countries = gountries.New()

// SendMail sends a mail
func (prv *Provider) SendMail(dest, subject, lang string, template string, data interface{}) (bool, error) {
	body := bytes.Buffer{}
	headers := []byte(fmt.Sprintf("Subject: %v\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n", subject))
	body.Write(headers)
	err := mailTemplates.ExecuteTemplate(&body, template+"."+lang+".gohtml", data)
	if err != nil {
		// Unfortunately we'll have to go the ugly way since
		// golang's templating thing is hardcoded with a fmt.Errorf
		if err.Error()[len(err.Error())-12:] == "is undefined" {
			fmt.Println("Unknown lang: " + lang)
			err := mailTemplates.ExecuteTemplate(&body, template+".en.gohtml", data)
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

// @TODO: Write this using an async queueing system
// SendValidationMail sends the validation mail for a user
func (prv *Provider) SendValidationMail(r *http.Request, dest, validationCode string) (bool, error) {
	url := prv.Config.WebURL + "app/validate/" + validationCode

	acceptLanguage, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		return false, nil
	}
	tag, _, _ := prv.languageMatcher.Match(acceptLanguage...)

	// It looks like golang can't get the 2 character code for the country
	// so yeah we'll fallback for now. >:(

	//return prv.SendMail(dest, translations.TLang(tag.String(), "registration.validation_email.subject"), tag.String(), struct {
	return prv.SendMail(dest, prv.Config.CustomBranding + ": Activate your account", tag.String(), "validation_email", struct {
		Branding string
		Url string
	}{
		Branding: prv.Config.CustomBranding,
		Url: url,
	})
}

func (prv *Provider) SendForgottenPasswordMail(r *http.Request, dest, validationCode string) (bool, error) {
	url := prv.Config.WebURL + "app/forgotten_password/" + validationCode

	acceptLanguage, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err != nil {
		return false, nil
	}
	tag, _, _ := prv.languageMatcher.Match(acceptLanguage...)

	// It looks like golang can't get the 2 character code for the country
	// so yeah we'll fallback for now. >:(

	//return prv.SendMail(dest, translations.TLang(tag.String(), "registration.validation_email.subject"), tag.String(), struct {
	return prv.SendMail(dest, prv.Config.CustomBranding + ": Reset your password", tag.String(), "reset_email", struct {
		Branding string
		Url string
	}{
		Branding: prv.Config.CustomBranding,
		Url: url,
	})
}
