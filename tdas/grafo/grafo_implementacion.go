package grafo

import (
	"fmt"
	TDAHash "tdas/diccionario"
)

type grafo[K comparable, V any] struct {
	dirigido bool
	vertices TDAHash.Diccionario[K, TDAHash.Diccionario[K, *V]]
	cantidad int
}

type grafoPesado[K comparable, V any] struct {
	grafo[K, *V]
}

type grafoNoPesado[K comparable, V any] struct {
	grafo[K, *V]
}

func CrearGrafoNoPesado[K comparable, V any](dirigido bool) GrafoNoPesado[K, V] {
	vertices := TDAHash.CrearHash[K, TDAHash.Diccionario[K, V]]()
	return &grafoNoPesado[K, V]{grafo[K, *V]{dirigido: dirigido, vertices: vertices}}
}

func CrearGrafoPesado[K comparable, V any](dirigido bool) GrafoPesado[K, V] {
	vertices := TDAHash.CrearHash[K, TDAHash.Diccionario[K, V]]()
	return &grafoPesado[K, V]{grafo[K, *V]{dirigido: dirigido, vertices: vertices}}
}

func (g *grafo[K, V]) EsDirigido() bool {
	return g.dirigido
}

func (g *grafo[K, V]) AgregarVertice(vertice K) {
	if !g.Existe(vertice) {
		g.vertices.Guardar(vertice, TDAHash.CrearHash[K, *V]())
	}
	g.cantidad++
}

func (g *grafo[K, V]) BorrarVertice(vertice K) {
	g.comprobarVertice(vertice)
	for iter := g.vertices.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, adyacentes := iter.VerActual()
		if adyacentes.Pertenece(vertice) {
			adyacentes.Borrar(vertice)
		}
	}
	g.vertices.Borrar(vertice)
	g.cantidad--
}

func (g *grafo[K, V]) BorrarArista(v1, v2 K) {
	g.comprobarVertice(v1)
	g.comprobarVertice(v2)
	g.comprobarArista(v1, v2)

	g.vertices.Obtener(v1).Borrar(v2)
	if g.dirigido {
		g.vertices.Obtener(v2).Borrar(v2)
	}
}

func (g grafo[K, V]) HayArista(v1, v2 K) bool {
	g.comprobarVertice(v1)
	g.comprobarVertice(v2)
	return g.vertices.Obtener(v1).Pertenece(v2)
}

func (g grafo[K, V]) Existe(vertice K) bool {
	return g.vertices.Pertenece(vertice)
}

func (g grafo[K, V]) ObtenerVertices() []K {
	var vertices []K
	for iter := g.vertices.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		v, _ := iter.VerActual()
		vertices = append(vertices, v)
	}
	return vertices
}

func (g grafo[K, V]) Cantidad() int {
	return g.cantidad
}

func (g *grafoNoPesado[K, V]) AgregarArista(v1, v2 K) {
	g.comprobarVertice(v1)
	g.comprobarVertice(v2)
	g.vertices.Obtener(v1).Guardar(v2, nil)
	if g.dirigido {
		g.vertices.Obtener(v2).Guardar(v1, nil)
	}

}

func (g *grafoPesado[K, V]) AgregarArista(v1, v2 K, peso V) {
	g.comprobarVertice(v1)
	g.comprobarVertice(v2)
	g.vertices.Obtener(v1).Guardar(v2, &peso)
	if g.dirigido {
		g.vertices.Obtener(v2).Guardar(v1, &peso)
	}
}

func (g grafoPesado[K, V]) VerPeso(v1, v2 K) V {
	return *(g.vertices.Obtener(v1).Obtener(v2))
}

func (g grafo[K, V]) comprobarVertice(vertice K) {
	if !g.Existe(vertice) {
		alerta := fmt.Sprintf("El vertice %v no pertenece al grafo", vertice)
		panic(alerta)
	}
}

func (grafo grafo[K, V]) comprobarArista(v1, v2 K) {
	if !grafo.HayArista(v1, v2) {
		alerta := fmt.Sprintf("No existe arista entre los vertices %v y %v", v1, v2)
		panic(alerta)
	}
}
