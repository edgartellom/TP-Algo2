package diccionario

import (
	TDAPila "tdas/pila"
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

type iterador[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
}
type iterAbb[K comparable, V any] struct {
	iterador[K, V]
}

type iterAbbRango[K comparable, V any] struct {
	iterador[K, V]
	desde *K
	hasta *K
}

/* ------------------------------------------ FUNCIONES DE CREACION ------------------------------------------ */

func crearNodoAbb[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = valor
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	arbol := new(abb[K, V])
	arbol.cmp = funcion_cmp
	return arbol
}

/* ------------------------------------------ FUNCIONES AUXILIARES ------------------------------------------ */

func (abb *abb[K, V]) obtenerPadreEHijo(padre, hijo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if hijo == nil {
		return padre, hijo
	}
	orientacion := abb.cmp(clave, hijo.clave)

	if orientacion < 0 {
		return abb.obtenerPadreEHijo(hijo, hijo.izquierdo, clave)
	}
	if orientacion > 0 {
		return abb.obtenerPadreEHijo(hijo, hijo.derecho, clave)
	}
	return padre, hijo
}

func (abb *abb[K, V]) agregarHijo(nodo *nodoAbb[K, V], clave K, valor V) {
	orientacion := abb.cmp(clave, nodo.clave)
	if orientacion < 0 {
		nodo.izquierdo = crearNodoAbb(clave, valor)
	} else {
		nodo.derecho = crearNodoAbb(clave, valor)
	}
}

func (abb *abb[K, V]) contarHijos(nodo *nodoAbb[K, V]) int {
	var contador int
	if nodo.izquierdo != nil {
		contador++
	}
	if nodo.derecho != nil {
		contador++
	}
	return contador
}

func (abb *abb[K, V]) borrar0Hijos(padre *nodoAbb[K, V], clave K) {
	orientacion := abb.cmp(clave, padre.clave)
	if orientacion < 0 {
		padre.izquierdo = nil
	} else {
		padre.derecho = nil
	}
}

func (abb *abb[K, V]) borrar1Hijo(padre, hijo *nodoAbb[K, V], clave K) {
	if padre == nil {
		if hijo.derecho != nil {
			abb.raiz = hijo.derecho
		} else {
			abb.raiz = hijo.izquierdo
		}
	} else {
		orientacion := abb.cmp(clave, padre.clave)
		if orientacion < 0 && hijo.derecho != nil {
			padre.izquierdo = hijo.derecho
		} else if orientacion < 0 && hijo.izquierdo != nil {
			padre.izquierdo = hijo.izquierdo
		} else if orientacion > 0 && hijo.derecho != nil {
			padre.derecho = hijo.derecho
		} else if orientacion > 0 && hijo.izquierdo != nil {
			padre.derecho = hijo.izquierdo
		}
	}
}

func (abb *abb[K, V]) obtenerElMasDerechoDelLadoIzquierdo(nodo *nodoAbb[K, V]) K {
	if nodo.derecho == nil {
		return nodo.clave
	}
	return abb.obtenerElMasDerechoDelLadoIzquierdo(nodo.derecho)
}

/* ------------------------------------------ PRIMITIVAS ------------------------------------------ */

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	padre, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	if hijo != nil {
		hijo.dato = valor
	} else {
		if padre == nil {
			abb.raiz = crearNodoAbb(clave, valor)
		} else {
			abb.agregarHijo(padre, clave, valor)
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
	abb.comprobar(hijo)
	return hijo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	padre, hijo := abb.obtenerPadreEHijo(nil, abb.raiz, clave)
	abb.comprobar(hijo)
	datoBorrado := (*hijo).dato
	if abb.cantidad == 1 {
		abb.raiz = nil
		abb.cantidad--
	} else {
		cantidadDeHijos := abb.contarHijos(hijo)
		switch cantidadDeHijos {
		case 0:
			abb.borrar0Hijos(padre, clave)
			abb.cantidad--
		case 1:
			abb.borrar1Hijo(padre, hijo, clave)
			abb.cantidad--
		case 2:
			claveDelReemplazante := abb.obtenerElMasDerechoDelLadoIzquierdo(hijo.izquierdo)
			valorDelReemplazante := abb.Borrar(claveDelReemplazante)
			if padre != nil {
				orientacion := abb.cmp(clave, padre.clave)
				if orientacion < 0 {
					padre.izquierdo.clave = claveDelReemplazante
					padre.izquierdo.dato = valorDelReemplazante
				} else {
					padre.derecho.clave = claveDelReemplazante
					padre.derecho.dato = valorDelReemplazante
				}
			} else {
				abb.raiz.clave = claveDelReemplazante
				abb.raiz.dato = valorDelReemplazante
			}
		}
	}
	return datoBorrado
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) _Iterar(nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	abb._Iterar(nodo.izquierdo, visitar)

	if !visitar(nodo.clave, nodo.dato) {
		return
	}

	abb._Iterar(nodo.derecho, visitar)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb._Iterar(abb.raiz, visitar)
}

func (abb *abb[K, V]) _IterarRango(nodo *nodoAbb[K, V], desde, hasta *K, visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	if abb.cmp(nodo.clave, *desde) >= 0 {
		abb._Iterar(nodo.izquierdo, visitar)
	}

	if !visitar(nodo.clave, nodo.dato) {
		return
	}

	if abb.cmp(nodo.clave, *hasta) <= 0 {
		abb._Iterar(nodo.derecho, visitar)
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb._IterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.arbol = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	// for actual := abb.raiz; actual != nil; actual = actual.izquierdo {
	// 	iter.pila.Apilar(actual)
	// }
	iter.apilarNodoEIzquierdosIterComun(abb.raiz)
	return iter
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbbRango[K, V])
	iter.arbol = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.apilarNodoEIzquierdosIterRango(abb.raiz)
	return iter
}

func (iter *iterador[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()

	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

/* ------------------------------------------ ITERADOR EXTERNO ------------------------------------------ */

func (iter *iterador[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia() // un elemento => pila.EstaVacia() = false => !pila.EstaVacia() = true
}

func (iter *iterador[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	desapilado := iter.pila.Desapilar()
	iter.apilarNodoEIzquierdosIterComun(desapilado.derecho)
}

func (iter *iterador[K, V]) apilarNodoEIzquierdosIterComun(nodo *nodoAbb[K, V]) {
	for actual := nodo; actual != nil; actual = actual.izquierdo {
		iter.pila.Apilar(actual)
	}
}

/* ------------------------------------ ITERADOR EXTERNO POR RANGOS ------------------------------------ */

func (iter *iterAbbRango[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia() && iter.arbol.cmp(iter.pila.VerTope().clave, *iter.hasta) <= 0
}

func (iter *iterAbbRango[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	desapilado := iter.pila.Desapilar()
	iter.apilarNodoEIzquierdosIterRango(desapilado.derecho)
}

func (iter *iterAbbRango[K, V]) apilarNodoEIzquierdosIterRango(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	iter.pila.Apilar(nodo)
	iter.apilarNodoEIzquierdosIterRango(nodo.izquierdo)
	for !iter.pila.EstaVacia() && iter.arbol.cmp(iter.pila.VerTope().clave, *iter.desde) < 0 {
		desapilado := iter.pila.Desapilar()
		iter.apilarNodoEIzquierdosIterRango(desapilado.derecho)
	}
}

/* ------------------------------------- FUNCIONES DE COMPROBACION ------------------------------------- */

func (iter *iterador[K, V]) comprobarIteradorFinalizo() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}

func (abb *abb[K, V]) comprobar(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic(PANIC_NO_PERTENECE)
	}
}
