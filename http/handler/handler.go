package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
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

func SaveImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Destination
	name := c.Param("name")
	dst, err := os.Create("../public/views/assets/img/products/" + name + ".jpg")
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
