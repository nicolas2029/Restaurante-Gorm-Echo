package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func GetRol(id uint) (model.Rol, error) {
	m := model.Rol{}
	err := storage.DB().First(&m).Error
	return m, err
}

func GetAllRol() ([]model.Rol, error) {
	ms := make([]model.Rol, 0)
	r := storage.DB().Find(&ms).Error
	return ms, r
}

func CreateRol(m *model.Rol) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateRol(m *model.Rol) error {
	return storage.DB().Save(m).Error
}

func DeleteRol(id uint) error {
	r := storage.DB().Delete(&model.Rol{}, id)
	return r.Error
}
