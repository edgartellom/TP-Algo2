package pila

const capInicial = 8
const factorRedimension = 2
const factorDesapilar = 4

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, capInicial)
	return pila
}

func redimensionarPila[T any](p *pilaDinamica[T], nuevaCap int) {
	s := make([]T, nuevaCap)
	copy(s, p.datos)
	p.datos = s
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == cap(p.datos) {
		nuevaCap := cap(p.datos) * factorRedimension
		redimensionarPila(p, nuevaCap)
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	tope := p.VerTope()
	p.cantidad--
	if p.cantidad*factorDesapilar <= cap(p.datos) {
		nuevaCap := cap(p.datos) / factorRedimension
		if nuevaCap < capInicial {
			nuevaCap = capInicial
		}
		redimensionarPila(p, nuevaCap)
	}
	return tope
}
