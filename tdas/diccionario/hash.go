package diccionario

import (
	"fmt"
)

const (
	LONGITUD_INICIAL   = 13
	PANIC_NO_PERTENECE = "La clave no pertenece al diccionario"
	PANIC_ITERADOR     = "El iterador termino de iterar"
	FACTOR_REDIMENSION = 2
	FACTOR_ACHICAR     = 4
)

type Estado int

const (
	VACIO Estado = iota
	BORRADO
	OCUPADO
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado Estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

type iterHash[K comparable, V any] struct {
	hashC    *hashCerrado[K, V]
	actual   *celdaHash[K, V]
	posicion int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tam = LONGITUD_INICIAL
	hash.tabla = make([]celdaHash[K, V], hash.tam)
	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hashC hashCerrado[K, V]) hash(clave K) int { // hash fnv-1a
	entrada := convertirABytes(clave)

	const (
		prime  = 1099511628211
		offset = 14695981039346656037
	)

	hash := uint64(offset)
	for _, b := range entrada {
		hash ^= uint64(b)
		hash *= prime
	}

	return int(hash) % hashC.tam
}

func (hash *hashCerrado[K, V]) avanzar(pos int) int {
	if pos > len(hash.tabla) {
		return 0
	}
	return pos + 1
}

func (hash *hashCerrado[K, V]) redimensionarTabla(nuevaCap int) {
	s := make([]celdaHash[K, V], nuevaCap)
	copy(s, hash.tabla)
	(*hash).tabla = s
}

func (hashC hashCerrado[K, V]) comprobarSiPertenece(clave K) {
	if !hashC.Pertenece(clave) {
		panic(PANIC_NO_PERTENECE)
	}
}

func (hashC hashCerrado[K, V]) obtenerPosicion(clave K) int {
	pos := hashC.hash(clave)
	for hashC.tabla[pos].clave != clave {
		pos = hashC.avanzar(pos)
	}
	return pos
}

func (hashC hashCerrado[K, V]) factorDeCarga() int {
	return hashC.cantidad + hashC.borrados/hashC.tam
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (hashC hashCerrado[K, V]) Pertenece(clave K) bool {
	pos := hashC.hash(clave)

	for hashC.tabla[pos].estado != VACIO {
		if hashC.tabla[pos].clave == clave {
			return true
		}
		pos = hashC.avanzar(pos)
	}
	return false
}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (hashC *hashCerrado[K, V]) Guardar(clave K, dato V) {
	pos := hashC.hash(clave)

	if hashC.factorDeCarga() == 1 {
		hashC.redimensionarTabla(hashC.tam * FACTOR_REDIMENSION)
	}

	if !hashC.Pertenece(clave) {
		(*hashC).tabla[pos].dato = dato
		(*hashC).tabla[pos].clave = clave
		(*hashC).cantidad++
	} else {
		for hashC.tabla[pos].estado != VACIO {
			if hashC.tabla[pos].clave == clave {
				(*hashC).tabla[pos].dato = dato
			}
			pos = hashC.avanzar(pos)
		}
	}
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (hashC hashCerrado[K, V]) Obtener(clave K) V {
	hashC.comprobarSiPertenece(clave)
	pos := hashC.obtenerPosicion(clave)
	return hashC.tabla[pos].dato
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (hashC *hashCerrado[K, V]) Borrar(clave K) V {
	hashC.comprobarSiPertenece(clave)
	pos := hashC.obtenerPosicion(clave)
	dato := hashC.tabla[pos].dato
	hashC.tabla[pos].estado = BORRADO
	hashC.borrados++
	hashC.cantidad--
	if hashC.factorDeCarga()*FACTOR_ACHICAR <= hashC.tam {
		nuevaCap := hashC.tam / FACTOR_REDIMENSION
		if nuevaCap < LONGITUD_INICIAL {
			nuevaCap = LONGITUD_INICIAL
		}
		hashC.redimensionarTabla(nuevaCap)
	}
	return dato

}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (hashC hashCerrado[K, V]) Cantidad() int {
	return hashC.cantidad
}

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
// mismo
func (hashC hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {

}

// Iterador devuelve un IterDiccionario para este Diccionario
func (hashC *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])

	return iter
}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
func (iter iterHash[K, V]) HaySiguiente() bool {
	return false
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter iterHash[K, V]) VerActual() (K, V) {
	return iter.actual.clave, iter.actual.dato
}

// Siguiente si HaySiguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter iterHash[K, V]) Siguiente() {

}
