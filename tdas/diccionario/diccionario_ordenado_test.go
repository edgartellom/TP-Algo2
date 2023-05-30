package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDAAbb "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	TAM_VOLUMEN        = []int{12500, 25000, 50000, 100000, 200000, 400000}
	CLAVES_ORDENADAS_1 = []int{1, 4, 6, 8, 10, 11, 13, 14, 15, 16}
	CLAVES_ORDENADAS_2 = []int{1, 4, 6, 8, 10, 11, 13, 16, 20, 24}
	VALORES_ORDENADOS  = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
)

func funcion_cmp_int(a, b int) int {
	return a - b
}

func funcion_cmp_str(a, b string) int {
	return strings.Compare(a, b)
}

func TestAbbVacio(t *testing.T) {
	t.Log("Comprueba que Abb vacio no tiene claves")
	abb := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { abb.Obtener("A") })
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { abb.Borrar("A") })
}

func TestAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Abb vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Obtener("") })
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Borrar("") })

	dicNum := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Abb con un elemento tiene esa Clave, unicamente")
	dic := TDAAbb.CrearABB[string, int](funcion_cmp_str)
	clave := "A"
	clave2 := "B"
	valor := 10

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.False(t, dic.Pertenece(clave2))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Obtener(clave2) })
}

func TestConClaveNumerica(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClavesVacias(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValoresNulos(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDAAbb.CrearABB[string, *int](funcion_cmp_str)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestAbbGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el Abb, y se comprueba que en todo momento funciona acorde")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatos(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatosHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDAAbb.CrearABB[int, int](funcion_cmp_int)
	var cantidad int = 500
	arr := make([]int, cantidad)

	for i := 0; i < cantidad; i++ {
		arr[i] = i
	}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	for i := 0; i < cantidad; i++ {
		dic.Guardar(arr[i], arr[i])
	}
	for i := 0; i < cantidad; i++ {
		dic.Guardar(arr[i], 2*arr[i])
	}
	ok := true
	for i := 0; i < cantidad && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el Abb, y se los borra, revisando que en todo momento " +
		"el Abb se comporte de manera adecuada")
	clave1, valor1 := "Gato", "miau"
	clave2, valor2 := "Perro", "guau"
	clave3, valor3 := "Vaca", "moo"
	claves, valores := []string{clave1, clave2, clave3}, []string{valor1, valor2, valor3}
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[1]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, TDAAbb.PANIC_NO_PERTENECE, func() { dic.Obtener(claves[1]) })
}

func comprobarClave[T comparable](clave T, claves []T) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestAbbIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDAAbb.CrearABB[string, *int](funcion_cmp_str)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, comprobarClave(cs[0], claves))
	require.NotEqualValues(t, -1, comprobarClave(cs[1], claves))
	require.NotEqualValues(t, -1, comprobarClave(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestAbbIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	claves := []string{"Perro", "Gato", "Vaca", "Burrito", "Hamster"}
	valores := []int{6, 2, 3, 4, 5}

	dic := TDAAbb.CrearABB[string, int](funcion_cmp_str)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestAbbIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	claves := []string{"Elefante", "Perro", "Gato", "Vaca", "Burrito", "Hamster"}
	valores := []int{7, 6, 2, 3, 4, 5}

	dic := TDAAbb.CrearABB[string, int](funcion_cmp_str)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])
	dic.Guardar(claves[5], valores[5])

	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenAbb(b *testing.B, n int) {
	dic := TDAAbb.CrearABB[string, int](funcion_cmp_str)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})

	for i := len(claves) - 1; i >= 0; i-- {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario Ordenado. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAM_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenAbb(b, n)
			}
		})
	}
}

func TestIteradorExternoEnAbbVacio(t *testing.T) {
	t.Log("Iterar sobre Abb vacio es simplemente tenerlo al final")
	dic := TDAAbb.CrearABB[string, int](funcion_cmp_str)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.Siguiente() })
}

