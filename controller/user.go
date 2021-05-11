package controller

import (
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/storage"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
	"golang.org/x/crypto/bcrypt"
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
	if !isPasswordValid(m.Password) {
		return sysError.ErrInvalidPassword
	}
	// Implementar confirmacion de correo
	pwd, err := hashAndSalt([]byte(m.Password))
	if err != nil {
		return err
	}
	m.Password = string(pwd)
	r := storage.DB().Create(m)
	return r.Error
}

func UpdateUser(m *model.User) error {
	if !isPasswordValid(m.Password) {
		return sysError.ErrInvalidPassword
	}

	pwd, err := hashAndSalt([]byte(m.Password))
	if err != nil {
		return err
	}
	m.Password = string(pwd)
	return storage.DB().Save(m).Error
}

func DeleteUser(id uint) error {
	r := storage.DB().Delete(&model.User{}, id)
	return r.Error
}

func Login(m *model.Login) (model.User, error) {
	user := model.User{}
	if !isPasswordValid(m.Password) {
		return user, sysError.ErrInvalidPassword
	}

	err := storage.DB().First(&user,
		&model.User{
			Email: m.Email,
		}).Error
	if err != nil {
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password)); err != nil {
		return model.User{}, err
	}
	return user, nil
}
