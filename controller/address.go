package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

// GetAddress return an Address by ID
func GetAddress(id uint) (model.Address, error) {
	m := model.Address{}
	err := storage.DB().First(&m, id).Error
	return m, err
}

// GetAllAddress return all addresses
func GetAllAddress() ([]model.Address, error) {
	ms := make([]model.Address, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

// CreateAddress create a new address
func CreateAddress(m *model.Address) error {
	r := storage.DB().Create(m)
	return r.Error
}

// UpdateAddress update an existing address
func UpdateAddress(m *model.Address) error {
	return storage.DB().Save(m).Error
}

// DeleteAddress use soft delete to remove an address
func DeleteAddress(id uint) error {
	r := storage.DB().Delete(&model.Address{}, id)
	return r.Error
}
