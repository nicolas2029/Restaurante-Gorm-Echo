package controller

import (
	"time"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func GetProduct(id uint) (model.Product, error) {
	p := model.Product{}
	err := storage.DB().First(&p, id).Error
	return p, err
}

// GetProducts return all products
func GetAllProduct() ([]model.Product, error) {
	ps := make([]model.Product, 0)
	r := storage.DB().Find(&ps, "updated = false")
	return ps, r.Error
}

func CreateProduct(m *model.Product) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateProduct(m *model.Product) error {
	r := storage.DB().Model(&model.Product{}).Where("id = ? and updated = false", m.ID).Update("updated", true)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected == 0 {
		return sysError.ErrProductAlreadyUpdated
	}
	m.ID = 0
	m.CreatedAt = time.Time{}
	m.UpdatedAt = time.Time{}
	m.Updated = false
	return CreateProduct(m)
}

func DeleteProduct(id uint) error {
	r := storage.DB().Model(&model.Product{}).Where("id = ? and updated = false", id).Update("updated", true)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected == 0 {
		return sysError.ErrProductAlreadyUpdated
	}
	return nil
}
