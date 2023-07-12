package biblioteca_grafos

import (
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func cmpPrim[K comparable](a, b Arista[K]) int {
	return a.peso - b.peso
}

func MstPrim[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, int]) TDAGrafo.GrafoPesado[K, int] {
	v := grafo.ObtenerVerticeAleatorio()
	visitado := TDADicc.CrearHash[K, bool]()
	visitado.Guardar(v, true)
	q := TDAHeap.CrearHeap[Arista[K]](cmpPrim[K])

	for _, w := range grafo.ObtenerAdyacentes(v) {
		q.Encolar(Arista[K]{v, w, grafo.VerPeso(v, w)})
	}
	arbol := TDAGrafo.CrearGrafoPesado[K, int](false)
	for _, v := range grafo.ObtenerVertices() {
		arbol.AgregarVertice(v)
	}

	for !q.EstaVacia() {
		aristaActual := q.Desencolar()
		v, w, peso := aristaActual.vertice, aristaActual.adyacente, aristaActual.peso
		if visitado.Pertenece(w) {
			continue
		}
		arbol.AgregarArista(v, w, peso)
		visitado.Guardar(w, true)

		for _, u := range grafo.ObtenerAdyacentes(w) {
			if !visitado.Pertenece(u) {
				q.Encolar(Arista[K]{w, u, grafo.VerPeso(w, u)})
			}
		}
	}
	return arbol
}
