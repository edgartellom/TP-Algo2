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
	return new(listaEnlazada[T])
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista listaEnlazada[T]) comprobarLista() {
	if lista.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
}

func (lista listaEnlazada[T]) VerPrimero() T {
	lista.comprobarLista()
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	lista.comprobarLista()
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.comprobarLista()
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.siguiente = nuevo
	}
	lista.ultimo = nuevo
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
	for actual := lista.primero; actual != nil && visitar(actual.dato); {
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

func (iter iterListaEnlazada[T]) comprobarIterador() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	iter.comprobarIterador()
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	iter.comprobarIterador()
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
	iter.comprobarIterador()
	dato := iter.actual.dato
	if iter.actual == iter.lista.primero {
		iter.lista.BorrarPrimero()
		iter.actual = iter.lista.primero
	} else {
		if iter.actual == iter.lista.ultimo {
			iter.lista.ultimo = iter.anterior
		}
		iter.anterior.siguiente = iter.actual.siguiente
		iter.actual = iter.actual.siguiente
		iter.lista.largo--
	}
	return dato
}
