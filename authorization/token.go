package authorization

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/model"
	"github.com/nicolas2029/Restaurante-Gorm-Echo/sysError"
)

// GenerateToken .
func GenerateToken(data *model.User) (string, error) {
	claim := model.Claim{
		UserID: data.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    "Gomez & Raygoza",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, sysError.ErrInvalidToken
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, sysError.ErrCannotGetClaim
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