func TestAbbIteradorExterno(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario Ordenado, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, comprobarClave(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, comprobarClave(segundo, claves))
	require.EqualValues(t, valores[comprobarClave(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, comprobarClave(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.Siguiente() })
}

func TestIteradorExternoNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, comprobarClave(primero, claves))
	require.NotEqualValues(t, -1, comprobarClave(segundo, claves))
	require.NotEqualValues(t, -1, comprobarClave(tercero, claves))
}

func TestIterardorExternoTrasBorrados(t *testing.T) {
	clave1, valor1 := "Perro", "guau"
	clave2, valor2 := "Gato", "miau"
	clave3, valor3 := "Vaca", "moo"

	dic := TDAAbb.CrearABB[string, string](funcion_cmp_str)
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	require.EqualValues(t, valor1, dic.Borrar(clave1))
	require.EqualValues(t, valor2, dic.Borrar(clave2))
	require.EqualValues(t, valor3, dic.Borrar(clave3))

	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.Siguiente() })

	dic.Guardar(clave1, "")

	iter = dic.Iterador()
	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDAAbb.PANIC_ITERADOR, func() { iter.Siguiente() })
}

func ejecutarPruebasVolumenIteradorExterno(b *testing.B, n int) {
	dic := TDAAbb.CrearABB[string, *int](funcion_cmp_str)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	rand.Shuffle(len(claves), func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorExterno(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario Ordenado. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAM_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorExterno(b, n)
			}
		})
	}
}

func TestVolumenIteradorConCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDAAbb.CrearABB[int, *int](funcion_cmp_int)
	cantidad := TAM_VOLUMEN[3] / 10
	elementos := make([]int, cantidad)
	claveDeCorte := elementos[cantidad/2]

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < cantidad; i++ {
		elementos[i] = i
	}

	rand.Shuffle(cantidad, func(i, j int) {
		elementos[i], elementos[j] = elementos[j], elementos[i]
	})

	for e := range elementos {
		dic.Guardar(e, nil)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(clave int, _ *int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if clave == claveDeCorte {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración se corte")
}

func compararArreglos[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func guardarClavesIntValorIterExterno[K int, V any](dic TDAAbb.DiccionarioOrdenado[K, V], desde *K, hasta *K) (clavesRes []K, valoresRes []V) {
	var i int
	for iter := dic.IteradorRango(desde, hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		clavesRes = append(clavesRes, clave)
		valoresRes = append(valoresRes, valor)
		i++
	}
	return clavesRes, valoresRes
}

func TestIteradorExternoRangosExcedentes(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	cs := []int{6, 1, 15, 10, 16, 4, 8, 13, 11, 14}
	vs := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(cs); i++ {
		dic.Guardar(cs[i], vs[i])
	}

	require.EqualValues(t, len(cs), dic.Cantidad())

	a, b := -9, 99
	newCsE, newVsE := guardarClavesIntValorIterExterno(dic, &a, &b)

	require.True(t, compararArreglos(CLAVES_ORDENADAS_1, newCsE))
	require.True(t, compararArreglos(VALORES_ORDENADOS, newVsE))
}

func TestIteradorExternoRangosAcotados(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	cs := []int{6, 1, 15, 10, 16, 4, 8, 13, 11, 14}
	vs := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(cs); i++ {
		dic.Guardar(cs[i], vs[i])
	}

	require.EqualValues(t, 10, dic.Cantidad())

	/* establezco los rangos, uno no incluido en el Abb y el otro incluido */
	a, b := 7, 13
	newCsE, newVsE := guardarClavesIntValorIterExterno(dic, &a, &b)

	require.True(t, compararArreglos(CLAVES_ORDENADAS_1[a/2:b/2+1], newCsE))
	require.True(t, compararArreglos(VALORES_ORDENADOS[a/2:b/2+1], newVsE))
}

func guardarClavesIntValorIterInterno[K int, V any](dic TDAAbb.DiccionarioOrdenado[K, V], desde *K, hasta *K, corte *K) (clavesRes []K, valoresRes []V) {

	var i int
	iPtr := &i
	dic.IterarRango(desde, hasta, func(c K, v V) bool {
		if corte != nil && c > *corte {
			return false
		}
		clavesRes = append(clavesRes, c)
		valoresRes = append(valoresRes, v)
		*iPtr++
		return true
	})
	return clavesRes, valoresRes
}

func TestIteradorInternoRangosExcedentesSinCorte(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	claves := []int{6, 1, 20, 10, 24, 4, 8, 13, 11, 16}
	valores := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	desde, hasta := -9, 99
	clavesRes, valoresRes := guardarClavesIntValorIterInterno(dic, &desde, &hasta, nil)

	require.EqualValues(t, len(claves), dic.Cantidad())
	require.True(t, compararArreglos(CLAVES_ORDENADAS_2, clavesRes))
	require.True(t, compararArreglos(VALORES_ORDENADOS, valoresRes))
}

func TestIteradorInternoRangosExcedentesConCorte(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	claves := []int{6, 1, 20, 10, 24, 4, 8, 13, 11, 16}
	valores := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	desde, hasta := -9, 99
	condicionDeCorte := len(claves)
	clavesRes, valoresRes := guardarClavesIntValorIterInterno(dic, &desde, &hasta, &condicionDeCorte)

	require.EqualValues(t, len(claves), dic.Cantidad())
	require.True(t, compararArreglos(CLAVES_ORDENADAS_2[:len(claves)/2], clavesRes))
	require.True(t, compararArreglos(VALORES_ORDENADOS[:len(valores)/2], valoresRes))
}

func TestIteradorInternoRangosAcotados(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	claves := []int{6, 1, 20, 10, 24, 4, 8, 13, 11, 16}
	valores := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	desde, hasta := 8, 16
	condicionDeCorte := len(claves)
	clavesRes, valoresRes := guardarClavesIntValorIterInterno(dic, &desde, &hasta, &condicionDeCorte)

	require.EqualValues(t, len(claves), dic.Cantidad())
	require.True(t, compararArreglos(CLAVES_ORDENADAS_2[len(claves)/3:len(claves)/2], clavesRes))
	require.True(t, compararArreglos(VALORES_ORDENADOS[len(valores)/3:len(valores)/2], valoresRes))
}

func TestAmbosIteradoresInorder(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	claves := []int{6, 1, 20, 10, 24, 4, 8, 13, 11, 16}
	valores := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	clavesIterInterno, valoresIterInterno := guardarClavesIntValorIterInterno(dic, nil, nil, nil)

	clavesIterExterno, valoresIterExterno := guardarClavesIntValorIterExterno(dic, nil, nil)

	require.EqualValues(t, 10, dic.Cantidad())
	require.True(t, compararArreglos(CLAVES_ORDENADAS_2, clavesIterExterno))
	require.True(t, compararArreglos(VALORES_ORDENADOS, valoresIterExterno))

	require.True(t, compararArreglos(CLAVES_ORDENADAS_2, clavesIterInterno))
	require.True(t, compararArreglos(VALORES_ORDENADOS, valoresIterInterno))

	require.True(t, compararArreglos(clavesIterInterno, clavesIterExterno))
	require.True(t, compararArreglos(valoresIterInterno, valoresIterExterno))
}

func TestIteradorInternoRangosConCorte(t *testing.T) {
	dic := TDAAbb.CrearABB[int, string](funcion_cmp_int)
	claves := []int{6, 1, 20, 10, 24, 4, 8, 13, 11, 16}
	valores := []string{"c", "a", "i", "e", "j", "b", "d", "g", "f", "h"}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	desde, hasta := 1, 6
	condicionDeCorte := len(claves)
	clavesRes, valoresRes := guardarClavesIntValorIterInterno(dic, &desde, &hasta, &condicionDeCorte)

	require.EqualValues(t, 10, dic.Cantidad())
	require.True(t, compararArreglos([]int{1, 4, 6}, clavesRes))
	require.True(t, compararArreglos([]string{"a", "b", "c"}, valoresRes))
}
