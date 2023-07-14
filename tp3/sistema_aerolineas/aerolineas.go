package sistema

import TDADicc "tdas/diccionario"

type (
	Ciudad string

	Codigo string

	indiceVuelo int

	indiceAeropuerto int

	// Aeropuerto tiene guardada la informacion completa de un aeropuerto, con su respectivo, codigo de aeropuerto, latitud, longitud y ciudad de este.
	Aeropuerto struct {
		Ciudad   Ciudad
		Codigo   Codigo
		Latitud  float64
		Longitud float64
	}

	// Vuelo tiene guardada la informacion de un vuelo programado, con su respectivo origen, destino, tiempo de vuelo, precio del vuelo y
	// cantidad de vuelos entre el origen y destino.
	Vuelo struct {
		AeropuertoOrigen  Codigo
		AeropuertoDestino Codigo
		Tiempo            float64
		Precio            float64
		Cant_vuelos       float64
	}

	// Ruta tiene guardada la informacion de una ruta de vuelo trazado desde su ciudad de origen hasta su destino.
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

	// GuardarAeropuerto guarda un aeropuerto indicado por parametro en el sistema.
	GuardarAeropuerto(Aeropuerto)

	// GuardarVuelo guarda un vuelo indicado en el sistema, estableciendo una unión entre aeropuertos.
	GuardarVuelo(Vuelo)

	// ObtenerCamino devuelve un arreglo de Aeropuertos que representan el camino minimo desde una Ciudad Origen hasta una Ciudad Destino, indicados por parametro
	// respectivamente, con el criterio que haya sido indicado por parametro.
	ObtenerCamino(string, Ciudad, Ciudad) []Aeropuerto

	// Pertenece devuelve si la Ciudad indicada por parametro, esta almacenada en el sistema.
	Pertenece(Ciudad) bool

	// ObtenerAeropuertosMasImportantes devuelve un Diccionario de clave Aeropuerto y valor Centralidad, siendo "Centralidad" la frecuencia con la que se pasa por
	// ese aeropuerto de los todos los aeropuertos almacenados en el sistema.
	ObtenerAeropuertosMasImportantes() TDADicc.Diccionario[Aeropuerto, float64]

	// ObtenerVuelosRutaMinima devuelve un arreglo de Vuelos necesarios para conectar todos los Aeropuertos con el costo minimo en relacion al precio.
	ObtenerVuelosRutaMinima() []Vuelo

	// ObtenerCaminosItinerario crea un "sistema" temporal de vuelos con los arreglos de Ciudades y Rutas indicados por parametro, y devuelve un orden especifico, en
	// forma de arreglo, para visitar todas las ciudades indicadas en el arreglo. Adicionalmente devuelve un arreglo de arreglos con los Aeropuertos necesarios a
	// visitar, para poder visitar todas las ciudades deseadas.
	ObtenerCaminosItinerario([]Ciudad, []Ruta) ([]Ciudad, [][]Aeropuerto)

	// ObtenerUltimaRutaSolicitada devuelve el ultimo arreglo de Aeropuertos obtenido, siendo este el que representa una ruta de vuelo.
	ObtenerUltimaRutaSolicitada() []Aeropuerto
}
