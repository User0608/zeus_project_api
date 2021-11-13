package errs_test

import (
	"errors"
	"testing"

	"github.com/User0608/zeus_project_api/errs"
)

func TestErrorWithoutMessage(t *testing.T) {
	want := "errs_test.TestError:error base, algo paso"
	myerr := errors.New("error base, algo paso")
	err := errs.Wrap(errs.Trc("errs_test", "TestError"), myerr)
	if err.Error() != want {
		t.Error("Tengo:", err.Error(), "Quiero:", want)
	}
}
func TestErrorCreate(t *testing.T) {
	want := "errs_test.TestError;message=el usuario ya existe"
	err := errs.Create(errs.Trc("errs_test", "TestError"), "el usuario ya existe")
	if err.Error() != want {
		t.Error("Tengo:", err.Error(), "Quiero:", want)
	}
}
func TestErrorWrapWithMessage(t *testing.T) {
	want := "errs_test.TestError:error base, algo paso;message=el usuario ya existe"
	myerr := errors.New("error base, algo paso")
	err := errs.WrapAndMessage(errs.Trc("errs_test", "TestError"), myerr, "el usuario ya existe")
	if err.Error() != want {
		t.Error("Tengo:", err.Error(), "Quiero:", want)
	}
}
func TestError(t *testing.T) {
	want := "services.CreateUser:errs_test.TestError:database.conn;message=no se pudo connectar;message=el usuario ya existe"
	//myerr := errors.New("error base, algo paso")
	err0 := errs.Create(errs.Trc("database", "conn"), "no se pudo connectar")
	err1 := errs.WrapAndMessage(errs.Trc("errs_test", "TestError"), err0, "el usuario ya existe")
	err := errs.Wrap(errs.Trc("services", "CreateUser"), err1)
	if err.Error() != want {
		t.Error("Tengo:", err.Error(), "Quiero:", want)
	}
}
func TestMesage(t *testing.T) {
	want := "services.CreateUser:errs_test.TestError:database.conn;message=no se pudo connectar;message=el usuario ya existe"
	err := errors.New(want)
	message := errs.RecuperarMessage(err)
	if message != "el usuario ya existe" {
		t.Error("Tengo:", message, "Quiero:", "el usuario ya existe")
	}
}
