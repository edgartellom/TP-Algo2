package sistema

type (
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
		Ciudad   string
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
	ObtenerCaminoMasBarato(string, string)
	ObtenerCaminoMasRapido(string, string)
	ObtenerCaminoConMenosEscalas(string, string)
	ObtenerAeropuertosMasImportantes(int)
	CrearRutaMinima(string)
	CrearItinerario(string)
	ExportarMapaCamino(string)
}
