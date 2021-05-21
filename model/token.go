package model

import "github.com/dgrijalva/jwt-go"

// Login .
type Login struct {
	Email    string `gorm:"type varchar(100); not null" json:"email"`
	Password string `gorm:"type varchar(64); not null" json:"password"`
}

// Claim .
type Claim struct {
	UserID uint `json:"uid"`
	jwt.StandardClaims
}

// GeneralResponse .
type GeneralResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

// LoginResponse .
type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type CodeVerification struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
