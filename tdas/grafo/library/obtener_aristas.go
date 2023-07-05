package grafo

import (
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

type Arista[K comparable, V any] struct {
	vertice   K
	adyacente K
	peso      V
}

func ObtenerAristas[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, V]) []Arista[K, V] {
	var aristas []Arista[K, V]
	visitados := TDADicc.CrearHash[K, bool]()

	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				aristas = append(aristas, Arista[K, V]{v, w, grafo.VerPeso(v, w)})
			}
		}
		visitados.Guardar(v, true)
	}
	return aristas
}
