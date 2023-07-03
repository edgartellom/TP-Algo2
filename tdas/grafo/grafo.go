package grafo

type Grafo[K comparable, V any] interface {
	EsDirigido() bool

	AgregarVertice(K)

	BorrarVertice(K)

	BorrarArista(K, K)

	HayArista(K, K) bool

	Existe(K) bool

	ObtenerVertices() []K

	ObtenerAdyacentes(K) []K

	Cantidad() int
}

type GrafoNoPesado[K comparable, V any] interface {
	Grafo[K, V]

	AgregarArista(K, K)
}

type GrafoPesado[K comparable, V any] interface {
	Grafo[K, V]

	AgregarArista(K, K, V)

	VerPeso(K, K) V
}
