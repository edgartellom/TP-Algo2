package biblioteca_grafos

import (
	"math"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func Centralidad[K comparable](grafo TDAGrafo.GrafoPesado[K, float64]) TDADicc.Diccionario[K, float64] {
	cent := TDADicc.CrearHash[K, float64]()
	vertices := grafo.ObtenerVertices()
	for _, v := range vertices {
		cent.Guardar(v, 0)
	}
	for _, v := range vertices {
		padres, distancias := CaminoMinimoDijkstra[K](grafo, v)
		cent_aux := TDADicc.CrearHash[K, float64]()
		for _, w := range vertices {
			cent_aux.Guardar(w, 0)
		}
		verticesOrdenados := filtrarYOrdenarVertices[K](distancias)
		for _, verticeDistancia := range verticesOrdenados {
			w := verticeDistancia.vertice
			z := padres.Obtener(w)
			if z != nil {
				cent_aux.Guardar(*z, cent_aux.Obtener(*z)+cent_aux.Obtener(w)+1)
			}
		}
		for _, w := range vertices {
			if w == v {
				continue
			}
			cent.Guardar(w, cent.Obtener(w)+cent_aux.Obtener(w))
		}
	}
	return cent
}

type verticeDistancia[K comparable] struct {
	vertice   K
	distancia float64
}

func cmpVertices[K comparable](v1, v2 verticeDistancia[K]) int {
	return int(v2.distancia - v1.distancia)
}

func filtrarYOrdenarVertices[K comparable](distancias TDADicc.Diccionario[K, float64]) []verticeDistancia[K] {
	var vertices []verticeDistancia[K]
	for iter := distancias.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		v, dist := iter.VerActual()
		if dist != math.MaxFloat64 {
			vertices = append(vertices, verticeDistancia[K]{v, dist})
		}
	}
	TDAHeap.HeapSort[verticeDistancia[K]](vertices, cmpVertices[K])
	return vertices
}
