package sistema

type (
	Ciudad string

	Codigo string

	indiceVuelo int

	indiceAeropuerto int

	Vuelo struct {
		AeropuertoOrigen  Codigo
		AeropuertoDestino Codigo
		Tiempo            int
		Precio            int
		Cant_vuelos       int
	}
	Aeropuerto struct {
		Ciudad   Ciudad
		Codigo   Codigo
		Latitud  float64
		Longitud float64
	}
)

const (
	ORIGEN indiceVuelo = iota
	DESTINO
	TIEMPO
	PRECIO
	CANT_VUELOS

	CIUDAD indiceAeropuerto = iota
	CODIGO
	LATITUD
	LONGITUD
)

type SistemaDeAerolineas interface {
	GuardarAeropuerto(Aeropuerto)
	GuardarVuelo(Vuelo)
	ObtenerCaminoMasBarato(Ciudad, Ciudad) []Aeropuerto
	ObtenerCaminoMasRapido(Ciudad, Ciudad) []Aeropuerto
	ObtenerCaminoConMenosEscalas(Ciudad, Ciudad) []Aeropuerto
	ObtenerAeropuertosMasImportantes(int)
	CrearRutaMinima(string)
	CrearItinerario(string)
	ExportarMapaCamino(string)
}
