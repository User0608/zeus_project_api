package auth

import (
	"fmt"
	"time"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/User0608/zeus_project_api/models"
	"github.com/dgrijalva/jwt-go"
)

type MyClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func GenerageToken(u models.Usuario) (string, error) {
	claim := MyClaim{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		Username: u.Username,
		Role:     u.OwnerEntity,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", errores.NewInternalf(fmt.Errorf("%s:%w", "auth.GenerageToken", err), errores.ErrSigningTokenString)
	}
	return signedToken, nil
}

func ValidateToken(t string) (MyClaim, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaim{}, verifyFunction)
	if err != nil {
		return MyClaim{}, errores.NewUnauthorizedf(err, errores.ErrInvalidToken)
	}
	if !token.Valid {
		return MyClaim{}, errores.NewUnauthorizedf(err, errores.ErrInvalidToken)
	}
	claim, ok := token.Claims.(*MyClaim)
	if !ok {
		return MyClaim{}, errores.NewUnauthorizedf(err, errores.ErrInvalidToken)
	}
	return *claim, nil
}
func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
