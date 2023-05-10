package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestEncolar(t *testing.T) {
	t.Log("Hacemos pruebas encolando y desencolando elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(2)
	cola.Encolar(7)
	cola.Encolar(4)
	require.EqualValues(t, 2, cola.VerPrimero())
	cola.Encolar(6)
	cola.Encolar(11)
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 7, cola.VerPrimero())
	cola.Encolar(9)
	cola.Encolar(5)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 7, cola.Desencolar())
	require.EqualValues(t, 4, cola.VerPrimero())
	require.EqualValues(t, 4, cola.Desencolar())
	require.EqualValues(t, 6, cola.Desencolar())
	require.EqualValues(t, 11, cola.Desencolar())
	require.EqualValues(t, 9, cola.Desencolar())
	require.EqualValues(t, 5, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar(13)
	require.EqualValues(t, 13, cola.VerPrimero())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	elem := 10000
	for i := 0; i < elem; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for j := 0; j < elem; j++ {
		require.EqualValues(t, j, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}

func TestBorde(t *testing.T) {
	t.Log("Prueba de desencolar al borde de la cola")
	cola := TDACola.CrearColaEnlazada[int]()
	//Comprobacion con cola recien creada
	require.True(t, cola.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	//Comprobacion luego de encolar
	cola.Encolar(7)
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestTiposDato(t *testing.T) {
	//Prueba con enteros
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaEnteros.EstaVacia())
	colaEnteros.Encolar(3)
	require.EqualValues(t, 3, colaEnteros.VerPrimero())
	//Prueba con cadenas
	colaCadenas := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaCadenas.EstaVacia())
	colaCadenas.Encolar("hola")
	require.EqualValues(t, "hola", colaCadenas.VerPrimero())
	//Prueba con booleanos
	colaBooleanos := TDACola.CrearColaEnlazada[bool]()
	require.True(t, colaBooleanos.EstaVacia())
	colaBooleanos.Encolar(5 == 5.5)
	require.EqualValues(t, false, colaBooleanos.VerPrimero())
	//Prueba con flotantes
	colaFlotantes := TDACola.CrearColaEnlazada[float32]()
	require.True(t, colaFlotantes.EstaVacia())
	colaFlotantes.Encolar(3.45)
	require.EqualValues(t, 3.45, colaFlotantes.VerPrimero())
}
