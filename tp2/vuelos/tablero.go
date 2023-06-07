package vuelos

import "os"

type Codigo string

type Claves struct {
	Fecha   string
	Codigo  Codigo
	Origen  string
	Destino string
}

type indice int

const (
	CODIGO indice = iota
	AEROLINEA
	ORIGEN
	DESTINO
	NUM_COLA
	PRIORIDAD
	FECHA
	DEMORA
	TIEMPO
	CANCELADO
)

const CANT_DATOS = CANCELADO + 1

type Vuelo [CANT_DATOS]string

type Tablero interface {
	ObtenerVuelos(K int, modo string, desde, hasta Claves) ([]Vuelo, error)
	ObtenerVuelo(codigo Codigo) (Vuelo, error)
	ObtenerVuelosPrioritarios(K int) []Vuelo
	SiguienteVuelo(origen, destino, fecha string) (Vuelo, error)
	ActualizarTablero(archivo *os.File)
	Borrar(desde, hasta Claves) ([]Vuelo, error)
}
