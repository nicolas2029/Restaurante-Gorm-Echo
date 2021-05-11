package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetPay(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetPay(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllPay(c echo.Context) error {
	ms, err := controller.GetAllPay()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreatePay(c echo.Context) error {
	var err error
	m := &model.Pay{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreatePay(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdatePay(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetPay(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(&m); err != nil {
		return err
	}

	err = controller.UpdatePay(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeletePay(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeletePay(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
