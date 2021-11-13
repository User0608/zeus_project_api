package auth

import (
	"github.com/User0608/zeus_project_api/errs"
	"github.com/labstack/echo/v4"
)

const (
	USERNAME_CONTEXT_KEY = "username_context_key"
	OWNER_CONTEXT_KEY    = "owner_context_key"
	IS_ADMIN_CONTEXT_KEY = "is_admin_context_key"
)

func jwtValidToken(token string, c echo.Context) error {
	if token == "" {
		err := errs.Create(errs.Trc("auth", "ValidateToken()"), errs.ErrTokenNull)
		return errs.GenErrEchoResponse(c, err)
	}
	clain, err := ValidateToken(token)
	if err != nil {
		return errs.GenErrEchoResponse(c, err)
	}
	c.Set(USERNAME_CONTEXT_KEY, clain.Id)
	c.Set(OWNER_CONTEXT_KEY, clain.Subject)
	return nil
}
func JWTMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if err := jwtValidToken(token, c); err != nil {
			er := errs.WrapAndMessage(errs.Trc("auth", "JWTMiddleware"), err, errs.ErrAuthorizationHeader)
			return errs.GenErrEchoResponse(c, er)
		}
		return f(c)
	}
}
