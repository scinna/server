package configuration

import (
	"crypto/tls"
	"net/smtp"
	"regexp"
)

// ServerType lets you choose the type of connection used to log in to SMTP server
type ServerType int

const (
	// NoAuth represents a server with plain-text auth (Most likely on port 25)
	NoAuth ServerType = 0

	// SMTPS represents a server with simple TLS auth (Port 465)
	SMTPS ServerType = 1

	// STARTTLS represents a server with StartTLS (Port 587, the most likely you'll encounter)
	STARTTLS ServerType = 2
)

// MailConfig represents the SMTP configuration
type MailConfig struct {
	Client smtp.Auth `yaml:"-"`

	Enabled bool `json:"Enabled"`

	ConnectionType ServerType `json:"ConnectionType"`

	Hostname string `json:"Hostname"`
	Port     string `json:"Port"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Sender   string `json:"Sender"`

	TestReceiver string `json:"TestReceiver" yaml:"-"`

	IsEmail *regexp.Regexp `yaml:"-"`
}

// SendMail sends a mail
func (mc *MailConfig) SendMail(dest, subject, body string) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "!\n"
	msg := []byte(subject + mime + "\n" + body)

	/**

		@TODO:
			Support non-starttls auth
			Add a way to stay connected to the SMTP server for X amt of time (If you have a huge userbase on your server you might not want to open a connection each email sent)

	**/
	mc.Client = smtp.PlainAuth("", mc.Username, mc.Password, mc.Hostname)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         mc.Hostname + ":" + mc.Port,
	}

	c, err := smtp.Dial(mc.Hostname + ":" + mc.Port)
	if err != nil {
		return false, err
	}

	err = c.StartTLS(tlsConfig)
	if err != nil {
		return false, err
	}

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
