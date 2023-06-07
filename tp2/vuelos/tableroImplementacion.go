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
const ASCENDENTE = "asc"
const DESCENDENTE = "desc"

type tablero struct {
	vuelosOrdenadosFechaAsc  TDADic.DiccionarioOrdenado[Claves, Vuelo]
	vuelosOrdenadosFechaDesc TDADic.DiccionarioOrdenado[Claves, Vuelo]
	tableroVuelos            TDADic.Diccionario[Codigo, Vuelo]
	// vuelos []*Vuelo
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

func cmpOrdenadosAsc(a, b Claves) int {
	if a.Fecha > b.Fecha {
		return 1
	}
	if a.Fecha == b.Fecha {
		return strings.Compare(string(a.Codigo), (string(b.Codigo)))
	}
	return -1
}

func cmpOrdenadosDesc(a, b Claves) int {
	if a.Fecha < b.Fecha {
		return 1
	}
	if a.Fecha == b.Fecha {
		return strings.Compare(string(a.Codigo), (string(b.Codigo)))
	}
	return -1
}

func CrearTablero() Tablero {
	vuelosOrdenadosFechaAsc := TDADic.CrearABB[Claves, Vuelo](cmpOrdenadosAsc)
	vuelosOrdenadosFechaDesc := TDADic.CrearABB[Claves, Vuelo](cmpOrdenadosDesc)
	tableroVuelos := TDADic.CrearHash[Codigo, Vuelo]()
	return &tablero{vuelosOrdenadosFechaAsc, vuelosOrdenadosFechaDesc, tableroVuelos}
}

func (tablero *tablero) ObtenerVuelos(K int, modo string, desde, hasta Claves) ([]Vuelo, error) {

	if K <= 0 || (modo != ASCENDENTE && modo != DESCENDENTE) || hasta.Fecha < desde.Fecha {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	var contador int
	contPtr := &contador
	if modo == DESCENDENTE {
		tablero.vuelosOrdenadosFechaDesc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
			if *contPtr >= K {
				return false
			}
			vuelos = append(vuelos, d)
			*contPtr++
			return true
		})

	} else {
		tablero.vuelosOrdenadosFechaAsc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
			if *contPtr >= K {
				return false
			}
			vuelos = append(vuelos, d)
			*contPtr++
			return true
		})
	}
	return vuelos, nil
}

func (tablero *tablero) ObtenerVuelo(codigo Codigo) (Vuelo, error) {
	if !tablero.tableroVuelos.Pertenece(codigo) {
		err := e.ErrorComando{}
		return Vuelo{}, err
	}
	vuelo := tablero.tableroVuelos.Obtener(codigo)
	return vuelo, nil
}

func (tablero *tablero) ObtenerVuelosPrioritarios(K int) []Vuelo {
	heap := TDAHeap.CrearHeap(cmpPrioridad)
	iter := tablero.tableroVuelos.Iterador()
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		heap.Encolar(&vuelo)
		iter.Siguiente()
	}
	var vuelos []Vuelo
	for j := 0; j < K; j++ {
		vuelos = append(vuelos, *heap.Desencolar())
	}
	return vuelos
}

func (tablero *tablero) SiguienteVuelo(origen, destino, Fecha string) (Vuelo, error) {
	return Vuelo{}, nil
}

func (tablero *tablero) ActualizarTablero(archivo *os.File) {
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		infoLinea := f.SepararEntrada(linea, ",")
		vuelo := Vuelo{infoLinea[CODIGO], infoLinea[AEROLINEA], infoLinea[ORIGEN], infoLinea[DESTINO], infoLinea[NUM_COLA], infoLinea[PRIORIDAD], infoLinea[FECHA], infoLinea[DEMORA], infoLinea[TIEMPO], infoLinea[CANCELADO]}
		tablero.guardar(Claves{Fecha: infoLinea[FECHA], Codigo: Codigo(infoLinea[CODIGO]), Origen: infoLinea[ORIGEN], Destino: infoLinea[DESTINO]}, vuelo)
	}

}

func (tablero *tablero) guardar(clave Claves, datos Vuelo) {
	tablero.vuelosOrdenadosFechaAsc.Guardar(clave, datos)
	tablero.vuelosOrdenadosFechaDesc.Guardar(clave, datos)
	tablero.tableroVuelos.Guardar(clave.Codigo, datos)
}

func (tablero *tablero) Borrar(desde, hasta Claves) ([]Vuelo, error) {
	if hasta.Fecha < desde.Fecha {
		err := e.ErrorComando{}
		return nil, err
	}
	var claves []Claves
	var vuelos []Vuelo
	tablero.vuelosOrdenadosFechaAsc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
		claves = append(claves, c)
		return true
	})
	for i := 0; i < len(vuelos); i++ {
		tablero.vuelosOrdenadosFechaAsc.Borrar(claves[i])
		tablero.vuelosOrdenadosFechaDesc.Borrar(claves[i])
		vuelos = append(vuelos, tablero.tableroVuelos.Borrar(claves[i].Codigo))
		//COMO ELIMINAR LOS VUELOS DE tablero.vuelos?
	}
	return vuelos, nil
}
