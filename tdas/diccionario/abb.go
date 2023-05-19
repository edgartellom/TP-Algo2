package diccionario

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
	abb    *abb[K, V]
	actual *nodoAbb[K, V]
	padre  *nodoAbb[K, V]
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
		// actual = nodo
		return padre, hijo
	}
	if abb.cmp(hijo.clave, clave) < 0 {

		return abb.buscarNodos(hijo, hijo.izquierdo, clave)
	}
	if abb.cmp(hijo.clave, clave) > 0 {

		return abb.buscarNodos(hijo, hijo.derecho, clave)
	}
	return padre, hijo

}

// func (abb *abb[K, V]) guardarNodo(nodo *nodoAbb[K, V], clave K, dato V, cmp func(K, K) int) {
// 	if cmp(clave, nodo.clave) < 0 {
// 		if nodo.izquierdo == nil {
// 			nodo.izquierdo = crearNodo(clave, dato)
// 			(*abb).cantidad++
// 		} else {
// 			abb.guardarNodo(nodo.izquierdo, clave, dato, cmp)
// 		}
// 	} else if cmp(clave, nodo.clave) > 0 {
// 		if nodo.derecho == nil {
// 			nodo.derecho = crearNodo(clave, dato)
// 			(*abb).cantidad++
// 		} else {
// 			abb.guardarNodo(nodo.derecho, clave, dato, cmp)
// 		}
// 	} else {
// 		nodo.dato = dato
// 	}
// }

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	padre, hijo := abb.buscarNodos(nil, abb.raiz, clave)

	if hijo == nil {
		if padre == nil {
			abb.raiz = crearNodoAbb(clave, dato)
		} else {
			if abb.cmp(padre.clave, clave) < 0 {
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

func (abb *abb[K, V]) Borrar(clave K) V {
	_, nodo := abb.buscarNodos(nil, abb.raiz, clave)
	comprobarSiNoPertenece(nodo)
	// posicion := abb.obtenerPosicion(clave)

	// abb.comprobarEstado(posicion)

	dato := (*nodo).dato
	// abb.tabla[posicion].estado = BORRADO
	// abb.cantidad--
	// abb.borrados++

	// nuevoTam := abb.tam / FACTOR_REDIMENSION
	// if abb.factorDeCarga() <= FACTOR_ACHICAR && nuevoTam < LONGITUD_INICIAL {
	// 	abb.redimensionarTabla(nuevoTam)
	// }

	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	iter.actual = abb.raiz
	// iter.actual, iter.posicion = obtenerCeldaOcupada(abb.tabla, 0)
	return iter
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	// iter.actual, iter.posicion = obtenerCeldaOcupada(abb.tabla, 0)
	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()

	return iter.actual.clave, iter.actual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	if iter.actual.izquierdo != nil {
		iter.actual = iter.actual.izquierdo
	}
	iter.Siguiente()
	if iter.actual.derecho != nil {
		iter.actual = iter.actual.derecho
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
