package diccionario

import (
	TDAPila "tdas/pila"
)

const (
	PANIC_NO_PERTENECE = "La clave no pertenece al diccionario"
	PANIC_ITERADOR     = "El iterador termino de iterar"
	COMPARADOR         = 0
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
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}

/* ------------------------------------------ FUNCIONES AUXILIARES ------------------------------------------ */

func (abb *abb[K, V]) obtenerPadreEHijo(padre, hijo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if hijo == nil {
		return padre, hijo
	}
	if abb.cmp(clave, hijo.clave) < COMPARADOR {
		return abb.obtenerPadreEHijo(hijo, hijo.izquierdo, clave)
	}
	if abb.cmp(clave, hijo.clave) > COMPARADOR {
		return abb.obtenerPadreEHijo(hijo, hijo.derecho, clave)
	}
	return padre, hijo
}

func (abb *abb[K, V]) agregarHijo(nodo *nodoAbb[K, V], clave K, dato V) {
	if abb.cmp(clave, nodo.clave) < COMPARADOR {
		nodo.izquierdo = crearNodoAbb(clave, dato)
	} else {
		nodo.derecho = crearNodoAbb(clave, dato)
	}
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

func (abb *abb[K, V]) borrar0Hijos(padre *nodoAbb[K, V], clave K) {
	if padre == nil {
		abb.raiz = nil
	} else {
		if abb.cmp(clave, padre.clave) < COMPARADOR {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}
	}
}

func (abb *abb[K, V]) borrar1Hijo(padre, hijo *nodoAbb[K, V], clave K) {
	if hijo.izquierdo != nil {
		if padre == nil {
			abb.raiz = hijo.izquierdo
		} else {
			if abb.cmp(clave, padre.clave) < COMPARADOR {
				padre.izquierdo = hijo.izquierdo
			} else {
				padre.derecho = hijo.izquierdo
			}
		}
	} else {
		if padre == nil {
			abb.raiz = hijo.derecho
		} else {
			if abb.cmp(clave, padre.clave) < COMPARADOR {
				padre.izquierdo = hijo.derecho
			} else {
				padre.derecho = hijo.derecho
			}
		}
	}
}

func (abb *abb[K, V]) borrar2Hijos(padre, hijo *nodoAbb[K, V], clave K) {
	claveDelReemplazante := abb.obtenerReemplazante(hijo.izquierdo)
	datoDelReemplazante := abb.Borrar(claveDelReemplazante)
	if padre == nil {
		abb.raiz.clave, abb.raiz.dato = claveDelReemplazante, datoDelReemplazante
	} else {
		if abb.cmp(clave, padre.clave) < COMPARADOR {
			padre.izquierdo.clave, padre.izquierdo.dato = claveDelReemplazante, datoDelReemplazante
		} else {
			padre.derecho.clave, padre.derecho.dato = claveDelReemplazante, datoDelReemplazante
		}
	}
}

/* ------------------------------------------ PRIMITIVAS ------------------------------------------ */

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	padre, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	if hijo != nil {
		hijo.dato = dato
	} else {
		if padre == nil {
			abb.raiz = crearNodoAbb(clave, dato)
		} else {
			abb.agregarHijo(padre, clave, dato)
		}
		abb.cantidad++
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	return hijo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	_, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	abb.comprobarExiste(hijo)
	return hijo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	padre, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	abb.comprobarExiste(hijo)
	datoBorrado := (*hijo).dato

	if hijo.izquierdo == nil && hijo.derecho == nil {
		abb.borrar0Hijos(padre, clave)
		abb.cantidad--
	} else if (hijo.izquierdo != nil && hijo.derecho == nil) || (hijo.izquierdo == nil && hijo.derecho != nil) {
		abb.borrar1Hijo(padre, hijo, clave)
		abb.cantidad--
	} else {
		abb.borrar2Hijos(padre, hijo, clave)
	}
	return datoBorrado
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iterar(nil, nil, visitar, nil)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterar(desde, hasta, visitar, abb.cmp)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.apilarNodos(abb.raiz)
	return iter
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
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
