package cola_prioridad

const (
	LARGO_INICIAL         = 10
	FACTOR_DE_REDIMENSION = 2
	FACTOR_DESENCOLAR     = 4
	PANIC_COLA_VACIA      = "La cola esta vacia"
	INICIO_DEL_ARREGLO    = 0
	COMPARADOR            = 0
)

type fcmpHeap[T comparable] func(T, T) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

/* ----------------------------------- FUNCIONES DE CREACION ----------------------------------- */

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, LARGO_INICIAL), cmp: funcion_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, len(arreglo))
	copy(arr, arreglo)
	heapify(arr, len(arr), funcion_cmp)
	return &heap[T]{datos: arr, cantidad: len(arr), cmp: funcion_cmp}
}

/* ------------------------------------ FUNCIONES AUXILIARES ----------------------------------- */

func swap[T comparable](elemento1, elemento2 *T) {
	*elemento1, *elemento2 = *elemento2, *elemento1
}

func obtenerIndHijoMayor[T comparable](arr *[]T, posHijoIzq, posHijoDer int, tam int, cmp fcmpHeap[T]) int {
	if posHijoDer >= tam || cmp((*arr)[posHijoIzq], (*arr)[posHijoDer]) > COMPARADOR {
		return posHijoIzq
	}
	return posHijoDer
}

func upheap[T comparable](arr *[]T, posElemento int, cmp fcmpHeap[T]) {
	if posElemento == INICIO_DEL_ARREGLO {
		return
	}
	i_padre := (posElemento - 1) / 2
	if cmp((*arr)[i_padre], (*arr)[posElemento]) < COMPARADOR {
		swap(&(*arr)[posElemento], &(*arr)[i_padre])
		upheap(arr, i_padre, cmp)
	}
}

func downheap[T comparable](arr *[]T, posElemento int, tam int, cmp fcmpHeap[T]) {
	if posElemento == tam-1 {
		return
	}
	posHijoIzq := 2*posElemento + 1
	posHijoDer := 2*posElemento + 2
	if posHijoIzq >= tam {
		return
	}
	posHijoMayor := obtenerIndHijoMayor(arr, posHijoIzq, posHijoDer, tam, cmp)
	if cmp((*arr)[posElemento], (*arr)[posHijoMayor]) < COMPARADOR {
		swap(&(*arr)[posElemento], &(*arr)[posHijoMayor])
		downheap(arr, posHijoMayor, tam, cmp)
	}
}

func heapify[T comparable](arr []T, tam int, cmp fcmpHeap[T]) {
	for posElemento := tam / 2; posElemento >= INICIO_DEL_ARREGLO; posElemento-- {
		downheap(&arr, posElemento, tam, cmp)
	}
}

/* ----------------------------------- ORDENAMIENTO HEAPSORT ----------------------------------- */

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)
	heapSort(elementos, len(elementos), funcion_cmp)
}

func heapSort[T comparable](elementos []T, tam int, funcion_cmp func(T, T) int) {
	for finDelArreglo := tam - 1; finDelArreglo >= INICIO_DEL_ARREGLO; finDelArreglo-- {
		swap(&elementos[INICIO_DEL_ARREGLO], &elementos[finDelArreglo])
		downheap(&elementos, INICIO_DEL_ARREGLO, finDelArreglo, funcion_cmp)
	}
}

/* ------------------------------- PRIMITIVAS COLA DE PRIORIDAD -------------------------------- */

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == COMPARADOR
}

func (heap heap[T]) VerMax() T {
	heap.comprobarEstaVacia()
	return heap.datos[INICIO_DEL_ARREGLO]
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) redimensionarHeap(nuevaCap int) {
	arr := make([]T, nuevaCap)
	copy(arr, heap.datos)
	heap.datos = arr
}

func (heap *heap[T]) Encolar(dato T) {
	nuevaCap := cap(heap.datos) * FACTOR_DE_REDIMENSION
	if cap(heap.datos) < LARGO_INICIAL {
		nuevaCap = LARGO_INICIAL
	}
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionarHeap(nuevaCap)
	}

	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	upheap(&heap.datos, heap.cantidad-1, heap.cmp)
}

func (heap *heap[T]) Desencolar() T {
	heap.comprobarEstaVacia()
	elemento := heap.datos[INICIO_DEL_ARREGLO]
	fin_del_arreglo := heap.cantidad - 1
	swap(&heap.datos[INICIO_DEL_ARREGLO], &heap.datos[fin_del_arreglo])
	heap.cantidad--
	downheap(&heap.datos, INICIO_DEL_ARREGLO, heap.cantidad, heap.cmp)

	nuevaCap := cap(heap.datos) / FACTOR_DE_REDIMENSION
	if nuevaCap < LARGO_INICIAL && cap(heap.datos) != LARGO_INICIAL {
		nuevaCap = LARGO_INICIAL
	}
	if heap.cantidad*FACTOR_DESENCOLAR <= cap(heap.datos) && nuevaCap >= LARGO_INICIAL {
		heap.redimensionarHeap(nuevaCap)
	}
	return elemento
}

func (heap *heap[T]) comprobarEstaVacia() {
	if heap.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}
}
