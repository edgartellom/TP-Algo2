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

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, LARGO_INICIAL), cmp: funcion_cmp}
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cantidad = len(arreglo)
	heap.cmp = funcion_cmp
	heapify(&arreglo, heap.cantidad, heap.cmp)
	heap.datos = arreglo
	return heap
}

func heapify[T comparable](arr *[]T, tam int, cmp fcmpHeap[T]) {
	for i_elemento := tam - 1; i_elemento >= INICIO_DEL_ARREGLO; i_elemento-- {
		downheap(arr, i_elemento, tam, cmp)
	}
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	heapify(&elementos, len(elementos), funcion_cmp)
	heapSort(elementos, len(elementos), funcion_cmp)
}

func heapSort[T comparable](elementos []T, tam int, funcion_cmp func(T, T) int) {
	if tam == 0 {
		return
	}
	fin_del_arreglo := tam - 1
	swap(&elementos[INICIO_DEL_ARREGLO], &elementos[fin_del_arreglo])
	downheap(&elementos, INICIO_DEL_ARREGLO, fin_del_arreglo, funcion_cmp)
	heapSort(elementos, fin_del_arreglo, funcion_cmp)
}

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

func upheap[T comparable](arr *[]T, i_elemento int, cmp fcmpHeap[T]) {
	if i_elemento == INICIO_DEL_ARREGLO {
		return
	}
	i_padre := (i_elemento - 1) / 2
	if cmp((*arr)[i_padre], (*arr)[i_elemento]) < COMPARADOR {
		swap(&(*arr)[i_elemento], &(*arr)[i_padre])
		upheap(arr, i_padre, cmp)
	}
}

func obtenerIndHijoMayor[T comparable](arr *[]T, i_h_izq, i_h_der int, tam int, cmp fcmpHeap[T]) int {
	if i_h_der >= tam || cmp((*arr)[i_h_izq], (*arr)[i_h_der]) > COMPARADOR {
		return i_h_izq
	}
	return i_h_der
}

func swap[T comparable](x, y *T) {
	*x, *y = *y, *x
}

func downheap[T comparable](arr *[]T, i_elemento int, tam int, cmp fcmpHeap[T]) {
	if i_elemento == tam-1 {
		return
	}
	i_h_izq := 2*i_elemento + 1
	i_h_der := 2*i_elemento + 2
	if i_h_izq >= tam && i_h_der >= tam {
		return
	}
	i_h_mayor := obtenerIndHijoMayor(arr, i_h_izq, i_h_der, tam, cmp)
	if cmp((*arr)[i_elemento], (*arr)[i_h_mayor]) < COMPARADOR {
		swap(&(*arr)[i_elemento], &(*arr)[i_h_mayor])
		downheap(arr, i_h_mayor, tam, cmp)
	}
}

func (heap *heap[T]) redimensionarHeap(nuevaCap int) {
	arr := make([]T, nuevaCap)
	copy(arr, heap.datos)
	heap.datos = arr
}

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensionarHeap(cap(heap.datos) * FACTOR_DE_REDIMENSION)
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
	if heap.cantidad*FACTOR_DESENCOLAR <= cap(heap.datos) {
		if nuevaCap < LARGO_INICIAL {
			nuevaCap = LARGO_INICIAL
		}
		heap.redimensionarHeap(nuevaCap)
	}
	return elemento
}

func (heap *heap[T]) comprobarEstaVacia() {
	if heap.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}
}
