package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Num     string      `json:"num,omitempty"`
	Page    string      `json:"page,omitempty"`
}

func OKToken(c echo.Context, token string, data interface{}) error {
	return c.JSON(http.StatusOK, &Response{Token: token, Data: data})
}
func OKResponse(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, &Response{Data: payload})
}
func OKMessage(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, &Response{Message: message})
}
func OKSuccess(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{Message: "operación realizada con éxito"})
}
