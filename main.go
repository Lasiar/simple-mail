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
	verbose                                        bool
)

func init() {
	flag.StringVar(&from, "from", "", "sender's mail, example: test@example.com")
	flag.StringVar(&to, "to", "", "recipient's mail, example: test@example.com")
	flag.StringVar(&subject, "subject", "", "")
	flag.StringVar(&password, "password", "", "password auth")
	flag.StringVar(&smtpHost, "host", "", "host smpt server")
	flag.StringVar(&message, "message", "", "")
	flag.BoolVar(&verbose, "verbose", false, "print debug info")
	flag.IntVar(&smtpPort, "port", 0, "port smtp server")
	flag.Parse()
}

func main() {
	if verbose {
		log.Println(from, to, subject, password, smtpHost, smtpPort)
	}
	m := mail.New(from, to, subject, password, smtpHost, smtpPort)
	m.Debug = true
	buf := new(bytes.Buffer)
	buf.WriteString(message)
	if err := m.SendMail(*buf); err != nil {
		log.Fatalf("err")
	}
}
