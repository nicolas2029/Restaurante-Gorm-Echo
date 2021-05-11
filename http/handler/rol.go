package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetRol(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetRol(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllRol(c echo.Context) error {
	ms, err := controller.GetAllRol()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateRol(c echo.Context) error {
	var err error
	m := &model.Rol{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreateRol(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdateRol(c echo.Context) error {
	var err error

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetRol(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.UpdateRol(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeleteRol(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteRol(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
