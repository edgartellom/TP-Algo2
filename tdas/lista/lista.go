package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista, al inicio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos se quita el primero de la lista,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si la lista tiene elementos se devuelve el valor del primero.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si la lista tiene elementos se devuelve el valor del ultimo.
	// Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve el tamaño de la lista
	Largo() int

	// Iterar(visitar func(T) bool)

	// Iterador() IteradorLista[T]
}
