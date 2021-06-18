package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

// GetRol return a role by ID
func GetRol(id uint) (model.Rol, error) {
	m := model.Rol{}
	err := storage.DB().Preload("Permissions").First(&m, id).Error
	return m, err
}

// GetAllRol return all roles
func GetAllRol() ([]model.Rol, error) {
	ms := make([]model.Rol, 0)
	r := storage.DB().Find(&ms).Error
	return ms, r
}

// CreateRol create a new role
func CreateRol(m *model.Rol) error {
	r := storage.DB().Create(m)
	return r.Error
}

// UpdateRole update an existing role
func UpdateRol(m *model.Rol) error {
	return storage.DB().Save(m).Error
}

// DeleteRol use soft delete to remove a role
func DeleteRol(id uint) error {
	r := storage.DB().Delete(&model.Rol{}, id)
	return r.Error
}
