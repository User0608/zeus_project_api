package models

import (
	"errors"
	"strings"
	"time"

	"github.com/User0608/zeus_project_api/errs"
)

type LogginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostUsuario struct {
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	OwnerEntity string `json:"owner"`
}

type Usuario struct {
	PostUsuario
	State     bool      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func (pu *PostUsuario) TableName() string { return "usuario" }

func (pu *PostUsuario) Valid() error {
	usernameLength := len(strings.TrimSpace(pu.Username))
	passwordLength := len(strings.TrimSpace(pu.Password))
	if usernameLength < 6 {
		return errors.New(errs.ErrUsername)
	}
	if passwordLength < 6 {
		return errors.New(errs.ErrPassword)
	}
	return nil
}
func (r *LogginRequest) Valid() error {
	usernameLength := len(strings.TrimSpace(r.Username))
	passwordLength := len(strings.TrimSpace(r.Password))
	if usernameLength < 6 {
		return errors.New(errs.ErrUsername)
	}
	if passwordLength < 6 {
		return errors.New(errs.ErrPassword)
	}
	return nil
}
