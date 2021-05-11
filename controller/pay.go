package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetPay(id uint) (model.Pay, error) {
	m := model.Pay{}
	err := storage.DB().First(&m).Error
	return m, err
}

func GetAllPay() ([]model.Pay, error) {
	ms := make([]model.Pay, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

func CreatePay(m *model.Pay) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdatePay(m *model.Pay) error {
	return storage.DB().Save(m).Error
}

func DeletePay(id uint) error {
	r := storage.DB().Delete(&model.Pay{}, id)
	return r.Error
}
