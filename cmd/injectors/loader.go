package injectors

import (
	"sync"

	"github.com/User0608/zeus_project_api/configs"
	"github.com/User0608/zeus_project_api/database"
	"github.com/User0608/zeus_project_api/handlers"
	"github.com/User0608/zeus_project_api/repos"
	"github.com/User0608/zeus_project_api/services"
	"gorm.io/gorm"
)

var ones sync.Once

func LoadConfig(configFilePath string) error {
	var errr error
	ones.Do(func() {
		if conf, err := configs.LoadDBConfigs(configFilePath); err != nil {
			errr = err
			return
		} else {
			con, er := database.GetDBConnextion(conf)
			if err != nil {
				errr = er
				return
			}
			initRepository(con)
			initServices()
			initHandlers()
		}
	})
	return errr
}
func initRepository(connextion *gorm.DB) {
	Connextion = connextion
	usuarioRepository = repos.NewUsuarioRepository(connextion)
	personaRepository = repos.NewPersonaRepository(connextion)
	codigoRepository = repos.NewCodigoRepository(connextion)
	convocatoriaRepository = repos.NewConvocatoriaRepository(connextion)
	moduloesbasRepository = repos.NewModuloEsbasRepository(connextion)
	documentoRepository = repos.NewDocumentoRepository(connextion)
	cronogramaRepository = repos.NewCronogramaRepository(connextion)
	reporteRepository = repos.NewReportesRepository(connextion)
}
func initServices() {
	usuarioService = services.NewUsuarioRepository(usuarioRepository)
	personaService = services.NewPersonaService(personaRepository)
	codigoService = services.NewCodigoService(codigoRepository)
	convocatoriaService = services.NewConvocatoriaService(convocatoriaRepository)
	moduloesbasService = services.NewModuloEsbasService(moduloesbasRepository)
	documentoService = services.NewDocumentoService(documentoRepository)
	cronogramaService = services.NewCronogramaService(cronogramaRepository)
	reporteService = services.NewReporteService(reporteRepository)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)
	personaHandler = handlers.NewPersonaHandler(personaService)
	codigoHandler = handlers.NewCodigoHandler(codigoService)
	convocatoriaHandler = handlers.NewConvocatoriaHandler(convocatoriaService)
	moduloesbasHandler = handlers.NewModuloEsbasHanlder(moduloesbasService)
	documentoHandler = handlers.NewDocumentoHandler(documentoService)
	cronogramaHandler = handlers.NewCronogramaHandler(cronogramaService)
	reporteHandler = handlers.NewRepositoyHandler(reporteService)
}
