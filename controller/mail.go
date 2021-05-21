package controller

import "github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"

func sendCodeConfirmationToEmail(add, token string) error {
	return authorization.ConfirmEmail(add, token)
}
