package vuelos

// type Fecha struct {
// 	fecha string
// }

type Claves struct {
	codigo  string
	fecha   string
	origen  string
	destino string
}

// Vuelo modela un Boleto de Avion, con la respectiva informacion escrita en el mismo.
type Vuelo interface {

	// Devuelve el valor de Prioridad del Vuelo.
	VerPrioridad() int

	// Devuelve la Fecha a la que esta programada el Vuelo.
	VerFecha() string

	// Devuelve el Codigo alfanumerico del Vuelo.
	VerCodigo() string

	// Devuelve un resumen de la informacion mas principal.
	VerInformacionPrincipal() Claves

	// Devuelve la Imformacion Completa del Vuelo.
	ObtenerInformacionDeVuelo() string
}
