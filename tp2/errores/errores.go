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
