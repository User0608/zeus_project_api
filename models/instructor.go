package models

type Instructor struct {
	Dni             string `gorm:"primaryKey" json:"dni"`
	State           bool   `json:"state"`
	FechaInstructor string `gorm:"->" json:"fecha_instructor"`
	Detalle         string `json:"detalles"`
}
type InstructorEntity struct {
	Entity
	State           bool   `json:"state"`
	FechaInstructor string `gorm:"->" json:"fecha_instructor"`
	Detalle         string `json:"detalle"`
}

func (i *InstructorEntity) GetInstructor() *Instructor {
	return &Instructor{
		Dni:             i.Dni,
		State:           i.State,
		FechaInstructor: i.FechaInstructor,
		Detalle:         i.Detalle,
	}
}
func (i *InstructorEntity) SetInstructor(ins *Instructor) {
	i.State = ins.State
	i.FechaInstructor = ins.FechaInstructor
	i.Detalle = ins.Detalle
}
