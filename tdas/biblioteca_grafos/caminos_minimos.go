package biblioteca_grafos

import (
	"math"
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func CaminoMinimoBFS[K comparable](grafo TDAGrafo.Grafo[K, int], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, math.MaxInt64)
	}
	visitados.Guardar(origen, true)
	padres.Guardar(origen, nil)
	distancias.Guardar(origen, 0)
	cola := TDACola.CrearColaEnlazada[K]()
	cola.Encolar(origen)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padres.Guardar(w, &v)
				distancias.Guardar(w, distancias.Obtener(v)+1)
				cola.Encolar(w)
			}
		}
	}
	return padres, distancias
}

type aristaDijkstra[K comparable, V any] struct {
	vertice           K
	distanciaAlOrigen V
}

func cmpDijkstra[K comparable](a, b aristaDijkstra[K, int]) int {
	return b.distanciaAlOrigen - a.distanciaAlOrigen
}

func cmpDijkstraFloat[K comparable](a, b aristaDijkstra[K, float64]) int {
	return int(b.distanciaAlOrigen - a.distanciaAlOrigen)
}

func CaminoMinimoDijkstra[K comparable](grafo TDAGrafo.GrafoPesado[K, int], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	heap := TDAHeap.CrearHeap[aristaDijkstra[K, int]](cmpDijkstra[K])
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, math.MaxInt64)
	}
	padres.Guardar(origen, nil)
	distancias.Guardar(origen, 0)
	heap.Encolar(aristaDijkstra[K, int]{origen, 0})
	for !heap.EstaVacia() {
		v := heap.Desencolar().vertice
		for _, w := range grafo.ObtenerAdyacentes(v) {
			distanciaActual := distancias.Obtener(v) + grafo.VerPeso(v, w)
			if distanciaActual < distancias.Obtener(w) {
				distancias.Guardar(w, distanciaActual)
				padres.Guardar(w, &v)
				heap.Encolar(aristaDijkstra[K, int]{w, distanciaActual})
			}
		}
	}
	return padres, distancias
}

func CaminoMinimoDijkstraFloat[K comparable](grafo TDAGrafo.GrafoPesado[K, float64], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, float64]) {
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, float64]()
	heap := TDAHeap.CrearHeap[aristaDijkstra[K, float64]](cmpDijkstraFloat[K])
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, math.MaxFloat64)
	}
	padres.Guardar(origen, nil)
	distancias.Guardar(origen, 0)
	heap.Encolar(aristaDijkstra[K, float64]{origen, 0})
	for !heap.EstaVacia() {
		v := heap.Desencolar().vertice
		for _, w := range grafo.ObtenerAdyacentes(v) {
			distanciaActual := distancias.Obtener(v) + grafo.VerPeso(v, w)
			if distanciaActual < distancias.Obtener(w) {
				distancias.Guardar(w, distanciaActual)
				padres.Guardar(w, &v)
				heap.Encolar(aristaDijkstra[K, float64]{w, distanciaActual})
			}
		}
	}
	return padres, distancias
}

func CaminoMinimoBellmanFord[K comparable, V int](grafo TDAGrafo.GrafoPesado[K, int], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	aristas := ObtenerAristas(grafo)
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, math.MaxInt64)
	}
	distancias.Guardar(origen, 0)
	padres.Guardar(origen, nil)
	for i := 0; i < grafo.Cantidad(); i++ {
		var cambio bool
		for _, arista := range aristas {
			origen, destino, peso := arista.vertice, arista.adyacente, arista.peso
			distanciaActual := distancias.Obtener(origen) + peso
			if distanciaActual < distancias.Obtener(destino) {
				cambio = true
				padres.Guardar(destino, &origen)
				distancias.Guardar(destino, distanciaActual)
			}
		}
		if !cambio {
			return padres, distancias
		}
	}
	for _, arista := range aristas {
		origen, destino, peso := arista.vertice, arista.adyacente, arista.peso
		if distancias.Obtener(origen)+peso < distancias.Obtener(destino) {
			panic("Hay un ciclo negativo")
		}
	}
	return padres, distancias
}
