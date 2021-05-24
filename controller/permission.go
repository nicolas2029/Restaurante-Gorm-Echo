package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetPermission(id uint) (model.Permission, error) {
	m := model.Permission{}
	err := storage.DB().First(&m, id).Error
	return m, err
}

func GetAllPermission() ([]model.Permission, error) {
	ms := make([]model.Permission, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

func CreatePermission(m *model.Permission) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdatePermission(m *model.Permission) error {
	return storage.DB().Save(m).Error
}

func DeletePermission(id uint) error {
	r := storage.DB().Delete(&model.Permission{}, id)
	return r.Error
}
