package acciones

import (
	exportar "flycombi/exportaciones_datos"
	aerolineas "flycombi/sistema_aerolineas"
	funciones "flycombi/validaciones_y_auxiliares"
	TDAHash "tdas/diccionario"
)

const (
	SEPARADOR_FLECHA = " -> "
	SEPARADOR_COMA   = ", "
	SALIDA_EXITOSA   = "OK"
	ESCALAS          = "esc"
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
	funciones.MostrarSalida(camino, SEPARADOR_FLECHA)
}

func CaminoEscalas(sistema aerolineas.SistemaDeAerolineas, origen, destino, _ string) {
	err := funciones.ComprobarEntradaCaminoEscalas(sistema, origen, destino)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	camino := sistema.ObtenerCamino(ESCALAS, aerolineas.Ciudad(origen), aerolineas.Ciudad(destino))
	funciones.MostrarSalida(camino, SEPARADOR_FLECHA)
}

func Centralidad(sistema aerolineas.SistemaDeAerolineas, n, _, _ string) {
	numero, err := funciones.ComprobarEntradaCentralidad(n)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	aeropuertoCentrales := sistema.ObtenerAeropuertosMasImportantes()
	masCentrales := funciones.ObtenerMasCentrales(numero, aeropuertoCentrales)
	funciones.MostrarSalida(masCentrales, SEPARADOR_COMA)
}

func NuevaAerolinea(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {
	vuelosOptimos := sistema.ObtenerVuelosRutaMinima()
	exportar.EscribirArchivoCSV(ruta, vuelosOptimos)
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func Itinerario(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {
	ciudades, rutas := funciones.ObtenerCiudadesYRutas(ruta)
	ordenTopo, caminos := sistema.ObtenerCaminosItinerario(ciudades, rutas)
	ordenTopoStr := make([]string, len(ordenTopo))
	for i, ciudad := range ordenTopo {
		ordenTopoStr[i] = string(ciudad)
	}
	funciones.MostrarMensaje(funciones.CrearMensaje(ordenTopoStr, SEPARADOR_COMA))
	for _, camino := range caminos {
		funciones.MostrarSalida(camino, SEPARADOR_FLECHA)
	}
}

func ExportarKML(sistema aerolineas.SistemaDeAerolineas, ruta, _, _ string) {
	camino := sistema.ObtenerUltimaRutaSolicitada()
	exportar.EscribirArchivoKML(ruta, camino)
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}
