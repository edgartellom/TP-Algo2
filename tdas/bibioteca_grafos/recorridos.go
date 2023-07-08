package bibioteca_grafos

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func Bfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	padres.Guardar(origen, nil)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)
	q := TDACola.CrearColaEnlazada[K]()
	q.Encolar(origen)

	for !q.EstaVacia() {
		v := q.Desencolar()

		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				padres.Guardar(w, &v)
				orden.Guardar(w, orden.Obtener(v)+1)
				visitados.Guardar(w, true)
				q.Encolar(w)
			}
		}
	}

	return padres, orden
}

func dfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V], v K, visitados TDADicc.Diccionario[K, bool], padres TDADicc.Diccionario[K, *K], orden TDADicc.Diccionario[K, int]) {
	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, true)
			padres.Guardar(w, &v)
			orden.Guardar(w, orden.Obtener(v)+1)
			dfs(grafo, w, visitados, padres, orden)
		}
	}
}

func RecorridoDfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	padres.Guardar(origen, nil)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)
	dfs(grafo, origen, visitados, padres, orden)
	return padres, orden
}

func RecorridoDfsCompleto[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()

	for _, v := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			padres.Guardar(v, nil)
			orden.Guardar(v, 0)
			dfs(grafo, v, visitados, padres, orden)
		}
	}
}
