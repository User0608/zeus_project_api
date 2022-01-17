package models

import "time"

type Aspirante struct {
	Dni                string    `json:"dni"`
	Delete             bool      `json:"delete"`
	FechaRegistro      time.Time `json:"fecha_registro"`
	Horas              int       `json:"horas"`
	Password           string    `json:"password"`
	IsPaswordDefault   bool      `json:"passup"`
	PasswordUpdate     time.Time `json:"password_update"`
	ConvocatoriaNombre string    `json:"convocatoria"`
}
