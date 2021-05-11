package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetAddress(id uint) (model.Address, error) {
	m := model.Address{}
	err := storage.DB().First(&m, id).Error
	return m, err
}

func GetAllAddress() ([]model.Address, error) {
	ms := make([]model.Address, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

func CreateAddress(m *model.Address) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateAddress(m *model.Address) error {
	return storage.DB().Save(m).Error
}

func DeleteAddress(id uint) error {
	r := storage.DB().Delete(&model.Address{}, id)
	return r.Error
}
