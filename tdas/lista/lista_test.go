package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// TEST LISTA ENLAZADA
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
}

func TestInsertar(t *testing.T) {
	t.Log("Hacemos pruebas insertando y borrando elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, lista.VerPrimero())
	lista.InsertarUltimo(7)
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
	require.EqualValues(t, 7, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 7, lista.BorrarPrimero())
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elem := 10000
	for i := 0; i < elem; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
	}
	for j := elem - 1; j >= 0; j-- {
		require.EqualValues(t, j, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestBorde(t *testing.T) {
	t.Log("Prueba de borrar al borde de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	//Comprobacion con lista recien creada
	require.True(t, lista.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
	//Comprobacion luego de insertar
	lista.InsertarUltimo(7)
	require.False(t, lista.EstaVacia())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.PANIC_LISTA_VACIA, func() { lista.VerUltimo() })
}

func TestTiposDato(t *testing.T) {
	t.Log("Prueba con distintos tipos de datos")
	//Prueba con enteros
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	require.True(t, listaEnteros.EstaVacia())
	listaEnteros.InsertarPrimero(3)
	require.EqualValues(t, 3, listaEnteros.VerPrimero())
	//Prueba con cadenas
	listaCadenas := TDALista.CrearListaEnlazada[string]()
	require.True(t, listaCadenas.EstaVacia())
	listaCadenas.InsertarUltimo("hola")
	require.EqualValues(t, "hola", listaCadenas.VerPrimero())
	//Prueba con booleanos
	listaBooleanos := TDALista.CrearListaEnlazada[bool]()
	require.True(t, listaBooleanos.EstaVacia())
	listaBooleanos.InsertarPrimero(5 == 5.5)
	require.EqualValues(t, false, listaBooleanos.VerPrimero())
	//Prueba con flotantes
	listaFlotantes := TDALista.CrearListaEnlazada[float32]()
	require.True(t, listaFlotantes.EstaVacia())
	listaFlotantes.InsertarPrimero(3.45)
	require.EqualValues(t, 3.45, listaFlotantes.VerPrimero())
}

// TEST ITERADOR EXTERNO
func TestAlPrincipio(t *testing.T) {
	t.Log("Prueba al principio de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
	iter.Insertar(8)
	iter.Insertar(3)
	iter.Insertar(5)
	require.EqualValues(t, 5, iter.VerActual())
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 8, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 5, iter.Borrar())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 8, lista.VerUltimo())
	iter.Siguiente()
	require.EqualValues(t, 8, iter.VerActual())
	iter.Siguiente()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.Borrar() })
}

func TestAlFinal(t *testing.T) {
	t.Log("Prueba al final de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	iter.Insertar(4)
	require.EqualValues(t, 4, iter.VerActual())
	iter.Siguiente()
	require.PanicsWithValue(t, TDALista.PANIC_ITERADOR, func() { iter.VerActual() })
	iter.Insertar(2)
	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, lista.VerUltimo())
	iter.Siguiente()
	iter.Insertar(7)
	iter.Borrar()
	require.EqualValues(t, 7, iter.VerActual())
	require.EqualValues(t, 7, lista.VerUltimo())
	require.EqualValues(t, 4, lista.VerPrimero())
}

func TestAlMedio(t *testing.T) {
	t.Log("Prueba al medio de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(11)
	iter.Insertar(9)
	iter.Insertar(6)
	require.EqualValues(t, 6, iter.VerActual())
	iter.Siguiente()
	iter.Insertar(3)
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, iter.Borrar())
	require.EqualValues(t, 6, iter.VerActual())
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
