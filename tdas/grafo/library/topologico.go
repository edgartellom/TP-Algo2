package grafo

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func GradosDeEntrada[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) TDADicc.Diccionario[K, int] {
	grados := TDADicc.CrearHash[K, int]()
	for _, vertice := range grafo.ObtenerVertices() {
		grados.Guardar(vertice, 0)
	}
	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}
	return grados
}

func TopologicoGrados[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) []K {
	grados := GradosDeEntrada(grafo)
	for _, v := range grafo.ObtenerVertices() {
		grados.Guardar(v, 0)
	}
	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}
	q := TDACola.CrearColaEnlazada[K]()

	for _, v := range grafo.ObtenerVertices() {
		if grados.Obtener(v) == 0 {
			q.Encolar(v)
		}
	}
	orden := make([]K, grafo.Cantidad())
	for i := 0; !q.EstaVacia(); i++ {
		v := q.Desencolar()
		orden[i] = v
		for _, w := range grafo.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)-1)
			if grados.Obtener(w) == 0 {
				q.Encolar(w)
			}
		}
	}
	if len(orden) == len(grafo.ObtenerVertices()) {
		return orden
	}
	return nil
}
