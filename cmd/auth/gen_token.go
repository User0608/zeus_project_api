package auth

import (
	"time"

	"github.com/User0608/zeus_project_api/errs"
	"github.com/User0608/zeus_project_api/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerageToken(u models.Usuario) (string, error) {
	claim := jwt.StandardClaims{
		Id:        u.Username,
		Subject:   u.OwnerEntity,
		Issuer:    "saucedo",
		Audience:  "web/mobile",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", errs.WrapAndMessage(errs.Trc("auth", "GenerageToken"), err, errs.ErrTokenSignedString)
	}
	return signedToken, nil
}

func ValidateToken(t string) (jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(t, &jwt.StandardClaims{}, verifyFunction)
	if err != nil {
		return jwt.StandardClaims{}, errs.Create(errs.Trc("auth", "ValidateToken()"), errs.ErrInvalidToken)
	}
	if !token.Valid {
		return jwt.StandardClaims{}, errs.Create(errs.Trc("auth", "ValidateToken()"), errs.ErrInvalidToken)
	}
	claim, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return jwt.StandardClaims{}, errs.Create(errs.Trc("auth", "ValidateToken()"), errs.ErrInvalidToken)
	}
	return *claim, nil
}
func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
