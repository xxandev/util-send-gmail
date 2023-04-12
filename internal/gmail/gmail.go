package gmail

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

// Configuration - gmail sender configuration
type Configuration interface {
	//GetFrom - return from email
	GetFrom() string

	//GetPass - return password email
	GetPass() string
}

type gmail struct{ c Configuration }

// New - return new gmail object
func New(c Configuration) *gmail { return &gmail{c: c} }

func (g gmail) Send(to, subject, body string) error {
	if len(to) < 8 {
		return errors.New("invalid mails to")
	}
	if len(body) < 1 {
		return errors.New("mails body cannot be empty")
	}
	return smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", g.c.GetFrom(), g.c.GetPass(), "smtp.gmail.com"),
		g.c.GetFrom(), strings.Split(to, ","), []byte(
			fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
				g.c.GetFrom(), to, subject, body,
			),
		),
	)
}
