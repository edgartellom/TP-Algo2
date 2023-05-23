package cola_prioridad

type fcmpHeap[T comparable] func(T, T) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

const LARGO_INICIAL = 10
const FACTOR_DE_REDIMENSION = 2

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	(*heap).datos = make([]T, LARGO_INICIAL)
	return heap
}

// func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

// }

func (heap *heap[T]) redimensionarHeap(nueva_capacidad int) {
	nuevoArray := make([]T, nueva_capacidad)
	copy(nuevoArray, heap.datos)
	heap.datos = nuevoArray
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) VerMax() T {
	heap.comprobarEstado()
	return heap.datos[0]
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionarHeap(cap(heap.datos) * FACTOR_DE_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	upHeap(&heap.datos, heap.cantidad-1, heap.cmp)
}

func upHeap[T comparable](arr *[]T, indice int, cmp fcmpHeap[T]) {
	if indice == 0 {
		return
	}
	posicionPadre := (indice - 1) / 2
	if cmp((*arr)[indice], (*arr)[posicionPadre]) > 0 {
		(*arr)[posicionPadre], (*arr)[indice] = (*arr)[indice], (*arr)[posicionPadre]
		upHeap(arr, posicionPadre, cmp)
	}
}

func (heap *heap[T]) Desencolar() T {
	heap.comprobarEstado()
	if heap.cantidad*4 <= cap(heap.datos) {
		heap.redimensionarHeap((cap(heap.datos) / FACTOR_DE_REDIMENSION))
	}
	dato := heap.datos[0]
	(*heap).datos[0], (*heap).datos[heap.cantidad-1] = (*heap).datos[heap.cantidad-1], (*heap).datos[0]
	heap.cantidad--
	downHeap(&heap.datos, 0, heap.cantidad-1, heap.cmp)
	return dato
}

func (heap *heap[T]) comprobarEstado() {
	if heap.EstaVacia() {
		panic("La cola está vacia")
	}
}

func downHeap[T comparable](arr *[]T, indice, tam int, cmp fcmpHeap[T]) {
	if indice == tam {
		return
	}
	posHijoIzquierdo := 2*indice + 1
	posHijoDerecho := 2*indice + 2
	posHijoMayor := obtenerPosicionDelMayor(*arr, posHijoIzquierdo, posHijoDerecho, cmp)
	if (*arr)[posHijoMayor] != (*arr)[indice] && cmp((*arr)[indice], (*arr)[posHijoMayor]) > 0 {
		(*arr)[posHijoMayor], (*arr)[indice] = (*arr)[indice], (*arr)[posHijoMayor]
		downHeap(arr, posHijoMayor, tam, cmp)
	}
}

func obtenerPosicionDelMayor[T comparable](arr []T, izquierdo, derecho int, cmp fcmpHeap[T]) int {
	if cmp(arr[izquierdo], arr[derecho]) > 0 {
		return izquierdo
	}
	return derecho
}
