package models

import "time"

type Informe struct {
	Codigo     string    `chk:"nosp nonil" json:"codigo"`
	ParteDel   string    `chk:"nonil" json:"parte_del"`
	DirigidoAl string    `chk:"nonil" json:"dirigido_al"`
	Asunto     string    `chk:"nonil" json:"asunto"`
	Fecha      time.Time `json:"fecha"`
	Contenido  string    `chk:"nonil" json:"contenido,omitempty"`
	CreateBy   string    `json:"created_by"`
	Filepath   string    `json:"path"`
}
