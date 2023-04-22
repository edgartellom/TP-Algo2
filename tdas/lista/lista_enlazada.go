package lista

const PANIC_LISTA_VACIA = "La lista esta vacia"
const PANIC_ITERADOR = "El iterador termino de iterar"

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero  *nodoLista[T]
	ultimo   *nodoLista[T]
	cantidad int
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func crearNodoLista[T any](dato T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil && l.ultimo == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodoLista(dato)
	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = l.primero
		l.primero = nuevoNodo
	}
	l.cantidad++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodoLista(dato)
	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
		l.ultimo = nuevoNodo
	}
	l.cantidad++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	primero := l.VerPrimero()
	if l.primero.siguiente == nil {
		l.primero = nil
		l.ultimo = nil
	} else {
		l.primero = l.primero.siguiente
	}
	l.cantidad--
	return primero
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.cantidad
}

// func Iterar(visitar func(T any) bool) {

// }

// func Iterador() IteradorLista[T] {
// 	return
// }
