package response

import "github.com/User0608/zeus_project_api/models"

type ResumenPrimerSegunJefe struct {
	PrimerJefe  *models.PrimerJefeEntity  `json:"primer_jefe,omitempty"`
	SegundoJefe *models.SegundoJefeEntity `json:"segundo_jefe,omitempty"`
}
