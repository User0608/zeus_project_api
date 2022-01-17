package models

import "time"

type Convocatoria struct {
	Nombre      string    `chk:"nonil" gorm:"primaryKey" json:"nombre"`
	Fecha       time.Time `json:"fecha"`
	Descripcion string    `json:"descripcion"`
	CreadoPr    string    `json:"created_por"`
}
