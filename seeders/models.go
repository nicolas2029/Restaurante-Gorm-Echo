package seeders

import (
	"strconv"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"syreclabs.com/go/faker"
)

func fakerAddress() model.Address {
	line2 := faker.Address().SecondaryAddress()
	m := model.Address{
		Line1:      faker.Address().StreetAddress(),
		Line2:      &line2,
		City:       faker.Address().City(),
		PostalCode: faker.Address().Postcode(),
		State:      faker.Address().State(),
		Country:    faker.Address().Country(),
	}
	return m
}

func fakerProduct() model.Product {
	des := faker.Lorem().Sentence(5)
	img := faker.Avatar().Url("jpg", 800, 822)
	color := faker.Commerce().Color()
	name := faker.Commerce().ProductName() + " " + color
	for len(name) < 25 {
		name = faker.Commerce().ProductName() + " " + color
	}
	m := model.Product{
		Name:        name,
		Price:       float64(faker.Commerce().Price()),
		Description: &des,
		Img:         &img,
	}
	return m
}

func fakerOrderProduct() (model.OrderProduct, error) {
	ProductID, err := strconv.ParseUint(faker.Number().Between(1, totalProduct), 10, 0)
	if err != nil {
		return model.OrderProduct{}, err
	}
	Amount, err := strconv.ParseUint(faker.Number().Between(1, 20), 10, 0)
	if err != nil {
		return model.OrderProduct{}, err
	}
	op := model.OrderProduct{
		ProductID: uint(ProductID),
		Amount:    uint(Amount),
	}
	return op, nil
}

func fakerOrder() (model.OrderOrderProduct, error) {
	payID, err := strconv.ParseUint(faker.Number().Between(1, 2), 10, 0)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	PayID := uint(payID)
	userID, err := strconv.ParseUint(faker.Number().Between(1, totalUser), 10, 0)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	UserID := uint(userID)
	EstablishmentID, err := strconv.ParseUint(faker.Number().Between(1, totalEstablishment), 10, 0)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	addresID, err := strconv.ParseUint(faker.Number().Between(1, totalAddress), 10, 0)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	AddresID := uint(addresID)
	Amount, err := strconv.ParseInt(faker.Number().Between(1, 20), 10, 0)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	order := model.Order{
		PayID:           &PayID,
		UserID:          &UserID,
		EstablishmentID: uint(EstablishmentID),
		AddressID:       &AddresID,
	}
	ops := make([]*model.OrderProduct, Amount)
	for i := 0; i < int(Amount); i++ {
		o, err := fakerOrderProduct()
		if err != nil {
			return model.OrderOrderProduct{}, err
		}
		ops[i] = &o
	}
	m := model.OrderOrderProduct{
		Order:        &order,
		OrderProduct: ops,
	}
	return m, nil
}

func FakerUser() (model.User, error) {
	var stID *uint
	var RolID *uint
	preRolID, err := strconv.ParseUint(faker.Number().Between(1, 5), 10, 0)
	if err != nil {
		return model.User{}, err
	}

	if preRolID == 1 {
		stID = nil
		RolID = nil
	} else if preRolID == 2 {
		stID = nil
		rid := uint(preRolID)
		RolID = &rid
	} else {
		if totalEstablishment > 0 {
			id, err := strconv.ParseUint(faker.Number().Between(1, totalEstablishment), 10, 0)
			if err != nil {
				return model.User{}, err
			}
			idU := uint(id)
			stID = &idU
			rid := uint(preRolID)
			RolID = &rid
		} else {
			stID = nil
			RolID = nil
		}
	}

	m := model.User{
		Email:           faker.Internet().FreeEmail(),
		Password:        faker.Internet().Password(10, 20),
		IsConfirmated:   true,
		RolID:           RolID,
		EstablishmentID: stID,
	}
	return m, nil
}
