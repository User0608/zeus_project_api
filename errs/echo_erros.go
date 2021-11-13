package errs

import (
	"log"
	"strings"

	"github.com/User0608/zeus_project_api/utils"
	"github.com/labstack/echo/v4"
)

func printLog(err error) {
	log.Println(err.Error())
}
func GenErrEchoResponse(c echo.Context, err error) error {
	go printLog(err)
	message := RecuperarMessage(err)
	code, ok := httpCodes[message]
	if !ok {
		code = 400
	}
	return c.JSON(code, utils.Response{Message: message})
}
func RecuperarMessage(err error) string {
	field := strings.Split(err.Error(), ";")
	end := len(field) - 1
	message := ErrNoDefined
	if end >= 0 {
		target := "message="
		cut := len(target)
		message = strings.TrimSpace(field[end])
		if len(message) > cut {
			if target == message[:cut] {
				message = strings.TrimSpace(field[end])[cut:]
			}
		}
	}
	return message
}
