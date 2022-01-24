package models

import "time"

type Programacion struct {
	ID           uint         `json:"id"`
	Nombre       string       `json:"nombre"`
	Detalle      string       `json:"detalle"`
	Fecha        time.Time    `json:"fecha"`
	CronogramaID uint         `json:"cronograma_id"`
	Actividades  []*Actividad `json:"actividades,omitempty"`
}
