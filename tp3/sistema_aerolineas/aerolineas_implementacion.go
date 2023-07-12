package sistema

import (
	"math"
	"strconv"
	"strings"

	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"

	// TDAPila "tdas/pila"
	BiblioGrafo "tdas/biblioteca_grafos"
)

const (
	TIPO_BARATO       = "barato"
	TIPO_RAPIDO       = "rapido"
	SEPARADOR_COMA    = ","
	SEPARADOR_ESPACIO = " "
)

type sistemaDeAerolineas struct {
	aeropuertosPorCiudad   TDADicc.Diccionario[Ciudad, []Aeropuerto]
	aeropuertosAlmacenados TDADicc.Diccionario[Codigo, Aeropuerto]
	aeropuertos            TDADicc.Diccionario[Codigo, Ciudad] //dudoso
	vuelosPorPrecio        TDAGrafo.GrafoPesado[Aeropuerto, int]
	vuelosPorTiempo        TDAGrafo.GrafoPesado[Aeropuerto, int]
	vuelos                 TDAGrafo.GrafoNoPesado[Aeropuerto, int]
	// caminosEncontrados     TDAPila.Pila[[]Aeropuerto]
}

func CrearAeropuerto(infoAeropuerto string) Aeropuerto {
	informacion := strings.Split(infoAeropuerto, SEPARADOR_COMA)
	latitud, longitud := convertirAFloat(informacion[LATITUD]), convertirAFloat(informacion[LONGITUD])
	return Aeropuerto{
		Ciudad:   Ciudad(informacion[CIUDAD]),
		Codigo:   Codigo(informacion[CODIGO]),
		Latitud:  latitud,
		Longitud: longitud,
	}
}

func CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, SEPARADOR_COMA)

	vuelo := new(Vuelo)
	vuelo.AeropuertoOrigen, vuelo.AeropuertoDestino = Codigo(informacion[ORIGEN]), Codigo(informacion[DESTINO])
	vuelo.Tiempo, vuelo.Precio, vuelo.Cant_vuelos = convertirAInt(informacion[TIEMPO]), convertirAInt(informacion[PRECIO]), convertirAInt(informacion[CANT_VUELOS])

	return *vuelo
}

func CrearSistema() SistemaDeAerolineas {
	sistema := new(sistemaDeAerolineas)
	sistema.aeropuertosPorCiudad = TDADicc.CrearHash[Ciudad, []Aeropuerto]()
	sistema.aeropuertosAlmacenados = TDADicc.CrearHash[Codigo, Aeropuerto]()
	sistema.aeropuertos = TDADicc.CrearHash[Codigo, Ciudad]() //dudoso
	sistema.vuelosPorPrecio = TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	sistema.vuelosPorTiempo = TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	sistema.vuelos = TDAGrafo.CrearGrafoNoPesado[Aeropuerto, int](false)
	return sistema
}

func (sistema *sistemaDeAerolineas) GuardarAeropuerto(aeropuerto Aeropuerto) {
	if !(*sistema).aeropuertosPorCiudad.Pertenece(aeropuerto.Ciudad) {
		(*sistema).aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, []Aeropuerto{})
	}
	aeropuertosEnCiudad := sistema.aeropuertosPorCiudad.Obtener(aeropuerto.Ciudad)
	(*sistema).aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, append(aeropuertosEnCiudad, aeropuerto))
	(*sistema).aeropuertosAlmacenados.Guardar(aeropuerto.Codigo, aeropuerto)
	(*sistema).aeropuertos.Guardar(aeropuerto.Codigo, aeropuerto.Ciudad) //dudoso
	(*sistema).vuelosPorPrecio.AgregarVertice(aeropuerto)
	(*sistema).vuelosPorTiempo.AgregarVertice(aeropuerto)
	(*sistema).vuelos.AgregarVertice(aeropuerto)
}

func (sistema *sistemaDeAerolineas) GuardarVuelo(vuelo Vuelo) {
	aeropuertoOrigen := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoOrigen)
	aeropuertoDestino := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoDestino)
	(*sistema).vuelosPorPrecio.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Precio)
	(*sistema).vuelosPorTiempo.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Tiempo)
	(*sistema).vuelos.AgregarArista(aeropuertoOrigen, aeropuertoDestino)
}

// func (sistema *sistemaDeAerolineas) ObtenerCaminoMasBarato(ciudadOrigen, ciudadDestino string) {

// }

// func (sistema *sistemaDeAerolineas) ObtenerCaminoMasRapido(ciudadOrigen, ciudadDestino string) {

// }

// func (sistema *sistemaDeAerolineas) ObtenerCaminoConMenosEscalas(ciudadOrigen, ciudadDestino string) {

