package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
)

func createCookie(r *http.Request, w *http.ResponseWriter, token string) error {
	s, err := sessionsCookie.Cookie().Get(r, "session")
	if err != nil {
		return err
	}
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   1800,
		HttpOnly: true,
	}
	s.Values["token"] = token
	err = s.Save(r, *w)
	return err
}
