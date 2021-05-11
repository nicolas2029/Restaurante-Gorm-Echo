package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetTable(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetTable(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllTable(c echo.Context) error {
	ms, err := controller.GetAllTable()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateTable(c echo.Context) error {
	var err error
	m := &model.Table{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreateTable(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdateTable(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetTable(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(&m); err != nil {
		return err
	}

	err = controller.UpdateTable(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeleteTable(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteTable(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
