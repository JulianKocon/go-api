package services

import (
	"crypto/tls"
	"example/go-api/initializers"

	"gopkg.in/gomail.v2"
)



type MailService interface {
	SendMail(htmlBody, subject, target string) error
}

type mailService struct{
	config initializers.MailServiceConfiguration
}

func NewMailService(config initializers.MailServiceConfiguration) MailService{
	return mailService{
		config: config,
	}
}

func (ms mailService) SendMail(htmlBody, subject, target string) error{
	message := gomail.NewMessage()
	d := gomail.NewDialer(ms.config.Host, ms.config.Port, ms.config.Username, ms.config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	message.SetBody("text/html", htmlBody)
	message.SetHeader("From", ms.config.DefaultSender)
	message.SetHeader("To", target)
	message.SetHeader("Subject", subject)
	if err := d.DialAndSend(message); err != nil{return err}
	return nil
}