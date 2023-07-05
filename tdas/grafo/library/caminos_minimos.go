package grafo

import (
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func CaminoMinimoBfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, int], TDADicc.Diccionario[K, *K]) {
	distancia := TDADicc.CrearHash[K, int]()
	padre := TDADicc.CrearHash[K, *K]()
	visitado := TDADicc.CrearHash[K, bool]()

	for _, v := range grafo.ObtenerVertices() {
		distancia.Guardar(v, 99999999999)
	}

	distancia.Guardar(origen, 0)
	padre.Guardar(origen, nil)
	visitado.Guardar(origen, true)
	q := TDACola.CrearColaEnlazada[K]()
	q.Encolar(origen)

	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitado.Pertenece(w) {
				distancia.Guardar(w, distancia.Obtener(v)+1)
				padre.Guardar(w, &v)
				visitado.Guardar(w, true)
				q.Encolar(w)
			}
		}
	}
	return distancia, padre
}

type aristaDistancia[K comparable, V any] struct {
	vertice          K
	distanciaAOrigen V
}

func cmpDijkstra[K comparable](a, b aristaDistancia[K, int]) int {
	return a.distanciaAOrigen - b.distanciaAOrigen
}

func CaminoMinimoDijkstra[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, int], origen K) (TDADicc.Diccionario[K, int], TDADicc.Diccionario[K, *K]) {
	distancia := TDADicc.CrearHash[K, int]()
	padre := TDADicc.CrearHash[K, *K]()

	for _, v := range grafo.ObtenerVertices() {
		distancia.Guardar(v, 99999999999)
	}

	distancia.Guardar(origen, 0)
	padre.Guardar(origen, nil)
	q := TDAHeap.CrearHeap[aristaDistancia[K, int]](cmpDijkstra[K])
	q.Encolar(aristaDistancia[K, int]{origen, 0})

	for !q.EstaVacia() {
		v := q.Desencolar().vertice
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if distancia.Obtener(v)+grafo.VerPeso(v, w) < distancia.Obtener(w) {
				distancia.Guardar(w, distancia.Obtener(v)+grafo.VerPeso(v, w))
				padre.Guardar(w, &v)
				q.Encolar(aristaDistancia[K, int]{w, distancia.Obtener(w)})
			}
		}
	}

	return distancia, padre
}
