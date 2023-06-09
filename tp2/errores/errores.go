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

type ErrorSiguienteVuelo struct {
	Origen  string
	Destino string
	Fecha   string
}

func (e ErrorSiguienteVuelo) Error() string {
	return fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", e.Origen, e.Destino, e.Fecha)
}
