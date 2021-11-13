package errs

import "net/http"

const (
	ErrInvalidJSON           = "estructura json invalida"
	ErrInvalidToken          = "el token es invalido"
	ErrTokenNull             = "no se encontro el token"
	ErrTokenSignedString     = "no se pudo authentificar"
	ErrNoDefined             = "hubo un error, no esperado"
	ErrDatabaseRequest       = "no se pudo realizar la operacion"
	ErrRecordNotFaund        = "no se encontrar el registro"
	ErrRecord                = "no se pudo guardar el registro"
	ErrUsernameExists        = "el usuario ya existe"
	ErrAuthorizationHeader   = "Authorization header no encontrado"
	ErrPassword              = "password debe de tener almenos 6 caracteres"
	ErrUsername              = "username debe de tener almenos 6 caracteres"
	ErrUserOrPasswordInvalid = "usuario o password incorrectos"
)

var httpCodes = map[string]int{
	ErrInvalidJSON:           http.StatusBadRequest,
	ErrAuthorizationHeader:   http.StatusForbidden,
	ErrInvalidToken:          http.StatusForbidden,
	ErrTokenNull:             http.StatusForbidden,
	ErrNoDefined:             http.StatusInternalServerError,
	ErrDatabaseRequest:       http.StatusInternalServerError,
	ErrRecordNotFaund:        http.StatusNotFound,
	ErrRecord:                http.StatusInternalServerError,
	ErrUsernameExists:        http.StatusBadRequest,
	ErrUserOrPasswordInvalid: http.StatusBadRequest,
}
