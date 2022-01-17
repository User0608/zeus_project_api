package models

import (
	"time"
)

type Memorando struct {
	Codigo     string    `json:"codigo"`
	ParteDel   string    `json:"parte_del"`
	DirigidoAl string    `json:"dirigido_al"`
	Asunto     string    `json:"asunto"`
	Fecha      time.Time `json:"fecha"`
	Contenido  string    `json:"contenido,omitempty"`
	CreateBy   string    `json:"created_by"`
	Filepath   string    `json:"path"`
}
