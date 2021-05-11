package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func GetOrder(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetOrder(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllOrder(c echo.Context) error {
	ms, err := controller.GetAllOrder()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllOrderByUser(c echo.Context) error {
	m, ok := c.Get("claim").(model.Claim)
	if !ok {
		return sysError.ErrCannotGetClaim
	}

	ms, err := controller.GetAllOrderByUser(m.UserID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllOrdersPendingByEstablishment(c echo.Context) error {
	m, ok := c.Get("claim").(model.Claim)
	if !ok {
		return sysError.ErrCannotGetClaim
	}

	u, err := controller.GetUser(m.UserID)
	if err != nil {
		return err
	}

	ms, err := controller.GetAllOrdersPendingByEstablishment(*u.EstablishmentID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllOrdersByEstablishment(c echo.Context) error {
	m, ok := c.Get("claim").(model.Claim)
	if !ok {
		return sysError.ErrCannotGetClaim
	}

	u, err := controller.GetUser(m.UserID)
	if err != nil {
		return err
	}

	ms, err := controller.GetAllOrdersByEstablishment(*u.EstablishmentID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func CreateOrder(c echo.Context) error {
	var err error
	m := &model.OrderOrderProduct{}
	if err = c.Bind(m); err != nil {
		return err
	}
	claim, ok := c.Get("claim").(model.Claim)
	if !ok {
		return sysError.ErrCannotGetClaim
	}
	m.Order.UserID = claim.UserID
	err = controller.CreateOrder(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, m)
}

func UpdateOrder(c echo.Context) error {
	var err error

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	m, err := controller.GetOrder(uint(id))
	if err != nil {
		return err
	}

	if err = c.Bind(m); err != nil {
		return err
	}

	err = controller.UpdateOrder(m.Order)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}

func DeleteOrder(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}

	err = controller.DeleteOrder(uint(id))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
