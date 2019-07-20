package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
	"log"
)

func main() {

	//{
	//	"Username": "User Name",
	//	"Password": "",
	//	"From": "Site name",  or your domine
	//	"Server": "",
	//	"Port": 587
	//}

	// using smpt
	log.Println(sendEmail("template.html", "to email", "Testing"))

	// no smtp
	//Nosmtp()

}

func sendEmail(htmlTemplate, to, subject string) bool {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "Sender email")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", htmlTemplate)
	dialer := gomail.NewPlainDialer("{server name}", 587, "{user name}", "{password}")
	err := dialer.DialAndSend(msg)
	if err != nil {
		log.Print("Sending mail to "+to+": ", err)
		return false
	}

	return true
}

func Nosmtp() {
	m := gomail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", "")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	s := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
		// Implements you email-sending function, for example by calling
		// an API, or running postfix, etc.
		fmt.Println("From:", from)
		fmt.Println("To:", to)
		return nil
	})

	if err := gomail.Send(s, m); err != nil {
		panic(err)
	}
}
