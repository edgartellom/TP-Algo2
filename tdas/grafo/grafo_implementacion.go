package grafo

import TDADiccionario "tdas/diccionario"

type grafo[K comparable, V any] struct {
	diccPrincipal TDADiccionario.Diccionario[K, TDADiccionario.Diccionario[K, *V]]
	esDirigido    bool
}

func CrearGrafo[K comparable, V any](esDirigido bool) Grafo[K, V] {
	diccionario := TDADiccionario.CrearHash[K, TDADiccionario.Diccionario[K, *V]]()
	return &grafo[K, V]{diccPrincipal: diccionario, esDirigido: esDirigido}
}

func (grafo grafo[K, V]) AgregarVertice(vertice K) {
	adyacentes := TDADiccionario.CrearHash[K, *V]()
	grafo.diccPrincipal.Guardar(vertice, adyacentes)
}

func (grafo grafo[K, V]) BorrarVertice(vertice K) {
	grafo.comprobarVertice(vertice)
	for iter := grafo.diccPrincipal.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, adyacentes := iter.VerActual()
		if adyacentes.Pertenece(vertice) {
			adyacentes.Borrar(vertice)
		}
	}
}

func (grafo grafo[K, V]) AgregarArista(vertice1, vertice2 K, peso V) {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)

	grafo.diccPrincipal.Obtener(vertice1).Guardar(vertice2, &peso)
	if grafo.esDirigido {
		grafo.diccPrincipal.Obtener(vertice2).Guardar(vertice1, &peso)
	}
}

func (grafo grafo[K, V]) BorrarArista(vertice1, vertice2 K) {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)
	grafo.comprobarArista(vertice1, vertice2)

	grafo.diccPrincipal.Obtener(vertice1).Borrar(vertice2)
	if grafo.esDirigido {
		grafo.diccPrincipal.Obtener(vertice2).Borrar(vertice1)
	}
}

func (grafo grafo[K, V]) Pertenece(vertice K) bool {
	return grafo.diccPrincipal.Pertenece(vertice)
}

func (grafo grafo[K, V]) HayArista(vertice1, vertice2 K) bool {
	grafo.comprobarVertice(vertice1)
	grafo.comprobarVertice(vertice2)
	return grafo.diccPrincipal.Obtener(vertice1).Pertenece(vertice2)
}

func (grafo grafo[K, V]) ObtenerPeso(vertice1, vertice2 K) V {
	return *(grafo.diccPrincipal.Obtener(vertice1)).Obtener(vertice2)
}

func (grafo grafo[K, V]) comprobarVertice(vertice K) {
	if !grafo.Pertenece(vertice) {
		panic("La arista no pertenece al grafo")
	}
}

func (grafo grafo[K, V]) comprobarArista(vertice1, vertice2 K) {
	if !grafo.HayArista(vertice1, vertice2) {
		panic("No existe arista entre los vertices")
	}
}
