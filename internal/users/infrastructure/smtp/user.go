package infrastructure

import (
	"bytes"
	"fmt"
	"github.com/ismailbayram/shopping/config"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"text/template"
)

type EmailSender struct {
	from     string
	host     string
	port     string
	password string
}

func NewEmailSender(cfg config.SMTPConfiguration) *EmailSender {
	return &EmailSender{
		from:     cfg.From,
		host:     cfg.Host,
		port:     cfg.Port,
		password: cfg.Password,
	}
}

func (es *EmailSender) SendWelcomeEmail(verificationToken string, email string) {
	auth := smtp.PlainAuth("", es.from, es.password, es.host)

	t, _ := template.ParseFiles("./internal/users/infrastructure/smtp/verification.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verify Your Account \n%s\n\n", mimeHeaders)))

	err := t.Execute(&body, map[string]string{
		"token": verificationToken,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"email": email,
		}).Error("SEND_WELCOME_EMAIL_TEMPLATE_ERROR")
		return
	}

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", es.host, es.port),
		auth,
		es.from,
		[]string{email},
		body.Bytes(),
	)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"auth":  auth,
			"email": email,
			"body":  body,
		}).Error("SEND_WELCOME_EMAIL_ERROR")
		return
	}
}
