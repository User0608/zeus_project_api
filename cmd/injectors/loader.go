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
}
func initServices() {
	usuarioService = services.NewUsuarioRepository(usuarioRepository)
}
func initHandlers() {
	usuarioHandler = handlers.NewUsuarioHandler(usuarioService)

}
