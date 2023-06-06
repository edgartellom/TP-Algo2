package vuelos

import "os"

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
	ObtenerVuelos(K int, modo string, desde, hasta *Vuelo) ([]Vuelo, error)
	ObtenerVuelo(codigo string) (Vuelo, error)
	ObtenerVuelosPrioritarios(K int) []Vuelo
	SiguienteVuelo(origen, destino, fecha string) (Vuelo, error)
	ActualizarTablero(archivo *os.File)
	Borrar(desde, hasta *Vuelo) ([]Vuelo, error)
}
