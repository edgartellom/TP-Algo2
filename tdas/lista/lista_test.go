package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}

func TestInsertar(t *testing.T) {
	t.Log("Hacemos pruebas encolando y desencolando elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(7)
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
	require.EqualValues(t, 7, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 7, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.EqualValues(t, 0, lista.Largo())
	// lista.Encolar(6)
	// lista.Encolar(11)
	// require.EqualValues(t, 2, lista.Desencolar())
	// require.EqualValues(t, 7, lista.VerPrimero())
	// lista.Encolar(9)
	// lista.Encolar(5)
	// require.False(t, lista.EstaVacia())
	// require.EqualValues(t, 7, lista.Desencolar())
	// require.EqualValues(t, 4, lista.VerPrimero())
	// require.EqualValues(t, 4, lista.Desencolar())
	// require.EqualValues(t, 6, lista.Desencolar())
	// require.EqualValues(t, 11, lista.Desencolar())
	// require.EqualValues(t, 9, lista.Desencolar())
	// require.EqualValues(t, 5, lista.Desencolar())
	// require.True(t, lista.EstaVacia())
	// require.PanicsWithValue(t, "La cola esta vacia", func() { lista.Desencolar() })
	// require.PanicsWithValue(t, "La cola esta vacia", func() { lista.VerPrimero() })
	// lista.Encolar(13)
	// require.EqualValues(t, 13, lista.VerPrimero())
}

// func TestVolumen(t *testing.T) {
// 	cola := TDALista.CrearListaEnlazada[int]()
// 	elem := 10000
// 	for i := 0; i < elem; i++ {
// 		cola.Encolar(i)
// 		require.EqualValues(t, 0, cola.VerPrimero())
// 	}
// 	for j := 0; j < elem; j++ {
// 		require.EqualValues(t, j, cola.VerPrimero())
// 		cola.Desencolar()
// 	}
// 	require.True(t, cola.EstaVacia())
// }

// func TestBorde(t *testing.T) {
// 	t.Log("Prueba de desapilar al borde de la pila")
// 	cola := TDALista.CrearListaEnlazada[int]()
// 	//Comprobacion con cola recien creada
// 	require.True(t, cola.EstaVacia())
// 	//Comprobacion de acciones invalidas
// 	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
// 	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
// 	//Comprobacion luego de encolar
// 	cola.Encolar(7)
// 	require.False(t, cola.EstaVacia())
// 	cola.Desencolar()
// 	require.True(t, cola.EstaVacia())
// 	//Comprobacion de acciones invalidas
// 	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
// 	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
// }

// func TestTiposDato(t *testing.T) {
// 	//Prueba con enteros
// 	colaEnteros := TDALista.CrearListaEnlazada[int]()
// 	require.True(t, colaEnteros.EstaVacia())
// 	colaEnteros.Encolar(3)
// 	require.EqualValues(t, 3, colaEnteros.VerPrimero())
// 	//Prueba con cadenas
// 	colaCadenas := TDALista.CrearListaEnlazada[string]()
// 	require.True(t, colaCadenas.EstaVacia())
// 	colaCadenas.Encolar("hola")
// 	require.EqualValues(t, "hola", colaCadenas.VerPrimero())
// 	//Prueba con booleanos
// 	pilaBooleanos := TDALista.CrearListaEnlazada[bool]()
// 	require.True(t, pilaBooleanos.EstaVacia())
// 	pilaBooleanos.Encolar(5 == 5.5)
// 	require.EqualValues(t, false, pilaBooleanos.VerPrimero())
// 	//Prueba con flotantes
// 	pilaFlotantes := TDALista.CrearListaEnlazada[float32]()
// 	require.True(t, pilaFlotantes.EstaVacia())
// 	pilaFlotantes.Encolar(3.45)
// 	require.EqualValues(t, 3.45, pilaFlotantes.VerPrimero())
// }
