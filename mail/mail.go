package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
)

func (m *Mail) getURL() string {
	return fmt.Sprintf("%s:%d", m.smtp.host, m.smtp.port)
}

// New construct Mail struct
func New(from, to, subject, password, smtpHost string, smtpPort int) *Mail {
	return &Mail{From: from, To: to, Subject: subject, Password: password, smtp: struct {
		host string
		port int
	}{smtpHost, smtpPort}}
}

// Mail main struct for send mail
type Mail struct {
	buffer   bytes.Buffer
	From     string
	To       string
	Subject  string
	Password string
	smtp     struct {
		host string
		port int
	}
}

func (m *Mail) writeFrom() error {
	_, err := m.buffer.WriteString(fmt.Sprintf("From: %s\n", m.From))
	return err
}

func (m *Mail) writeTo() error {
	_, err := m.buffer.WriteString(fmt.Sprintf("To: %s\n", m.To))
	return err
}

func (m *Mail) writeSubject() error {
	_, err := m.buffer.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	return err
}

func (m *Mail) writeEndHeader() error { _, err := m.buffer.WriteString("\n"); return err }

func (m *Mail) writeHeader() error {
	if err := m.writeFrom(); err != nil {
		return err
	}

	if err := m.writeTo(); err != nil {
		return err
	}

	if err := m.writeSubject(); err != nil {
		return err
	}

	m.buffer.WriteString("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n")

	if err := m.writeEndHeader(); err != nil {
		return nil
	}
	return nil
}

// SendMail send mail
func (m *Mail) SendMail(body bytes.Buffer) error {

	if err := m.writeHeader(); err != nil {
		return err
	}

	if _, err := body.WriteTo(&m.buffer); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", m.From, m.Password, m.smtp.host)

	return smtp.SendMail(m.getURL(), auth, m.From, []string{m.To}, m.buffer.Bytes())
}
