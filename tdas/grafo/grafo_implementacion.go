package grafo

import TDAHash "tdas/diccionario"

type grafo[K comparable, V any] struct {
	diccVertices TDAHash.Diccionario[K, TDAHash.Diccionario[K, *V]]
	esDirigido   bool
}

type grafoPesado[K comparable, V any] struct {
	grafo[K, V]
}

type grafoNoPesado[K comparable, V any] struct {
	grafo[K, V]
}

func CrearGrafoPesado[K comparable, V any](esDirigido bool) GrafoPesado[K, V] {
	vertices := TDAHash.CrearHash[K, TDAHash.Diccionario[K, *V]]()
	return &grafoPesado[K, V]{grafo[K, V]{diccVertices: vertices, esDirigido: esDirigido}}
}

func CrearGrafoNoPesado[K comparable, V any](esDirigido bool) GrafoNoPesado[K, V] {
	vertices := TDAHash.CrearHash[K, TDAHash.Diccionario[K, *V]]()
	return &grafoNoPesado[K, V]{grafo[K, V]{diccVertices: vertices, esDirigido: esDirigido}}
}

func (grafo grafo[K, V]) EsDirigido() bool {
	return grafo.esDirigido
}

func (grafo grafo[K, V]) AgregarVertice(vertice K) {
	if !grafo.Existe(vertice) {
		grafo.diccVertices.Guardar(vertice, TDAHash.CrearHash[K, *V]())
	}
}

func (grafo grafo[K, V]) BorrarVertice(vertice K) {
	grafo.comprobarVertice(vertice)
	for iter := grafo.diccVertices.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, adyacentes := iter.VerActual()
		if adyacentes.Pertenece(vertice) {
			adyacentes.Borrar(vertice)
		}
	}
	grafo.diccVertices.Borrar(vertice)
}

func (grafo grafo[K, V]) BorrarArista(vertice1, vertice2 K) {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)
	grafo.comprobarArista(vertice1, vertice2)

	grafo.diccVertices.Obtener(vertice1).Borrar(vertice2)
	if !grafo.esDirigido {
		grafo.diccVertices.Obtener(vertice2).Borrar(vertice1)
	}
}

func (grafo grafo[K, V]) Existe(vertice K) bool {
	return grafo.diccVertices.Pertenece(vertice)
}

func (grafo grafo[K, V]) HayArista(vertice1, vertice2 K) bool {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)
	return grafo.diccVertices.Obtener(vertice1).Pertenece(vertice2)
}

func (grafo grafo[K, V]) Cantidad() int {
	return grafo.diccVertices.Cantidad()
}

func (grafo grafoPesado[K, V]) AgregarArista(vertice1, vertice2 K, peso V) {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)

	grafo.diccVertices.Obtener(vertice1).Guardar(vertice2, &peso)
	if !grafo.esDirigido {
		grafo.diccVertices.Obtener(vertice2).Guardar(vertice1, &peso)
	}
}

func (grafo grafoPesado[K, V]) VerPeso(vertice1, vertice2 K) V {
	return *grafo.diccVertices.Obtener(vertice1).Obtener(vertice2)
}

func (grafo grafoNoPesado[K, V]) AgregarArista(vertice1, vertice2 K) {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)

	grafo.diccVertices.Obtener(vertice1).Guardar(vertice2, nil)
	if !grafo.esDirigido {
		grafo.diccVertices.Obtener(vertice2).Guardar(vertice1, nil)
	}
}

func (grafo grafo[K, V]) ObtenerAdyacentes(vertice1 K) []K {
	adyacentes := make([]K, grafo.diccVertices.Obtener(vertice1).Cantidad())
	for iter, i := grafo.diccVertices.Obtener(vertice1).Iterador(), 0; iter.HaySiguiente(); iter.Siguiente() {
		adyacente, _ := iter.VerActual()
		adyacentes[i] = adyacente
		i++
	}
	return adyacentes
}

func (grafo grafo[K, V]) ObtenerVertices() []K {
	vertices := make([]K, grafo.Cantidad())
	for iter, i := grafo.diccVertices.Iterador(), 0; iter.HaySiguiente(); iter.Siguiente() {
		vertice, _ := iter.VerActual()
		vertices[i] = vertice
		i++
	}
	return vertices
}

func (grafo grafo[K, V]) comprobarVertice(vertice K) {
	if !grafo.Existe(vertice) {
		panic("El vertice no pertenece al grafo")
	}
}

func (grafo grafo[K, V]) comprobarArista(vertice1, vertice2 K) {
	if !grafo.HayArista(vertice1, vertice2) {
		panic("No existe arista entre los vertices")
	}
}
