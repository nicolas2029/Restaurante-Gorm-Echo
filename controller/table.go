package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

// GetTable return a table by ID
func GetTable(id uint) (model.Table, error) {
	m := model.Table{}
	err := storage.DB().First(&m, id).Error
	return m, err
}

// GetAllTable return all tables
func GetAllTable() ([]model.Table, error) {
	ms := make([]model.Table, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

// CreateTable create a new table
func CreateTable(m *model.Table) error {
	r := storage.DB().Create(m)
	return r.Error
}

// UpdateTable update an existing table
func UpdateTable(m *model.Table) error {
	return storage.DB().Save(m).Error
}

// DeleteTable use soft delete to remove a table
func DeleteTable(id uint) error {
	r := storage.DB().Delete(&model.Table{}, id)
	return r.Error
}

// updateTableStatus update the status of a table
func updateTableStatus(id uint, status bool) error {
	return storage.DB().Model(&model.Table{}).Where("id = ? AND is_avalaible = ?", id, !status).Update("is_avalaible", status).Error
}
