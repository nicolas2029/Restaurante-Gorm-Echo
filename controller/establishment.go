package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

// GetEstablishment return an Establishment by ID
func GetEstablishment(id uint) (model.Establishment, error) {
	m := model.Establishment{}
	err := storage.DB().Preload("Address").Preload("Tables").First(&m, id).Error
	return m, err
}

// GetAllEstablishment return all Establishments
func GetAllEstablishment() ([]model.Establishment, error) {
	ms := make([]model.Establishment, 0)
	r := storage.DB().Preload("Address").Find(&ms)
	return ms, r.Error
}

// CreateEstablishment create a new Establishment
func CreateEstablishment(m *model.Establishment) error {
	r := storage.DB().Create(m)
	return r.Error
}

// CreateEstablishmentWithTables create a new Establishment with tables
func CreateEstablishmentWithTables(m *model.Establishment, amount int) error {
	r := storage.DB().Create(m)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected == 0 {
		return sysError.ErrEmptyResult
	}
	if amount <= 0 {
		return nil
	}

	if amount == 1 {
		return storage.DB().Create(model.Table{EstablishmentID: m.ID}).Error
	}
	ms := make([]model.Table, amount)
	t := model.Table{}
	t.EstablishmentID = m.ID
	t.IsAvalaible = true
	for i := 0; i < amount; i++ {
		ms[i] = t
	}
	return storage.DB().CreateInBatches(ms, amount).Error
}

// UpdateEstablishment update an existing Establishment
func UpdateEstablishment(m *model.Establishment) error {
	return storage.DB().Save(m).Error
}

// DeleteEstablishment use soft delete to remove an Establishment
func DeleteEstablishment(id uint) error {
	r := storage.DB().Delete(&model.Establishment{}, id)
	return r.Error
}
