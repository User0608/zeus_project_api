package response

import "github.com/User0608/zeus_project_api/models"

type ResumenInstructor struct {
	JefeInstruccion *models.JefeInstruccionEntity `json:"jefe_instrucion"`
	Instructores    []*models.InstructorEntity    `json:"instructores"`
}
