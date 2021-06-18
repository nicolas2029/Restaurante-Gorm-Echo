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

// GetUser return an user by ID
func GetUser(id uint) (model.User, error) {
	p := model.User{}
	err := storage.DB().First(&p, id).Error
	return p, err
}

// UpdateUserRolByEmail update a user's role by email
func UpdateUserRolByEmail(email string, rolId uint) error {
	m := model.Rol{}
	m.ID = rolId
	err := storage.DB().First(&m).Error
	if err != nil {
		return err
	}
	res := storage.DB().Model(&model.User{}).Where("email = ? AND rol > 0", email).Update("rol_id = ?", rolId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// UpdateUserRolInEstablishmentByEmail update the role of a user and assign an establishment by email
func UpdateUserRolInEstablishmentByEmail(email string, rolId, establishmentId uint) error {
	m := model.Rol{}
	m.ID = rolId
	err := storage.DB().First(&m).Error
	if err != nil {
		return err
	}
	res := storage.DB().Model(&model.User{}).Where("email = ? AND establishment_id = ? AND rol > 0", email, establishmentId).Update("rol_id = ?", rolId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// HireEmployeeAndSetRol hire an employee and set role by Email
func HireEmployeeAndSetRol(email string, rolId uint) error {
	mRol := model.Rol{}
	mRol.ID = rolId
	err := storage.DB().First(&mRol).Error
	if err != nil {
		return err
	}
	res := storage.DB().Model(&model.User{}).Where("email = ? AND rol_id IS NULL", email).Updates(map[string]interface{}{
		"rol_id": rolId,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// HireEmployeeInEstablishmentAndSetRol hire an employee, set role and assign an establishment by Email
func HireEmployeeInEstablishmentAndSetRol(email string, rolId, establishmentId uint) error {
	m := model.Establishment{}
	m.ID = establishmentId
	err := storage.DB().First(&m).Error
	if err != nil {
		return err
	}

	mRol := model.Rol{}
	mRol.ID = rolId
	err = storage.DB().First(&mRol).Error
	if err != nil {
		return err
	}
	res := storage.DB().Model(&model.User{}).Where("email = ? AND rol_id IS NULL", email).Updates(map[string]interface{}{
		"rol_id":           rolId,
		"establishment_id": establishmentId,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// FireEmployeeByEmail update user's role to null and update user's establishment to null by Email
func FireEmployeeByEmail(email string, rolId uint) error {
	res := storage.DB().Model(&model.User{}).Where("email = ? AND rol_id > 0 AND rol_id > ?", email, rolId).Updates(map[string]interface{}{
		"rol_id":           nil,
		"establishment_id": nil,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// FireEmployeeByEmail update user's role to null and update user's establishment to null by Email and Establishment
func FireEmployeeInEstablishmentByEmail(email string, rolId, establishmentId uint) error {
	res := storage.DB().Model(&model.User{}).Where("email = ? AND establishment_id = ? AND rol_id > 0 AND rol_id > ?", email, establishmentId, rolId).Updates(
		map[string]interface{}{
			"rol_id":           nil,
			"establishment_id": nil,
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return sysError.ErrUserNotFound
	}
	return nil
}

// GetAllUser return all users
func GetAllUser() ([]model.User, error) {
	ps := make([]model.User, 0)
	r := storage.DB().Find(&ps)
	return ps, r.Error
}

// CreateUser create a new user, encrypt the password and send a confirmation code to the email
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

// UpdateUser Update an existing user
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

// ValidateUser receives a token, validates it and updates the user status to confirmed
func ValidateUser(token string) error {
	claim, err := authorization.ValidateCodeVerification(token)
	if err != nil {
		return err
	}
	return storage.DB().Model(&model.User{}).Where("email = ?", claim.Email).Update("is_confirmated", true).Error
}

// DeleteUser use soft delete to remove an user
func DeleteUser(id uint) error {
	r := storage.DB().Delete(&model.User{}, id)
	return r.Error
}

// Login Receive the username and password of a user, confirm that the credentials are correct and return a user
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

// isEmailValid return true if the email is valid, else return false
func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// isEmailAndPasswordValid return an error if the password or emails are invalid
func isEmailAndPasswordValid(email, password string) error {
	if !isEmailValid(email) {
		return sysError.ErrInvalidEmail
	}
	if !isPasswordValid(password) {
		return sysError.ErrInvalidPassword
	}
	return nil
}

// UpdateUserPassword update a user's password by ID
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

// UpdateUserEmailAndPassowrd Update a user's email and password
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
