package vuelos

import (
	errores "algueiza/errores"
	"bufio"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

const COMPARADOR = 0

type tablero struct {
	abb    TDADicc.DiccionarioOrdenado[Claves, Vuelo]
	hash   TDADicc.Diccionario[string, Vuelo]
	vuelos []Vuelo
}

func cmpPrioridad(a, b Vuelo) int {
	if a.VerPrioridad() > b.VerPrioridad() {
		return 1
	} else if a.VerPrioridad() == b.VerPrioridad() {
		return strings.Compare(b.VerCodigo(), a.VerCodigo())
	}
	return -1
}

func cmpTablero(a, b Claves) int {
	superior := strings.Compare(a.fecha, b.fecha)
	if superior == COMPARADOR {
		return strings.Compare(b.codigo, a.codigo)
	}
	return superior
	// if a.VerFecha() > b.VerFecha() {
	// 	return 1
	// } else if a.VerFecha() == b.VerFecha() {
	// 	return strings.Compare(b.VerCodigo(), a.VerCodigo())
	// }
	// return -1
}

func CrearTablero() Tablero {
	abb := TDADicc.CrearABB[Claves, Vuelo](cmpTablero)
	hash := TDADicc.CrearHash[string, Vuelo]()
	return &tablero{abb: abb, hash: hash}
}

func (tablero *tablero) CargarInformacion(ruta string) {
	// archivo := funciones.AbrirArchivo(ruta)
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	var vuelos []Vuelo
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		infoDeVuelo := scanner.Text()
		vuelo := CrearVuelo(infoDeVuelo)
		(*tablero).hash.Guardar(vuelo.VerCodigo(), vuelo)
		(*tablero).abb.Guardar(vuelo.VerInformacionPrincipal(), vuelo)
		vuelos = append(vuelos, vuelo)
	}
	(*tablero).vuelos = vuelos
}

func (tablero *tablero) ObtenerVuelosEntreRango(k int, desde, hasta string) []Vuelo {
	vuelos := make([]Vuelo, k)
	// var vuelos []Vuelo
	FechaDeSalida := CrearInformacionPrincipal("", desde, "", "")
	FechaDeLlegada := CrearInformacionPrincipal("", hasta, "", "")
	for iter, i := tablero.abb.IteradorRango(&FechaDeSalida, &FechaDeLlegada), 0; iter.HaySiguiente() && i != k; iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelos[i] = vuelo
	}
	return vuelos
}

// func (tablero *tablero) ObtenerVuelos(K int, modo string, desde, hasta Claves) ([]Vuelo, error) {

// 	if K <= 0 || (modo != ASCENDENTE && modo != DESCENDENTE) || hasta.Fecha < desde.Fecha {
// 		err := e.ErrorComando{}
// 		return nil, err
// 	}
// 	var vuelos []Vuelo
// 	var contador int
// 	contPtr := &contador
// 	if modo == DESCENDENTE {
// 		tablero.vuelosOrdenadosFechaDesc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
// 			if *contPtr >= K {
// 				return false
// 			}
// 			vuelos = append(vuelos, d)
// 			*contPtr++
// 			return true
// 		})

// 	} else {
// 		tablero.vuelosOrdenadosFechaAsc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
// 			if *contPtr >= K {
// 				return false
// 			}
// 			vuelos = append(vuelos, d)
// 			*contPtr++
// 			return true
// 		})
// 	}
// 	return vuelos, nil
// }

func (tablero *tablero) ObtenerVuelo(codigo string) (*Vuelo, error) {
	if !tablero.hash.Pertenece(codigo) {
		err := errores.ErrorComando{}
		return nil, err
	}
	vuelo := tablero.hash.Obtener(codigo)
	return &vuelo, nil
}

func (tablero *tablero) ObtenerVuelosPrioritarios(k int) []Vuelo {
	var vuelos []Vuelo
	for iter := tablero.hash.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelos = append(vuelos, vuelo)
	}
	TDAHeap.HeapSort(vuelos, cmpPrioridad)
	return vuelos[:k+1]
}
