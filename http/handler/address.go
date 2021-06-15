package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func isEmptyAddress(m *model.Address) bool {
	return len(m.City) == 0 || len(m.Line1) == 0 || len(m.Country) == 0 || len(m.PostalCode) == 0 || len(m.State) == 0
}

func GetAddress(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetAddress(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllAddress(c echo.Context) error {
	ms, err := controller.GetAllAddress()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateAddress(c echo.Context) error {
	var err error
	m := &model.Address{}
	if err = c.Bind(m); err != nil {
		return err
	}
	if isEmptyAddress(m) {
		return sysError.ErrEmptyAddress
	}
	err = controller.CreateAddress(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m.ID)
}

func UpdateAddress(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetAddress(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(&m); err != nil {
		return err
	}

	err = controller.UpdateAddress(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeleteAddress(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteAddress(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
