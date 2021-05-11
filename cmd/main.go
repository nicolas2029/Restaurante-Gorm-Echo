package main

import (
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/route"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	storage.New(storage.Postgres)
	sessionsCookie.NewCookieStore("certificates/cookieKey")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessionsCookie.Cookie()))
	route.All(e)
	e.Start(":8080")
}
