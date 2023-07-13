package acciones

import (
	aerolineas "flycombi/sistema_aerolineas"
	funciones "flycombi/validaciones_y_auxiliares"
	TDAHash "tdas/diccionario"
)

const (
	SALIDA_EXITOSA = "OK"
	ESCALAS        = "esc"
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
	camino := sistema.ObtenerCamino(tipo, aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	funciones.MostrarCamino(camino)
}

func CaminoEscalas(sistema aerolineas.SistemaDeAerolineas, origen, destino, _ string) {
	err := funciones.ComprobarEntradaCaminoEscalas(sistema, origen, destino)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	camino := sistema.ObtenerCamino(ESCALAS, aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	funciones.MostrarCamino(camino)
}

func Centralidad(sistema aerolineas.SistemaDeAerolineas, n, _, _ string) {

}

func NuevaAerolinea(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func Itinerario(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {

}

func ExportarKML(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {
	camino := sistema.ObtenerUltimaRutaSolicitada()
	funciones.ExportarUltimoCamino(camino, ruta)
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}
