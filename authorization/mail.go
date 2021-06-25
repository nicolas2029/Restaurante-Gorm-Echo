package authorization

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"sync"
	"text/template"
)

type jsonEmail struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type send struct {
	Url string
}

var (
	auth     smtp.Auth
	config   *jsonEmail
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
	config = &jsonEmail{}
	err = json.Unmarshal(m, config)
	if err != nil {
		return err
	}
	auth = LoginAuth(config.User, config.Password)
	return nil
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func sendMessage(message string, to, from mail.Address) error {
	if config.Host == "smtp.gmail.com" {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.gmail.com",
		}

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
	} else {
		return smtp.SendMail(config.Host+":"+config.Port, auth, config.Email, []string{to.Address}, []byte(message))
	}
}

func confirmEmail(dest, token string) error {
	from := mail.Address{
		Name:    "Beer para Creer",
		Address: config.Email,
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
	message += "\r\n"
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
	message += buf.String() + "\r\n"

	return sendMessage(message, to, from)
}

func ConfirmEmail(dest, token string) error {
	return confirmEmail(dest, token)
}
