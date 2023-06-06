package acciones

import (
	"fmt"
	"os"
	"strings"

	e "algueiza/errores"
	f "algueiza/funciones"
	"algueiza/vuelos"
)

type comando int

const (
	SALIDA_EXITOSA = "OK"
)

const (
	AGREGAR_ARCHIVO comando = iota
	VER_TABLERO
	INFO_VUELO
	PRIORIDAD_VUELOS
	SIGUIENTE_VUELO
	BORRAR
)

var LISTA_COMANDOS = []string{"agregar_archivo", "ver_tablero", "info_vuelo", "prioridad_vuelos", "siguiente_vuelo", "borrar"}

var tablero = vuelos.CrearTablero()

func AgregarArchivo(ruta string) {
	archivo, err := os.Open(ruta)
	comando := LISTA_COMANDOS[AGREGAR_ARCHIVO]
	defer archivo.Close()
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	tablero.ActualizarTablero(archivo)
	f.MostrarSalida(SALIDA_EXITOSA)
}

func InfoVuelo(codigo string) {
	vuelo, err := tablero.ObtenerVuelo(codigo)
	comando := LISTA_COMANDOS[INFO_VUELO]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	mensaje := fmt.Sprintf("%s ", strings.Join(vuelo.Datos, " "))
	f.MostrarSalida(mensaje)
	f.MostrarSalida(SALIDA_EXITOSA)
}

func PrioridadVuelos(K int) {
	vuelos := tablero.ObtenerVuelosPrioritarios(K)
	for _, vuelo := range vuelos {
		mensaje := fmt.Sprintf("%d - %s", vuelo.Prioridad, vuelo.Codigo)
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)
}
