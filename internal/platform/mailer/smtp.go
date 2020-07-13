package mailer

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strconv"

	"github.com/matcornic/hermes/v2"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/gomail.v2"

	generator "Gogin/internal/helpers/emailTemplate"
)

type SmtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// SendOptions are options for sending an email
type SendOptions struct {
	To      string
	Subject string
}

// send sends the email
func send(smtpConfig SmtpAuthentication, options SendOptions, htmlBody string, txtBody string) error {

	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}

func SendEmail(email, subject string, template hermes.Email) {

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	password := os.Getenv("SMTP_PASSWORD")
	SMTPUser := os.Getenv("SMTP_USERNAME")
	if password == "" {
		fmt.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
		bytePassword, _ := terminal.ReadPassword(0)
		password = string(bytePassword)
	}
	smtpConfig := SmtpAuthentication{
		Server:         os.Getenv("SMTP_HOST"),
		Port:           port,
		SenderEmail:    SMTPUser,
		SenderIdentity: os.Getenv("PRODUCT_NAME"),
		SMTPPassword:   password,
		SMTPUser:       SMTPUser,
	}
	options := SendOptions{
		To: email,
	}

	options.Subject = subject
	htmlBytes, txtBytes := generator.Export(template)

	err := send(smtpConfig, options, string(htmlBytes), string(txtBytes))
	if err != nil {
		fmt.Println(err)
	}
}
