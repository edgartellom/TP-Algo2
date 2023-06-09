package acciones

import (
	"fmt"
	"math"
	"os"
	"strconv"

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

var hastaMax = math.Inf(1)

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
	mensaje := mostrarInfoVuelo(*vuelo)
	f.MostrarSalida(mensaje)
	f.MostrarSalida(SALIDA_EXITOSA)
}

func PrioridadVuelos(K string) {
	cantidad, err := strconv.Atoi(K)
	comando := LISTA_COMANDOS[PRIORIDAD_VUELOS]
	if err != nil || cantidad < 1 {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	vuelos := tablero.ObtenerVuelosPrioritarios(cantidad)
	for _, vuelo := range vuelos {
		mensaje := fmt.Sprintf("%d - %s", vuelo.Prioridad, vuelo.Claves.Codigo)
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)
}

func VerTablero(K string, modo string, desde, hasta string) {
	cantidad, err := strconv.Atoi(K)
	fechaDesde := f.ConvertirCadenaAFecha(desde)
	fechaHasta := f.ConvertirCadenaAFecha(hasta)
	claveDesde := v.Claves{Fecha: fechaDesde}
	claveHasta := v.Claves{Fecha: fechaHasta, Codigo: "999999"}
	vuelos, err := tablero.ObtenerVuelos(cantidad, modo, claveDesde, claveHasta)
	comando := LISTA_COMANDOS[VER_TABLERO]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	for _, vuelo := range vuelos {
		cadenaFecha := f.ConvertirFechaACadena(vuelo.Claves.Fecha)
		mensaje := fmt.Sprintf("%s - %s", cadenaFecha, vuelo.Claves.Codigo)
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)

}

func SiguienteVuelo(origen, destino string, fecha string) {
	fechaCorrecta := f.ConvertirCadenaAFecha(fecha)
	claveFecha := v.Claves{Fecha: fechaCorrecta}
	vuelo, panic := tablero.SiguienteVuelo(origen, destino, claveFecha)
	if panic != nil {
		f.MostrarSalida(panic.Error())
		f.MostrarSalida(SALIDA_EXITOSA)
		return
	}
	mensaje := mostrarInfoVuelo(*vuelo)
	f.MostrarSalida(mensaje)
	f.MostrarSalida(SALIDA_EXITOSA)
}

func Borrar(desde, hasta string) {
	fechaDesde := f.ConvertirCadenaAFecha(desde)
	fechaHasta := f.ConvertirCadenaAFecha(hasta)
	claveDesde := v.Claves{Fecha: fechaDesde}
	claveHasta := v.Claves{Fecha: fechaHasta, Codigo: "999999"}
	vuelos, err := tablero.BorrarVuelos(claveDesde, claveHasta)
	comando := LISTA_COMANDOS[BORRAR]
	if err != nil {
		err = e.ErrorComando{Comando: comando}
		f.MostrarError(err)
		return
	}
	for _, vuelo := range vuelos {
		mensaje := mostrarInfoVuelo(vuelo)
		f.MostrarSalida(mensaje)
	}
	f.MostrarSalida(SALIDA_EXITOSA)

}

func mostrarInfoVuelo(vuelo v.Vuelo) string {
	cadenaFecha := f.ConvertirFechaACadena(vuelo.Claves.Fecha)
	mensaje := fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d",
		vuelo.Claves.Codigo, vuelo.Aerolinea, vuelo.Origen, vuelo.Destino,
		vuelo.NumCola, vuelo.Prioridad, cadenaFecha, vuelo.Demora,
		vuelo.Tiempo, vuelo.Cancelado,
	)
	return mensaje
}
