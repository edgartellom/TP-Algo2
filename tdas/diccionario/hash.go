package diccionario

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Estado int

const (
	LONGITUD_INICIAL   = 13
	METRICA            = 0.7
	FACTOR_AGRANDAR    = 2
	FACTOR_ACHICAR     = 4
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

/* ----------------- PARA FUNCION DE HASHING -------------------- */

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func f_hashing[K comparable](clave K, longitud int) int {
	bytes := convertirABytes(clave)

	const (
		offset = 14695981039346656037
		prime  = 1099511628211
	)

	hash := uint64(offset)
	for _, b := range bytes {
		hash ^= uint64(b)
		hash *= prime
	}
	resultado := int(hash) % longitud
	resultado = int(math.Abs(float64(resultado)))
	return resultado
}

func mumurHash123(clave string, longitud int) int {
	key := convertirABytes(clave)

	const (
		m = 0x5bd1e995
		r = 24
	)

	var (
		h    = uint32(len(key))
		data = key
	)

	for len(data) >= 4 {
		k := binary.LittleEndian.Uint32(data)
		k *= m
		k ^= k >> r
		k *= m

		h *= m
		h ^= k

		data = data[4:]
	}

	switch len(data) {
	case 3:
		h ^= uint32(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(data[0])
		h *= m
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	posicion := int(h) % longitud
	return posicion
}

/* -------------------------------- FUNCIONES CREAR ---------------------------------- */

func crearCelda[K comparable, V any](clave K, valor V) celdasHash[K, V] {
	celda := new(celdasHash[K, V])
	celda.clave = clave
	celda.dato = valor
	celda.estado = OCUPADO
	return *celda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tam = LONGITUD_INICIAL
	tabla := make([]celdasHash[K, V], hash.tam)
	// for i := range tabla {
	// 	tabla[i] = crearCelda[K, V]()
	// }
	hash.tabla = tabla
	return hash
}

/* --------------------------- PRIMITIVAS AUXILIARES ---------------------------- */

func (hash *hashCerrado[K, V]) factorDeCarga() float64 {
	return float64(hash.cantidad+hash.borrados) / float64(hash.tam)
}

func (hash *hashCerrado[K, V]) comprobarSiPertenece(clave K) {
	if !hash.Pertenece(clave) {
		panic(PANIC_NO_PERTENECE)
	}
}

func (hash *hashCerrado[K, V]) avanzarPosicion(posicion int) int {
	if posicion == hash.tam-1 {
		return 0
	}
	return posicion + 1
}

func (hash *hashCerrado[K, V]) redimensionarTabla(nuevoTam int) {
	tablaActual := copiarTabla(hash.tabla)

	hash.tam = nuevoTam
	hash.tabla = make([]celdasHash[K, V], nuevoTam)

	for actual, pos := obtenerCeldaOcupada(tablaActual, 0); actual != nil; {
		clave, valor := actual.clave, actual.dato
		hash.Guardar(clave, valor)
		actual, pos = obtenerCeldaOcupada(tablaActual, pos+1)
	}

}

func (hash *hashCerrado[K, V]) obtenerPosicionValida(clave K) int {
	posicion := f_hashing(clave, hash.tam)
	for ; hash.tabla[posicion].estado != VACIO; posicion = hash.avanzarPosicion(posicion) {
	}
	return posicion
}

// PENSAR EN LAS CELDAS BORRADAS; REVISARIAS LA CLAVE DE UNA CELDA BORRADA
func (hash *hashCerrado[K, V]) obtenerPosicionDeClave(clave K) int {
	posicion := f_hashing(clave, hash.tam)
	for hash.tabla[posicion].clave != clave {
		posicion = hash.avanzarPosicion(posicion)
	}
	return posicion
}

/* ---------------------------------- FUNCIONES AUXILIARES ------------------------------ */

func copiarTabla[K comparable, V any](tabla []celdasHash[K, V]) []celdasHash[K, V] {
	copia := make([]celdasHash[K, V], len(tabla))
	indice := 0
	for actual, pos := obtenerCeldaOcupada(tabla, 0); actual != nil; indice++ {
		clave, valor := actual.clave, actual.dato
		celda := crearCelda(clave, valor)
		copia[indice] = celda
		actual, pos = obtenerCeldaOcupada(tabla, pos+1)
	}
	return copia
}

func obtenerCeldaOcupada[K comparable, V any](tabla []celdasHash[K, V], posicion int) (*celdasHash[K, V], int) {
	for i := posicion; i < len(tabla); i++ {
		if tabla[i].estado == OCUPADO {
			return &tabla[i], i
		}
	}
	return nil, len(tabla)
}

/* ----------------------------------- PRIMITIVAS PRINCIPALES ------------------------------------- */

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	for pos := f_hashing(clave, hash.tam); hash.tabla[pos].estado != VACIO; pos = hash.avanzarPosicion(pos) {
		if hash.tabla[pos].estado == OCUPADO && hash.tabla[pos].clave == clave {
			return true
		}
	}
	return false
}

func (hash *hashCerrado[K, V]) Guardar(clave K, valor V) {
	if hash.factorDeCarga() > METRICA {
		hash.redimensionarTabla(hash.tam * FACTOR_AGRANDAR)
	}

	if !hash.Pertenece(clave) {
		posicion := hash.obtenerPosicionValida(clave)
		celda := crearCelda(clave, valor)
		(*hash).tabla[posicion] = celda
		hash.cantidad++
	} else {
		posicion := hash.obtenerPosicionDeClave(clave)
		(*hash).tabla[posicion].dato = valor
	}
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	hash.comprobarSiPertenece(clave)

	posicion := f_hashing(clave, hash.tam)
	for ; hash.tabla[posicion].clave != clave; posicion = hash.avanzarPosicion(posicion) {
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	hash.comprobarSiPertenece(clave)

	// SE PODRIA REDUCIR UNA LINEA???
	posicion := f_hashing(clave, hash.tam)
	for hash.tabla[posicion].clave != clave {
		posicion = hash.avanzarPosicion(posicion)
	}
	dato := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = BORRADO
	hash.borrados++
	hash.cantidad--

	if hash.factorDeCarga()*FACTOR_ACHICAR < METRICA {
		nuevoTam := hash.tam / FACTOR_AGRANDAR
		if nuevoTam < LONGITUD_INICIAL {
			nuevoTam = LONGITUD_INICIAL
		}
		hash.redimensionarTabla(nuevoTam)
	}

	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

/* ------------------------------------ ITERADOR INTERNO ----------------------------------------- */
func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for actual, pos := obtenerCeldaOcupada(hash.tabla, 0); actual != nil && visitar(actual.clave, actual.dato); {
		actual, pos = obtenerCeldaOcupada(hash.tabla, pos+1)
	}
}

/* ------------------------------------ ITERADOR EXTERNO ----------------------------------------- */
func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterHash[K, V])
	iter.hash = hash
	iter.actual, iter.posicion = obtenerCeldaOcupada(hash.tabla, iter.posicion)
	return iter
}

func (iter *iterHash[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterHash[K, V]) VerActual() (K, V) {
	iter.iteradorFueraDeRango()
	return iter.actual.clave, iter.actual.dato
}

func (iter *iterHash[K, V]) Siguiente() {
	iter.iteradorFueraDeRango()
	for iter.hash.tabla[iter.posicion].estado != OCUPADO {
		iter.posicion++
	}
	if iter.posicion < len(iter.hash.tabla) {
		iter.actual = &iter.hash.tabla[iter.posicion]
		iter.posicion++
	} else {
		iter.actual = nil
	}
}

func (iter *iterHash[K, V]) iteradorFueraDeRango() {
	if !iter.HaySiguiente() {
		panic(PANIC_ITERADOR)
	}
}
