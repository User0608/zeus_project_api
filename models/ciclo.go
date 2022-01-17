package models

import "time"

type Ciclo struct {
	Nombre   string    `json:"nombre"`
	Fecha    time.Time `json:"fecha"`
	Detalles string    `json:"detalles"`
}
