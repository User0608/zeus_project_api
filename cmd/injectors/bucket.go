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
	usuarioRepository      *repos.UsuarioRepository
	personaRepository      *repos.PersonaRepository
	codigoRepository       *repos.CodigoRepository
	convocatoriaRepository *repos.ConvocatoriaRepository
	moduloesbasRepository  *repos.ModuloEsbasRepository
	documentoRepository    *repos.DocumentoRepository
	cronogramaRepository   *repos.CronogramaRepository

	//services
	usuarioService      *services.UsuarioService
	personaService      *services.PersonaService
	codigoService       *services.CodigoService
	convocatoriaService *services.ConvocatoriaService
	moduloesbasService  *services.ModuloEsbasService
	documentoService    *services.DocumentoService
	cronogramaService   *services.CronogramaService

	//handlers
	usuarioHandler      *handlers.UsuarioHandler
	personaHandler      *handlers.PersonaHandler
	codigoHandler       *handlers.CodigoHandler
	convocatoriaHandler *handlers.ConvocatoriaHandler
	moduloesbasHandler  *handlers.ModuloEsbasHandler
	documentoHandler    *handlers.DocumentoHandler
	cronogramaHandler   *handlers.CronogramaHandler
)
