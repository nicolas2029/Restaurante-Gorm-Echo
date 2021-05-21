package controller_test

import (
	"log"
	"testing"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
)

func TestCreateOrder(t *testing.T) {
	//m := model.OrderOrderProduct{
	//	Order: &model.Order{},
	//}
	//controller.HashPassword(&m)
	//log.Fatalf("pass: %s", m.Password)
}

func TestGetAllOrderByUser(t *testing.T) {
	m, err := controller.GetAllOrderByUser(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatalf("%+v", m)
}
