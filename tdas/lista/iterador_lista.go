package lista

type IteradorLista[T any] interface {
	
	VerActual() T

	HaySiguiente() bool

	Siguiente()

	Insertar(T)
	
	Borrar() T
}
