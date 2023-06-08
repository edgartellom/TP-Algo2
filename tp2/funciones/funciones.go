package funciones

import (
	"algueiza/errores"
	"algueiza/vuelos"
	"fmt"
	"os"
	"strconv"
	TDAPila "tdas/pila"
)

const (
	MODO_ASCENDETE   = "asc"
	MODO_DESCENDENTE = "desc"
)

func AbrirArchivo(ruta string) *os.File {
	archivo, _ := os.Open(ruta)
	// if err != nil {
	// 	MostrarError(new(errores.ErrorLeerArchivo))
	// }
	return archivo
}

func MostrarSalida(mensaje string) {
	fmt.Fprintf(os.Stdout, "%s\n", mensaje)
}

func ComprobarEntradaVerTablero(cantidad, modo, desde, hasta string) (int, error) {
	cant, err := strconv.Atoi(cantidad)
	if (err != nil) || (cant <= 0) || (modo != MODO_ASCENDETE && modo != MODO_DESCENDENTE) || (hasta < desde) {
		err = errores.ErrorComando{Comando: "VerTablero"}
		return -1, err
	}
	return cant, nil
}

func InvertirOrden(arreglo []vuelos.Vuelo) {
	pilaAux := TDAPila.CrearPilaDinamica[vuelos.Vuelo]()
	for _, e := range arreglo {
		pilaAux.Apilar(e)
	}
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = pilaAux.Desapilar()
	}
}
