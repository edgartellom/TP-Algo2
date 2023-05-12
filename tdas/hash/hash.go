package diccionario

import (
	"encoding/binary"
	"fmt"
)

const (
	LONGITUD_INICIAL    = 13
	PANIC_NO_PERTENECE  = "La clave no pertenece al diccionario"
	PANIC_ITERADOR      = "El iterador termino de iterar"
	FACTOR_REDIMENSION  = 2
	FACTOR_CARGA_ARRIBA = 0.7
	FACTOR_CARGA_ABAJO  = 0.2
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

/* ---------------------------------- FUNCIONES DE CREACION ------------------------------ */
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

/* ---------------------------------- FUNCIONES AUXILIARES DE REDIMENSION ------------------------------ */
func (hashC *hashCerrado[K, V]) redimensionarTabla(nuevoTam int) {
	tablaActual := copiarTabla(hashC.tabla)

	hashC.tam = nuevoTam
	hashC.tabla = make([]celdaHash[K, V], nuevoTam)
	hashC.cantidad = 0

	for actual, pos := obtenerCeldaOcupada(tablaActual, 0); actual != nil; {
		clave, valor := actual.clave, actual.dato
		hashC.Guardar(clave, valor)
		actual, pos = obtenerCeldaOcupada(tablaActual, pos+1)
	}
}

func copiarTabla[K comparable, V any](tabla []celdaHash[K, V]) []celdaHash[K, V] {
	copia := make([]celdaHash[K, V], len(tabla))
	indice := 0
	for actual, pos := obtenerCeldaOcupada(tabla, 0); actual != nil; indice++ {
		clave, valor := actual.clave, actual.dato
		celda := crearCelda(clave, valor)
		copia[indice] = celda
		actual, pos = obtenerCeldaOcupada(tabla, pos+1)
	}
	return copia
}

/* ---------------------------------- FUNCIONES AUXILIARES DE HASHING ------------------------------ */
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func f_hash[K comparable](clave K) uint32 { // murmur hash
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

	return h1

}

/* ---------------------------------- FUNCIONES AUXILIARES HASH CERRADO ------------------------------ */
func (hashC *hashCerrado[K, V]) avanzarPosicion(pos int) int {
	if pos == hashC.tam-1 {
		return 0
	}
	return pos + 1
}

func (hashC hashCerrado[K, V]) factorDeCarga() float64 {
	return float64(hashC.cantidad+hashC.borrados) / float64(hashC.tam)
}

func (hashC hashCerrado[K, V]) obtenerPosicionHashing(clave K) int {
	hash := f_hash(clave)
	pos := int(hash) % hashC.tam
	return pos
}

func (hashC hashCerrado[K, V]) obtenerPosicionDeBusqueda(clave K) int {
	pos := hashC.obtenerPosicionHashing(clave)
	for hashC.tabla[pos].clave != clave || hashC.tabla[pos].estado != OCUPADO {
		pos = hashC.avanzarPosicion(pos)
	}
	return pos
}

func (hashC hashCerrado[K, V]) obtenerPosicionDeGuardado(clave K) int {
	pos := hashC.obtenerPosicionHashing(clave)
	for hashC.tabla[pos].estado != VACIO {
		pos = hashC.avanzarPosicion(pos)
	}
	return pos
}

/* ---------------------------------- FUNCION AUXILIAR ITERADORA ------------------------------ */
func obtenerCeldaOcupada[K comparable, V any](tabla []celdaHash[K, V], posicion int) (*celdaHash[K, V], int) {
	for i := posicion; i < len(tabla); i++ {
		if tabla[i].estado == OCUPADO {
			return &tabla[i], i
		}
	}
	return nil, len(tabla)
}

/* ---------------------------------- PRIMITIVAS HASH CERRADO ------------------------------ */
func (hashC hashCerrado[K, V]) Cantidad() int {
	return hashC.cantidad
}

func (hashC hashCerrado[K, V]) Pertenece(clave K) bool {
	pos := hashC.obtenerPosicionHashing(clave)

	for hashC.tabla[pos].estado != VACIO {
		if hashC.tabla[pos].clave == clave && hashC.tabla[pos].estado == OCUPADO {
			return true
		}
		pos = hashC.avanzarPosicion(pos)
	}
	return false
}

func (hashC hashCerrado[K, V]) Obtener(clave K) V {
	hashC.comprobarSiPertenece(clave)
	pos := hashC.obtenerPosicionDeBusqueda(clave)
	return hashC.tabla[pos].dato
}

func (hashC *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hashC.factorDeCarga() >= FACTOR_CARGA_ARRIBA && hashC.factorDeCarga() < 1 {
		hashC.redimensionarTabla(hashC.tam * FACTOR_REDIMENSION)
	}

	if !hashC.Pertenece(clave) {
		pos := hashC.obtenerPosicionDeGuardado(clave)
		celda := crearCelda(clave, dato)
		(*hashC).tabla[pos] = celda
		(*hashC).cantidad++
	} else {
		pos := hashC.obtenerPosicionDeBusqueda(clave)
		(*hashC).tabla[pos].dato = dato
	}

}

func (hashC *hashCerrado[K, V]) Borrar(clave K) V {
	hashC.comprobarSiPertenece(clave)
	pos := hashC.obtenerPosicionDeBusqueda(clave)
	dato := hashC.tabla[pos].dato
	(*hashC).tabla[pos].estado = BORRADO
	(*hashC).borrados++
	(*hashC).cantidad--
	if hashC.factorDeCarga() <= FACTOR_CARGA_ABAJO && hashC.factorDeCarga() > 0 {
		nuevoTam := hashC.tam / FACTOR_REDIMENSION
		if nuevoTam < LONGITUD_INICIAL {
			nuevoTam = LONGITUD_INICIAL
		}
		hashC.redimensionarTabla(nuevoTam)
	}
	return dato
}

func (hashC hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for actual, pos := obtenerCeldaOcupada(hashC.tabla, 0); actual != nil && visitar(actual.clave, actual.dato); {
		actual, pos = obtenerCeldaOcupada(hashC.tabla, pos+1)
	}
}

func (hashC *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])
	iter.hashC = hashC
	iter.actual, iter.posicion = obtenerCeldaOcupada(hashC.tabla, 0)
	return iter
}

func (iter iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

/* ---------------------------------- PRIMITIVAS ITERADOR ------------------------------ */
func (iter iterHash[K, V]) VerActual() (K, V) {
	iter.comprobarIteradorFinalizo()
	return iter.actual.clave, iter.actual.dato
}

func (iter *iterHash[K, V]) Siguiente() {
	iter.comprobarIteradorFinalizo()
	(*iter).actual, (*iter).posicion = obtenerCeldaOcupada(iter.hashC.tabla, iter.posicion+1)

}

/* ---------------------------------- COMPROBADORES DE PANICS ------------------------------ */
func (hashC hashCerrado[K, V]) comprobarSiPertenece(clave K) {
	if !hashC.Pertenece(clave) {
		panic(PANIC_NO_PERTENECE)
	}
}

func (iter iterHash[K, V]) comprobarIteradorFinalizo() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}
