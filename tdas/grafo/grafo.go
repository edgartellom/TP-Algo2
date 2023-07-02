package grafo

type Grafo[K comparable, V any] interface {
	EsDirigido() bool
	AgregarVertice(vertice K)
	BorrarVertice(vertice K)
	BorrarArista(v1, v2 K)
	HayArista(v1, v2 K) bool
	Existe(vertice K) bool
	ObtenerVertices() []K
	Cantidad() int
}

type GrafoNoPesado[K comparable, V any] interface {
	Grafo[K, V]
	AgregarArista(v1, v2 K)
}

type GrafoPesado[K comparable, V any] interface {
	Grafo[K, V]
	AgregarArista(v1, v2 K, peso V)
	VerPeso(v1, v2 K) V
}

// Estan disponibles las funciones:
// * CrearGrafoPesado(dirigido bool) GrafoPesado
// * CrearGrafoNoPesado(dirigido bool) GrafoNoPesado
