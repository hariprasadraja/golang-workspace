package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"io"
	"fmt"
)

func main() {

}
func sendEmail(htmlTemplate, to, subject string) bool {
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
