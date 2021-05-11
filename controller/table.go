package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetTable(id uint) (model.Table, error) {
	m := model.Table{}
	err := storage.DB().First(&m).Error
	return m, err
}

func GetAllTable() ([]model.Table, error) {
	ms := make([]model.Table, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

func CreateTable(m *model.Table) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateTable(m *model.Table) error {
	return storage.DB().Save(m).Error
}

func DeleteTable(id uint) error {
	r := storage.DB().Delete(&model.Table{}, id)
	return r.Error
}
