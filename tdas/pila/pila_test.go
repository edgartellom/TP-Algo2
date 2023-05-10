package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	elem := 10000
	for i := 0; i < elem; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for j := elem - 1; j >= 0; j-- {
		require.EqualValues(t, j, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestApilar(t *testing.T) {
	t.Log("Hacemos pruebas apilando y desapilando elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(2)
	pila.Apilar(7)
	pila.Apilar(4)
	require.EqualValues(t, 4, pila.VerTope())
	pila.Apilar(6)
	pila.Apilar(11)
	require.EqualValues(t, 11, pila.Desapilar())
	require.EqualValues(t, 6, pila.VerTope())
	pila.Apilar(9)
	pila.Apilar(5)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 9, pila.VerTope())
	require.EqualValues(t, 9, pila.Desapilar())
	require.EqualValues(t, 6, pila.Desapilar())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	pila.Apilar(13)
	require.EqualValues(t, 13, pila.VerTope())
}

func TestBorde(t *testing.T) {
	t.Log("Prueba de desapilar al borde de la pila")
	pila := TDAPila.CrearPilaDinamica[int]()
	//Comprobacion con pila recien creada
	require.True(t, pila.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	//Comprobacion luego de apilar
	pila.Apilar(7)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	//Comprobacion de acciones invalidas
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestTiposDato(t *testing.T) {
	//Prueba con enteros
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaEnteros.EstaVacia())
	pilaEnteros.Apilar(3)
	require.EqualValues(t, 3, pilaEnteros.VerTope())
	//Prueba con cadenas
	pilaCadenas := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaCadenas.EstaVacia())
	pilaCadenas.Apilar("hola")
	require.EqualValues(t, "hola", pilaCadenas.VerTope())
	//Prueba con booleanos
	pilaBooleanos := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pilaBooleanos.EstaVacia())
	pilaBooleanos.Apilar(5 == 5.5)
	require.EqualValues(t, false, pilaBooleanos.VerTope())
	//Prueba con flotantes
	pilaFlotantes := TDAPila.CrearPilaDinamica[float32]()
	require.True(t, pilaFlotantes.EstaVacia())
	pilaFlotantes.Apilar(3.45)
	require.EqualValues(t, 3.45, pilaFlotantes.VerTope())
}
