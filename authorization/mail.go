package authorization

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"sync"
	"text/template"
)

type jsonEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type send struct {
	Url string
}

var (
	auth     smtp.Auth
	email    string
	onceMail sync.Once
)

// LoadFiles .
func LoadMail(file string) error {
	var err error
	onceMail.Do(func() {
		err = loadMail(file)
	})

	return err
}

func loadMail(file string) error {
	m, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	u := &jsonEmail{}
	err = json.Unmarshal(m, u)
	if err != nil {
		return err
	}
	auth = smtp.PlainAuth("", u.Email, u.Password, "smtp.gmail.com")
	email = u.Email
	return nil
}

func confirmEmail(dest, token string) error {
	from := mail.Address{
		Name:    "Beer para Creer",
		Address: email,
	}
	to := mail.Address{Address: dest}
	subject := "Validacion de correo electronico con Beerparacreer"

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("../public/template/confirm.html")
	if err != nil {
		return err
	}
	s := send{Url: "http://localhost/api/v1/user/validate/" + token}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, s)
	if err != nil {
		return err
	}

	message += buf.String()
	//auth := smtp.PlainAuth("", "beerparacreergr@gmail.com", "qytxdkmwebrolpom", "smtp.gmail.com")
	//
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.gmail.com",
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", "smtp.gmail.com:465", tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, "smtp.gmail.com")
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()
	//

	return nil
}

func ConfirmEmail(dest, token string) error {
	return confirmEmail(dest, token)
}
