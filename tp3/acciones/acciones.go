package acciones

import (
	aerolineas "flycombi/sistema_aerolineas"
	funciones "flycombi/validaciones_y_auxiliares"
	TDAHash "tdas/diccionario"
)

type comandos func(*aerolineas.SistemaDeAerolineas, string, string, string)

const CANT_COMANDOS = 6

var ACCIONES = [CANT_COMANDOS]comandos{CaminoMas, CaminoEscalas, Centralidad, NuevaAerolinea, Itinerario, ExportarKML}
var COMANDOS = [CANT_COMANDOS]string{"camino_mas", "camino_escalas", "centralidad", "nueva_aerolinea", "itinerario", "exportar_kml"}

func CrearBaseDeDatos() aerolineas.SistemaDeAerolineas {
	return aerolineas.CrearSistema()
}

func CrearOpciones() TDAHash.Diccionario[string, comandos] {
	opciones := TDAHash.CrearHash[string, comandos]()
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

func CaminoMas(sistema *aerolineas.SistemaDeAerolineas, tipo, origen, destino string) {
	var aeropuertos []aerolineas.Aeropuerto
	if tipo == "barato" {
		aeropuertos = (*sistema).ObtenerCaminoMasBarato(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	}
	if tipo == "rapido" {
		aeropuertos = (*sistema).ObtenerCaminoMasRapido(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	}
	funciones.ImprimirCamino(aeropuertos)
}

func CaminoEscalas(sistema *aerolineas.SistemaDeAerolineas, origen, destino, _ string) {
	aeropuertos := (*sistema).ObtenerCaminoConMenosEscalas(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	funciones.ImprimirCamino(aeropuertos)
}

func Centralidad(sistema *aerolineas.SistemaDeAerolineas, n, _, _ string) {

}

func NuevaAerolinea(sistema *aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func Itinerario(sistema *aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func ExportarKML(sistema *aerolineas.SistemaDeAerolineas, archivo, _, _ string) {

}
