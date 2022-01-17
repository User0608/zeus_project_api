package models

import (
	"time"
)

type LogginRequest struct {
	Username string `chk:"nonil" json:"username"`
	Password string `chk:"nonil" json:"password"`
}

type PostUsuario struct {
	Username    string `chk:"nosp min=6" json:"username"`
	Password    string `chk:"min=6" json:"password,omitempty"`
	OwnerEntity string `json:"owner"`
}

type Usuario struct {
	PostUsuario
	State     bool      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func (pu *PostUsuario) TableName() string { return "usuario" }
