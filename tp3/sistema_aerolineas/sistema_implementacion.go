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
	aeropuertosAlmacenados TDADicc.Diccionario[Codigo, Aeropuerto]
	vuelosPorPrecio        TDAGrafo.GrafoPesado[Aeropuerto, int]
	vuelosPorTiempo        TDAGrafo.GrafoPesado[Aeropuerto, int]
	caminosEncontrados     TDAPila.Pila[[]Aeropuerto]
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

func (sistema *sistemaDeAerolineas) CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, SEPARADOR_COMA)

	tiempo, precio, cant_vuelos := convertirAInt(informacion[TIEMPO]), convertirAInt(informacion[PRECIO]), convertirAInt(informacion[CANT_VUELOS])

	codAeropuertoOrigen, codAeropuertoDestino := Codigo(informacion[ORIGEN]), Codigo(informacion[DESTINO])

	var aeropuertoOrigen Aeropuerto
	var aeropuertoDestino Aeropuerto

	if (*sistema).aeropuertosAlmacenados.Pertenece(codAeropuertoOrigen) {
		aeropuertoOrigen = (*sistema).aeropuertosAlmacenados.Obtener(codAeropuertoOrigen)
	} else {
		aeropuertoOrigen = Aeropuerto{Codigo: codAeropuertoOrigen}
	}

	if (*sistema).aeropuertosAlmacenados.Pertenece(codAeropuertoDestino) {
		aeropuertoDestino = (*sistema).aeropuertosAlmacenados.Obtener(codAeropuertoDestino)
	} else {
		aeropuertoDestino = Aeropuerto{Codigo: codAeropuertoDestino}
	}

	return Vuelo{
		AeropuertoOrigen:  aeropuertoOrigen,
		AeropuertoDestino: aeropuertoDestino,
		Tiempo:            tiempo,
		Precio:            precio,
		Cant_vuelos:       cant_vuelos,
	}
}

func CrearSistema() SistemaDeAerolineas {
	diccAeropuertos := TDADicc.CrearHash[Codigo, Aeropuerto]()
	grafoConPrecios := TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	grafoConTiempos := TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	return &sistemaDeAerolineas{
		aeropuertosAlmacenados: diccAeropuertos,
		vuelosPorPrecio:        grafoConPrecios,
		vuelosPorTiempo:        grafoConTiempos,
	}
}

func (sistema *sistemaDeAerolineas) GuardarAeropuerto(aeropuerto Aeropuerto) {
	sistema.aeropuertosAlmacenados.Guardar(aeropuerto.Codigo, aeropuerto)
	sistema.vuelosPorPrecio.AgregarVertice(aeropuerto)
	sistema.vuelosPorTiempo.AgregarVertice(aeropuerto)
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

func (sistema *sistemaDeAerolineas) ExportarMapaCamino(rutaSalida string)
