package biblioteca_grafos

import (
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	"tdas/pila"
)

type Arista[K comparable] struct {
	vertice   K
	adyacente K
	peso      float64
}

func ObtenerAristas[K comparable, V float64](grafo TDAGrafo.GrafoPesado[K, float64]) []Arista[K] {
	var aristas []Arista[K]
	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			aristas = append(aristas, Arista[K]{vertice: v, adyacente: w, peso: grafo.VerPeso(v, w)})
		}
	}
	return aristas
}

func GradosDeEntrada[K comparable, V any](grafo TDAGrafo.Grafo[K, V]) TDADicc.Diccionario[K, int] {
	grados := TDADicc.CrearHash[K, int]()
	vertices := grafo.ObtenerVertices()
	for _, vertice := range vertices {
		grados.Guardar(vertice, 0)
	}
	for _, v := range vertices {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			grados.Guardar(w, grados.Obtener(w)+1)
		}
	}
	return grados
}

func ReconstruirCamino[K comparable](padres TDADicc.Diccionario[K, *K], destino K) []K {
	var camino []K
	pilaAux := pila.CrearPilaDinamica[K]()
	for actual := &destino; actual != nil; actual = padres.Obtener(*actual) {
		pilaAux.Apilar(*actual)
	}
	for !pilaAux.EstaVacia() {
		camino = append(camino, pilaAux.Desapilar())
	}
	return camino
}

func MergeSort[K comparable](vertices []K, distancias TDADicc.Diccionario[K, float64]) []K {
	if len(vertices) <= 1 {
		return vertices
	}

	medio := len(vertices) / 2
	izq := MergeSort(vertices[:medio], distancias)
	der := MergeSort(vertices[medio:], distancias)

	return Merge(izq, der, distancias)
}

func Merge[K comparable](izq, der []K, distancias TDADicc.Diccionario[K, float64]) []K {
	resul := make([]K, 0)

	i, j := 0, 0
	for i < len(izq) && j < len(der) {
		if distancias.Obtener(izq[i]) < distancias.Obtener(der[j]) {
			resul = append(resul, izq[i])
			i++
		} else {
			resul = append(resul, der[j])
			j++
		}
	}

	resul = append(resul, izq[i:]...)
	resul = append(resul, der[j:]...)

	return resul
}
