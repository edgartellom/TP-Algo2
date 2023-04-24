package lista

const PANIC_LISTA_VACIA = "La lista esta vacia"
const PANIC_ITERADOR = "El iterador termino de iterar"

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	(*nodo).dato = dato
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevo
		lista.ultimo = lista.primero
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = lista.ultimo.siguiente
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := crearNodo(elemento)
	nuevo.siguiente = lista.primero
	lista.primero = nuevo
	if lista.EstaVacia() {
		lista.ultimo = lista.primero
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iterListaEnlazada[T])
	(*iter).lista = lista
	(*iter).actual = lista.primero
	return iter
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	if iter.actual == iter.lista.primero {
		iter.lista.InsertarPrimero(elemento)
		iter.actual = iter.lista.primero
	} else if iter.actual == iter.lista.ultimo.siguiente {
		iter.lista.InsertarUltimo(elemento)
		iter.actual = iter.lista.ultimo
	} else {
		nuevo := crearNodo(elemento)
		iter.anterior.siguiente = nuevo
		nuevo.siguiente = iter.actual
		iter.actual = nuevo
		iter.lista.largo++
	}
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
	dato := iter.actual.dato
	if iter.actual == iter.lista.primero {
		iter.lista.BorrarPrimero()
		iter.actual = iter.lista.primero
	} else if iter.actual == iter.lista.ultimo {
		iter.anterior.siguiente = iter.actual.siguiente
		iter.lista.ultimo = iter.anterior
		iter.actual = iter.actual.siguiente
		iter.lista.largo--
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
		iter.actual = iter.actual.siguiente
		iter.lista.largo--
	}
	return dato
}
