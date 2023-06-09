package vuelos

type (
	Codigo string

	indice int

	CamposComparables struct {
		Prioridad int
		Fecha     string
		Codigo    Codigo
	}

	Vuelo struct {
		InfoComparable      CamposComparables
		Origen              string
		Destino             string
		DemoraDeDespegue    int
		TiempoDeVuelo       int
		Cancelacion         int
		InformacionCompleta string
	}
)

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

type Tablero interface {
	GuardarVuelo(Vuelo)
	Pertenece(string) bool
	ObtenerVuelo(string) Vuelo
	ObtenerVuelosPrioritarios(K int) []Vuelo
	ObtenerVuelosEntreRango(int, string, string) []Vuelo

	ObtenerSiguienteVuelo(origen, destino, fecha string) *Vuelo
	Borrar(desde, hasta string) []Vuelo
}
