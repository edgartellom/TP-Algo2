package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 1000

func TestPilaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
}

func TestPilaConUnElemento(t *testing.T) {
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
	for i := 0; i < 10; i++ {
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
	for i := valorInicial; i < 10; i++ {
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

	var valorInicial int = 0
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

func TestDeIteradorInternoSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	var valorActual int

	var comprobanteDeSumatoria int
	sumatoria := 0
	sum_ptr := &sumatoria

	var comprobanteDeLongitud int
	longitud := 0
	long_ptr := &longitud

	for i := 0; i <= 20; i++ {
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
	var valorActual int

	var comprobanteDeSumatoria int
	sumatoria := 0
	sum_ptr := &sumatoria

	var comprobanteDeLongitud int
	longitud := 0
	long_ptr := &longitud

	for i := 0; i <= 20; i++ {
		if i <= 10 {
			comprobanteDeSumatoria += i
		}
		comprobanteDeLongitud += 1
		lista.InsertarUltimo(i)
		i++
	}

	lista.Iterar(func(v int) bool {
		require.EqualValues(t, valorActual, v)
		valorActual += 2
		return v != 10
	})

	lista.Iterar(func(v int) bool {
		*sum_ptr += v
		return v != 10
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

	for i := 0; i < 5; i++ {
		iter.Insertar(i)
	}

	for i := 4; i >= 0; i-- {
		require.EqualValues(t, i, iter.VerActual())
		iter.Siguiente()
	}

	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestIteradorInsertandoAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(9)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(12)
	require.EqualValues(t, 12, lista.VerUltimo())
	require.EqualValues(t, 12, iter.VerActual())
	require.EqualValues(t, 4, lista.Largo())

	iter.Siguiente()
	iter.Insertar(15)
	require.EqualValues(t, 15, lista.VerUltimo())
	require.EqualValues(t, 15, iter.VerActual())
	require.EqualValues(t, 5, lista.Largo())

	iter.Siguiente()
	iter.Insertar(18)
	require.EqualValues(t, 18, lista.VerUltimo())
	require.EqualValues(t, 18, iter.VerActual())
	require.EqualValues(t, 6, lista.Largo())
}

func TestBorrarPrimeroConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 10, lista.Largo())

	iter := lista.Iterador()
	for i := 0; i < 10; i++ {
		require.EqualValues(t, i, iter.Borrar())
	}

	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestBorrarEnElMedioConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 10, lista.Largo())

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()

	for i := 2; i < 10; i++ {
		require.EqualValues(t, i, iter.Borrar())
	}
	require.EqualValues(t, 2, lista.Largo())

	iter2 := lista.Iterador()
	require.EqualValues(t, 0, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 1, iter2.VerActual())
}

func TestBorrarAnteultimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 10, lista.Largo())

	iter := lista.Iterador()

	for iter.VerActual() != 8 {
		iter.Siguiente()
	}
	require.EqualValues(t, 10, lista.Largo())

	require.EqualValues(t, 8, iter.Borrar())
	require.EqualValues(t, 9, iter.Borrar())

	require.EqualValues(t, 8, lista.Largo())

	iter2 := lista.Iterador()
	for i := 0; i < 8; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func TestBorrarUltimoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 10, lista.Largo())

	iter := lista.Iterador()

	for iter.VerActual() != 9 {
		iter.Siguiente()
	}
	require.EqualValues(t, 10, lista.Largo())

	require.EqualValues(t, 9, iter.Borrar())

	require.EqualValues(t, 9, lista.Largo())
	require.EqualValues(t, 8, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := 0; i < 9; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

func TestBorrarUltimoVariasVecesConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i)
	}
	require.EqualValues(t, 5, lista.Largo())

	iter := lista.Iterador()
	for iter.VerActual() != 4 {
		iter.Siguiente()
	}

	require.EqualValues(t, 5, lista.Largo())

	require.EqualValues(t, 4, iter.Borrar())

	require.EqualValues(t, 4, lista.Largo())
	require.EqualValues(t, 3, lista.VerUltimo())

	iter2 := lista.Iterador()
	for i := 0; i < 4; i++ {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}

	//
	iter3 := lista.Iterador()
	for iter3.VerActual() != 3 {
		iter3.Siguiente()
	}

	require.EqualValues(t, 4, lista.Largo())

	require.EqualValues(t, 3, iter3.Borrar())

	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 2, lista.VerUltimo())

	iter4 := lista.Iterador()
	for i := 0; i < 3; i++ {
		require.EqualValues(t, i, iter4.VerActual())
		iter4.Siguiente()
	}

	//
	iter5 := lista.Iterador()
	for iter5.VerActual() != 2 {
		iter5.Siguiente()
	}

	require.EqualValues(t, 3, lista.Largo())

	require.EqualValues(t, 2, iter5.Borrar())

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 1, lista.VerUltimo())

	iter6 := lista.Iterador()
	for i := 0; i < 2; i++ {
		require.EqualValues(t, i, iter6.VerActual())
		iter6.Siguiente()
	}

	//
	iter7 := lista.Iterador()
	for iter7.VerActual() != 1 {
		iter7.Siguiente()
	}

	require.EqualValues(t, 2, lista.Largo())

	require.EqualValues(t, 1, iter7.Borrar())

	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())

	iter8 := lista.Iterador()
	for i := 0; i < 1; i++ {
		require.EqualValues(t, i, iter8.VerActual())
		iter8.Siguiente()
	}

	//
	iter9 := lista.Iterador()
	for iter9.VerActual() != 0 {
		iter9.Siguiente()
	}

	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, 0, iter9.Borrar())

	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })

	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestDelPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })

	iter.Insertar(9)
	iter.Insertar(6)
	iter.Insertar(3)

	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 3, iter.Borrar())

	require.EqualValues(t, 6, lista.VerPrimero())
	require.EqualValues(t, 6, iter.VerActual())
	require.EqualValues(t, 2, lista.Largo())

	iter.Siguiente()

	require.EqualValues(t, 9, iter.VerActual())

	iter.Siguiente()

	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}
