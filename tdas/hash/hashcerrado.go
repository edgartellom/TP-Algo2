package diccionario

import "fmt"

type Estado int

const (
	LONGITUD_INICIAL   = 13
	PANIC_NO_PERTENECE = "La clave no pertenece al diccionario"
	PANIC_ITERADOR     = "El iterador termino de iterar"
)

const (
	VACIO Estado = iota
	OCUPADO
	BORRADO
)

type celdasHash[K comparable, V any] struct {
	estado Estado
	clave  K
	dato   V
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdasHash[K, V]
	cantidad int
	borrados int
	tam      int
}

type iterHash[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	actual   *celdasHash[K, V]
	posicion int
}

/* ----------------- FUNCION HASHING -------------- */
func f_hashing[K comparable](clave K) int {
	return 0
}

/* ---------- FUNCION PARA CONVERTIR A BYTES --------- */
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

/* ----------------- ES NECESARIO ??? ------------------*/
func crearCelda[K comparable, V any]() celdasHash[K, V] {
	return *new(celdasHash[K, V])
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tam = LONGITUD_INICIAL
	tabla := make([]celdasHash[K, V], hash.tam)
	for i := range tabla {
		tabla[i] = crearCelda[K, V]()
	}
	hash.tabla = tabla
	return hash
}

/*
ESTARIA BIEN HACER:

	"for posicion := f_hashing(clave); hash.tabla[posicion].estado == OCUPADO; hash.avanzar(posicion) {"

NOSE SI QUEDARIA MUY LARGA Y DESPROLIJO
*/
func (hash *hashCerrado[K, V]) comprobarSiPertenece(clave K) {
	if !hash.Pertenece(clave) {
		panic(PANIC_NO_PERTENECE)
	}
}

func (hash *hashCerrado[K, V]) avanzar(posicion int) int {
	if posicion > len(hash.tabla) {
		return 0
	}
	return posicion + 1
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	posicion := f_hashing(clave)
	for hash.tabla[posicion].estado == OCUPADO {
		if hash.tabla[posicion].clave == clave {
			return true
		}
		posicion = hash.avanzar(posicion)
	}
	return false
}

/* TAL VEZ HACER QUE "(*hash).tabla[posicion].dato = valor" PASE 1 VEZ, OSEA DESPUES DEL "if" Y "else"...*/
func (hash *hashCerrado[K, V]) Guardar(clave K, valor V) {
	posicion := f_hashing(clave)
	if !hash.Pertenece(clave) {
		for hash.tabla[posicion].estado != VACIO {
			posicion = hash.avanzar(posicion)
		}
		(*hash).tabla[posicion].estado = OCUPADO
		(*hash).tabla[posicion].clave = clave
		(*hash).tabla[posicion].dato = valor
	} else {
		for hash.tabla[posicion].clave != clave {
			posicion = hash.avanzar(posicion)
		}
		(*hash).tabla[posicion].dato = valor
	}
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	hash.comprobarSiPertenece(clave)
	posicion := f_hashing(clave)
	for hash.tabla[posicion].clave != clave {
		posicion = hash.avanzar(posicion)
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	hash.comprobarSiPertenece(clave)
	posicion := f_hashing(clave)
	for hash.tabla[posicion].clave != clave {
		posicion = hash.avanzar(posicion)
	}
	dato := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = BORRADO
	hash.borrados++
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {

}

/* --------------- ESTA BIEN USAR UNA FUNCION DEL ITERADOR EN EL CREARITERADOR ??? ---------- */
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])
	iter.hash = hash
	iter.actual = iter.proximoClaveValor()
	return iter
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	return iter.actual.clave, iter.actual.dato
}

func (iter *iterHash[K, V]) Siguiente() {
	iter.actual = iter.proximoClaveValor()
}

/* ------------------------- NOMBRE DUDOSO, FUNCION MAS DUDOSA -----------------------------*/
func (iter *iterHash[K, V]) proximoClaveValor() *celdasHash[K, V] {
	for iter.posicion < len(iter.hash.tabla) {
		if iter.hash.tabla[iter.posicion].estado == OCUPADO {
			actual := iter.hash.tabla[iter.posicion]
			return &actual
		}
		iter.posicion++
	}
	return nil
}
