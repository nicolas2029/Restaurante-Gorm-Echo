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

func HavePermission(userId, permissionId uint) (uint, uint, error) {
	var eId uint
	m := model.Rol{}
	user := model.User{}
	err := storage.DB().First(&user, userId).Error
	if err != nil {
		return 0, 0, err
	}
	if user.RolID == nil {
		return 0, 0, sysError.ErrUserWhitoutRol
	}
	id := *user.RolID
	err = storage.DB().Preload("Permissions", "id = ? OR id = 1", permissionId).First(&m, id).Error

	if err != nil {
		return 0, 0, err
	}

	if len(m.Permissions) == 0 {
		return 0, 0, sysError.ErrYouAreNotAutorized
	}

	eId = 0
	if user.EstablishmentID != nil {
		eId = *user.EstablishmentID
	}
	return id, eId, nil
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

func IsTableInEstablishment(tableID, establishmentID uint) (model.Table, error) {
	m := &model.Table{}
	err := storage.DB().Where("id = ? AND establishment_id = ?", tableID, establishmentID).First(m).Error
	if err != nil {
		return model.Table{}, err
	}
	return *m, nil
}
