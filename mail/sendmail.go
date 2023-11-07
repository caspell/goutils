package mail

import (
	"crypto/tls"
	"net"
	"net/mail"
	"net/smtp"
)

type SmtpConfig struct {
	SmtpServer      string `toml:"SMTP_SERVER"`
	AuthUsername    string `toml:"AUTH_USERNAME"`
	AuthPassword    string `toml:"AUTH_PASSWORD"`
	EmailSenderName string `toml:"EMAIL_SENDER_NAME"`
	EmailSender     string `toml:"EMAIL_SENDER"`
}

type MailSender struct {
	Config       SmtpConfig
	HtmlTemplate string
}

type MailMessage struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Body    string
	IsHtml  bool
}

func (m *MailMessage) SetFrom(name, email string) {
	m.From = mail.Address{
		Name:    name,
		Address: email,
	}
}

func (m *MailMessage) SetTo(name, email string) {
	m.To = mail.Address{
		Name:    name,
		Address: email,
	}
}

func (m *MailMessage) ToString() string {
	return "From: " + m.From.String() + "\r\n" +
		"To: " + m.To.String() + "\r\n" +
		"Subject: " + m.Subject + "\r\n\r\n" +
		m.Body
}

func (m *MailMessage) ToHtml() []byte {
	subject := "From: " + m.From.String() + "\r\n" +
		"To: " + m.To.String() + "\r\n" +
		"Subject: " + m.Subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	body := m.Body
	return []byte(subject + mime + "\n" + body)
}

func (s *MailSender) GetEmailAddress() string {
	return s.Config.EmailSender
}

func (s *MailSender) Send(message *MailMessage) error {

	// Connect to the SMTP Server
	servername := s.Config.SmtpServer

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", s.Config.AuthUsername, s.Config.AuthPassword, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		return err
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(message.From.Address); err != nil {
		return err
	}

	if err = c.Rcpt(message.To.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	if message.IsHtml {
		_, err = w.Write(message.ToHtml())
		if err != nil {
			return err
		}
	} else {
		_, err = w.Write([]byte(message.ToString()))
		if err != nil {
			return err
		}
	}

	defer w.Close()
	defer c.Quit()

	return nil
}
