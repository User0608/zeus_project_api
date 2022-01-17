package auth

import (
	"errors"

	"github.com/User0608/zeus_project_api/errores"
	"github.com/labstack/echo/v4"
)

const (
	USERNAME_CONTEXT_KEY = "username_context_key"
	OWNER_CONTEXT_KEY    = "owner_context_key"
	IS_ADMIN_CONTEXT_KEY = "is_admin_context_key"
)

func jwtValidToken(token string, c echo.Context) error {
	if token == "" {
		return errores.NewBadRequestf(errors.New("auth.jwtValidToken"), errores.ErrTokenNull)
	}
	_, err := ValidateToken(token)
	if err != nil {
		return err
	}
	// c.Set(USERNAME_CONTEXT_KEY, clain.Id)
	// c.Set(OWNER_CONTEXT_KEY, clain.Subject)
	return nil
}
func JWTMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if err := jwtValidToken(token, c); err != nil {
			return errores.ErrorResponse(c, err)
		}
		return f(c)
	}
}
