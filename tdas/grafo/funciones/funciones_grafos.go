package funciones

import (
	"fmt"
	"math"
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
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

/* ---------------------------------------------------------- ORDEN TOPOLOGICO ---------------------------------------------------------- */

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

func convertirPilaArreglo[K comparable](pila TDAPila.Pila[K]) []K {
	var resultante []K
	for !pila.EstaVacia() {
		resultante = append(resultante, pila.Desapilar())
	}
	return resultante
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

/* ---------------------------------------------------------- CAMINOS MINIMOS ---------------------------------------------------------- */

func cmpDefault[K comparable](a, b K) int {
	if a != b {
		return 1
	} else if a == b {
		return -1
	}
	return 0
}

func CaminoMinimoDijkstra[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	heap := TDAHeap.CrearHeap[K](cmpDefault[K])
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, int(math.Inf(+1)))
	}
	padres.Guardar(origen, nil)
	distancias.Guardar(origen, 0)
	heap.Encolar(origen) // heap.Encolar((vertice, distancia))
	for !heap.EstaVacia() {
		v := heap.Desencolar() // v, _ := heap.Desencolar()

		/* ---------> Con un destino pasado por parámetro <--------- */

		// if v == destino {
		// 	return padres, distancias
		// }

		for _, w := range grafo.ObtenerAdyacentes(v) {
			distanciaActual := distancias.Obtener(v) // + grafo.VerPeso(v, w)
			if distanciaActual < distancias.Obtener(w) {
				distancias.Guardar(w, distanciaActual)
				padres.Guardar(w, &v)
				heap.Encolar(w) // heap.Encolar((vertice, distancia))
			}
		}
	}
	return padres, distancias
}

type arista[K comparable, V any] struct {
	vertice   K
	adyacente K
	peso      V
}

func ObtenerAristas[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, V]) []arista[K, V] {
	var aristas []arista[K, V]
	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			aristas = append(aristas, arista[K, V]{vertice: v, adyacente: w, peso: grafo.VerPeso(v, w)})
		}
	}
	return aristas
}

func CaminoMinimoBellmanFord[K comparable, V any](grafo TDAGrafo.GrafoPesado[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	aristas := ObtenerAristas(grafo)
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, int(math.Inf(+1)))
	}
	distancias.Guardar(origen, 0)
	padres.Guardar(origen, nil)
	for i := 0; i < grafo.Cantidad(); i++ {
		var cambio bool
		for _, arista := range aristas {
			origen, destino, peso := arista.vertice, arista.adyacente, arista.peso
			distanciaActual := distancias.Obtener(origen) // + peso
			if distanciaActual < distancias.Obtener(destino) {
				cambio = true
				padres.Guardar(destino, &origen)
				distancias.Guardar(destino, distanciaActual)
				fmt.Println(peso) // -----------------------------------------------------> BORRAR
			}
		}
		if !cambio {
			return padres, distancias
		}
	}
	for _, arista := range aristas {
		origen, destino, peso := arista.vertice, arista.adyacente, arista.peso
		if distancias.Obtener(origen) /*+ peso*/ < distancias.Obtener(destino) {
			panic("Hay un ciclo negativo")
		}
		fmt.Println(peso) // -----------------------------------------------------> BORRAR
	}
	return padres, distancias
}

func CaminoMinimoBFS[K comparable, V any](grafo TDAGrafo.Grafo[K, V], origen K) (TDADicc.Diccionario[K, *K], TDADicc.Diccionario[K, int]) {
	visitados := TDADicc.CrearHash[K, bool]()
	padres := TDADicc.CrearHash[K, *K]()
	distancias := TDADicc.CrearHash[K, int]()
	for _, v := range grafo.ObtenerVertices() {
		distancias.Guardar(v, int(math.Inf(+1)))
	}
	visitados.Guardar(origen, true)
	padres.Guardar(origen, nil)
	distancias.Guardar(origen, 0)
	cola := TDACola.CrearColaEnlazada[K]()
	cola.Encolar(origen)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitados.Pertenece(w) {
				visitados.Guardar(w, true)
				padres.Guardar(w, &v)
				distancias.Guardar(w, distancias.Obtener(v)+1)
				cola.Encolar(w)
			}
		}
	}
	return padres, distancias
}
