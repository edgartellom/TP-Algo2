package lista

const PANIC_LISTA_VACIA = "La lista esta vacía"
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

type iteradorLista[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return l.primero.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	return l.ultimo.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic(PANIC_LISTA_VACIA)
	}
	dato := l.primero.dato
	l.primero = l.primero.siguiente
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := crearNodo(elemento)
	if l.EstaVacia() {
		l.primero = nuevo
		l.ultimo = l.primero
	} else {
		l.ultimo.siguiente = nuevo
		l.ultimo = l.ultimo.siguiente
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := crearNodo(elemento)
	nuevo.siguiente = l.primero
	l.primero = nuevo
	if l.EstaVacia() {
		l.ultimo = l.primero
	}
	l.largo++
}

// func Iterar(visitar func(T any) bool) {

// }

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorLista[T])
	iter.lista = l
	iter.actual = l.primero
	return iter
}

func (iter *iteradorLista[T]) VerActual() T {
	if iter.actual == nil {
		panic(PANIC_ITERADOR)
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual.siguiente != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	if iter.actual == nil {
		panic(PANIC_ITERADOR)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	if iter.actual == iter.lista.primero {
		iter.lista.InsertarPrimero(elemento)
		iter.actual = iter.lista.primero
	} else if iter.actual == iter.lista.ultimo.siguiente {
		iter.lista.InsertarUltimo(elemento)
		iter.actual = iter.lista.ultimo
	} else {
		nuevo := crearNodo(elemento)
		nuevo.siguiente = iter.actual
		iter.actual = nuevo
		iter.anterior.siguiente = nuevo
		iter.lista.largo++
	}
}

func (iter *iteradorLista[T]) Borrar() T {
	if iter.actual == nil {
		panic(PANIC_ITERADOR)
	}
	dato := iter.actual.dato
	if iter.actual == iter.lista.primero {
		iter.lista.BorrarPrimero()
		iter.actual = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual
		iter.actual = iter.anterior
		iter.lista.largo--
	}
	return dato
}
