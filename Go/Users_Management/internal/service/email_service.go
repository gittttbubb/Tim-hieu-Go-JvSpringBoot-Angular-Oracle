package service

import (
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	Send(to, subject, body string) error
}
type emailService struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewEmailService(
	host string,
	port int,
	username string,
	password string,
	from string,
) EmailService {
	return &emailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (s *emailService) Send(to string, subject string, body string,) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(
		s.host,
		s.port,
		s.username,
		s.password,
	)
	return d.DialAndSend(m)
}