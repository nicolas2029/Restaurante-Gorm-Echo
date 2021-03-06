package route

import (
	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/handler"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/middleware"
)

const (
	withoutRestrictions                = 1
	crudEstablishment                  = 2
	setRol                             = 3
	crudProduct                        = 4
	makeOrderEstablishment             = 5
	confirmPay                         = 6
	showOrdersIncompleteByStablishment = 7
	crudPayMethod                      = 8
	showInvoice                        = 9
	setEmployeeInEstablishment         = 10
	showEmployeesInEstablishment       = 11
	showAllEmployee                    = 12
	crudRol                            = 13
	crudTable                          = 14
	showOrdersIncompleteByUser         = 15
	setRolInEstablishment              = 16
)

func Product(e *echo.Echo) {
	g := e.Group("api/v1/product")
	g.GET("/:id", handler.GetProduct)
	g.GET("/", handler.GetAllProduct)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateProduct, crudProduct))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateProduct, crudProduct))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteProduct, crudProduct))
	g.POST("/img/:name", handler.SaveImage)
}

func Address(e *echo.Echo) {
	g := e.Group("api/v1/address")
	g.GET("/:id", middleware.AuthorizeIsLogin(handler.GetAddress))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllAddress, withoutRestrictions))
	g.POST("/", middleware.AuthorizeIsLogin(handler.CreateAddress))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateAddress, withoutRestrictions))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteAddress, withoutRestrictions))

}

func Permission(e *echo.Echo) {
	g := e.Group("api/v1/permission")
	g.GET("/:id", middleware.AuthorizeIsLogin(handler.GetPermission))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllPermission, crudRol))
	g.POST("/", middleware.AuthorizeWithRol(handler.CreatePermission, crudRol))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdatePermission, crudRol))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeletePermission, crudRol))

}

func Rol(e *echo.Echo) {
	g := e.Group("api/v1/rol")
	g.GET("/:id", middleware.AuthorizeIsLogin(handler.GetRol))
	g.GET("/", middleware.AuthorizeIsLogin(handler.GetAllRol))
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateRol, crudRol))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateRol, crudRol))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteRol, crudRol))

}

func User(e *echo.Echo) {
	g := e.Group("api/v1/user")
	g.GET("/:id", middleware.AuthorizeIsUser(handler.GetUser))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllUser, withoutRestrictions))
	g.PATCH("/password/", middleware.AuthorizeIsLogin(handler.UpdateUserPassword))
	g.POST("/", handler.CreateUser)
	g.POST("/login/", handler.LoginUser)
	g.GET("/login/", middleware.AuthorizeIsLogin(handler.GetMyUser))
	g.GET("/logout/", middleware.DeleteSession)
	g.GET("/validate/:code", handler.ValidateCodeConfirmation)
	g.PUT("/:id", middleware.AuthorizeIsUser(handler.UpdateUser))

	g.DELETE("/:id", middleware.AuthorizeIsUser(handler.DeleteUser))
	g.PATCH("/hire/:email", middleware.AuthorizeWithRol(handler.HireEmployeeInEstablishmentAndSetRol, setRolInEstablishment))
	g.PATCH("/fire/:email", middleware.AuthorizeWithRol(handler.FireEmployeeInEstablishmentByEmail, setRolInEstablishment))
	g.PATCH("/rol/:email", middleware.AuthorizeWithRol(handler.UpdateEmployeeInEstablishmentByEmail, setRolInEstablishment))
}

func Admin(e *echo.Echo) {
	g := e.Group("api/v1/admin")
	g.PATCH("/hire/:email", middleware.AuthorizeWithRol(handler.HireEmployeeAndSetRol, setRol))
	g.PATCH("/fire/:email", middleware.AuthorizeWithRol(handler.FireEmployeeByEmail, setRol))
	g.PATCH("/rol/:email", middleware.AuthorizeWithRol(handler.UpdateUserRol, setRol))
}

func Pay(e *echo.Echo) {
	g := e.Group("api/v1/pay")
	g.GET("/:id", handler.GetPay)
	g.GET("/", handler.GetAllPay)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreatePay, crudPayMethod))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdatePay, crudPayMethod))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeletePay, crudPayMethod))

}

func Table(e *echo.Echo) {
	g := e.Group("api/v1/table")
	g.GET("/:id", handler.GetTable)
	g.GET("/", handler.GetAllTable)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateTable, crudTable))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateTable, crudTable))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteTable, crudTable))

}

func Establishment(e *echo.Echo) {
	g := e.Group("api/v1/establishment")
	g.GET("/:id", handler.GetEstablishment)
	g.GET("/", handler.GetAllEstablishment)
	g.GET("/order/", handler.GetAllEstablishment)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateEstablishment, crudEstablishment))
	g.POST("/:id", middleware.AuthorizeWithRol(handler.CreateEstablishmentWithTables, crudEstablishment))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateEstablishment, crudEstablishment))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteEstablishment, crudEstablishment))

}

func OrderRemote(e *echo.Echo) {
	g := e.Group("api/v1/order/remote")
	g.POST("/", middleware.AuthorizeIsLogin(handler.CreateOrderRemote))
	g.PATCH("/:id", middleware.AuthorizeWithRol(handler.CompleteOrderRemote, showOrdersIncompleteByStablishment))
}

func Order(e *echo.Echo) {
	g := e.Group("api/v1/order")
	g.GET("/", middleware.AuthorizeIsLogin(handler.GetAllOrderByUser))
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateOrder, makeOrderEstablishment))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.AddProductsToOrder, makeOrderEstablishment))
	g.GET("/establishment/", middleware.AuthorizeWithRol(handler.GetAllOrdersPendingByEstablishment, showOrdersIncompleteByStablishment))
	g.GET("/establishment/all", middleware.AuthorizeWithRol(handler.GetAllOrdersByEstablishment, showInvoice))
	g.PATCH("/:id", middleware.AuthorizeWithRol(handler.CompleteOrder, makeOrderEstablishment))
	g.GET("/user/", middleware.AuthorizeWithRol(handler.GetAllOrdersPendingByUser, makeOrderEstablishment))

}

func View(e *echo.Echo) {
	route := "public/views/"
	e.Static("/", route)
}

func All(e *echo.Echo) {
	e.Use(middleware.SwitchResponse)
	Product(e)
	Address(e)
	Permission(e)
	Rol(e)
	User(e)
	OrderRemote(e)
	Pay(e)
	Table(e)
	Establishment(e)
	Order(e)
	View(e)
	Admin(e)
}
