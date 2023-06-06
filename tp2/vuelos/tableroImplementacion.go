package vuelos

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	e "algueiza/errores"
	f "algueiza/funciones"
	TDAHeap "tdas/cola_prioridad"
	TDADic "tdas/diccionario"
)

const COMPARADOR = 0

type tablero struct {
	abb    TDADic.DiccionarioOrdenado[*Vuelo, Vuelo]
	hash   TDADic.Diccionario[string, Vuelo]
	vuelos []*Vuelo
}

func cmpPrioridad(a, b *Vuelo) int {
	prioridad1, _ := strconv.Atoi((*a)[PRIORIDAD])
	prioridad2, _ := strconv.Atoi((*b)[PRIORIDAD])
	if prioridad1 > prioridad2 {
		return 1
	}
	if prioridad1 == prioridad2 {
		return strings.Compare((*b)[CODIGO], (*a)[CODIGO])
	}
	return -1
}

func cmpTablero(a, b *Vuelo) int {
	if (*a)[FECHA] > (*b)[FECHA] {
		return 1
	}
	if (*a)[FECHA] == (*b)[FECHA] {
		return strings.Compare((*b)[CODIGO], (*a)[CODIGO])
	}
	return -1
}

func CrearTablero() Tablero {
	abb := TDADic.CrearABB[*Vuelo, Vuelo](cmpTablero)
	hash := TDADic.CrearHash[string, Vuelo]()
	return &tablero{abb: abb, hash: hash}
}

func (tablero *tablero) ObtenerVuelos(K int, modo string, desde, hasta *Vuelo) ([]Vuelo, error) {
	if K <= 0 || (modo != "asc" && modo != "desc") || (*hasta)[FECHA] < (*desde)[FECHA] {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	var contador int
	contPtr := &contador
	tablero.abb.IterarRango(&desde, &hasta, func(c *Vuelo, d Vuelo) bool {
		if *contPtr >= K {
			return false
		}

		vuelos = append(vuelos, d)
		*contPtr++
		return true
	})
	return vuelos, nil
}

func (tablero *tablero) ObtenerVuelo(codigo string) (Vuelo, error) {
	if !tablero.hash.Pertenece(codigo) {
		err := e.ErrorComando{}
		return Vuelo{}, err
	}
	vuelo := tablero.hash.Obtener(codigo)
	return vuelo, nil
}

func (tablero *tablero) ObtenerVuelosPrioritarios(K int) []Vuelo {
	heap := TDAHeap.CrearHeapArr(tablero.vuelos, cmpPrioridad)
	var vuelos []Vuelo
	for j := 0; j < K; j++ {
		vuelos = append(vuelos, *heap.Desencolar())
	}
	return vuelos
}

func (tablero *tablero) SiguienteVuelo(origen, destino, fecha string) (Vuelo, error) {
	return Vuelo{}, nil
}

func (tablero *tablero) ActualizarTablero(archivo *os.File) {
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		infoLinea := f.SepararEntrada(linea, ",")
		vuelo := Vuelo{infoLinea[CODIGO], infoLinea[AEROLINEA], infoLinea[ORIGEN], infoLinea[DESTINO], infoLinea[NUM_COLA], infoLinea[PRIORIDAD], infoLinea[FECHA], infoLinea[DEMORA], infoLinea[TIEMPO], infoLinea[CANCELADO]}
		tablero.guardar(vuelo, vuelo)
		tablero.vuelos = append(tablero.vuelos, &vuelo)
	}

}

func (tablero *tablero) guardar(clave Vuelo, datos Vuelo) {
	tablero.abb.Guardar(&clave, datos)
	tablero.hash.Guardar(clave[CODIGO], datos)
}

func (tablero *tablero) Borrar(desde, hasta *Vuelo) ([]Vuelo, error) {
	if hasta[FECHA] < desde[FECHA] {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	tablero.abb.IterarRango(&desde, &hasta, func(c *Vuelo, d Vuelo) bool {
		vuelos = append(vuelos, d)
		return true
	})
	for i := 0; i < len(vuelos); i++ {
		tablero.abb.Borrar(&vuelos[i])
		tablero.hash.Borrar(vuelos[i][CODIGO])
		//COMO ELIMINAR LOS VUELOS DE tablero.vuelos?
	}
	return vuelos, nil
}
