package initializers

import (
	"log"
	"os"
	"strconv"
)

type MailServiceConfiguration struct{
	Host 				string 
	Port 				int
	Username 			string
	Password 			string
	DefaultSender 		string
}

func ExtractMailConfig() MailServiceConfiguration{
	var config MailServiceConfiguration
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil{
		log.Fatal("Invalid smtp port number: " + os.Getenv("SMTP_PORT"))
	}
	config.Host = os.Getenv("SMPT_HOST")
	config.Port =  port
	config.Username = os.Getenv("SMTP_USERNAME")
	config.Password = os.Getenv("SMTP_PASSWORD")
	config.DefaultSender = os.Getenv("SMTP_DEFAULT_SENDER")
	return config
}