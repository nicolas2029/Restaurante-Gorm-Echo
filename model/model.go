package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type varchar(50); not null"`
	Price       float64 `gorm:"float; not null"`
	Description *string `gorm:"type varchar(100)"`
}

type Table struct {
	gorm.Model
	IsAvalaible     bool `gorm:"type bool; not null; default true"`
	Orders          []Order
	EstablishmentID uint
}

type Pay struct {
	gorm.Model
	Name   string `gorm:"type varchar(50); not null"`
	Orders []Order
}

type Address struct {
	gorm.Model
	Line1      string  `gorm:"type varchar(100); not null"`
	Line2      *string `gorm:"type varchar(100);"`
	City       string  `gorm:"type varchar(100); not null"`
	PostalCode string  `gorm:"type varchar(15); not null"`
	State      string  `gorm:"type varchar(100); not null"`
	Country    string  `gorm:"type varchar(100); not null"`
}

type Establishment struct {
	gorm.Model
	AddressID uint
	Address   Address
	Tables    []Table
	Users     []User
}

type Permission struct {
	gorm.Model
	Name string `gorm:"type varchar(50); not null"`
}

type Rol struct {
	gorm.Model
	Permissions []Permission `gorm:"many2many:rol_permissions"`
	Users       []User
}

type User struct {
	gorm.Model
	Email           string `gorm:"type varchar(100); not null; unique"`
	Password        string `gorm:"type varchar(256); not null"`
	RolID           *uint
	Orders          []*Order `gorm:"many2many:order_users"`
	EstablishmentID *uint
}

type Order struct {
	gorm.Model
	PayID     uint
	Addresses []Address `gorm:"many2many:order_addresses"`
	Products  []Product `gorm:"many2many:order_products"`
	Users     []*User   `gorm:"many2many:order_users"`
	TableID   *uint
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Amount    uint `gorm:"type uint"`
}
