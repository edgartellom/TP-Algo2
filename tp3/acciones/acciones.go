package acciones

import (
	aerolineas "flycombi/sistema_aerolineas"
	funciones "flycombi/validaciones_y_auxiliares"
	TDAHash "tdas/diccionario"
)

type comando func(aerolineas.SistemaDeAerolineas, string, string, string)

var ACCIONES = [funciones.CANT_COMANDOS]comando{CaminoMas, CaminoEscalas, Centralidad, NuevaAerolinea, Itinerario, ExportarKML}

func CrearBaseDeDatos() aerolineas.SistemaDeAerolineas {
	return aerolineas.CrearSistema()
}

func CrearOpciones() TDAHash.Diccionario[string, comando] {
	opciones := TDAHash.CrearHash[string, comando]()
	for i := 0; i < funciones.CANT_COMANDOS; i++ {
		opciones.Guardar(funciones.COMANDOS[i], ACCIONES[i])
	}
	return opciones
}

func GuardarInformacion(sistema aerolineas.SistemaDeAerolineas, rutaAeropuertos, rutaVuelos string) {
	aeropuertos := funciones.ObtenerAeropuertos(rutaAeropuertos)
	vuelos := funciones.ObtenerVuelos(rutaVuelos)

	for _, aeropuerto := range aeropuertos {
		sistema.GuardarAeropuerto(aeropuerto)
	}
	for _, vuelo := range vuelos {
		sistema.GuardarVuelo(vuelo)
	}
}

func CaminoMas(sistema aerolineas.SistemaDeAerolineas, tipo, origen, destino string) {
	err := funciones.ComprobarEntradaCaminoMas(sistema, tipo, origen, destino)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	var camino []aerolineas.Aeropuerto
	if tipo == "barato" {
		camino = sistema.ObtenerCaminoMasBarato(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	} else {
		camino = sistema.ObtenerCaminoMasRapido(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	}
	funciones.MostrarAeropuertos(camino, " -> ")
}

func CaminoEscalas(sistema aerolineas.SistemaDeAerolineas, origen, destino, _ string) {
	err := funciones.ComprobarEntradaCaminoEscalas(sistema, origen, destino)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	camino := sistema.ObtenerCaminoConMenosEscalas(aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	funciones.MostrarAeropuertos(camino, " -> ")
}

func Centralidad(sistema aerolineas.SistemaDeAerolineas, n, _, _ string) {
	err := funciones.ComprobarEntradaCentralidad(sistema, n)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	cantidad := aerolineas.ConvertirAInt(n)
	aeropuertos := sistema.ObtenerAeropuertosMasImportantes(cantidad)
	funciones.MostrarAeropuertos(aeropuertos, ",")
}

func NuevaAerolinea(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func Itinerario(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func ExportarKML(sistema aerolineas.SistemaDeAerolineas, archivo, _, _ string) {

}
