package model

import (
	"html/template"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Product struct {
	Model
	Name        string  `gorm:"type varchar(50); not null" json:"name"`
	Price       float64 `gorm:"float; not null" json:"price"`
	Description *string `gorm:"type varchar(100)" json:"description"`
	Img         *string `gorm:"type varchar(100)" json:"img"`
}

type Table struct {
	Model
	IsAvalaible     bool    `gorm:"type bool; not null; default true" json:"is_avalaible"`
	Orders          []Order `json:"orders"`
	EstablishmentID uint    `json:"establishment_id"`
}

type Pay struct {
	Model
	Name   string  `gorm:"type varchar(50); not null" json:"name"`
	Orders []Order `json:"-"`
}

type Address struct {
	Model
	Line1      string  `gorm:"type varchar(100); not null" json:"line1"`
	Line2      *string `gorm:"type varchar(100);" json:"line2"`
	City       string  `gorm:"type varchar(100); not null" json:"city"`
	PostalCode string  `gorm:"type varchar(15); not null" json:"postal_code"`
	State      string  `gorm:"type varchar(100); not null" json:"state"`
	Country    string  `gorm:"type varchar(100); not null" json:"country"`
}

type Establishment struct {
	Model
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
	Model
	Name        string       `gorm:"type varchar(50); not null" json:"name"`
	Permissions []Permission `gorm:"many2many:rol_permissions" json:"permissions"`
	Users       []User       `json:"-"`
}

type User struct {
	Model
	Email           string  `gorm:"type varchar(100); not null; unique" json:"email"`
	Password        string  `gorm:"type varchar(64); not null" json:"password"`
	RolID           *uint   `json:"rol_id"`
	Orders          []Order `json:"-"`
	EstablishmentID *uint   `json:"establishment_id"`
	IsConfirmated   bool    `json:"-" gorm:"not null; default false"`
}

type Order struct {
	Model
	PayID           uint      `json:"pay_id"`
	UserID          uint      `json:"user_id"`
	EstablishmentID uint      `json:"establishment_id"`
	Products        []Product `gorm:"many2many:order_products" json:"-"`
	AddressID       *uint     `json:"address_id"`
	TableID         *uint     `json:"table_id"`
}

type OrderProduct struct {
	ID uint `gorm:"primarykey" json:"id"`
	//Model
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

/*func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}*/
