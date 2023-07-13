package errores

import (
	"fmt"
)

type ErrorComando struct {
	Comando string
}

func (e ErrorComando) Error() string {
	return fmt.Sprintf("Error en comando %s", e.Comando)
}

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}
