package diccionario

import (
	"encoding/binary"
	"fmt"
)

type estado int

const (
	LONGITUD_INICIAL   = 13
	FACTOR_REDIMENSION = 2
	FACTOR_AGRANDAR    = 0.7
	FACTOR_ACHICAR     = 0.2
	// PANIC_NO_PERTENECE = "La clave no pertenece al diccionario"
	// PANIC_ITERADOR     = "El iterador termino de iterar"
)

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)

type celdaHash[K comparable, V any] struct {
	estado estado
	clave  K
	dato   V
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	borrados int
	tam      int
}

type iterHash[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	actual   *celdaHash[K, V]
	posicion int
}

/* ---------------------------------- FUNCIONES AUXILIARES DE HASHING ---------------------------------- */

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionDeHashing[K comparable](clave K) int { // <--- MUMUR HASH
	entrada := convertirABytes(clave)

	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	var (
		h1    = uint32(len(entrada))
		k1    uint32
		chunk uint32
	)

	for len(entrada) >= 4 {
		chunk = binary.LittleEndian.Uint32(entrada)
		k1 = chunk

		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1))
		k1 *= c2

		h1 ^= k1
		h1 = (h1 << r2) | (h1 >> (32 - r2))
		h1 = h1*m + n

		entrada = entrada[4:]
	}

	k1 = 0
	switch len(entrada) {
	case 3:
		k1 ^= uint32(entrada[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(entrada[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(entrada[0])
		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1))
		k1 *= c2
		h1 ^= k1
	}

	h1 ^= uint32(len(entrada))
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return int(h1)
}

/* --------------------------------------- FUNCIONES DE CREACION --------------------------------------- */

func crearCelda[K comparable, V any](clave K, dato V) celdaHash[K, V] {
	celda := new(celdaHash[K, V])
	celda.clave = clave
	celda.dato = dato
	celda.estado = OCUPADO
	return *celda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tam = LONGITUD_INICIAL
	tabla := make([]celdaHash[K, V], hash.tam)
	hash.tabla = tabla
	return hash
}

/* --------------------------------- FUNCIONES AUXILIARES HASH CERRADO --------------------------------- */

func (hash *hashCerrado[K, V]) factorDeCarga() float64 {
	return float64(hash.cantidad+hash.borrados) / float64(hash.tam)
}

func (hash *hashCerrado[K, V]) avanzarPosicion(posicion int) int {
	if posicion == hash.tam-1 {
		return 0
	}
	return posicion + 1
}

func (hash hashCerrado[K, V]) obtenerPosicionHashing(clave K) int {
	numeroDeHash := funcionDeHashing(clave)
	return numeroDeHash % hash.tam
}

func (hash *hashCerrado[K, V]) obtenerPosicion(clave K) int {
	pos := hash.obtenerPosicionHashing(clave)
	for ; hash.tabla[pos].estado != VACIO; pos = hash.avanzarPosicion(pos) {
		if hash.tabla[pos].estado == OCUPADO && hash.tabla[pos].clave == clave {
			return pos
		}
	}
	return pos
}

/* ------------------------------------- FUNCION AUXILIAR ITERADORA ------------------------------------ */

func obtenerCeldaOcupada[K comparable, V any](tabla []celdaHash[K, V], posicion int) (*celdaHash[K, V], int) {
	for indice := posicion; indice < len(tabla); indice++ {
		if tabla[indice].estado == OCUPADO {
			return &tabla[indice], indice
		}
	}
	return nil, len(tabla)
}

/* -------------------------------- FUNCIONES AUXILIARES DE REDIMENSION -------------------------------- */

func (hash *hashCerrado[K, V]) redimensionarTabla(nuevoTam int) {
	tablaActual := copiarTabla(hash.tabla)

	hash.tam = nuevoTam
	hash.tabla = make([]celdaHash[K, V], nuevoTam)
	hash.cantidad = 0

	for actual, pos := obtenerCeldaOcupada(tablaActual, 0); actual != nil; actual, pos = obtenerCeldaOcupada(tablaActual, pos+1) {
		clave, valor := actual.clave, actual.dato
		hash.Guardar(clave, valor)
	}
}

func copiarTabla[K comparable, V any](tabla []celdaHash[K, V]) []celdaHash[K, V] {
	copia := make([]celdaHash[K, V], len(tabla))
	var indice int
	for actual, pos := obtenerCeldaOcupada(tabla, 0); actual != nil; actual, pos = obtenerCeldaOcupada(tabla, pos+1) {
		clave, valor := actual.clave, actual.dato
		celda := crearCelda(clave, valor)
		copia[indice] = celda
		indice++
	}
	return copia
}

/* -------------------------------------- PRIMITIVAS HASH CERRADO -------------------------------------- */

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	posicion := hash.obtenerPosicion(clave)
	return hash.tabla[posicion].estado == OCUPADO && hash.tabla[posicion].clave == clave
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.factorDeCarga() >= FACTOR_AGRANDAR {
		hash.redimensionarTabla(hash.tam * FACTOR_REDIMENSION)
	}

	posicion := hash.obtenerPosicion(clave)

	if hash.tabla[posicion].estado == VACIO {
		celda := crearCelda(clave, dato)
		(*hash).tabla[posicion] = celda
		hash.cantidad++
	} else {
		(*hash).tabla[posicion].dato = dato
	}
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.obtenerPosicion(clave)

	hash.comprobarEstado(posicion)

	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion := hash.obtenerPosicion(clave)

	hash.comprobarEstado(posicion)

	dato := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = BORRADO
	hash.cantidad--
	hash.borrados++

	nuevoTam := hash.tam / FACTOR_REDIMENSION
	if hash.factorDeCarga() <= FACTOR_ACHICAR && nuevoTam < LONGITUD_INICIAL {
		hash.redimensionarTabla(nuevoTam)
	}

	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for actual, pos := obtenerCeldaOcupada(hash.tabla, 0); actual != nil && visitar(actual.clave, actual.dato); {
		actual, pos = obtenerCeldaOcupada(hash.tabla, pos+1)
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])
	iter.hash = hash
	iter.actual, iter.posicion = obtenerCeldaOcupada(hash.tabla, 0)
	return iter
}

/* ---------------------------------------- PRIMITIVAS ITERADOR ---------------------------------------- */

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()

	return iter.actual.clave, iter.actual.dato
}

func (iter *iterHash[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()

	(*iter).actual, (*iter).posicion = obtenerCeldaOcupada(iter.hash.tabla, iter.posicion+1)
}

/* -------------------------------------- COMPROBADORES DE PANICS -------------------------------------- */

func (iter *iterHash[K, V]) comprobarIteradorFinalizo() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}

func (hash *hashCerrado[K, V]) comprobarEstado(posicion int) {
	if hash.tabla[posicion].estado == VACIO {
		panic(PANIC_NO_PERTENECE)
	}
}
