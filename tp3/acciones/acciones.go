package acciones

import (
	aerolineas "flycombi/sistema_aerolineas"
	funciones "flycombi/validaciones_y_auxiliares"
	TDAHash "tdas/diccionario"
)

type comando func(string, string, string)

const CANT_COMANDOS = 6

var ACCIONES = [CANT_COMANDOS]comando{CaminoMas, CaminoEscalas, Centralidad, NuevaAerolinea, Itinerario, ExportarKML}
var COMANDOS = [CANT_COMANDOS]string{"camino_mas", "camino_escalas", "centralidad", "nueva_aerolinea", "itinerario", "exportar_kml"}

func CrearBaseDeDatos() aerolineas.SistemaDeAerolineas {
	return aerolineas.CrearSistema()
}

func CrearOpciones() TDAHash.Diccionario[string, comando] {
	opciones := TDAHash.CrearHash[string, comando]()
	for i := 0; i < CANT_COMANDOS; i++ {
		opciones.Guardar(COMANDOS[i], ACCIONES[i])
	}
	return opciones
}

func GuardarInformacion(sistema aerolineas.SistemaDeAerolineas, rutaAeropuertos, rutaVuelos string) {
	aeropuertos := funciones.ObtenerAeropuertos(rutaAeropuertos)
	vuelos := funciones.ObtenerVuelos(rutaAeropuertos)

	for _, aeropuerto := range aeropuertos {
		sistema.GuardarAeropuerto(aeropuerto)
	}
	for _, vuelo := range vuelos {
		sistema.GuardarVuelo(vuelo)
	}
}

func CaminoMas(tipo, origen, destino string) {

}

func CaminoEscalas(origen, destino, _ string) {

}

func Centralidad(n, _, _ string) {

}

func NuevaAerolinea(ruta, _, _ string) {

}

func Itinerario(ruta, _, _ string) {

}

func ExportarKML(archivo, _, _ string) {

}
