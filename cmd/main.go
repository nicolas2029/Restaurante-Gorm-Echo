package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/route"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/seeders"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
)

func seedAll() error {
	var establishments, addresses, products, users, orders int
	fmt.Println("Enter the number of establishments: ")
	_, err := fmt.Scan(&establishments)
	if err != nil {
		return err
	}
	fmt.Println("Enter the number of addresses: ")
	_, err = fmt.Scan(&addresses)
	if err != nil {
		return err
	}
	fmt.Println("Enter the number of products: ")
	_, err = fmt.Scan(&products)
	if err != nil {
		return err
	}
	fmt.Println("Enter the number of users: ")
	_, err = fmt.Scan(&users)
	if err != nil {
		return err
	}
	fmt.Println("Enter the number of orders: ")
	_, err = fmt.Scan(&orders)
	if err != nil {
		return err
	}

	return seeders.SeederAll(establishments, addresses, products, users, orders)
}

func newUser() error {
	var rolID uint
	user := model.User{}
	rolID = 1
	fmt.Println("Please enter your email: ")
	fmt.Scanln(&user.Email)
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&user.Password)
	user.RolID = &rolID
	return controller.CreateUser(&user)
}

func newPermissions() error {
	permissions := make([]*model.Permission, 15)
	names := []string{"Sin restricciones", "CRUD estaclimientos", "Asignar roles de menor jerarquia", "CRUD productos",
		"Tomar orden en un establecimiento", "Confirmar Pago", "Ver todas las ordenes sin completar del local", "CRUD metodos de pago",
		"Ver facturas", "Asignar empleado a un establecimiento", "Mostrar todos los empleados del establecimiento", "Mostrar todos los empleados",
		"CRUD Roles", "CRUD Tables", "Mostrar ordenes pendientes del empleado", "Asignar rol de menor jerarquia de empleado dentro de un establecimiento"}
	for i := 0; i < 15; i++ {
		permissions[i] = &model.Permission{Name: names[i]}
	}
	return storage.DB().CreateInBatches(&permissions, 15).Error
}

func newRoles() error {
	roles := make([]*model.Rol, 5)
	ids := [][]uint{{1}, {2, 3, 4, 8, 11, 12, 13}, {9, 10, 11, 16}, {5, 6, 15}, {7}}
	names := []string{"Dueño", "Admin", "Gerente", "Mesero", "Cocina"}
	for i, val := range names {
		roles[i] = &model.Rol{Name: val}
		for _, id := range ids[i] {
			p := model.Permission{}
			p.ID = id
			roles[i].Permissions = append(roles[i].Permissions, p)
		}
	}
	return storage.DB().CreateInBatches(&roles, 5).Error
}

func newDB() error {
	err := storage.DB().First(&model.Permission{}).Error
	if err != nil {
		err = newPermissions()
		if err != nil {
			return err
		}
	}

	err = storage.DB().First(&model.Rol{}).Error
	if err != nil {
		err = newRoles()
		if err != nil {
			return err
		}
	}

	err = storage.DB().First(&model.User{}, "is_confirmated = true and rol_id = 1").Error
	if err != nil {
		err = newUser()
		if err != nil {
			return err
		}
		isSeeder := ""
		for isSeeder != "1" && isSeeder != "0" {
			fmt.Println("Do you want to use the seeders(1.-Yes / 2.- No)?: ")
			fmt.Scan(&isSeeder)
			if isSeeder == "1" {
				seedAll()
			}
		}

	}
	return nil
}

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	storage.New("certificates/db.json")
	err = sessionsCookie.NewCookieStore("certificates/cookieKey")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para cookies: %v", err)
	}

	err = authorization.LoadMail("certificates/email.json")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados para email: %v", err)
	}

	err = storage.DB().AutoMigrate(
		&model.Product{},
		&model.Address{},
		&model.Permission{},
		&model.Establishment{},
		&model.Rol{},
		&model.User{},
		&model.Table{},
		&model.Pay{},
		&model.Order{},
		&model.OrderProduct{},
	)

	if err != nil {
		log.Fatalf("no se pudo realizar las migraciones: %v", err)
	}

	err = newDB()
	if err != nil {
		log.Fatalf("no se pudo poblar la base de datos: %v", err)
	}
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessionsCookie.Cookie()))
	route.All(e)
	err = e.Start(":80")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
