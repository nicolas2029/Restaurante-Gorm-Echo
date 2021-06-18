package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

// GetPay return a pay by ID
func GetPay(id uint) (model.Pay, error) {
	m := model.Pay{}
	err := storage.DB().First(&m, id).Error
	return m, err
}

// GetAllPay return all payments
func GetAllPay() ([]model.Pay, error) {
	ms := make([]model.Pay, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

// CreatePay create a new pay
func CreatePay(m *model.Pay) error {
	r := storage.DB().Create(m)
	return r.Error
}

// UpdatePay update an existing pay
func UpdatePay(m *model.Pay) error {
	return storage.DB().Save(m).Error
}

// DeletePay use soft delete to remove a pay
func DeletePay(id uint) error {
	r := storage.DB().Delete(&model.Pay{}, id)
	return r.Error
}
