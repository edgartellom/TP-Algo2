package vuelos

type Tablero interface {
	CargarInformacion(string)
	ObtenerVuelo(string) (*Vuelo, error)
	ObtenerVuelosEntreRango(int, string, string) []Vuelo
	ObtenerVuelosPrioritarios(int) []Vuelo
	// ObtenerVuelos(K int, modo string, desde, hasta *Vuelo) ([]Vuelo, error)
	// ObtenerVuelo(codigo string) (Vuelo, error)
	// ObtenerVuelosPrioritarios(K int) []Vuelo
	// SiguienteVuelo(origen, destino, fecha string) (Vuelo, error)
	// ActualizarTablero(archivo *os.File)
	// Borrar(desde, hasta *Vuelo) ([]Vuelo, error)
}
