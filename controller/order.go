package controller

import (
	"log"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func getOrderProducts(id uint) ([]*model.OrderProduct, error) {
	products := []*model.OrderProduct{}
	return products, storage.DB().Where("order_id = ?", id).Find(&products).Error
}

func GetOrder(id uint) (model.OrderOrderProduct, error) {
	order := &model.Order{}
	err := storage.DB().Find(order, id).Error
	if err != nil {
		log.Fatal("aqui")
		return model.OrderOrderProduct{}, err
	}
	if order.ID == 0 {
		return model.OrderOrderProduct{}, sysError.ErrEmptyResult
	}
	//log.Fatal("aqui 2")
	products, err := getOrderProducts(order.ID)
	if err != nil {
		return model.OrderOrderProduct{}, err
	}
	ms := model.OrderOrderProduct{
		Order:        order,
		OrderProduct: products,
	}

	return ms, nil
}

func GetAllOrder() ([]model.Order, error) {
	ms := make([]model.Order, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

func getAllOrderOrderPorducts(orders []model.Order) ([]model.OrderOrderProduct, error) {
	if len(orders) == 0 {
		return []model.OrderOrderProduct{}, sysError.ErrEmptyResult
	}

	ms := make([]model.OrderOrderProduct, len(orders))
	for i, v := range orders {
		m, err := getOrderProducts(v.ID)
		if err != nil {
			return []model.OrderOrderProduct{}, err
		}
		ms[i].Order = &v
		ms[i].OrderProduct = m
	}

	return ms, nil
}

func GetAllOrderByUser(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Find(&orders, "user_id = ?", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

func GetAllOrdersPendingByEstablishment(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Find(&orders, "establishment_id = ? AND is_done = false", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

func GetAllOrdersByEstablishment(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Find(&orders, "establishment_id = ?", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

func CreateOrder(m *model.OrderOrderProduct) error {
	r := storage.DB().Create(m.Order).Error
	if r != nil {
		return r
	}
	for _, v := range m.OrderProduct {
		v.OrderID = m.Order.ID
	}
	return storage.DB().CreateInBatches(m.OrderProduct, len(m.OrderProduct)).Error
}

func UpdateOrder(m *model.Order) error {
	return storage.DB().Save(m).Error
}

func DeleteOrder(id uint) error {
	r := storage.DB().Delete(&model.Order{}, id)
	return r.Error
}
