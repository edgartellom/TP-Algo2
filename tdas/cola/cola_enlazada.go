package cola

const panicColaVacia = "La cola esta vacia"

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func crearNodoCola[T any](dato T) *nodoCola[T] {
	nuevoNodo := new(nodoCola[T])
	nuevoNodo.dato = dato
	return nuevoNodo
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(panicColaVacia)
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodoCola(dato)
	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
		c.ultimo = nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(panicColaVacia)
	}
	primero := c.VerPrimero()
	if c.primero.siguiente == nil {
		c.primero = nil
		c.ultimo = nil
	} else {
		c.primero = c.primero.siguiente
	}
	return primero
}
