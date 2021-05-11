package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
)

func GetUser(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetUser(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllUser(c echo.Context) error {
	ms, err := controller.GetAllUser()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateUser(c echo.Context) error {
	var err error
	m := &model.User{}
	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.CreateUser(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdateUser(c echo.Context) error {
	var err error

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetUser(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(&m); err != nil {
		return err
	}

	err = controller.UpdateUser(&m)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeleteUser(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteUser(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func LoginUser(c echo.Context) error {
	var err error
	m := &model.Login{}
	if err = c.Bind(m); err != nil {
		return err
	}

	user, err := controller.Login(m)
	if err != nil {
		return err
	}

	token, err := authorization.GenerateToken(&user)
	if err != nil {
		return err
	}

	err = createCookie(c.Request(), &c.Response().Writer, token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