// }

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasBarato(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minPrecio := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoOrigen  Aeropuerto
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		padres, distancias := BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorPrecio, aeropuertoOrigen)
		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				suma := BiblioGrafo.SumarDistancias[Aeropuerto](distancias)
				resul.Guardar(suma, Result{padres, aeropuertoOrigen, aeropuertoDestino})
				if suma < minPrecio {
					minPrecio = suma
				}
			}
		}
	}
	minimo := resul.Obtener(minPrecio)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minimo.padres, minimo.aeropuertoOrigen, minimo.aeropuertoDestino)
	return camino
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasRapido(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minTiempo := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoOrigen  Aeropuerto
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		padres, distancias := BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorTiempo, aeropuertoOrigen)

		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				suma := BiblioGrafo.SumarDistancias[Aeropuerto](distancias)
				resul.Guardar(suma, Result{padres, aeropuertoOrigen, aeropuertoDestino})
				if suma < minTiempo {
					minTiempo = suma
				}
			}
		}
	}
	minimo := resul.Obtener(minTiempo)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minimo.padres, minimo.aeropuertoOrigen, minimo.aeropuertoDestino)
	return camino
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoConMenosEscalas(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minOrden := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoOrigen  Aeropuerto
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		padres, orden := BiblioGrafo.CaminoMinimoBFS[Aeropuerto](sistema.vuelos, aeropuertoOrigen)
		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				resul.Guardar(orden.Obtener(aeropuertoDestino), Result{padres, aeropuertoOrigen, aeropuertoDestino})
				if orden.Obtener(aeropuertoDestino) < minOrden {
					minOrden = orden.Obtener(aeropuertoDestino)
				}
			}
		}
	}
	minimo := resul.Obtener(minOrden)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minimo.padres, minimo.aeropuertoOrigen, minimo.aeropuertoDestino)
	return camino
}

func (sistema sistemaDeAerolineas) Pertenece(ciudad Ciudad) bool {
	return sistema.aeropuertosPorCiudad.Pertenece(ciudad)
}

// func (sistema *sistemaDeAerolineas) ObtenerCamino(tipo string, ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
// 	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
// 	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
// 	tiempoMinimoActual := math.MaxInt64
// 	type Result struct {
// 		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
// 		aeropuertoOrigen  Aeropuerto
// 		aeropuertoDestino Aeropuerto
// 	}
// 	var res *Result
// 	for _, aeropuertoDeOrigen := range aeropuertosOrigen {
// 		padres, distancias := sistema.usarCaminoSegunTipo(tipo, aeropuertoDeOrigen)
// 		for _, aeropuertoDeDestino := range aeropuertosDestino {
// 			if padres.Pertenece(aeropuertoDeDestino) {
// 				suma := BiblioGrafo.SumarDistancias[Aeropuerto](distancias)
// 				if suma < tiempoMinimoActual {
// 					res = &Result{padres, aeropuertoDeOrigen, aeropuertoDeDestino}
// 				}
// 			}
// 		}
// 	}
// 	var camino []Aeropuerto
// 	if res != nil && (*res).padres.Cantidad() != 0 {
// 		camino = BiblioGrafo.ReconstruirCamino[Aeropuerto]((*res).padres, (*res).aeropuertoOrigen, (*res).aeropuertoDestino)
// 	}
// 	return camino
// }

// func (sistema *sistemaDeAerolineas) usarCaminoSegunTipo(tipo string, origen Aeropuerto) (TDADicc.Diccionario[Aeropuerto, *Aeropuerto], TDADicc.Diccionario[Aeropuerto, int]) {
// 	if tipo == TIPO_BARATO {
// 		return BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorPrecio, origen)
// 	} else if tipo == TIPO_RAPIDO {
// 		return BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorTiempo, origen)
// 	}
// 	return BiblioGrafo.CaminoMinimoBFS[Aeropuerto](sistema.vuelos, origen)
// }

func (sistema *sistemaDeAerolineas) ObtenerAeropuertosMasImportantes(cantidad int) {

}

func (sistema *sistemaDeAerolineas) CrearRutaMinima(rutaSalida string) {

}

func (sistema *sistemaDeAerolineas) CrearItinerario(rutaEntrada string) {

}

func (sistema *sistemaDeAerolineas) ExportarMapaCamino(rutaSalida string) {

}

/* -------------------------------------------------- FUNCIONES AUXILIARES -------------------------------------------------- */

func convertirAInt(cadena string) int {
	numero, _ := strconv.Atoi(cadena)
	return numero
}

func convertirAFloat(cadena string) float64 {
	float, _ := strconv.ParseFloat(cadena, 64)
	return float
}
