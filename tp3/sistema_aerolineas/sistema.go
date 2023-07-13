package sistema

import TDADicc "tdas/diccionario"

type (
	Ciudad string

	Codigo string

	indiceVuelo int

	indiceAeropuerto int

	Aeropuerto struct {
		Ciudad   Ciudad
		Codigo   Codigo
		Latitud  float64
		Longitud float64
	}

	Vuelo struct {
		AeropuertoOrigen  Codigo
		AeropuertoDestino Codigo
		Tiempo            float64
		Precio            float64
		Cant_vuelos       float64
	}

	Ruta struct {
		CiudadOrigen  Ciudad
		CiudadDestino Ciudad
	}
)

const (
	ORIGEN indiceVuelo = iota
	DESTINO
	TIEMPO
	PRECIO
	CANT_VUELOS
)

const (
	CIUDAD indiceAeropuerto = iota
	CODIGO
	LATITUD
	LONGITUD
)

type SistemaDeAerolineas interface {
	GuardarAeropuerto(Aeropuerto)
	GuardarVuelo(Vuelo)
	ObtenerCamino(string, Ciudad, Ciudad) []Aeropuerto
	Pertenece(Ciudad) bool
	ObtenerAeropuertosMasImportantes() TDADicc.Diccionario[Aeropuerto, float64]
	ObtenerVuelosRutaMinima() []Vuelo
	ObtenerCaminosItinerario([]Ciudad, []Ruta) ([]Ciudad, [][]Aeropuerto)
	ObtenerUltimaRutaSolicitada() []Aeropuerto
}
