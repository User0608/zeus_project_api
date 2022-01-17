package dberrs

import (
	"encoding/json"
	"fmt"

	"github.com/User0608/zeus_project_api/errores"
)

type DBError struct {
	Message string `json:"message"`
}

func GetJSONFromText(data string) (string, bool) {
	start := -1
	end := -1
	for i, d := range data {
		if start == -1 {
			if d == '{' {
				start = i
			}
		} else {
			if d == '}' && end == -1 {
				end = i + 1
				break
			}
		}
	}
	if start < end {
		return data[start:end], true
	}
	return "", false
}
func DBBind(err error, location string) error {
	dberr := &DBError{}
	res, ok := GetJSONFromText(err.Error())
	if !ok {
		return errores.NewInternalf(fmt.Errorf("%s:%w", location, err), errores.ErrDatabaseRequest)
	}
	if err := json.Unmarshal([]byte(res), dberr); err != nil {
		return errores.NewInternalf(fmt.Errorf("%s:%w", location, err), errores.ErrDatabaseRequest)
	} else {
		return errores.NewBadRequestf(fmt.Errorf("%s", location), dberr.Message)
	}

}
