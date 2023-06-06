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

type Vuelo struct {
	Codigo    string
	Prioridad int
	Origen    string
	Destino   string
	Fecha     string
	Datos     []string
}

type Tablero interface {
	ObtenerVuelos(K int, modo string, desde, hasta *Vuelo) ([]Vuelo, error)
	ObtenerVuelo(codigo string) (Vuelo, error)
	ObtenerVuelosPrioritarios(K int) []Vuelo
	SiguienteVuelo(origen, destino, fecha string) (Vuelo, error)
	ActualizarTablero(archivo *os.File)
	Borrar(desde, hasta *Vuelo) ([]Vuelo, error)
}
