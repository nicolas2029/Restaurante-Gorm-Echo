package controller_test

import (
	"errors"
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func TestCreateProduct(t *testing.T) {
	d := "El tercer producto en la base de datos"
	storage.New(storage.Postgres)
	m := model.Product{
		Name:        "Producto 3",
		Price:       333.33,
		Description: &d,
	}
	err := controller.CreateProduct(&m)
	if err != nil {
		t.Errorf("AddProduct(m) = %d", err)
	}
}

func TestGetProducts(t *testing.T) {
	storage.New(storage.Postgres)
	ps, err := controller.GetAllProduct()
	if err != nil {
		t.Errorf("GetProducts() = %w", err)
	}
	if len(ps) != 3 {
		t.Errorf("GetProducts() = %w %d", errors.New("se esperaban 3 se obtuvo"), len(ps))
	}

}
