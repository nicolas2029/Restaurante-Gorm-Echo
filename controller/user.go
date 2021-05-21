package controller

import (
	"net/mail"

	"github.com/nicolas2029/Restaurante-Gorm-Echo/authorization"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUser(id uint) (model.User, error) {
	p := model.User{}
	err := storage.DB().First(&p, id).Error
	return p, err
}

// GetUsers return all products
func GetAllUser() ([]model.User, error) {
	ps := make([]model.User, 0)
	r := storage.DB().Find(&ps)
	return ps, r.Error
}

func CreateUser(m *model.User) error {
	var err error
	if err = isEmailAndPasswordValid(m.Email, m.Password); err != nil {
		return err
	}
	token, err := authorization.GenerateCodeVerification(m.Email)
	if err != nil {
		return err
	}
	u := &model.User{}
	result := storage.DB().Where("email = ?", m.Email).First(u)
	if result.RowsAffected != 0 {
		if !u.IsConfirmated {
			err = sendCodeConfirmationToEmail(m.Email, token)
			if err != nil {
				return err
			}
		}
		return sysError.ErrEmailAlreadyInUsed
	}
	err = result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	pwd, err := hashAndSalt([]byte(m.Password))
	if err != nil {
		return err
	}
	err = sendCodeConfirmationToEmail(m.Email, token)
	if err != nil {
		return err
	}
	m.Password = string(pwd)
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateUser(m *model.User) error {
	var err error
	if err = isEmailAndPasswordValid(m.Email, m.Password); err != nil {
		return err
	}

	pwd, err := hashAndSalt([]byte(m.Password))
	if err != nil {
		return err
	}
	m.Password = string(pwd)
	return storage.DB().Save(m).Error
}

func ValidateUser(token string) error {
	claim, err := authorization.ValidateCodeVerification(token)
	if err != nil {
		return err
	}
	return storage.DB().Model(&model.User{}).Where("email = ?", claim.Email).Update("is_confirmated", true).Error
}

func DeleteUser(id uint) error {
	r := storage.DB().Delete(&model.User{}, id)
	return r.Error
}

func Login(m *model.Login) (model.User, error) {
	user := model.User{}
	var err error
	if err = isEmailAndPasswordValid(m.Email, m.Password); err != nil {
		return user, err
	}

	err = storage.DB().First(&user,
		&model.User{
			Email: m.Email,
		}).Error
	if err != nil {
		return model.User{}, err
	}
	if !user.IsConfirmated {
		return model.User{}, sysError.ErrUserNotConfirm
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password)); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isEmailAndPasswordValid(email, password string) error {
	if !isEmailValid(email) {
		return sysError.ErrInvalidEmail
	}
	if !isPasswordValid(password) {
		return sysError.ErrInvalidPassword
	}
	return nil
}

func UpdateUserPassword(id uint, password string) error {
	if !isPasswordValid(password) {
		return sysError.ErrInvalidPassword
	}
	m := &model.User{}
	m.ID = id

	pwd, err := hashAndSalt([]byte(password))
	if err != nil {
		return err
	}

	return storage.DB().Model(m).Updates(model.User{
		Password: string(pwd),
	}).Error
}

// FALTA ACTUALIZAR EL METODO DE CONFIRMACION
func UpdateUserEmailAndPassword(id uint, email, password string) error {
	err := isEmailAndPasswordValid(email, password)
	if err != nil {
		return err
	}
	m := &model.User{}
	m.ID = id
	token, err := authorization.GenerateToken(m)
	if err != nil {
		return err
	}
	err = storage.DB().First(&model.User{Email: email}).Error
	if err != gorm.ErrRecordNotFound {
		if err == nil {
			return sysError.ErrEmailAlreadyInUsed
		}
		return err
	}

	pwd, err := hashAndSalt([]byte(password))
	if err != nil {
		return err
	}
	err = sendCodeConfirmationToEmail(m.Email, token)
	if err != nil {
		return err
	}
	return storage.DB().Model(m).Updates(model.User{
		Email:         email,
		Password:      string(pwd),
		IsConfirmated: false,
	}).Error
}
