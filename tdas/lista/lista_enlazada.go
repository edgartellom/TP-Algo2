package lista

const PANIC_LISTA_VACIA = "La lista esta vacía"

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func crearNodo[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	nodo := new(nodoLista[T])
	(*nodo).dato = dato
	(*nodo).siguiente = siguiente
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
	nuevo := crearNodo(elemento, nil)
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
	nuevo := crearNodo(elemento, lista.primero)
	lista.primero = nuevo
	if lista.EstaVacia() {
		lista.ultimo = lista.primero
	}
	lista.largo++
}
