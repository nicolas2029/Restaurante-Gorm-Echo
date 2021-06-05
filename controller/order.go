package controller

import (
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
		return model.OrderOrderProduct{}, err
	}
	if order.ID == 0 {
		return model.OrderOrderProduct{}, sysError.ErrEmptyResult
	}
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
		order := v
		ms[i].Order = &order
		ms[i].OrderProduct = m
	}
	return ms, nil
}

func GetAllOrderByUser(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Preload("Products", "deleted_at IS NOT NULL").Find(&orders, "user_id = ? AND table_id IS NULL", id).Error
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

func GetAllOrdersPendingByUser(userID uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Find(&orders, "user_id = ? AND is_done = false AND table_id IS NOT NULL", userID).Error
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
	if len(m.OrderProduct) < 1 {
		return sysError.ErrEmptyOrder
	}
	r := storage.DB().Create(m.Order).Error
	if r != nil {
		return r
	}
	for _, v := range m.OrderProduct {
		v.OrderID = m.Order.ID
		v.IsDone = false
	}
	return storage.DB().CreateInBatches(m.OrderProduct, len(m.OrderProduct)).Error
}

func CompleteOrder(orderID, userID uint) error {
	//err := updateTableStatus()
	order := &model.Order{}
	err := storage.DB().First(order, "id = ?", orderID).Error
	if err != nil {
		return err
	}
	if order.UserID == nil || *order.UserID != userID {
		return sysError.ErrYouAreNotAutorized
	}
	//log.Fatal(order)
	if order.IsDone {
		return sysError.ErrOrderAlreadyCompleted
	}
	err = updateTableStatus(*order.TableID, true)
	if err != nil {
		return err
	}
	return storage.DB().Model(&model.Order{}).Where("id = ?", orderID).Update("is_done", true).Error
}

func UpdateOrder(m *model.Order) error {
	return storage.DB().Save(m).Error
}

func DeleteOrder(id uint) error {
	r := storage.DB().Delete(&model.Order{}, id)
	return r.Error
}

func AddProductsToOrder(m []*model.OrderProduct, orderID, userID uint) error {
	order := &model.Order{}
	table := &model.Table{}
	err := storage.DB().Where("id = ? AND is_done = false AND user_id = ?", orderID, userID).First(order).Error
	if err != nil {
		return err
	}
	err = storage.DB().Where("id = ? AND is_avalaible = false", order.TableID).First(table).Error
	if err != nil {
		return err
	}
	for _, v := range m {
		v.OrderID = orderID
		v.IsDone = false
	}
	return storage.DB().CreateInBatches(m, len(m)).Error
}
