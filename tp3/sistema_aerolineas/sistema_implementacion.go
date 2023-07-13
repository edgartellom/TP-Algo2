package sistema

import (
	"math"
	"strconv"
	"strings"

	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"

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
	aeropuertos            TDADicc.Diccionario[Codigo, Ciudad]
	vuelosPorPrecio        TDAGrafo.GrafoPesado[Aeropuerto, float64]
	vuelosPorTiempo        TDAGrafo.GrafoPesado[Aeropuerto, float64]
	vuelosPorFrecuencia    TDAGrafo.GrafoPesado[Aeropuerto, float64]
	ultimoComando          []Aeropuerto
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
	vuelo.Tiempo, vuelo.Precio, vuelo.Cant_vuelos = convertirAFloat(informacion[TIEMPO]), convertirAFloat(informacion[PRECIO]), convertirAFloat(informacion[CANT_VUELOS])

	return *vuelo
}

func CrearSistema() SistemaDeAerolineas {
	sistema := new(sistemaDeAerolineas)
	sistema.aeropuertosPorCiudad = TDADicc.CrearHash[Ciudad, []Aeropuerto]()
	sistema.aeropuertosAlmacenados = TDADicc.CrearHash[Codigo, Aeropuerto]()
	sistema.aeropuertos = TDADicc.CrearHash[Codigo, Ciudad]()
	sistema.vuelosPorPrecio = TDAGrafo.CrearGrafoPesado[Aeropuerto, float64](false)
	sistema.vuelosPorTiempo = TDAGrafo.CrearGrafoPesado[Aeropuerto, float64](false)
	sistema.vuelosPorFrecuencia = TDAGrafo.CrearGrafoPesado[Aeropuerto, float64](false)
	return sistema
}

func (sistema *sistemaDeAerolineas) GuardarAeropuerto(aeropuerto Aeropuerto) {
	if !(*sistema).aeropuertosPorCiudad.Pertenece(aeropuerto.Ciudad) {
		(*sistema).aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, []Aeropuerto{})
	}
	aeropuertosEnCiudad := sistema.aeropuertosPorCiudad.Obtener(aeropuerto.Ciudad)
	(*sistema).aeropuertosPorCiudad.Guardar(aeropuerto.Ciudad, append(aeropuertosEnCiudad, aeropuerto))
	(*sistema).aeropuertosAlmacenados.Guardar(aeropuerto.Codigo, aeropuerto)
	(*sistema).aeropuertos.Guardar(aeropuerto.Codigo, aeropuerto.Ciudad)
	(*sistema).vuelosPorPrecio.AgregarVertice(aeropuerto)
	(*sistema).vuelosPorTiempo.AgregarVertice(aeropuerto)
	(*sistema).vuelosPorFrecuencia.AgregarVertice(aeropuerto)
}

func (sistema *sistemaDeAerolineas) GuardarVuelo(vuelo Vuelo) {
	aeropuertoOrigen := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoOrigen)
	aeropuertoDestino := sistema.aeropuertosAlmacenados.Obtener(vuelo.AeropuertoDestino)
	(*sistema).vuelosPorPrecio.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Precio)
	(*sistema).vuelosPorTiempo.AgregarArista(aeropuertoOrigen, aeropuertoDestino, vuelo.Tiempo)
	(*sistema).vuelosPorFrecuencia.AgregarArista(aeropuertoOrigen, aeropuertoDestino, 1/vuelo.Cant_vuelos)
}

