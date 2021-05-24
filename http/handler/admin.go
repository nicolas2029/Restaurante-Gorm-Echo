package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

// HireEmployeeAndSetRol hire an employee and assign him to an establishment
func HireEmployeeAndSetRol(c echo.Context) error {
	var err error
	m := model.User{}
	userRolId := c.Get("rolId").(uint)

	if err = c.Bind(&m); err != nil {
		return err
	}
	if m.RolID == nil {
		return sysError.ErrInvalidRole
	}
	if userRolId >= *m.RolID {
		return sysError.ErrYouAreNotAutorized
	}
	email := c.Param("email")
	if m.EstablishmentID == nil {
		return controller.HireEmployeeAndSetRol(email, *m.RolID) //This can be upgrade with other hantler and new permission SetRoleWithoutEstablishment
	}
	return controller.HireEmployeeInEstablishmentAndSetRol(email, *m.RolID, *m.EstablishmentID)
}

func FireEmployeeByEmail(c echo.Context) error {
	var err error
	m := model.User{}
	userRolId := c.Get("rolId").(uint)

	if err = c.Bind(&m); err != nil {
		return err
	}
	email := c.Param("email")
	return controller.FireEmployeeByEmail(email, userRolId)
}

func UpdateUserRol(c echo.Context) error {
	var err error
	m := model.User{}
	userRolId := c.Get("rolId").(uint)

	if err = c.Bind(&m); err != nil {
		return err
	}
	if m.RolID == nil {
		return sysError.ErrInvalidRole
	}
	if userRolId >= *m.RolID {
		return sysError.ErrYouAreNotAutorized
	}

	email := c.Param("email")
	return controller.UpdateUserRolByEmail(email, *m.RolID)
}
