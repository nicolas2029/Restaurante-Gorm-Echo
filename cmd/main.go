package main

import (
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/route"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	storage.New("certificates/db.json")
	err = sessionsCookie.NewCookieStore("certificates/cookieKey")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para cookies: %v", err)
	}

	err = authorization.LoadMail("certificates/email.json")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para email: %v", err)
	}
	storage.DB().AutoMigrate(
		&model.Product{},
		&model.Address{},
		&model.Permission{},
		&model.Establishment{},
		&model.Rol{},
		&model.User{},
		&model.Table{},
		&model.Pay{},
		&model.Order{},
		&model.OrderProduct{},
	)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessionsCookie.Cookie()))
	route.All(e)
	err = e.Start(":80")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
