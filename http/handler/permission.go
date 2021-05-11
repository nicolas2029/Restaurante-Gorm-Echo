package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetPermission(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetPermission(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllPermission(c echo.Context) error {
	ms, err := controller.GetAllPermission()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreatePermission(c echo.Context) error {
	var err error
	m := &model.Permission{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreatePermission(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdatePermission(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetPermission(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(&m); err != nil {
		return err
	}

	err = controller.UpdatePermission(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeletePermission(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeletePermission(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
