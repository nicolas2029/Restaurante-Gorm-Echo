package model

/*
import "gorm.io/gorm"

func (m *OrderProduct) BeforeCreate(tx *gorm.DB) (err error) {
	p := &Product{}
	err = tx.Select("price").Where("id = ?", m.ProductID).First(p).Error
	if err != nil {
		return err
	}
	m.Price = p.Price
	return nil
}*/
