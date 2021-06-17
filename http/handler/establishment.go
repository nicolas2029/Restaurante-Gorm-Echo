package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetEstablishment(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetEstablishment(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllEstablishment(c echo.Context) error {
	ms, err := controller.GetAllEstablishment()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateEstablishment(c echo.Context) error {
	var err error
	m := &model.Establishment{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreateEstablishment(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func CreateEstablishmentWithTables(c echo.Context) error {
	var err error
	m := &model.Establishment{}
	if err = c.Bind(m); err != nil {
		return err
	}
	amount, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = controller.CreateEstablishmentWithTables(m, amount)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdateEstablishment(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	st, err := controller.GetEstablishment(uint(id))
	if err != nil {
		return err
	}
	m := st.Address

	if err = c.Bind(&m); err != nil {
		return err
	}
	err = controller.UpdateAddress(&m)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func DeleteEstablishment(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteEstablishment(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
