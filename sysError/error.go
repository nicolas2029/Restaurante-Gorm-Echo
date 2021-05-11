package sysError

import "errors"

var (
	ErrUserNotLogin       = errors.New("user not logged in")
	ErrUserNotFound       = errors.New("user not found")
	ErrNotFound           = errors.New("not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrCannotGetClaim     = errors.New("cannot get claim")
	ErrYouAreNotAutorized = errors.New("you are not autorized")
	ErrUserWhitoutRol     = errors.New("user whitout rol")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrEmptyResult        = errors.New("empty result")
)
