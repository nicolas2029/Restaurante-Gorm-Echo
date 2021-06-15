package sysError

import "errors"

var (
	ErrUserNotLogin          = errors.New("user not logged in")
	ErrUserNotFound          = errors.New("user not found")
	ErrNotFound              = errors.New("not found")
	ErrInvalidToken          = errors.New("invalid token")
	ErrCannotGetClaim        = errors.New("cannot get claim")
	ErrYouAreNotAutorized    = errors.New("you are not autorized")
	ErrUserWhitoutRol        = errors.New("user whitout role")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrEmptyResult           = errors.New("empty result")
	ErrUserNotConfirm        = errors.New("user not confirmed")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrEmailAlreadyInUsed    = errors.New("email already in use")
	ErrInvalidRole           = errors.New("invalid role")
	ErrCannotGetData         = errors.New("cannot get data")
	ErrTableNotAvailable     = errors.New("table not available")
	ErrEmptyOrder            = errors.New("empty order")
	ErrOrderAlreadyCompleted = errors.New("order already completed")
	ErrEmptyAddress          = errors.New("empty address")
)
