package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func main() {
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	// log.Fatalln(err)
	// }
	// err = tpl.Execute(os.Stdout, 5*5)
	// if err != nil {
	// 	// log.Fatalln(err)
	// }
	// t := template.New("action")

	var err error
	t, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, 5*5); err != nil {
		panic(err)
	}

	result := tpl.String()
	fmt.Println(result)

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
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Kết quả thi đấu WorldCup" + "!\n"
	err = smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"olivertran732169@gmail.com",
		users,
		[]byte(subject+mime+"\n"+result),
	)

}
