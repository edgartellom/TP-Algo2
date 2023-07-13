package biblioteca_grafos

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
)

/* ---------------------------------------------------------- ORDEN TOPOLOGICO ---------------------------------------------------------- */

func TopologicoGrados[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) []K {
	grados := GradosDeEntrada(grafo)
	orden := make([]K, grafo.Cantidad())
	cola := TDACola.CrearColaEnlazada[K]()
	for iter := grados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		vertice, grado := iter.VerActual()
		if grado == 0 {
			cola.Encolar(vertice)
		}
	}
	for i := 0; !cola.EstaVacia(); i++ {
		v := cola.Desencolar()
		orden[i] = v
		for _, w := range grafo.ObtenerAdyacentes(v) {
			nuevoGrado := grados.Obtener(w) - 1
			grados.Guardar(w, nuevoGrado)
			if nuevoGrado == 0 {
				cola.Encolar(w)
			}
		}
	}
	return orden
}

func _TopologicoDfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V], v K, visitados TDADicc.Diccionario[K, bool], pila TDAPila.Pila[K]) {
	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados.Pertenece(w) {
			visitados.Guardar(w, true)
			_TopologicoDfs(grafo, w, visitados, pila)
		}
	}
	pila.Apilar(v)
}

func TopologicoDfs[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) []K {
	pila := TDAPila.CrearPilaDinamica[K]()
	visitados := TDADicc.CrearHash[K, bool]()
	for _, v := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			_TopologicoDfs(grafo, v, visitados, pila)
		}
	}
	return convertirPilaArreglo[K](pila)
}
