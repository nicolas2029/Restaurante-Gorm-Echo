package model

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type varchar(50); not null" json:"name"`
	Price       float64 `gorm:"float; not null" json:"price"`
	Description *string `gorm:"type varchar(100)" json:"description"`
}

type Table struct {
	gorm.Model
	IsAvalaible     bool    `gorm:"type bool; not null; default true" json:"is_avalaible"`
	Orders          []Order `json:"orders"`
	EstablishmentID uint    `json:"establishment_id"`
}

type Pay struct {
	gorm.Model
	Name   string  `gorm:"type varchar(50); not null"`
	Orders []Order `json:"orders"`
}

type Address struct {
	gorm.Model
	Line1      string  `gorm:"type varchar(100); not null" json:"line1"`
	Line2      *string `gorm:"type varchar(100);" json:"line2"`
	City       string  `gorm:"type varchar(100); not null" json:"city"`
	PostalCode string  `gorm:"type varchar(15); not null" json:"postal_code"`
	State      string  `gorm:"type varchar(100); not null" json:"state"`
	Country    string  `gorm:"type varchar(100); not null" json:"country"`
}

type Establishment struct {
	gorm.Model
	AddressID uint    `json:"address_id"`
	Address   Address `json:"address"`
	Tables    []Table `json:"-"`
	Users     []User  `json:"-"`
	Orders    []Order `json:"-"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"type varchar(50); not null" json:"name"`
}

type Rol struct {
	gorm.Model
	Name        string       `gorm:"type varchar(50); not null" json:"name"`
	Permissions []Permission `gorm:"many2many:rol_permissions" json:"-"`
	Users       []User       `json:"-"`
}

type User struct {
	gorm.Model
	Email           string  `gorm:"type varchar(100); not null; unique" json:"email"`
	Password        string  `gorm:"type varchar(64); not null" json:"password"`
	RolID           *uint   `json:"rol_id"`
	Orders          []Order `json:"-"`
	EstablishmentID *uint   `json:"establishment_id"`
}

type Order struct {
	gorm.Model
	PayID           uint      `json:"pay_id"`
	UserID          uint      `json:"user_id"`
	EstablishmentID uint      `json:"establishment_id"`
	Products        []Product `gorm:"many2many:order_products" json:"-"`
	AddressID       *uint     `json:"address_id"`
	TableID         *uint     `json:"table_id"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   uint `gorm:"primaryKey" json:"order_id"`
	ProductID uint `gorm:"primaryKey" json:"product_id"`
	Amount    uint `gorm:"type uint" json:"amount"`
	IsDone    bool `gorm:"type bool" json:"is_done"`
}

type Template struct {
	Templates *template.Template
}

type OrderOrderProduct struct {
	Order        *Order          `json:"order"`
	OrderProduct []*OrderProduct `json:"order_products"`
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
