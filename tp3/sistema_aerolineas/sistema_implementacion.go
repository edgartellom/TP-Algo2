package sistema

import (
	"math"
	"strconv"
	"strings"

	BiblioGrafo "tdas/biblioteca_grafos"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

const (
	SEPARADOR_COMA    = ","
	SEPARADOR_ESPACIO = " "
	COMPARADOR        = 0
)

type sistemaDeAerolineas struct {
	aeropuertosAlmacenados TDADicc.Diccionario[Codigo, Aeropuerto]
	aeropuertosPorCiudad   TDADicc.Diccionario[Ciudad, []Aeropuerto]
	vuelosPorPrecio        TDAGrafo.GrafoPesado[Aeropuerto, int]
	vuelosPorTiempo        TDAGrafo.GrafoPesado[Aeropuerto, int]
	vuelos                 TDAGrafo.GrafoNoPesado[Aeropuerto, int]
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
		Ciudad:   Ciudad(informacion[CIUDAD]),
		Codigo:   Codigo(informacion[CODIGO]),
		Latitud:  latitud,
		Longitud: longitud,
	}
}

func CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, SEPARADOR_COMA)

	codAeropuertoOrigen, codAeropuertoDestino := Codigo(informacion[ORIGEN]), Codigo(informacion[DESTINO])

	tiempo, precio, cant_vuelos := convertirAInt(informacion[TIEMPO]), convertirAInt(informacion[PRECIO]), convertirAInt(informacion[CANT_VUELOS])

	return Vuelo{
		AeropuertoOrigen:  codAeropuertoOrigen,
		AeropuertoDestino: codAeropuertoDestino,
		Tiempo:            tiempo,
		Precio:            precio,
		Cant_vuelos:       cant_vuelos,
	}
}

func CrearSistema() SistemaDeAerolineas {
	diccAeropuertos := TDADicc.CrearHash[Codigo, Aeropuerto]()
	diccAeropuertosPorCiudad := TDADicc.CrearHash[Ciudad, []Aeropuerto]()
	grafoConPrecios := TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	grafoConTiempos := TDAGrafo.CrearGrafoPesado[Aeropuerto, int](false)
	pilaCaminos := TDAPila.CrearPilaDinamica[[]Aeropuerto]()
	return &sistemaDeAerolineas{
		aeropuertosAlmacenados: diccAeropuertos,
		aeropuertosPorCiudad:   diccAeropuertosPorCiudad,
		vuelosPorPrecio:        grafoConPrecios,
		vuelosPorTiempo:        grafoConTiempos,
		caminosEncontrados:     pilaCaminos,
	}
}

func (sistema *sistemaDeAerolineas) GuardarAeropuerto(aeropuerto Aeropuerto) {
	sistema.aeropuertosAlmacenados.Guardar(aeropuerto.Codigo, aeropuerto)

	if !sistema.aeropuertosPorCiudad.Pertenece(aeropuerto.Ciudad) {
		sistema.aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, []Aeropuerto{})
	}
	aeropuertos := sistema.aeropuertosPorCiudad.Obtener(aeropuerto.Ciudad)
	aeropuertos = append(aeropuertos, aeropuerto)
	sistema.aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, aeropuertos)

	sistema.vuelosPorPrecio.AgregarVertice(aeropuerto)
	sistema.vuelosPorTiempo.AgregarVertice(aeropuerto)
}

func (sistema *sistemaDeAerolineas) GuardarVuelo(vuelo Vuelo) {
	if sistema.aeropuertosAlmacenados.Pertenece(vuelo.AeropuertoOrigen) && sistema.aeropuertosAlmacenados.Pertenece(vuelo.AeropuertoDestino) {
		aeropuertoOrigen := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoOrigen)
		aeropuertoDestino := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoDestino)
		sistema.vuelosPorPrecio.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Precio)
		sistema.vuelosPorPrecio.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Tiempo)
	}
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasBarato(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minPrecio := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		distancias, padres := BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorPrecio, aeropuertoOrigen)

		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				suma := BiblioGrafo.SumarDistancias[Aeropuerto](distancias)
				resul.Guardar(suma, Result{padres, aeropuertoDestino})
				if suma < minPrecio {
					minPrecio = suma
				}
			}
		}
	}
	minimo := resul.Obtener(minPrecio)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minimo.padres, minimo.aeropuertoDestino)
	return camino
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoMasRapido(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minTiempo := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		distancias, padres := BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorTiempo, aeropuertoOrigen)

		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				suma := BiblioGrafo.SumarDistancias[Aeropuerto](distancias)
				resul.Guardar(suma, Result{padres, aeropuertoDestino})
				if suma < minTiempo {
					minTiempo = suma
				}
			}
		}
	}
	minimo := resul.Obtener(minTiempo)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minimo.padres, minimo.aeropuertoDestino)
	return camino
}

func (sistema *sistemaDeAerolineas) ObtenerCaminoConMenosEscalas(ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	minOrden := math.MaxInt64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoDestino Aeropuerto
	}
	resul := TDADicc.CrearHash[int, Result]()
	for _, aeropuertoOrigen := range aeropuertosOrigen {
		padres, orden := BiblioGrafo.RecorridoDfs[Aeropuerto](sistema.vuelos, aeropuertoOrigen)
		for _, aeropuertoDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDestino) {
				resul.Guardar(orden.Obtener(aeropuertoDestino), Result{padres, aeropuertoDestino})
				if orden.Obtener(aeropuertoDestino) < minOrden {
					minOrden = orden.Obtener(aeropuertoDestino)
				}
			}
		}
	}
	minino := resul.Obtener(minOrden)
	camino := BiblioGrafo.ReconstruirCamino[Aeropuerto](minino.padres, minino.aeropuertoDestino)
	return camino
}

func (sistema *sistemaDeAerolineas) ObtenerAeropuertosMasImportantes(cantidad int) {

}

func (sistema *sistemaDeAerolineas) CrearRutaMinima(rutaSalida string) {

}

func (sistema *sistemaDeAerolineas) CrearItinerario(rutaEntrada string) {

}

func (sistema *sistemaDeAerolineas) ExportarMapaCamino(rutaSalida string) {

}
