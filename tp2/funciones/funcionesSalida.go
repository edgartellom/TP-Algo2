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

// func MensajeInfoVuelo(vuelo vuelos.Vuelo) string{
// 	codigo := fmt.Sprintf("%s", vuelo.Codigo)
// 	aerolinea := fmt.Sprintf("%s", vuelo.Datos.Aerolinea)
// 	return fmt.Sprintf("%s %s %s %s", partido.nombreDelPartido, partido.votosActuales[tipo], cantidad)
// }
