package injectors

import (
	"github.com/User0608/zeus_project_api/handlers"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/User0608/zeus_project_api/services"
	"gorm.io/gorm"
)

var ( //db connextion
	Connextion *gorm.DB
	//repository
	usuarioRepository *repos.UsuarioRepository

	//services
	usuarioService *services.UsuarioService

	//handlers
	usuarioHandler *handlers.UsuarioHandler
)
