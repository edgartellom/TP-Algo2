package sistema

import (
	"strconv"
	"strings"

	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

const (
	SEPARADOR_COMA    = ","
	SEPARADOR_ESPACIO = " "
)

type sistemaDeAerolineas struct {
	aeropuertosAlmacenados TDADicc.Diccionario[string, []Codigo] // El Hash debe ser de clave: Ciudades, valor: []Codigo; para obtener todos los codigos de aeropuertos asociados a una ciudad, porque una ciudad puede tener mas de un aeropuerto.
	vuelosPorPrecio        TDAGrafo.GrafoPesado[Codigo, int]     // Ambos Grafos deben ser de vertice: Codigo, peso: Int; para que al hacer el camino minimo podamos hacer:
	vuelosPorTiempo        TDAGrafo.GrafoPesado[Codigo, int]     // desde "X" codigo_de_aeropuerto, hasta "Y" codigo_de_aeropuerto, sino creo que va a ser necesario crear un struct
	caminosEncontrados     TDAPila.Pila[[]Aeropuerto]            // para hacer los caminos minimos, porque CREO que sino no podemos acceder a los vertices de origen y destino, y si creamos los structs los demas campos estaran vacíos.
}

func convertirAInt(cadena string) int {
	numero, _ := strconv.Atoi(cadena)
	return numero
}

func convertirAFloat(cadena string) float64 {
	float, _ := strconv.ParseFloat(cadena, 64)
	return float
}

func CrearAeropuerto(infoAeropuerto string) Aeropuerto {
	informacion := strings.Split(infoAeropuerto, SEPARADOR_COMA)
	latitud, longitud := convertirAFloat(informacion[LATITUD]), convertirAFloat(informacion[LONGITUD])
	return Aeropuerto{
		Ciudad:   informacion[CIUDAD],
		Codigo:   Codigo(informacion[CODIGO]),
		Latitud:  latitud,
		Longitud: longitud,
	}
}

func CrearVuelo(infoDeVuelo string) Vuelo { // Desde mi punto de vista CrearVuelo no deberia ser una primitiva de vuelo, es como el TP1, el vuelo en si mismo no tiene vida/acciones
	informacion := strings.Split(infoDeVuelo, SEPARADOR_COMA)

	tiempo, precio, cant_vuelos := convertirAInt(informacion[TIEMPO]), convertirAInt(informacion[PRECIO]), convertirAInt(informacion[CANT_VUELOS])

	codAeropuertoOrigen, codAeropuertoDestino := Codigo(informacion[ORIGEN]), Codigo(informacion[DESTINO])

	// NO ESTOY SEGURO, pero CREO, que podemos asumir que todos los vuelos en el archivo no se repiten, y si se repite se pisa el peso anterior... CREO.

	return Vuelo{
		AeropuertoOrigen:  codAeropuertoOrigen,
		AeropuertoDestino: codAeropuertoDestino,
		Tiempo:            tiempo,
		Precio:            precio,
		Cant_vuelos:       cant_vuelos,
	}
}

func CrearSistema() SistemaDeAerolineas {
	diccAeropuertos := TDADicc.CrearHash[string, []Codigo]()
	grafoConPrecios := TDAGrafo.CrearGrafoPesado[Codigo, int](false)
	grafoConTiempos := TDAGrafo.CrearGrafoPesado[Codigo, int](false)
	return &sistemaDeAerolineas{
		aeropuertosAlmacenados: diccAeropuertos,
		vuelosPorPrecio:        grafoConPrecios,
		vuelosPorTiempo:        grafoConTiempos,
	}
}

func (sistema *sistemaDeAerolineas) GuardarAeropuerto(aeropuerto Aeropuerto) {
	aeropuertosEnCiudad := sistema.aeropuertosAlmacenados.Obtener(aeropuerto.Ciudad)
	aeropuertosEnCiudad = append(aeropuertosEnCiudad, aeropuerto.Codigo)
	sistema.aeropuertosAlmacenados.Guardar(aeropuerto.Ciudad, aeropuertosEnCiudad)
	sistema.vuelosPorPrecio.AgregarVertice(aeropuerto.Codigo)
	sistema.vuelosPorTiempo.AgregarVertice(aeropuerto.Codigo)
}

func (sistema *sistemaDeAerolineas) GuardarVuelo(vuelo Vuelo) {
	sistema.vuelosPorPrecio.AgregarArista(vuelo.AeropuertoOrigen, vuelo.AeropuertoDestino, vuelo.Precio)
	sistema.vuelosPorPrecio.AgregarArista(vuelo.AeropuertoOrigen, vuelo.AeropuertoDestino, vuelo.Tiempo)
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasBarato(ciudadOrigen, ciudadDestino string) {

}

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasRapido(ciudadOrigen, ciudadDestino string) {

}

func (sistema *sistemaDeAerolineas) ObtenerCaminoConMenosEscalas(ciudadOrigen, ciudadDestino string) {

}

func (sistema *sistemaDeAerolineas) ObtenerAeropuertosMasImportantes(cantidad int) {

}

func (sistema *sistemaDeAerolineas) CrearRutaMinima(rutaSalida string) {

}

func (sistema *sistemaDeAerolineas) CrearItinerario(rutaEntrada string) {

}

func (sistema *sistemaDeAerolineas) ExportarMapaCamino(rutaSalida string) {

}
