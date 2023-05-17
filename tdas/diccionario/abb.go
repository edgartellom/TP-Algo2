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

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
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
	nodo := abb.raiz.buscarNodo(clave, abb.cmp)
	return nodo != nil
}

func (nodo *nodoAbb[K, V]) buscarNodo(clave K, cmp func(K, K) int) *nodoAbb[K, V] {
	if nodo == nil {
		return nodo
	}
	if cmp(clave, nodo.clave) < 0 {
		return nodo.izquierdo.buscarNodo(clave, cmp)
	}
	if cmp(clave, nodo.clave) > 0 {
		return nodo.derecho.buscarNodo(clave, cmp)
	}
	return nodo

}

func (abb *abb[K, V]) guardarNodo(nodo *nodoAbb[K, V], clave K, dato V, cmp func(K, K) int) {
	if cmp(clave, nodo.clave) < 0 {
		if nodo.izquierdo == nil {
			nodo.izquierdo = crearNodo(clave, dato)
			(*abb).cantidad++
		} else {
			abb.guardarNodo(nodo.izquierdo, clave, dato, cmp)
		}
	} else if cmp(clave, nodo.clave) > 0 {
		if nodo.derecho == nil {
			nodo.derecho = crearNodo(clave, dato)
			(*abb).cantidad++
		} else {
			abb.guardarNodo(nodo.derecho, clave, dato, cmp)
		}
	} else {
		nodo.dato = dato
	}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz == nil {
		abb.raiz = crearNodo(clave, dato)
		(*abb).cantidad++
		return
	}
	abb.guardarNodo(abb.raiz, clave, dato, abb.cmp)
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.raiz.buscarNodo(clave, abb.cmp)
	comprobarSiNoPertenece(nodo)
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := abb.raiz.buscarNodo(clave, abb.cmp)
	comprobarSiNoPertenece(nodo)
	// posicion := abb.obtenerPosicion(clave)

	// abb.comprobarEstado(posicion)

	dato := nodo.dato
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
