package controller

import (
	"unicode"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// hashPassword .
func HashPassword(m *model.Login) error {
	pwd, err := hashAndSalt([]byte(m.Password))
	if err != nil {
		return err
	}
	m.Password = string(pwd)
	return nil
}

func HavePermission(userId, permissionId uint) error {
	m := model.Rol{}
	user := model.User{}
	err := storage.DB().First(&user, userId).Error
	if err != nil {
		return err
	}
	if user.RolID == nil {
		return sysError.ErrUserWhitoutRol
	}
	id := *user.RolID
	err = storage.DB().Preload("Permissions", "id = ? OR id = 1", permissionId).First(&m, id).Error

	if err != nil {
		return err
	}

	if len(m.Permissions) == 0 {
		return sysError.ErrYouAreNotAutorized
	}
	return nil
}

func isPasswordValid(pwd string) bool {
	if len(pwd) < 8 {
		return false
	}

	var (
		hasUpperCase bool
		hasSpecial   bool
		hasNumber    bool
		hasLower     bool
	)

	for _, v := range pwd {
		if hasLower && hasNumber && hasSpecial && hasUpperCase {
			return true
		}
		switch {
		case unicode.IsLower(v):
			hasLower = true
		case unicode.IsUpper(v):
			hasUpperCase = true
		case unicode.IsNumber(v):
			hasNumber = true
		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			hasSpecial = true
		}
	}

	return hasLower && hasNumber && hasSpecial && hasUpperCase
}