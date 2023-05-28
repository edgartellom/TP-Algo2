package cola_prioridad

const (
	PANIC_COLA_VACIA      = "La cola esta vacia"
	LARGO_INICIAL         = 10
	FACTOR_DE_REDIMENSION = 2
	FACTOR_DESENCOLAR     = 4
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
	elementos := make([]T, len(arreglo))
	copy(elementos, arreglo)
	heapify(elementos, len(arreglo), funcion_cmp)
	return &heap[T]{datos: elementos, cantidad: len(arreglo), cmp: funcion_cmp}
}

/* ------------------------------------ FUNCIONES AUXILIARES ----------------------------------- */

func swap[T comparable](elemento1, elemento2 *T) {
	*elemento1, *elemento2 = *elemento2, *elemento1
}

func obtenerPosHijoMayor[T comparable](elementos *[]T, posHijoIzq, posHijoDer int, tam int, cmp fcmpHeap[T]) int {
	if posHijoDer >= tam || cmp((*elementos)[posHijoIzq], (*elementos)[posHijoDer]) > COMPARADOR {
		return posHijoIzq
	}
	return posHijoDer
}

func upHeap[T comparable](elementos *[]T, posActual int, cmp fcmpHeap[T]) {
	if posActual == INICIO_DEL_ARREGLO {
		return
	}
	posPadre := (posActual - 1) / 2
	if cmp((*elementos)[posPadre], (*elementos)[posActual]) < COMPARADOR {
		swap(&(*elementos)[posActual], &(*elementos)[posPadre])
		upHeap(elementos, posPadre, cmp)
	}
}

func downHeap[T comparable](elementos *[]T, posActual, tam int, cmp fcmpHeap[T]) {
	if posActual >= tam-1 {
		return
	}
	posHijoIzq, posHijoDer := (2*posActual + 1), (2*posActual + 2)

	if posHijoIzq >= tam {
		return
	}
	posHijoMayor := obtenerPosHijoMayor(elementos, posHijoIzq, posHijoDer, tam, cmp)
	if cmp((*elementos)[posActual], (*elementos)[posHijoMayor]) < COMPARADOR {
		swap(&((*elementos)[posHijoMayor]), &((*elementos)[posActual]))
		downHeap(elementos, posHijoMayor, tam, cmp)
	}
}

func heapify[T comparable](arreglo []T, tam int, cmp fcmpHeap[T]) {
	for posActual := tam - 1; posActual >= INICIO_DEL_ARREGLO; posActual-- {
		downHeap(&arreglo, posActual, tam, cmp)
	}
}

/* ----------------------------------- ORDENAMIENTO HEAPSORT ----------------------------------- */

func HeapSort[T comparable](arreglo []T, funcion_cmp func(T, T) int) {
	heapify(arreglo, len(arreglo), funcion_cmp)
	heapSort(arreglo, len(arreglo), funcion_cmp)
}

func heapSort[T comparable](elementos []T, tam int, funcion_cmp func(T, T) int) {
	if tam == 0 {
		return
	}
	inicioDelArreglo, finDelArreglo := INICIO_DEL_ARREGLO, tam-1
	swap(&elementos[inicioDelArreglo], &elementos[finDelArreglo])
	downHeap(&elementos, inicioDelArreglo, finDelArreglo, funcion_cmp)
	heapSort(elementos, finDelArreglo, funcion_cmp)
}

/* ------------------------------- PRIMITIVAS COLA DE PRIORIDAD -------------------------------- */

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap heap[T]) VerMax() T {
	heap.comprobarEstaVacia()
	return heap.datos[INICIO_DEL_ARREGLO]
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == COMPARADOR
}

func (heap *heap[T]) Encolar(dato T) {
	if cap(heap.datos) < LARGO_INICIAL {
		heap.redimensionarHeap(LARGO_INICIAL)
	} else if heap.cantidad == cap(heap.datos) {
		heap.redimensionarHeap(cap(heap.datos) * FACTOR_DE_REDIMENSION)
	}

	heap.datos[heap.cantidad] = dato
	heap.cantidad++
	upHeap(&heap.datos, heap.cantidad-1, heap.cmp)
}

func (heap *heap[T]) Desencolar() T {
	heap.comprobarEstaVacia()
	elemento := heap.datos[INICIO_DEL_ARREGLO]
	swap(&((*heap).datos[INICIO_DEL_ARREGLO]), &((*heap).datos[heap.cantidad-1]))
	heap.cantidad--

	downHeap(&heap.datos, INICIO_DEL_ARREGLO, heap.cantidad, heap.cmp)

	nuevaCap := cap(heap.datos) / FACTOR_DE_REDIMENSION
	if nuevaCap < LARGO_INICIAL && cap(heap.datos) != LARGO_INICIAL {
		nuevaCap = LARGO_INICIAL
	}
	if heap.cantidad*FACTOR_DESENCOLAR <= cap(heap.datos) && nuevaCap >= LARGO_INICIAL {
		heap.redimensionarHeap(nuevaCap)
	}

	return elemento
}

func (heap *heap[T]) redimensionarHeap(nuevaCapacidad int) {
	nuevoArreglo := make([]T, nuevaCapacidad)
	copy(nuevoArreglo, heap.datos)
	heap.datos = nuevoArreglo
}

func (heap *heap[T]) comprobarEstaVacia() {
	if heap.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}
}
