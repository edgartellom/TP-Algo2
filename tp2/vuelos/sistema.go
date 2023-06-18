package vuelos

type (
	Codigo string

	indice int

	// CamposComparables tiene guardada la informacion necesaria para ordenar los vuelos en diferentes aspectos.
	CamposComparables struct {
		Prioridad int
		Codigo    Codigo
		Fecha     string
	}

	// Vuelo tiene guardada la información principal de un vuelo programado.
	// Además incluye la información completa de este, con el formato conveniente para ser mostrado.
	Vuelo struct {
		InfoComparable      CamposComparables
		Origen              string
		Destino             string
		InformacionCompleta string
	}
)

const (
	CODIGO indice = iota
	_
	ORIGEN
	DESTINO
	_
	PRIORIDAD
	FECHA
	DEMORA
)

// Sistema modela una base de datos de un Aeropuerto, en la cual podremos almacenar y borrar información, además de poder acceder a esta información de diferentes formas.
type SistemaDeVuelos interface {

	//GuardarVuelo nos permite almacenar un Vuelo indicado en la base de datos.
	GuardarVuelo(Vuelo)

	//Pertenece nos permite saber si un Vuelo, con el código de vuelo indicado, está almacenado (o no) en la base de datos.
	Pertenece(Codigo) bool

	//ObtenerVuelo nos permite obtener un Vuelo, guardado en el sistema, con el código de vuelo indicado.
	ObtenerVuelo(string) Vuelo

	//ObtenerVuelosPrioritarios nos permite obtener una cantidad indicada de Vuelos con mayor prioridad entre todos los vuelos almacenados.
	ObtenerVuelosPrioritarios(int) []Vuelo

	//ObtenerVuelosEntreRango nos permite obtener los Vuelos que están contenidos entre las fechas indicadas.
	ObtenerVuelosEntreRango(string, string) []Vuelo

	//ObtenerSiguienteVuelo nos permite obtener (si hay) el Vuelo más próximo con el origen, destino y desde la fecha indicada.
	ObtenerSiguienteVuelo(string, string, string) *Vuelo

	//BorrarVuelos nos permite eliminar y obtener los Vuelos entre las fechas indicadas.
	BorrarVuelos(string, string) []Vuelo
}
