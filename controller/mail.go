package controller

import "github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"

// sendCodeConfirmationToEmail send a confirmation code to an email
func sendCodeConfirmationToEmail(add, token string) error {
	return authorization.ConfirmEmail(add, token)
}
