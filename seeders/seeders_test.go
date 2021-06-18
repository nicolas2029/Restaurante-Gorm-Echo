package seeders_test

import (
	"log"
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/seeders"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func TestSeederAll(t *testing.T) {
	p := `C:/Users/ULTIMATE/Desktop/Go/restaurante-gorm-echo/cmd/certificates`
	err := authorization.LoadFiles(p+"/app.rsa", p+"/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	storage.New(p + "/db.json")
	err = sessionsCookie.NewCookieStore(p + "/cookieKey")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para cookies: %v", err)
	}

	err = authorization.LoadMail(p + "/email.json")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para email: %v", err)
	}
	seeders.SeederAll(10, 30, 30, 30, 40)
}

func TestFakerUser(t *testing.T) {
	m, _ := seeders.FakerUser()
	if m.EstablishmentID != nil {
		log.Fatal("e: ", *m.EstablishmentID)
	}
	if m.RolID != nil {
		if m.EstablishmentID != nil {
			log.Fatal("a: ", *m.EstablishmentID)
		}
		log.Fatal("b: ", *m.RolID)
	}
	log.Fatal("c:", m.RolID, m.EstablishmentID)
}
