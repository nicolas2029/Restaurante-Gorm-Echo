package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/handler"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/middleware"
)

const (
	withoutRestrictions          = 1
	crudEstablishment            = 2
	setRol                       = 3
	crudProduct                  = 4
	makeOrderEstablishment       = 5
	confirmPay                   = 6
	showOrdersIncomplete         = 7
	crudPayMethod                = 8
	showInvoice                  = 9
	setEmployeeInEstablishment   = 10
	showEmployeesInEstablishment = 11
	showAllEmployee              = 12
	crudRol                      = 13
	crudTable                    = 14
)

func Product(e *echo.Echo) {
	g := e.Group("api/v1/product")
	g.GET("/:id", handler.GetProduct)
	g.GET("/", handler.GetAllProduct)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateProduct, crudProduct))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateProduct, crudProduct))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteProduct, crudProduct))

}

func ViewProduct(e *echo.Echo) {
	g := e.Group("productos")
	g.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "products.html", "lista de productos")
	})

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
	g.GET("/:id", middleware.AuthorizeWithRol(handler.GetPermission, crudRol))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllPermission, crudRol))
	g.POST("/", middleware.AuthorizeWithRol(handler.CreatePermission, crudRol))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdatePermission, crudRol))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeletePermission, crudRol))

}

func Rol(e *echo.Echo) {
	g := e.Group("api/v1/rol")
	g.GET("/:id", middleware.AuthorizeWithRol(handler.GetRol, crudRol))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllRol, crudRol))
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateRol, crudRol))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateRol, crudRol))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteRol, crudRol))

}

func User(e *echo.Echo) {
	g := e.Group("api/v1/user")
	//sessions.Cookie().Get(e)
	g.GET("/:id", middleware.AuthorizeIsUser(handler.GetUser))
	g.GET("/", middleware.AuthorizeWithRol(handler.GetAllUser, withoutRestrictions))
	g.POST("/", handler.CreateUser)
	g.POST("/login/", handler.LoginUser)
	g.GET("/login/", middleware.AuthorizeIsLogin(func(c echo.Context) error {
		return nil
	}))
	g.PUT("/:id", middleware.AuthorizeIsUser(handler.UpdateUser))
	g.DELETE("/:id", middleware.AuthorizeIsUser(handler.DeleteUser))

}

func Pay(e *echo.Echo) {
	g := e.Group("api/v1/pay")
	//sessions.Cookie().Get(e)
	g.GET("/:id", handler.GetPay)
	g.GET("/", handler.GetAllPay)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreatePay, crudPayMethod))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdatePay, crudPayMethod))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeletePay, crudPayMethod))

}

func Table(e *echo.Echo) {
	g := e.Group("api/v1/table")
	//sessions.Cookie().Get(e)
	g.GET("/:id", handler.GetTable)
	g.GET("/", handler.GetAllTable)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateTable, crudTable))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateTable, crudTable))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteTable, crudTable))

}

func Establishment(e *echo.Echo) {
	g := e.Group("api/v1/establishment")
	//sessions.Cookie().Get(e)
	g.GET("/:id", handler.GetEstablishment)
	g.GET("/", handler.GetAllEstablishment)
	g.GET("/order/", handler.GetAllEstablishment)
	g.POST("/", middleware.AuthorizeWithRol(handler.CreateEstablishment, crudEstablishment))
	g.PUT("/:id", middleware.AuthorizeWithRol(handler.UpdateEstablishment, crudEstablishment))
	g.DELETE("/:id", middleware.AuthorizeWithRol(handler.DeleteEstablishment, crudEstablishment))

}

func OrderRemote(e *echo.Echo) {
	g := e.Group("api/v1/order/remote")
	g.GET("/:id", handler.GetOrder)
	g.GET("/", handler.GetAllOrder)
	g.POST("/", middleware.AuthorizeIsLogin(handler.CreateOrder))
	//g.PUT("/:id", middleware.AuthorizeIsLogin(handler.UpdateOrder))
	//g.DELETE("/:id", middleware.AuthorizeIsLogin(handler.DeleteOrder))

}

func Order(e *echo.Echo) {
	g := e.Group("api/v1/order")
	//g.GET("/:id", handler.GetOrder)
	g.GET("/", middleware.AuthorizeIsLogin(handler.GetAllOrderByUser))
	g.POST("/", middleware.AuthorizeIsLogin(handler.CreateOrder))
	g.GET("/establishment/", middleware.AuthorizeWithRol(handler.GetAllOrdersPendingByEstablishment, showOrdersIncomplete))
	g.GET("/establishment/all", middleware.AuthorizeWithRol(handler.GetAllOrdersByEstablishment, showInvoice))
	//g.PUT("/:id", middleware.AuthorizeIsLogin(handler.UpdateOrder))
	//g.DELETE("/:id", middleware.AuthorizeIsLogin(handler.DeleteOrder))
}

func View(e *echo.Echo) {
	route := "../public/views/"
	e.Static("/", route)
	//e.Renderer = handler.NewRender()
	e.Static("/table/", route+"table/")
	/*e.GET("/table/", func(c echo.Context) error {
		template.ParseFiles(route + "table/")
		return c.Render(http.StatusOK, route+"table/", "ok")
	})*/
	//e.Static("/", route+"index.html")

}

func All(e *echo.Echo) {
	e.Use(middleware.SwitchResponse)
	Product(e)
	ViewProduct(e)
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
}
