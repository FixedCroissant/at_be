package mail

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendMailOurMail() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "from@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "to@example.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "This is your test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is the body of your message")

	// Settings for SMTP server
	//d := gomail.NewDialer("smtp.127.0.0.1", 1025, "from@gmail.com", "")
	//Localhost
	d := gomail.Dialer{Host: "localhost", Port: 1025}

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}

//Send custom mailer
func SendMailCustom(bodyMessage string) {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "from@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "to@example.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "This is your test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyMessage)

	// Settings for SMTP server
	//d := gomail.NewDialer("smtp.127.0.0.1", 1025, "from@gmail.com", "")
	//Localhost
	d := gomail.Dialer{Host: "localhost", Port: 1025}

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
