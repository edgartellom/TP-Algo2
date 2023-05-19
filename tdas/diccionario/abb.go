package diccionario

import TDAPila "tdas/pila"

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

func (abb *abb[K, V]) _obtenerPadreEHijo(padre, hijo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if padre == nil && hijo == nil {
		return nil, nil
	}
	orientacion := abb.cmp(hijo.clave, clave)
	if (orientacion < 0 && hijo.izquierdo == nil) || (orientacion > 0 && hijo.derecho == nil) {
		return hijo, nil
	}
	if orientacion < 0 {
		return abb._obtenerPadreEHijo(hijo, hijo.izquierdo, clave)
	} else if orientacion > 0 {
		return abb._obtenerPadreEHijo(hijo, hijo.derecho, clave)
	}
	return padre, hijo
}

func (abb *abb[K, V]) obtenerPadreEHijo(clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	return abb._obtenerPadreEHijo(nil, abb.raiz, clave)
}

func (abb *abb[K, V]) agregarHijo(nodo *nodoAbb[K, V], clave K, valor V) {
	orientacion := abb.cmp(nodo.clave, clave)
	if orientacion < 0 {
		nodo.izquierdo = crearNodoAbb(clave, valor)
	} else {
		nodo.derecho = crearNodoAbb(clave, valor)
	}
	abb.cantidad++
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	padre, hijo := abb.obtenerPadreEHijo(clave)
	if padre == nil && hijo == nil {
		abb.raiz = crearNodoAbb(clave, valor)
		abb.cantidad++
	} else if padre != nil && hijo == nil {
		abb.agregarHijo(padre, clave, valor)
	} else {
		hijo.dato = valor
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, hijo := abb.obtenerPadreEHijo(clave)
	return hijo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	_, hijo := abb.obtenerPadreEHijo(clave)
	abb.comprobar(hijo)
	return hijo.dato
}

func (abb *abb[K, V]) borrar0Hijos(padre *nodoAbb[K, V], clave K) {
	orientacion := abb.cmp(padre.clave, clave)
	if orientacion < 0 {
		padre.izquierdo = nil
	} else {
		padre.derecho = nil
	}
	abb.cantidad--
}

func (abb *abb[K, V]) borrar1Hijo(padre, hijo *nodoAbb[K, V], clave K) {
	orientacion := abb.cmp(padre.clave, clave)
	if orientacion < 0 {
		if hijo.derecho != nil {
			padre.izquierdo = hijo.derecho
		} else {
			padre.izquierdo = hijo.izquierdo
		}
	} else {
		if hijo.derecho != nil {
			padre.derecho = hijo.derecho
		} else {
			padre.derecho = hijo.izquierdo
		}
	}
	abb.cantidad--
}

func (abb *abb[K, V]) Borrar(clave K) V {
	padre, hijo := abb.obtenerPadreEHijo(clave)
	abb.comprobar(hijo)
	if abb.cantidad == 1 {
		abb.raiz = nil
		abb.cantidad--
	} else {
		cantidadDeHijos := abb.contarHijos(hijo)
		switch cantidadDeHijos {
		case 0:
			abb.borrar0Hijos(padre, clave)
		case 1:
			abb.borrar1Hijo(padre, hijo, clave)
		case 2:
		}
	}
	return (*hijo).dato
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

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

/* --------------------------------- NO FUNCIONA --------------------------------- */

// func (abb *abb[K, V]) obtenerActual(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
// 	if nodo.izquierdo != nil {
// 		nodo = abb.obtenerActual(nodo.izquierdo)
// 	} else if nodo.derecho != nil {
// 		nodo = abb.obtenerActual(nodo.derecho)
// 	}
// 	return nodo
// }

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for actual := abb.raiz; actual != nil && visitar(actual.clave, actual.dato); {
		if actual.izquierdo != nil {
			actual = actual.izquierdo
		} else {
			actual = actual.derecho
		}
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

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
	return iter
}

func (iter *iterador[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()

	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

/* ------------------------------------------ ITERADOR EXTERNO ------------------------------------------ */

func (iter *iterador[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
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
	booleano := iter.pila.EstaVacia()
	if !booleano {
		return iter.arbol.cmp(iter.pila.VerTope().clave, *iter.hasta) > 0
	}
	return !booleano
}

func (iter *iterAbbRango[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	iter.pila.Desapilar()

}

// func (iter *iterAbbRango[K, V]) apilarNodoEIzquierdosIterRango(nodo *nodoAbb[K, V]) {
// 	for actual := nodo; actual != nil; actual = actual.izquierdo {
// 		if iter.arbol.cmp(actual.clave, *iter.desde) > 0 {
// 			iter.pila.Apilar(actual)
// 		}
// 	}
// }

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
