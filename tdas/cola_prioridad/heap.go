package cola_prioridad

const (
	LARGO_INICIAL         = 10
	FACTOR_DE_REDIMENSION = 2
	FACTOR_DESENCOLAR     = 4
	PANIC_COLA_VACIA      = "La cola esta vacia"
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
	heap.datos = *heapify(&arreglo, heap.cantidad, heap.cmp)
	return heap
}

func heapify[T comparable](arr *[]T, tam int, cmp fcmpHeap[T]) *[]T {
	for i_elemento := tam - 1; i_elemento > -1; i_elemento-- {
		downheap(arr, i_elemento, tam, cmp)
	}
	return arr
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	elementos = *heapify(&elementos, len(elementos), funcion_cmp)
	heapSort(elementos, len(elementos), funcion_cmp)
}

func heapSort[T comparable](elementos []T, tam int, funcion_cmp func(T, T) int) {
	if tam == 0 {
		return
	}
	swap(&elementos[0], &elementos[tam-1])
	downheap(&elementos, 0, tam-1, funcion_cmp)
	heapSort(elementos, tam-1, funcion_cmp)
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap heap[T]) VerMax() T {
	heap.comprobarEstaVacia()
	return heap.datos[0]
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func upheap[T comparable](arr *[]T, i_elemento int, cmp fcmpHeap[T]) {
	if i_elemento == 0 {
		return
	}
	i_padre := (i_elemento - 1) / 2
	if cmp((*arr)[i_padre], (*arr)[i_elemento]) < 0 {
		swap(&(*arr)[i_elemento], &(*arr)[i_padre])
		upheap(arr, i_padre, cmp)
	}
}

func obtenerIndHijoMayor[T comparable](arr *[]T, i_h_izq, i_h_der int, tam int, cmp fcmpHeap[T]) int {
	if i_h_izq >= tam {
		return i_h_der
	}
	if i_h_der >= tam {
		return i_h_izq
	}
	if cmp((*arr)[i_h_izq], (*arr)[i_h_der]) > 0 {
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
	if cmp((*arr)[i_elemento], (*arr)[i_h_mayor]) < 0 {
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
	elemento := heap.datos[0]
	swap(&heap.datos[0], &heap.datos[heap.cantidad-1])
	heap.cantidad--
	downheap(&heap.datos, 0, heap.cantidad, heap.cmp)

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
