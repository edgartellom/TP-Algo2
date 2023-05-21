package diccionario

import TDAPila "tdas/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	(*nodo).clave = clave
	(*nodo).dato = dato
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, nodo := abb.buscarNodos(nil, abb.raiz, clave)
	return nodo != nil
}

func (abb *abb[K, V]) buscarNodos(padre, hijo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {

	if hijo == nil {
		return padre, hijo
	}

	if abb.cmp(clave, hijo.clave) < 0 {
		return abb.buscarNodos(hijo, hijo.izquierdo, clave)
	}

	if abb.cmp(clave, hijo.clave) > 0 {
		return abb.buscarNodos(hijo, hijo.derecho, clave)
	}

	return padre, hijo

}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	padre, hijo := abb.buscarNodos(nil, abb.raiz, clave)

	if hijo == nil {
		if padre == nil {
			abb.raiz = crearNodoAbb(clave, dato)
		} else {
			if abb.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = crearNodoAbb(clave, dato)
			} else {
				padre.derecho = crearNodoAbb(clave, dato)
			}

		}
		(*abb).cantidad++
	} else {
		hijo.dato = dato
	}
}

func (abb *abb[K, V]) Obtener(clave K) V {
	_, nodo := abb.buscarNodos(nil, abb.raiz, clave)
	comprobarSiNoPertenece(nodo)
	return (*nodo).dato
}

func (abb *abb[K, V]) borrar0Hijos(padre *nodoAbb[K, V], clave K) {
	if padre == nil {
		abb.raiz = nil
	} else {
		if abb.cmp(clave, padre.clave) < 0 {
			padre.izquierdo = nil
		} else {
			padre.derecho = nil
		}
	}
}

func (abb *abb[K, V]) borrar1Hijo(padre, nodo *nodoAbb[K, V], clave K) {
	if nodo.izquierdo != nil && nodo.derecho == nil {
		if padre == nil {
			abb.raiz = nodo.izquierdo
		} else {
			if abb.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = nodo.izquierdo
			} else {
				padre.derecho = nodo.izquierdo
			}
		}

	} else if nodo.izquierdo == nil && nodo.derecho != nil {
		if padre == nil {
			abb.raiz = nodo.derecho
		} else {
			if abb.cmp(clave, padre.clave) < 0 {
				padre.izquierdo = nodo.derecho
			} else {
				padre.derecho = nodo.derecho
			}
		}
	}

}

func (abb *abb[K, V]) borrar2Hijos(padre, nodo *nodoAbb[K, V], clave K) {
	reemplazante := nodo.izquierdo
	for reemplazante.derecho != nil {
		reemplazante = reemplazante.derecho
	}
	clave_reemplazante := reemplazante.clave
	dato_reemplazante := abb.Borrar(reemplazante.clave)
	if padre == nil {
		abb.raiz.clave = clave_reemplazante
		abb.raiz.dato = dato_reemplazante
	} else {
		nodo.clave = clave_reemplazante
		nodo.dato = dato_reemplazante
	}

}

func (abb *abb[K, V]) Borrar(clave K) V {
	padre, nodo := abb.buscarNodos(nil, abb.raiz, clave)
	comprobarSiNoPertenece(nodo)
	dato := (*nodo).dato

	if nodo.izquierdo == nil && nodo.derecho == nil {
		abb.borrar0Hijos(padre, clave)

	} else if nodo.izquierdo == nil || nodo.derecho == nil {
		abb.borrar1Hijo(padre, nodo, clave)

	} else {
		abb.borrar2Hijos(padre, nodo, clave)
	}
	abb.cantidad--
	return dato

}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) iterar(nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	if nodo.izquierdo != nil {
		abb.iterar(nodo.izquierdo, visitar)
	}
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	if nodo.derecho != nil {
		abb.iterar(nodo.derecho, visitar)
	}
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.iterar(abb.raiz, visitar)

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()

	for nodo := abb.raiz; nodo != nil; nodo = nodo.izquierdo {
		iter.pila.Apilar(nodo)
	}
	return iter
}

func (abb *abb[K, V]) estaEnelRango(nodo *nodoAbb[K, V], desde, hasta *K) bool {
	return (abb.cmp(nodo.clave, *desde) >= 0 && abb.cmp(nodo.clave, *hasta) <= 0)
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde, hasta *K, visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	if nodo.izquierdo != nil && abb.estaEnelRango(nodo.izquierdo, desde, hasta) {
		abb.iterar(nodo.izquierdo, visitar)
	}
	if !visitar(nodo.clave, nodo.dato) && abb.estaEnelRango(nodo, desde, hasta) {
		return
	}
	if nodo.derecho != nil && abb.estaEnelRango(nodo.derecho, desde, hasta) {
		abb.iterar(nodo.derecho, visitar)
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta

	for nodo := abb.raiz; nodo != nil && abb.cmp(nodo.clave, *iter.hasta) <= 0; nodo = nodo.izquierdo {
		if abb.cmp(nodo.clave, *iter.desde) > 0 {
			iter.pila.Apilar(nodo)
		}
	}
	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	estaVacia := iter.pila.EstaVacia()
	if !estaVacia && iter.hasta != nil {
		return !(iter.abb.cmp(iter.pila.VerTope().clave, *iter.hasta) > 0)
	}
	return !estaVacia
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	for nodo := iter.pila.Desapilar().derecho; nodo != nil; nodo = nodo.izquierdo {
		if iter.desde != nil {
			if iter.abb.cmp(nodo.clave, *iter.desde) > 0 {
				iter.pila.Apilar(nodo)
			}
		} else {
			iter.pila.Apilar(nodo)
		}
	}
}

func comprobarSiNoPertenece[K comparable, V any](nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic(PANIC_NO_PERTENECE)
	}
}

func (iter *iterAbb[K, V]) comprobarIteradorFinalizo() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}
