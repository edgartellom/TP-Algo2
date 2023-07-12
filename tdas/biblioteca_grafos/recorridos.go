package biblioteca_grafos

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

/* ---------------------------------------------------------- RECORRIDOS CON ORIGEN ---------------------------------------------------------- */

func Dfs_con_origen[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	visitados.Guardar(origen, true)
	padres.Guardar(origen, nil)
	orden.Guardar(origen, 0)
	_Dfs_con_origen[K, V](grafo, origen, visitados, padres, orden)
	return padres, orden
}

func _Dfs_con_origen[K comparable, V any](grafo TDAGrafo.Grafo[K, V], v K, visitados TDADicc.Diccionario[K, bool], padres TDADicc.Diccionario[K, *K], orden TDADicc.Diccionario[K, int]) {
	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, true)
			padres.Guardar(w, &v)
			orden.Guardar(w, orden.Obtener(v)+1)
			_Dfs_con_origen[K, V](grafo, w, visitados, padres, orden)
		}
	}
}

func Bfs_con_origen[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	visitados.Guardar(origen, true)
	padres.Guardar(origen, nil)
	orden.Guardar(origen, 0)
	cola := TDACola.CrearColaEnlazada[K]()
	cola.Encolar(origen)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padres.Guardar(w, &v)
				orden.Guardar(w, orden.Obtener(v)+1)
				cola.Encolar(w)
			}
		}
	}
	return padres, orden
}

/* ---------------------------------------------------------- RECORRIDOS SIN ORIGEN ---------------------------------------------------------- */

func Dfs_completo[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	for _, v := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			padres.Guardar(v, nil)
			orden.Guardar(v, 0)
			_Dfs_completo[K, V](grafo, v, visitados, padres, orden)
		}
	}
	return padres, orden
}

func _Dfs_completo[K comparable, V any](grafo TDAGrafo.Grafo[K, V], v K, visitados TDADicc.Diccionario[K, bool], padres TDADicc.Diccionario[K, *K], orden TDADicc.Diccionario[K, int]) {
	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, true)
			padres.Guardar(w, &v)
			orden.Guardar(w, orden.Obtener(v)+1)
			_Dfs_con_origen[K, V](grafo, w, visitados, padres, orden)
		}
	}
}

func Bfs_completo[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	orden := TDADicc.CrearHash[K, int]()
	cola := TDACola.CrearColaEnlazada[K]()
	for _, v := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			padres.Guardar(v, nil)
			orden.Guardar(v, 0)
			cola.Encolar(v)
			for !cola.EstaVacia() {
				v := cola.Desencolar()
				for _, w := range grafo.ObtenerAdyacentes(v) {
					if !visitados.Pertenece(w) {
						visitados.Guardar(w, true)
						padres.Guardar(w, &v)
						orden.Guardar(w, orden.Obtener(v)+1)
						cola.Encolar(w)
					}
				}
			}
		}
	}
	return padres, orden
}
