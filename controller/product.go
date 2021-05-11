package controller

import (
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
	r := storage.DB().Find(&ps)
	return ps, r.Error
}

func CreateProduct(m *model.Product) error {
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateProduct(m *model.Product) error {
	return storage.DB().Save(m).Error
}

func DeleteProduct(id uint) error {
	r := storage.DB().Delete(&model.Product{}, id)
	return r.Error
}
