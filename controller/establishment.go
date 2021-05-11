package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetEstablishment(id uint) (model.Establishment, error) {
	m := model.Establishment{}
	err := storage.DB().Preload("Address").First(&m, id).Error
	return m, err
}

func GetAllEstablishment() ([]model.Establishment, error) {
	ms := make([]model.Establishment, 0)
	r := storage.DB().Preload("Address").Find(&ms)
	return ms, r.Error
}

func CreateEstablishment(m *model.Establishment) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateEstablishment(m *model.Establishment) error {
	return storage.DB().Save(m).Error
}

func DeleteEstablishment(id uint) error {
	r := storage.DB().Delete(&model.Establishment{}, id)
	return r.Error
}
