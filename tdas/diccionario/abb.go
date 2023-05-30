package diccionario

import (
	TDAPila "tdas/pila"
)

const (
	PANIC_NO_PERTENECE    = "La clave no pertenece al diccionario"
	PANIC_ITERADOR        = "El iterador termino de iterar"
	COMPARADOR            = 0
	MAXIMA_CANTIDAD_HIJOS = 2
)

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      funcCmp[K]
	cantidad int
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

/* ------------------------------------------ FUNCIONES DE CREACION ------------------------------------------ */

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

/* ------------------------------------------ FUNCIONES AUXILIARES ------------------------------------------ */

func (abb *abb[K, V]) obtenerVinculo(vinculo **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	actual := *vinculo
	if actual == nil {
		return vinculo
	}
	if abb.cmp(clave, (actual).clave) < COMPARADOR {
		return abb.obtenerVinculo(&(actual).izquierdo, clave)
	}
	if abb.cmp(clave, (actual).clave) > COMPARADOR {
		return abb.obtenerVinculo(&(actual).derecho, clave)
	}
	return vinculo
}

func (abb *abb[K, V]) obtenerReemplazante(nodo *nodoAbb[K, V]) K {
	if nodo.derecho == nil {
		return nodo.clave
	}
	return abb.obtenerReemplazante(nodo.derecho)
}

func (nodo *nodoAbb[K, V]) iterar(desde, hasta *K, visitar func(clave K, dato V) bool, cmp funcCmp[K]) bool {
	var condicionDeCorte bool
	if nodo == nil {
		return condicionDeCorte
	}

	if !condicionDeCorte && nodo.comprobarDesde(desde, cmp) {
		condicionDeCorte = nodo.izquierdo.iterar(desde, hasta, visitar, cmp)
	}
	if !condicionDeCorte && nodo.comprobarEnRango(desde, hasta, cmp) {
		condicionDeCorte = !visitar(nodo.clave, nodo.dato)
	}
	if !condicionDeCorte && nodo.comprobarHasta(hasta, cmp) {
		condicionDeCorte = nodo.derecho.iterar(desde, hasta, visitar, cmp)
	}
	return condicionDeCorte
}

func (nodo *nodoAbb[K, V]) comprobarDesde(desde *K, cmp funcCmp[K]) bool {
	return ((desde == nil) || (desde != nil && cmp(nodo.clave, *desde) >= 0))
}

func (nodo *nodoAbb[K, V]) comprobarHasta(hasta *K, cmp funcCmp[K]) bool {
	return ((hasta == nil) || (hasta != nil && cmp(nodo.clave, *hasta) <= 0))
}

func (nodo *nodoAbb[K, V]) comprobarEnRango(desde, hasta *K, cmp funcCmp[K]) bool {
	return nodo.comprobarDesde(desde, cmp) && nodo.comprobarHasta(hasta, cmp)
}

func (iter *iterAbb[K, V]) apilarNodos(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if iter.desde != nil && iter.abb.cmp(nodo.clave, *iter.desde) < COMPARADOR {
		iter.apilarNodos(nodo.derecho)
	} else {
		iter.pila.Apilar(nodo)
		iter.apilarNodos(nodo.izquierdo)
	}
}

func (abb *abb[K, V]) borrarCon0o1Hijo(vinculo **nodoAbb[K, V]) {
	actual := *vinculo
	cantidadDeHijos := abb.contarHijos(*vinculo)
	if cantidadDeHijos == COMPARADOR {
		*vinculo = nil
	} else {
		if actual.izquierdo != nil {
			*vinculo = actual.izquierdo
		} else {
			*vinculo = actual.derecho
		}
	}
	abb.cantidad--
}

func (abb *abb[K, V]) contarHijos(actual *nodoAbb[K, V]) int {
	var cantidadDeHijos int
	if actual.izquierdo != nil {
		cantidadDeHijos++
	}
	if actual.derecho != nil {
		cantidadDeHijos++
	}
	return cantidadDeHijos
}

/* ------------------------------------------ PRIMITIVAS ------------------------------------------ */

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	vinculo := abb.obtenerVinculo(&abb.raiz, clave)
	actual := *vinculo
	if actual != nil {
		actual.dato = dato
	} else {
		*vinculo = crearNodoAbb(clave, dato)
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	actual := *abb.obtenerVinculo(&abb.raiz, clave)
	return actual != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	actual := *abb.obtenerVinculo(&abb.raiz, clave)
	abb.comprobarExiste(actual)
	return actual.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	vinculo := abb.obtenerVinculo(&abb.raiz, clave)
	actual := *vinculo
	abb.comprobarExiste(actual)
	datoBorrado := actual.dato

	cantidadDeHijos := abb.contarHijos(actual)
	if cantidadDeHijos < MAXIMA_CANTIDAD_HIJOS {
		abb.borrarCon0o1Hijo(vinculo)
	} else {
		claveReemplazante := abb.obtenerReemplazante(actual.izquierdo)
		datoReemplazante := abb.Borrar(claveReemplazante)
		(*vinculo).clave, (*vinculo).dato = claveReemplazante, datoReemplazante
	}
	return datoBorrado
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterar(desde, hasta, visitar, abb.cmp)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := &iterAbb[K, V]{abb, pila, desde, hasta}
	iter.apilarNodos(abb.raiz)
	return iter
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()

	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	if iter.hasta != nil {
		return !iter.pila.EstaVacia() && iter.abb.cmp(iter.pila.VerTope().clave, *iter.hasta) <= COMPARADOR
	}
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	desapilado := iter.pila.Desapilar()
	iter.apilarNodos(desapilado.derecho)
}

/* ------------------------------------- FUNCIONES DE COMPROBACION ------------------------------------- */

func (iter *iterAbb[K, V]) comprobarIteradorFinalizo() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}

func (abb *abb[K, V]) comprobarExiste(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic(PANIC_NO_PERTENECE)
	}
}
