package seeders

import (
	"errors"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"syreclabs.com/go/faker"
)

var (
	totalEstablishment = 0
	totalUser          = 0
	totalProduct       = 0
	totalAddress       = 0
)

func SeederAddress() error {
	m := fakerAddress()
	err := controller.CreateAddress(&m)
	return err
}

func SeederEstablishment() error {
	address := fakerAddress()
	establishment := model.Establishment{Address: address}
	amountTables := faker.RandomInt(5, 20)
	return controller.CreateEstablishmentWithTables(&establishment, amountTables)
}

func SeederProduct() error {
	m := fakerProduct()
	return controller.CreateProduct(&m)
}

func SeederUser() error {
	m, err := FakerUser()
	if err != nil {
		return err
	}
	return storage.DB().Create(&m).Error
}

func SeederPay() error {
	m1 := model.Pay{
		Name: `Efectivo`,
	}
	m2 := model.Pay{
		Name: `Paypal`,
	}
	err := controller.CreatePay(&m1)
	if err != nil {
		return err
	}
	return controller.CreatePay(&m2)
}

func SeederOrder() error {
	m, err := fakerOrder()
	if err != nil {
		return err
	}
	return controller.CreateOrder(&m)
}

func SeederAll(establishments, addresses, products, users, orders int) error {
	totalEstablishment = establishments
	totalAddress = establishments + addresses
	totalProduct = products
	totalUser = users
	if totalUser < 0 || totalAddress < 0 || totalProduct < 0 || totalEstablishment < 0 {
		return errors.New("ingresa un valor positivo mayor a uno al utilizar seeders")
	}
	SeederPay()
	for i := 0; i < establishments; i++ {
		SeederEstablishment()
	}
	for i := 0; i < addresses; i++ {
		SeederAddress()
	}
	for i := 0; i < products; i++ {
		SeederProduct()
	}
	for i := 0; i < users; i++ {
		SeederUser()
	}
	for i := 0; i < orders; i++ {
		SeederOrder()
	}
	return nil
}
