package controller

import (
	"time"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
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
	err := storage.DB().Model(&model.Product{}).Where("id = ?", m.ID).Update("updated", true).Error
	if err != nil {
		return err
	}
	m.ID = 0
	m.CreatedAt = time.Time{}
	m.UpdatedAt = time.Time{}
	m.Updated = false
	return CreateProduct(m)
}

func DeleteProduct(id uint) error {
	return storage.DB().Model(&model.Product{}).Where("id = ?", id).Update("updated", true).Error
}
