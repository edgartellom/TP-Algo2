package grafo

type Grafo[K comparable, V any] interface {
	AgregarVertice(vertice K)

	BorrarVertice(vertice K)

	AgregarArista(vertice1, vertice2 K, peso V)

	BorrarArista(vertice1, vertice2 K)

	Pertenece(vertice K) bool

	HayArista(vertice1, vertice2 K) bool

	ObtenerPeso(vertice1, vertice2 K) V
}
