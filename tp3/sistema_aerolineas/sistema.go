package sistema

type Ciudad string

type Codigo string

type indiceVuelo int

type indiceAeropuerto int

type Aeropuerto struct {
	Ciudad   Ciudad
	Codigo   Codigo
	Latitud  float64
	Longitud float64
}

type Vuelo struct {
	AeropuertoOrigen  Codigo
	AeropuertoDestino Codigo
	Tiempo            int
	Precio            int
	Cant_vuelos       int
}

const (
	CIUDAD indiceAeropuerto = iota
	CODIGO
	LATITUD
	LONGITUD
)

const (
	ORIGEN indiceVuelo = iota
	DESTINO
	TIEMPO
	PRECIO
	CANT_VUELOS
)

type SistemaDeAerolineas interface {
	GuardarAeropuerto(Aeropuerto)
	GuardarVuelo(Vuelo)
	ObtenerCaminoMasBarato(Ciudad, Ciudad) []Aeropuerto
	ObtenerCaminoMasRapido(Ciudad, Ciudad) []Aeropuerto
	ObtenerCaminoConMenosEscalas(Ciudad, Ciudad) []Aeropuerto
	Pertenece(Ciudad) bool
	ObtenerAeropuertosMasImportantes(int) []Aeropuerto
	CrearRutaMinima(string)
	CrearItinerario(string)
	ExportarMapaCamino(string)
}
