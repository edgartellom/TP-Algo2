package acciones

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	e "algueiza/errores"
	f "algueiza/funciones"
	v "algueiza/vuelos"
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

const CANT_COMANDOS = BORRAR + 1

var LISTA_COMANDOS = [CANT_COMANDOS]string{"agregar_archivo", "ver_tablero", "info_vuelo", "prioridad_vuelos", "siguiente_vuelo", "borrar"}

var tablero = v.CrearTablero()

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
	vuelo, err := tablero.ObtenerVuelo(v.Codigo(codigo))
	comando := LISTA_COMANDOS[INFO_VUELO]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	mensaje := fmt.Sprintf("%s ", strings.Join(vuelo[:], " "))
	f.MostrarSalida(mensaje)
	f.MostrarSalida(SALIDA_EXITOSA)
}

func PrioridadVuelos(K string) {
	cantidad, err := strconv.Atoi(K)
	comando := LISTA_COMANDOS[PRIORIDAD_VUELOS]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	vuelos := tablero.ObtenerVuelosPrioritarios(cantidad)
	for _, vuelo := range vuelos {
		mensaje := fmt.Sprintf("%s - %s", vuelo[v.PRIORIDAD], vuelo[v.CODIGO])
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)
}

func VerTablero(K string, modo string, desde, hasta string) {
	cantidad, err := strconv.Atoi(K)
	claveDesde := v.Claves{Fecha: desde}
	claveHasta := v.Claves{Fecha: hasta}
	vuelos, err := tablero.ObtenerVuelos(cantidad, modo, claveDesde, claveHasta)
	comando := LISTA_COMANDOS[VER_TABLERO]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	for _, vuelo := range vuelos {
		mensaje := fmt.Sprintf("%s - %s", vuelo[v.FECHA], vuelo[v.CODIGO])
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)

}
