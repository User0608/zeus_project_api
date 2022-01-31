package models

import "time"

type Oficio struct {
	Codigo     string    `chk:"nosp nonil" json:"codigo"`
	DirigidoAl string    `chk:"nonil" json:"dirigido_al"`
	Asunto     string    `chk:"nonil" json:"asunto"`
	Fecha      time.Time `json:"fecha"`
	Contenido  string    `chk:"nonil" json:"contenido,omitempty"`
	CreateBy   string    `json:"created_by"`
	Filepath   string    `json:"path"`
}
