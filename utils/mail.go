package utils

import (
	"fmt"
	"net/smtp"
	"regexp"

	"github.com/oxodao/scinna/configuration"
)

// MailClient handles everything needed to connect to a SMTP user and send a mail
type MailClient struct {
	smtpHost   string
	smtpUser   string
	SMTPClient smtp.Auth
	IsEmail    *regexp.Regexp
}

// LoadMail loads the mail client
func LoadMail(cfg configuration.Configuration) MailClient {

	var smtpAuth smtp.Auth

	if len(cfg.SMTPUser) > 0 && len(cfg.SMTPPass) > 0 {
		smtpAuth = smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPPass)
	}

	reg, err := regexp.Compile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	if err != nil {
		fmt.Println(err)
	}

	return MailClient{SMTPClient: smtpAuth, smtpUser: cfg.SMTPSender, smtpHost: cfg.SMTPHost, IsEmail: reg}
}

// SendMail sends a mail
func (mc *MailClient) SendMail(dest, subject, body string) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "!\n"
	msg := []byte(subject + mime + "\n" + body)

	if err := smtp.SendMail(mc.smtpHost, mc.SMTPClient, mc.smtpUser, []string{dest}, msg); err != nil {
		return false, err
	}
	return true, nil
}

// SendValidationMail sends the validation mail for a user
func (mc *MailClient) SendValidationMail(url, dest, validationCode string) (bool, error) {
	if url[len(url)-1:] != "/" {
		url = url + "/"
	}

	return mc.SendMail(dest, "Scinna: Activate your account", `
		Please validate your account.
		`+url+`auth/register/`+validationCode)
}
