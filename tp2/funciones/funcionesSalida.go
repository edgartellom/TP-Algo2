package funciones

import (
	// "algueiza/vuelos"
	"fmt"
	"os"
)

func MostrarError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
}

func MostrarSalida(mensaje string) {
	fmt.Fprintf(os.Stdout, "%s\n", mensaje)
}
