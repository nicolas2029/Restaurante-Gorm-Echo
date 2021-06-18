package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

// getOrderProducts return a slice of OrderProduct by Order ID
func getOrderProducts(id uint) ([]*model.OrderProduct, error) {
	products := []*model.OrderProduct{}
	return products, storage.DB().Where("order_id = ?", id).Find(&products).Error
}

// GetOrder return an Order with Products by ID
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

// GetAllOrder return all orders
func GetAllOrder() ([]model.Order, error) {
	ms := make([]model.Order, 0)
	r := storage.DB().Find(&ms)
	return ms, r.Error
}

// getAllOrderOrderPorducts return all Order with products By IDs in slice of Orders
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

// GetAllOrderByUser return all orders with product by User ID
func GetAllOrderByUser(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Preload("Products", "updated = true").Find(&orders, "user_id = ? AND table_id IS NULL", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}
	return getAllOrderOrderPorducts(orders)
}

// GetAllOrdersPendingByEstablishment return all orders pending with products by Establishment ID
func GetAllOrdersPendingByEstablishment(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Preload("Products", "updated = true").Find(&orders, "establishment_id = ? AND is_done = false", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

// GetAllOrdersPendingByUser return all orders with products pending by User ID
func GetAllOrdersPendingByUser(userID uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Preload("Products", "updated = true").Find(&orders, "user_id = ? AND is_done = false AND table_id IS NOT NULL", userID).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

// GetAllOrdersByEstablishment return all orders with products by Establishment ID
func GetAllOrdersByEstablishment(id uint) ([]model.OrderOrderProduct, error) {
	orders := []model.Order{}
	err := storage.DB().Preload("Products", "updated = true").Find(&orders, "establishment_id = ?", id).Error
	if err != nil {
		return []model.OrderOrderProduct{}, err
	}

	return getAllOrderOrderPorducts(orders)
}

// CreateOrder create a new Order
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

// CompleteOrder Updates the status of an order to completed and set a paymethod
func CompleteOrder(orderID, userID, payID uint) error {
	//err := updateTableStatus()
	order := &model.Order{}
	err := storage.DB().First(order, "id = ?", orderID).Error
	if err != nil {
		return err
	}
	if order.UserID == nil || *order.UserID != userID {
		return sysError.ErrYouAreNotAutorized
	}
	if order.IsDone {
		return sysError.ErrOrderAlreadyCompleted
	}
	err = updateTableStatus(*order.TableID, true)
	if err != nil {
		return err
	}
	return storage.DB().Model(&model.Order{}).Where("id = ?", orderID).Updates(model.Order{IsDone: true, PayID: &payID}).Error
}

// CompleteOrderRemote Updates the status of an order remote to completed
func CompleteOrderRemote(orderID, establishmentID uint) error {
	order := &model.Order{}
	err := storage.DB().First(order, "id = ? AND table_id IS NULL", orderID).Error
	if err != nil {
		return err
	}
	if order.IsDone {
		return sysError.ErrOrderAlreadyCompleted
	}

	return storage.DB().Model(&model.Order{}).Where("id = ?", orderID).Update("is_done", true).Error
}

// UpdateOrder update an existing order
func UpdateOrder(m *model.Order) error {
	return storage.DB().Save(m).Error
}

// DeleteOrder use soft delete to remove an order
func DeleteOrder(id uint) error {
	r := storage.DB().Delete(&model.Order{}, id)
	return r.Error
}

// AddProductsToOrder add new products to an existing and unfilled order
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
