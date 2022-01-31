package models

type Actividad struct {
	ID             uint   `json:"id"`
	Titulo         string `chk:"nonil" json:"titulo"`
	Detalle        string `json:"detalle"`
	HoraInicio     uint   `json:"hora_inicio"`
	HoraFin        uint   `json:"hora_fin"`
	OwnerDni       string `json:"owner_dni"`
	ProgramacionID uint   `json:"-"`
}
