package injectors

import "github.com/User0608/zeus_project_api/handlers"

func GetUsuarioHandler() *handlers.UsuarioHandler {
	return usuarioHandler
}
func GetPersonaHandler() *handlers.PersonaHandler {
	return personaHandler
}
func GetCodigoHandler() *handlers.CodigoHandler {
	return codigoHandler
}
func GetConvocatoria() *handlers.ConvocatoriaHandler {
	return convocatoriaHandler
}
func GetModuloesbasHandler() *handlers.ModuloEsbasHandler {
	return moduloesbasHandler
}
func GetDocumentoHandler() *handlers.DocumentoHandler {
	return documentoHandler
}
