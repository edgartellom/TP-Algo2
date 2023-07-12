package errores

import (
	"fmt"
)

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type ErrorComando struct {
	Comando string
}

func (e ErrorComando) Error() string {
	return fmt.Sprintf("Error en comando %s", e.Comando)
}
