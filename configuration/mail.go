package configuration

import (
	"net/smtp"
	"regexp"
)

// MailConfig represents the SMTP configuration
type MailConfig struct {
	Client smtp.Auth `yaml:"-"`

	Hostname string
	Port     string
	Username string
	Password string
	Sender   string

	IsEmail *regexp.Regexp `yaml:"-"`
}

// SendMail sends a mail
func (mc *MailConfig) SendMail(dest, subject, body string) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "!\n"
	msg := []byte(subject + mime + "\n" + body)

	if err := smtp.SendMail(mc.Hostname+":"+mc.Port, mc.Client, mc.Sender, []string{dest}, msg); err != nil {
		return false, err
	}
	return true, nil
}

// SendValidationMail sends the validation mail for a user
func (mc *MailConfig) SendValidationMail(url, dest, validationCode string) (bool, error) {
	if url[len(url)-1:] != "/" {
		url = url + "/"
	}

	return mc.SendMail(dest, "Scinna: Activate your account", `
		Please validate your account.
		`+url+`auth/register/`+validationCode)
}
