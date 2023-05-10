package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 1000
const LARGO_TEST = 10
const CANTIDAD_A_AVANZAR = 3

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
}

func TestListaConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var valor int = 9
	lista.InsertarPrimero(valor)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, valor, lista.VerPrimero())
	require.EqualValues(t, valor, lista.VerUltimo())
	require.EqualValues(t, valor, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestInsertarYQuitarPrimeroSimultaneamente(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < LARGO_TEST; i++ {
		lista.InsertarPrimero(i)
		valorFrente := lista.VerPrimero()
		require.EqualValues(t, valorFrente, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestComportamientoAlVaciarLaLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var valorInicial int
	var contador int = 0
	for i := valorInicial; i < LARGO_TEST; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, valorInicial, lista.VerPrimero())
		require.EqualValues(t, i, lista.VerUltimo())
		contador++
	}

	require.False(t, lista.EstaVacia())

	for !lista.EstaVacia() {
		require.EqualValues(t, lista.VerPrimero(), lista.BorrarPrimero())
		contador--
		require.EqualValues(t, contador, lista.Largo())
	}

	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.True(t, lista.EstaVacia())
}

func TestListaDeDiferentesTipos(t *testing.T) {
	var numero int = 333
	listaInts := TDALista.CrearListaEnlazada[int]()
	listaInts.InsertarPrimero(numero)
	require.EqualValues(t, 333, listaInts.VerPrimero())

	var palabra string = "Holis"
	listaStrings := TDALista.CrearListaEnlazada[string]()
	listaStrings.InsertarPrimero(palabra)
	require.EqualValues(t, "Holis", listaStrings.VerPrimero())

	var booleano bool = true
	listaBooleans := TDALista.CrearListaEnlazada[bool]()
	listaBooleans.InsertarPrimero(booleano)
	require.EqualValues(t, true, listaBooleans.VerPrimero())
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var valorEnlistado int
	var valorInicial int

	for i := valorInicial; i < VOLUMEN/2; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, valorInicial, lista.VerPrimero())
	}

	require.False(t, lista.EstaVacia())

	var valorLimite int = VOLUMEN / 4
	for i := 0; i < valorLimite; i++ {
		valorEnlistado = i
		require.EqualValues(t, valorEnlistado, lista.VerPrimero())
		require.EqualValues(t, valorEnlistado, lista.BorrarPrimero())
	}

	require.False(t, lista.EstaVacia())

	valorInicial = valorLimite
	for i := VOLUMEN / 2; i < VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, valorInicial, lista.VerPrimero())
	}

	for i := valorEnlistado + 1; !lista.EstaVacia(); i++ {
		valorEnlistado = i
		require.EqualValues(t, valorEnlistado, lista.VerPrimero())
		require.EqualValues(t, valorEnlistado, lista.BorrarPrimero())
	}

	require.True(t, lista.EstaVacia())
}

// TEST ITERADORES INTERNOS

func TestInteradorInternoEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })

	contador := 0
	contador_ptr := &contador

	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})

	require.EqualValues(t, 0, contador)
}

func TestDeIteradorInternoSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var valorActual int

	var comprobanteDeSumatoria int
	sumatoria := 0
	sum_ptr := &sumatoria

	var comprobanteDeLongitud int
	longitud := 0
	long_ptr := &longitud

	for i := 0; i <= LARGO_TEST*2; i++ {
		comprobanteDeSumatoria += i
		comprobanteDeLongitud += 1
		lista.InsertarUltimo(i)
		i++
	}

	lista.Iterar(func(v int) bool {
		require.EqualValues(t, valorActual, v)
		valorActual += 2
		return true
	})

	lista.Iterar(func(v int) bool {
		*sum_ptr += v
		return true
	})

	lista.Iterar(func(v int) bool {
		*long_ptr += 1
		return true
	})

	require.EqualValues(t, comprobanteDeSumatoria, sumatoria)
	require.EqualValues(t, comprobanteDeLongitud, longitud)
}

func TestIteradorInternoConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var numeroLimite int = LARGO_TEST * 2
	var valorActual int

	var comprobanteDeSumatoria int
	var sumatoria int
	sum_ptr := &sumatoria

	var comprobanteDeLongitud int
	var longitud int
	long_ptr := &longitud

	for i := 0; i <= numeroLimite; i++ {
		if i <= numeroLimite/2 {
			comprobanteDeSumatoria += i
		}
		comprobanteDeLongitud += 1
		lista.InsertarUltimo(i)
		i++
	}

	lista.Iterar(func(v int) bool {
		require.EqualValues(t, valorActual, v)
		valorActual += 2
		return v != numeroLimite/2
	})

	lista.Iterar(func(v int) bool {
		*sum_ptr += v
		return v != numeroLimite/2
	})

	lista.Iterar(func(v int) bool {
		*long_ptr += 1
		return true
	})

	require.EqualValues(t, comprobanteDeSumatoria, sumatoria)
	require.EqualValues(t, comprobanteDeLongitud, longitud)
}

// TEST ITERADORES EXTERNOS

func TestIteradorDeListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestIteradorInsertaEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	for i := 0; i <= LARGO_TEST; i++ {
		iter.Insertar(i)
	}

	for i := LARGO_TEST; i >= 0; i-- {
		require.EqualValues(t, i, iter.VerActual())
		iter.Siguiente()
	}

	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestIteradorInsertandoAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int
	var InsertarValor int
	InsertarValor += CANTIDAD_A_AVANZAR
	lista.InsertarUltimo(InsertarValor)
	largoActual++

	InsertarValor += CANTIDAD_A_AVANZAR
	lista.InsertarUltimo(InsertarValor)
	largoActual++

	InsertarValor += CANTIDAD_A_AVANZAR
	lista.InsertarUltimo(InsertarValor)
	largoActual++

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	InsertarValor += CANTIDAD_A_AVANZAR
	iter.Insertar(InsertarValor)
	largoActual++

	require.EqualValues(t, InsertarValor, lista.VerUltimo())
	require.EqualValues(t, InsertarValor, iter.VerActual())
	require.EqualValues(t, largoActual, lista.Largo())

	iter.Siguiente()
	InsertarValor += CANTIDAD_A_AVANZAR
	iter.Insertar(InsertarValor)
	largoActual++

	require.EqualValues(t, InsertarValor, lista.VerUltimo())
	require.EqualValues(t, InsertarValor, iter.VerActual())
	require.EqualValues(t, largoActual, lista.Largo())

	iter.Siguiente()
	InsertarValor += CANTIDAD_A_AVANZAR
	iter.Insertar(InsertarValor)
	largoActual++
	require.EqualValues(t, InsertarValor, lista.VerUltimo())
	require.EqualValues(t, InsertarValor, iter.VerActual())
	require.EqualValues(t, largoActual, lista.Largo())
}

func TestBorrarPrimeroConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int = LARGO_TEST
	for i := 0; i < LARGO_TEST; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter := lista.Iterador()
	for i := 0; i < LARGO_TEST; i++ {
		require.EqualValues(t, i, iter.Borrar())
		largoActual--
	}

	require.EqualValues(t, largoActual, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestBorrarEnElMedioConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int = LARGO_TEST
	for i := 0; i < LARGO_TEST; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	for i := 2; i < LARGO_TEST; i++ {
		require.EqualValues(t, i, iter.Borrar())
		largoActual--
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter2 := lista.Iterador()
	require.EqualValues(t, 0, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 1, iter2.VerActual())
}

func TestBorrarAnteultimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int = LARGO_TEST
	var ultimoValor int = largoActual - 1
	var anteultimoValor int = ultimoValor - 1
	for i := 0; i < LARGO_TEST; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter := lista.Iterador()

	for iter.VerActual() != anteultimoValor {
		iter.Siguiente()
	}
	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, anteultimoValor, iter.Borrar())
	largoActual--
	require.EqualValues(t, ultimoValor, iter.Borrar())
	largoActual--

	require.EqualValues(t, largoActual, lista.Largo())

	iter2 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func TestBorrarUltimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int = LARGO_TEST
	var ultimoValor int = largoActual - 1
	for i := 0; i < LARGO_TEST; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter := lista.Iterador()

	for iter.VerActual() != ultimoValor {
		iter.Siguiente()
	}
	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter.Borrar())
	largoActual--
	ultimoValor--

	require.EqualValues(t, largoActual, lista.Largo())
	require.EqualValues(t, ultimoValor, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func TestBorrarUltimoVariasVecesConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var largoActual int = LARGO_TEST / 2
	var ultimoValor int = largoActual - 1
	for i := 0; i < largoActual; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, largoActual, lista.Largo())

	iter := lista.Iterador()
	for iter.VerActual() != ultimoValor {
		iter.Siguiente()
	}

	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter.Borrar())
	largoActual--
	ultimoValor--

	require.EqualValues(t, largoActual, lista.Largo())
	require.EqualValues(t, ultimoValor, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

	//
	iter3 := lista.Iterador()
	for iter3.VerActual() != ultimoValor {
		iter3.Siguiente()
	}

	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter3.Borrar())
	largoActual--
	ultimoValor--

	require.EqualValues(t, largoActual, lista.Largo())
	require.EqualValues(t, ultimoValor, lista.VerUltimo())

	iter4 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter4.VerActual())
		iter4.Siguiente()
	}

	//
	iter5 := lista.Iterador()
	for iter5.VerActual() != ultimoValor {
		iter5.Siguiente()
	}

	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter5.Borrar())
	largoActual--
	ultimoValor--

	require.EqualValues(t, largoActual, lista.Largo())
	require.EqualValues(t, ultimoValor, lista.VerUltimo())

	iter6 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter6.VerActual())
		iter6.Siguiente()
	}

	//
	iter7 := lista.Iterador()
	for iter7.VerActual() != ultimoValor {
		iter7.Siguiente()
	}

	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter7.Borrar())
	largoActual--
	ultimoValor--

	require.EqualValues(t, largoActual, lista.Largo())
	require.EqualValues(t, ultimoValor, lista.VerPrimero())
	require.EqualValues(t, ultimoValor, lista.VerUltimo())

	iter8 := lista.Iterador()
	for i := 0; i < largoActual; i++ {
		require.EqualValues(t, i, iter8.VerActual())
		iter8.Siguiente()
	}

	//
	iter9 := lista.Iterador()
	for iter9.VerActual() != ultimoValor {
		iter9.Siguiente()
	}

	require.EqualValues(t, largoActual, lista.Largo())

	require.EqualValues(t, ultimoValor, iter9.Borrar())
	largoActual--

	require.EqualValues(t, largoActual, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })

	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}
