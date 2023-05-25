package cola_prioridad_test

import (
	"fmt"
	"strings"

	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAM_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func func_cmp_int(a, b int) int {
	return a - b
}

func func_cmp_str(a, b string) int {
	return strings.Compare(a, b)
}

func TestHeapVacio(t *testing.T) {
	t.Log("Comprueba que Heap vacio no tiene elementos")
	heap := TDAHeap.CrearHeap(func_cmp_str)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Heap con un elemento tiene ese elemento, unicamente")
	heap := TDAHeap.CrearHeap(func_cmp_str)
	heap.Encolar("A")
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "A", heap.VerMax())
	require.EqualValues(t, "A", heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapEncolar(t *testing.T) {
	t.Log("Encola algunos pocos elementos en el heap, y se comprueba que en todo momento funciona acorde")
	arr := []int{6, 4, 2, 6, 5, 1, 0, 9}

	heap := TDAHeap.CrearHeap(func_cmp_int)

	heap.Encolar(arr[0])
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[1])
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[2])
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[3])
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[4])
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[5])
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[6])
	require.EqualValues(t, 7, heap.Cantidad())
	require.EqualValues(t, 6, heap.VerMax())
	heap.Encolar(arr[7])
	require.EqualValues(t, 8, heap.Cantidad())
	require.EqualValues(t, 9, heap.VerMax())
}

func TestHeapDesencolar(t *testing.T) {
	t.Log("Encola algunos pocos elementos en el heap, y los desencola, revisando que en todo momento " +
		"el heap se comporte de manera adecuada")
	arr := []int{6, 4, 2, 6, 5, 1, 0, 9}
	heap := TDAHeap.CrearHeap(func_cmp_int)

	for i, elem := range arr {
		heap.Encolar(elem)
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	require.EqualValues(t, 9, heap.Desencolar())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 4, heap.Desencolar())
	require.EqualValues(t, 2, heap.Desencolar())
	require.EqualValues(t, 1, heap.Desencolar())
	require.EqualValues(t, 0, heap.Desencolar())
}

func TestConStrings(t *testing.T) {
	t.Log("Valida que no solo funcione con enteros")
	heap := TDAHeap.CrearHeap(func_cmp_str)
	arr := []string{"Gatito", "Perro", "Loro", "Sapo", "Raton", "Leon", "Vaca", "Burro"}

	for i, elem := range arr {
		heap.Encolar(elem)
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	require.EqualValues(t, "Vaca", heap.Desencolar())
	require.EqualValues(t, "Sapo", heap.Desencolar())
	require.EqualValues(t, "Raton", heap.Desencolar())
	require.EqualValues(t, "Perro", heap.Desencolar())
	require.EqualValues(t, "Loro", heap.Desencolar())
	require.EqualValues(t, "Leon", heap.Desencolar())
	require.EqualValues(t, "Gatito", heap.Desencolar())
	require.EqualValues(t, "Burro", heap.Desencolar())
}

func TestHeapify(t *testing.T) {
	t.Log("Valida que se cree un heap apartir de un arreglo y que funcione correctamente")
	arr := []int{6, 4, 2, 6, 5, 1, 0, 9}

	heap := TDAHeap.CrearHeapArr(arr, func_cmp_int)

	require.EqualValues(t, 9, heap.VerMax())
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, heap.VerMax(), heap.Desencolar())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	t.Log("Valida que HeapSort ordene el arreglo correctamente y sea in-place")
	arr := []int{6, 4, 2, 6, 5, 1, 0, 9}

	TDAHeap.HeapSort(arr, func_cmp_int)
	require.EqualValues(t, []int{0, 1, 2, 4, 5, 6, 6, 9}, arr)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	heap := TDAHeap.CrearHeap(func_cmp_int)

	/* Inserta 'n' elementos en el heap */
	for i := 0; i < n; i++ {
		heap.Encolar(i)
	}

	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que VerMax y Desencolar devuelva los valores correctos */
	ok := true
	for i := n - 1; i >= 0; i-- {
		ok = heap.VerMax() == i
		if !ok {
			break
		}
		ok = heap.Desencolar() == i
		if !ok {
			break
		}
	}

	require.True(b, ok, "VerMax y Desencolar con muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, heap.Cantidad(), "La cantidad de elementos es incorrecta")

}

func BenchmarkHeap(b *testing.B) {
	b.Log("Prueba de stress del Heap. Prueba encolando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos ver maximo y que al desencolar todos los elementos generados, " +
		"devuelva el elemento correspondiente sin problemas")
	for _, n := range TAM_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}
