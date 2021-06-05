package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func GetEstablishment(id uint) (model.Establishment, error) {
	m := model.Establishment{}
	err := storage.DB().Preload("Address").Preload("Tables").First(&m, id).Error
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

func UpdateEstablishment(m *model.Establishment) error {
	return storage.DB().Save(m).Error
}

func DeleteEstablishment(id uint) error {
	r := storage.DB().Delete(&model.Establishment{}, id)
	return r.Error
}
