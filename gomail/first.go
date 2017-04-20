package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"io"
	"fmt"
)

func main() {
	//{
	//	"Username": "Autoemail@bevopos.com",
	//	"Password": "Bevopos@2015",
	//	"From": "Zenpepper<support@zenpepper.com>",
	//	"Server": "west.EXCH031.serverdata.net",
	//	"Port": 587
	//}

	// using smpt
	//log.Println(sendEmail("template.html","hariprasadcsmails@gmail.com","Testing"))

	// no smtp
	//Nosmtp()

}
func sendEmail(htmlTemplate, to, subject string) bool {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "Hariprasad<hariprasadcsmails@gmail.com>")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", htmlTemplate)
	dialer := gomail.NewPlainDialer("west.EXCH031.serverdata.net",587, "Autoemail@bevopos.com","Bevopos@2015")
	err := dialer.DialAndSend(msg)
	if err != nil {
		log.Print("Sending mail to "+to+": ", err)
		return false
	}

	return true
}
func Nosmtp(){
	m := gomail.NewMessage()
	m.SetHeader("From", "hariprasadcsmails@gmail.com	")
	m.SetHeader("To", "Hariprasad@benseron.com")
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