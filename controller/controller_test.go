package controller_test

import (
	"log"
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func TestHashAndSalt(t *testing.T) {
	m := model.Login{
		Email:    "nicolas2029",
		Password: "123456",
	}
	controller.HashPassword(&m)
	log.Fatalf("pass: %s", m.Password)
}

func TestHavePermission(t *testing.T) {
	storage.New(storage.Postgres)
	err := controller.HavePermission(1, 1)
	if err != nil {
		t.Fatalf("No se tiene permiso %v", err)
	}
	t.Fatal("Se tieme permiso: ")
}
