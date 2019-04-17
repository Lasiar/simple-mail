package main

import (
	"bytes"
	"flag"
	"log"
	"mail-sender/mail"
)

var (
	from, to, subject, password, smtpHost, message string
	smtpPort                                       int
)

func init() {
	flag.StringVar(&from, "from", "", "")
	flag.StringVar(&to, "to", "", "")
	flag.StringVar(&subject, "subject", "", "")
	flag.StringVar(&password, "password", "", "")
	flag.StringVar(&smtpHost, "host", "", "")
	flag.StringVar(&message, "message", "", "")
	flag.IntVar(&smtpPort, "port", 0, "")
	flag.Parse()
}

func main() {
	m := mail.New(from, to, subject, password, smtpHost, smtpPort)
	buf := new(bytes.Buffer)
	buf.WriteString(message)
	if err := m.SendMail(*buf); err != nil {
		log.Fatalf("err")
	}
}
