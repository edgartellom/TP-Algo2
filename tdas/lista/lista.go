package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista, al inicio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la lista,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si la lista tiene elementos se devuelve el valor del primero.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si la lista tiene elementos se devuelve el valor del ultimo.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la longitud de la lista.
	Largo() int

	// Permite iterar la lista.
	Iterar(visitar func(T) bool)

	// Iterador instancia al iterador externo de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual del iterador de la lista.
	VerActual() T

	// HaySiguiente devuelve verdadero si existe un elemento siguiente al actual del iterador, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza al siguiente elemento en el iterador de la lista.
	Siguiente()

	// Insertar agrega un nuevo elemento a la lista en la posicion del elemento que apunta el iterador, desplazando los elementos.
	Insertar(T)

	// Borrar quita un elemento a la lista en la posicion del elemento actual del iterador y devuelve dicho elemento.
	Borrar() T
}
