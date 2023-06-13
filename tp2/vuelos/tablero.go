package vuelos

import (
	"os"
	"time"
)

type Codigo string

type Claves struct {
	Codigo Codigo
	Fecha  time.Time
}

type Vuelo struct {
	Claves    Claves
	Aerolinea string
	Origen    string
	Destino   string
	NumCola   string
	Prioridad int
	Demora    int
	Tiempo    int
	Cancelado int
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

// const CANT_DATOS = CANCELADO + 1

// type Vuelo [CANT_DATOS]string

type Tablero interface {
	ObtenerVuelos(K int, modo string, desde, hasta Claves) ([]Vuelo, error)
	ObtenerVuelo(codigo Codigo) (*Vuelo, error)
	ObtenerVuelosPrioritarios(K int) []Vuelo
	SiguienteVuelo(origen, destino string, fecha Claves) (*Vuelo, error)
	ActualizarTablero(archivo *os.File)
	BorrarVuelos(desde, hasta Claves) ([]Vuelo, error)
}
