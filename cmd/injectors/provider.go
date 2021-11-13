package injectors

import "github.com/User0608/zeus_project_api/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandler {
	return usuarioHandler
}
