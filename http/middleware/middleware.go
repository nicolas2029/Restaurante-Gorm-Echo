package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/controller"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/http/sessionsCookie"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

func authorizeLogin(c echo.Context) (model.Claim, error) {
	var err error
	if v := c.Request().Header.Get("authorization"); v != "" {
		m, err := authorization.ValidateToken(v)
		if err != nil {
			return model.Claim{}, err
		}
		return m, nil
	}
	cookie := sessionsCookie.Cookie()
	sess, err := cookie.Get(c.Request(), "session")
	if err != nil {
		return model.Claim{}, err
	}
	v, f := sess.Values["token"]
	if !f {
		return model.Claim{}, sysError.ErrUserNotLogin
	}
	m, err := authorization.ValidateToken(v.(string))
	if err != nil {
		return model.Claim{}, err
	}
	c.Request().Header.Set("authorization", v.(string))
	return m, nil
}

func AuthorizeWithRol(next echo.HandlerFunc, permission uint) echo.HandlerFunc {
	return func(c echo.Context) error {
		m, err := authorizeLogin(c)
		if err != nil {
			return err
		}
		err = controller.HavePermission(m.UserID, permission)
		if err != nil {
			return err
		}
		return next(c)
	}
}

func AuthorizeIsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m, err := authorizeLogin(c)
		if err != nil {
			return err
		}
		c.Set("claim", m)
		return next(c)
	}
}

func AuthorizeIsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		m, err := authorizeLogin(c)
		if err != nil {
			return err
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			return err
		}
		if m.UserID != uint(id) {
			return sysError.ErrYouAreNotAutorized
		}
		return next(c)
	}
}

func getMapErr(err error) map[string]interface{} {
	return map[string]interface{}{
		"message_type": "error",
		"message":      fmt.Sprint(err),
	}
}

func SwitchResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		//log.Fatal("aqui:", err)
		switch err {
		case nil:
			return nil
		case sysError.ErrInvalidPassword:
			return c.JSON(http.StatusBadRequest, getMapErr(err))
		case sysError.ErrInvalidToken:
			return c.JSON(http.StatusBadRequest, getMapErr(err))
		case sysError.ErrCannotGetClaim:
			return c.JSON(http.StatusInternalServerError, getMapErr(err))
		case sysError.ErrUserNotLogin:
			return c.JSON(http.StatusMethodNotAllowed, getMapErr(sysError.ErrUserNotLogin))
		case sysError.ErrUserWhitoutRol:
			return c.JSON(http.StatusMethodNotAllowed, getMapErr(err))
		case sysError.ErrYouAreNotAutorized:
			return c.JSON(http.StatusMethodNotAllowed, getMapErr(err))
		case sysError.ErrEmptyResult:
			return c.JSON(http.StatusNoContent, nil)
		default:
			return err
		}
	}
}