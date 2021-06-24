package controller_test

import (
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func TestLoginUser(t *testing.T) {
	err := authorization.LoadFiles("../cmd/certificates/app.rsa", "../cmd/certificates/app.rsa.pub")
	if err != nil {
		t.Errorf("Error in LoadFiles: %v", err)
	}
	storage.New("../cmd/certificates/db.json")
	m := model.Login{
		Email:    "nivoh11811@bbsaili.com",
		Password: "Contra12345.",
	}
	_, err = controller.Login(&m)
	if err != nil {
		t.Errorf("Error in Login: %v", err)
	}
}

func TestLoginUserWithInvalidPassword(t *testing.T) {
	err := authorization.LoadFiles("../cmd/certificates/app.rsa", "../cmd/certificates/app.rsa.pub")
	if err != nil {
		t.Errorf("Error in LoadFiles: %v", err)
	}
	storage.New("../cmd/certificates/db.json")
	m := model.Login{
		Email:    "nivoh11811@bbsaili.com",
		Password: "Contra12345.1",
	}
	_, err = controller.Login(&m)
	if err == nil {
		t.Errorf("Error in Login: %v", err)
	}
}

func TestHavePermission(t *testing.T) {
	storage.New("../cmd/certificates/db.json")
	_, _, err := controller.HavePermission(1, 1)
	if err != nil {
		t.Errorf("Error in HavePermission: %v", err)
	}
}

func TestCreateOrder(t *testing.T) {
	storage.New("../cmd/certificates/db.json")
	m := model.OrderOrderProduct{
		Order:        &model.Order{},
		OrderProduct: make([]*model.OrderProduct, 0),
	}
	err := controller.CreateOrder(&m)
	if err == nil {
		t.Errorf("Error in CreateOrder empty")
	}
}

func TestGetAllOrderByUser(t *testing.T) {
	storage.New("../cmd/certificates/db.json")
	_, err := controller.GetAllOrderByUser(1)
	if err != nil {
		t.Errorf("Error in GetAllOrderByUser: %v", err)
	}
}

func TestCreateProduct(t *testing.T) {
	d := "El tercer producto en la base de datos"
	storage.New("../cmd/certificates/db.json")
	m := model.Product{
		Name:        "Producto 3",
		Price:       333.33,
		Description: &d,
	}
	err := controller.CreateProduct(&m)
	if err != nil {
		t.Errorf("Error in AddProduct = %v", err)
	}
}

func TestGetProducts(t *testing.T) {
	storage.New("../cmd/certificates/db.json")
	_, err := controller.GetAllProduct()
	if err != nil {
		t.Errorf("GetProducts() = %w", err)
	}
}
