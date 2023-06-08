package errores

import (
	"fmt"
)

// type ErrorLeerArchivo struct{}

// func (e ErrorLeerArchivo) Error() string {
// 	return "ERROR: Lectura de archivos"
// }

type ErrorComando struct {
	Comando string
}

func (e ErrorComando) Error() string {
	return fmt.Sprintf("Error en comando %s", e.Comando)
}