func (sistema *sistemaDeAerolineas) ObtenerCamino(tipo string, ciudadOrigen, ciudadDestino Ciudad) []Aeropuerto {
	aeropuertosOrigen := sistema.aeropuertosPorCiudad.Obtener(ciudadOrigen)
	aeropuertosDestino := sistema.aeropuertosPorCiudad.Obtener(ciudadDestino)
	tiempoMinimoActual := math.MaxFloat64
	type Result struct {
		padres            TDADicc.Diccionario[Aeropuerto, *Aeropuerto]
		aeropuertoDestino Aeropuerto
	}
	var res *Result
	for _, aeropuertoDeOrigen := range aeropuertosOrigen {
		padres, distancias := sistema.usarCaminoSegunTipo(tipo, aeropuertoDeOrigen)
		for _, aeropuertoDeDestino := range aeropuertosDestino {
			if padres.Pertenece(aeropuertoDeDestino) {
				if distancias.Obtener(aeropuertoDeDestino) < tiempoMinimoActual {
					res = &Result{padres, aeropuertoDeDestino}
					tiempoMinimoActual = distancias.Obtener(aeropuertoDeDestino)
				}
			}
		}
	}
	var camino []Aeropuerto
	if res != nil && (*res).padres.Cantidad() != 0 {
		camino = BiblioGrafo.ReconstruirCamino[Aeropuerto]((*res).padres, (*res).aeropuertoDestino)
	}
	sistema.ultimoComando = camino
	return camino
}

func (sistema *sistemaDeAerolineas) usarCaminoSegunTipo(tipo string, origen Aeropuerto) (TDADicc.Diccionario[Aeropuerto, *Aeropuerto], TDADicc.Diccionario[Aeropuerto, float64]) {
	if tipo == TIPO_BARATO {
		return BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorPrecio, origen)
	} else if tipo == TIPO_RAPIDO {
		return BiblioGrafo.CaminoMinimoDijkstra[Aeropuerto](sistema.vuelosPorTiempo, origen)
	}
	return BiblioGrafo.CaminoMinimoBFS[Aeropuerto](sistema.vuelosPorFrecuencia, origen)
}

func (sistema sistemaDeAerolineas) Pertenece(ciudad Ciudad) bool {
	return sistema.aeropuertosPorCiudad.Pertenece(ciudad)
}

func (sistema *sistemaDeAerolineas) ObtenerAeropuertosMasImportantes(cantidad int) {

}

func (sistema *sistemaDeAerolineas) ObtenerVuelosMST() []Vuelo {
	var vuelos []Vuelo
	arbol := BiblioGrafo.MstPrim[Aeropuerto, float64](sistema.vuelosPorPrecio)
	for _, aeropuerto := range arbol.ObtenerVertices() {
		for _, adyacente := range arbol.ObtenerAdyacentes(aeropuerto) {
			aeropuertoOrigen := adyacente.Codigo
			aeropuertoDeDestino := aeropuerto.Codigo
			tiempo := sistema.vuelosPorTiempo.VerPeso(adyacente, aeropuerto)
			precio := sistema.vuelosPorPrecio.VerPeso(adyacente, aeropuerto)
			cantVuelos := math.Round(1 / sistema.vuelosPorFrecuencia.VerPeso(adyacente, aeropuerto))
			vuelos = append(vuelos, Vuelo{AeropuertoOrigen: aeropuertoOrigen, AeropuertoDestino: aeropuertoDeDestino, Tiempo: tiempo, Precio: precio, Cant_vuelos: cantVuelos})
		}
	}
	return vuelos
}

func (sistema *sistemaDeAerolineas) ObtenerCaminosItinerario(ciudades []Ciudad, rutas []Ruta) ([]Ciudad, [][]Aeropuerto) {
	grafo := TDAGrafo.CrearGrafoNoPesado[Ciudad, float64](true)
	for _, ciudad := range ciudades {
		grafo.AgregarVertice(ciudad)
	}
	for _, ruta := range rutas {
		grafo.AgregarArista(ruta.CiudadOrigen, ruta.CiudadDestino)
	}
	ordenTopo := BiblioGrafo.TopologicoGrados[Ciudad, float64](grafo)
	var caminos [][]Aeropuerto
	for i := 1; i < len(ordenTopo); i++ {
		camino := sistema.ObtenerCamino("TIPO_RAPIDO", ordenTopo[i-1], ordenTopo[i])
		caminos = append(caminos, camino)
	}

	return ordenTopo, caminos
}

func (sistema *sistemaDeAerolineas) ObtenerUltimaRutaSolicitada() []Aeropuerto {
	return sistema.ultimoComando
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
