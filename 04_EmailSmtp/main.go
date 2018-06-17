// package main

// import (
// 	"bytes"
// 	"log"
// 	"net/smtp"
// )

// func main() {
// 	// Connect to the remote SMTP server.
// 	c, err := smtp.Dial("smtp.gmail.com:25")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()
// 	// Set the sender and recipient.
// 	c.Mail("congluc19297@gmail.com")
// 	c.Rcpt("congluc19297@gmail.com")
// 	// Send the email body.
// 	wc, err := c.Data()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer wc.Close()
// 	buf := bytes.NewBufferString("This is the email body.")
// 	if _, err = buf.WriteTo(wc); err != nil {
// 		log.Fatal(err)
// 	}
// }
package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	users := []string{"congluc19297@gmail.com", "1511917@hcmut.edu.vn"}

	auth := smtp.PlainAuth(
		"",
		"olivertran732169@gmail.com",
		"ttcnpm_nhom21",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"olivertran732169@gmail.com",
		users,
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
