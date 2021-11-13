package errs

import (
	"fmt"
	"strings"
)

type Trace string

// Wrap recibe la ubicacion actual y el error actual, no hay mensaje para el usuario
func Wrap(tracer Trace, err error) error {
	return fmt.Errorf("%s:%v", tracer, err)
}

// WrapAndMessage recibe tanto la ubicacion actual, el error y inserta un mensaje
func WrapAndMessage(tracer Trace, errInternal error, clientMessage string) error {
	clientMessage = strings.TrimSpace(clientMessage)
	return fmt.Errorf("%s:%v;message=%s", tracer, errInternal, clientMessage)
}

// Create :genera un error juncto con el mensaje y la ubicacion
func Create(tracer Trace, clientMessage string) error {
	clientMessage = strings.TrimSpace(clientMessage)
	return fmt.Errorf("%s;message=%s", tracer, clientMessage)
}

// Trc :genera la ubicacion, package.struc.proced <-
func Trc(current, location string) Trace {
	location = strings.TrimSpace(location)
	current = strings.TrimSpace(current)
	return Trace(fmt.Sprintf("%s.%s", current, location))
}
func UsuarioNotFount(username string) string {
	return fmt.Sprintf("el registro con username = %s no existe", username)
}
