package biblioteca_grafos

import (
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func Centralidad[K comparable](grafo TDAGrafo.GrafoPesado[K, float64]) TDADicc.Diccionario[K, float64] {
	cent := TDADicc.CrearHash[K, float64]()
	var verticesOrdenados []K
	for _, v := range grafo.ObtenerVertices() {
		cent.Guardar(v, 0)
	}
	for _, v := range grafo.ObtenerVertices() {
		padres, distancias := CaminoMinimoDijkstra(grafo, v)
		cent_aux := TDADicc.CrearHash[K, float64]()
		for _, w := range grafo.ObtenerVertices() {
			cent_aux.Guardar(w, 0)
			verticesOrdenados = ordenarVertices(grafo, distancias)
		}
		for _, w := range verticesOrdenados {
			cent_aux.Guardar(*padres.Obtener(w), cent_aux.Obtener(*padres.Obtener(w))+1+cent_aux.Obtener(w))
		}
		for _, w := range grafo.ObtenerVertices() {
			if w == v {
				continue
			}
			cent.Guardar(w, cent.Obtener(w)+cent_aux.Obtener(w))
		}
	}

	return cent

}

func ordenarVertices[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, V], distancias TDADicc.Diccionario[K, float64]) []K {
	return MergeSort(grafo.ObtenerVertices(), distancias)
}
