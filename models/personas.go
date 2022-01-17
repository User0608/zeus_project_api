package models

import (
	"time"
)

type EstadoCivil struct {
	ID     uint   `json:"id"`
	Estado string `json:"estado"`
}
type NivelEstudio struct {
	ID    uint   `json:"id"`
	Nivel string `json:"nivel"`
}
type InformacionBaseEntity struct {
	EstadosCiviles []EstadoCivil  `json:"estados_civiles"`
	NivelesEstudio []NivelEstudio `json:"niveles_estudio"`
}
type Entity struct {
	Dni             string    `chk:"num len=8" gorm:"primaryKey" json:"dni"`
	Nombre          string    `chk:"stxt" json:"nombre"`
	ApellidoPaterno string    `chk:"stxt" json:"ap_paterno"`
	ApellidoMaterno string    `chk:"stxt" json:"ap_materno"`
	Direccion       string    `chk:"nosp" json:"direccion"`
	Telefono        string    `chk:"nonil num max=15" json:"tel"`
	Email           string    `chk:"email" json:"email"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	NivelEstudioID  uint      `json:"nivel_estudio"`
	EstadoCivilID   uint      `json:"estado_civil"`
	FechaRegistro   string    `json:"-" gorm:"->"`
}
type PrimerJefe struct {
	Dni           string
	State         bool
	FechaRegistro time.Time `json:"-" gorm:"->"`
}
type PrimerJefeEntity struct {
	Entity        `gorm:"foreignKey:dni"`
	State         bool      `json:"-"`
	FechaRegistro time.Time `json:"-" gorm:"->"`
}

func (pje *PrimerJefeEntity) GetPrimerJefeObjetct() *PrimerJefe {

	return &PrimerJefe{Dni: pje.Dni, State: pje.State, FechaRegistro: pje.FechaRegistro}
}
func (*PrimerJefeEntity) TableName() string { return "primer_jefe" }

///////////////////////////////////////////////
type SegundoJefe struct {
	Dni           string
	State         bool
	FechaRegistro time.Time `json:"-" gorm:"->"`
}

type SegundoJefeEntity struct {
	Entity
	State         bool      `json:"-"`
	FechaRegistro time.Time `json:"-" gorm:"->"`
}

func (pje *SegundoJefeEntity) GetSegundoJefeObjetct() *SegundoJefe {

	return &SegundoJefe{Dni: pje.Dni, State: pje.State, FechaRegistro: pje.FechaRegistro}
}
func (*SegundoJefeEntity) TableName() string { return "segundo_jefe" }

type JefeInstruccion struct {
	Dni           string `gorm:"primaryKey"`
	State         bool
	Detalle       string
	FechaRegistro time.Time `json:"-" gorm:"->"`
}
type JefeInstruccionEntity struct {
	Entity
	State         bool      `json:"state"`
	Detalle       string    `json:"detalle"`
	FechaRegistro time.Time `json:"-" gorm:"->"`
}

func (ji *JefeInstruccionEntity) GetJefeInstructor() *JefeInstruccion {
	return &JefeInstruccion{
		Dni:           ji.Dni,
		State:         ji.State,
		Detalle:       ji.Detalle,
		FechaRegistro: ji.FechaRegistro,
	}
}
