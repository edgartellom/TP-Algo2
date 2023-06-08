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
	vuelosFechaAsc  TDADic.DiccionarioOrdenado[Claves, Vuelo]
	vuelosFechaDesc TDADic.DiccionarioOrdenado[Claves, Vuelo]
	vuelos          TDADic.Diccionario[Codigo, Vuelo]
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
	return -cmpOrdenadosAsc(a, b)
}

func CrearTablero() Tablero {
	vuelosFechaAsc := TDADic.CrearABB[Claves, Vuelo](cmpOrdenadosAsc)
	vuelosFechaDesc := TDADic.CrearABB[Claves, Vuelo](cmpOrdenadosDesc)
	tableroVuelos := TDADic.CrearHash[Codigo, Vuelo]()
	return &tablero{vuelosFechaAsc, vuelosFechaDesc, tableroVuelos}
}

func (tablero *tablero) ObtenerVuelos(K int, modo string, desde, hasta Claves) ([]Vuelo, error) {

	if K <= 0 || (modo != ASCENDENTE && modo != DESCENDENTE) || hasta.Fecha < desde.Fecha {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	var contador int
	contPtr := &contador
	if modo == ASCENDENTE {
		tablero.vuelosFechaAsc.IterarRango(&desde, &hasta, func(c Claves, d Vuelo) bool {
			if *contPtr >= K {
				return false
			}
			vuelos = append(vuelos, d)
			*contPtr++
			return true
		})

	} else {
		tablero.vuelosFechaDesc.IterarRango(&hasta, &desde, func(c Claves, d Vuelo) bool {
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
	if !tablero.vuelos.Pertenece(codigo) {
		err := e.ErrorComando{}
		return Vuelo{}, err
	}
	vuelo := tablero.vuelos.Obtener(codigo)
	return vuelo, nil
}

func (tablero *tablero) ObtenerVuelosPrioritarios(K int) []Vuelo {
	heap := TDAHeap.CrearHeap(cmpPrioridad)
	iter := tablero.vuelos.Iterador()
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

func (tablero *tablero) SiguienteVuelo(origen, destino string, fecha Claves) (Vuelo, error) {
	var vuelo *Vuelo
	tablero.vuelosFechaAsc.IterarRango(&fecha, nil, func(c Claves, d Vuelo) bool {
		if d[ORIGEN] == origen && d[DESTINO] == destino {
			vuelo = &d
			return false
		}
		return true
	})
	if vuelo == nil {
		err := e.ErrorSiguienteVuelo{Origen: origen, Destino: destino, Fecha: fecha.Fecha}
		return Vuelo{}, err
	}
	return *vuelo, nil
}

func (tablero *tablero) ActualizarTablero(archivo *os.File) {
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		infoLinea := f.SepararEntrada(linea, ",")
		vuelo := Vuelo{infoLinea[CODIGO], infoLinea[AEROLINEA], infoLinea[ORIGEN], infoLinea[DESTINO], infoLinea[NUM_COLA], infoLinea[PRIORIDAD], infoLinea[FECHA], infoLinea[DEMORA], infoLinea[TIEMPO], infoLinea[CANCELADO]}
		tablero.guardar(Claves{Fecha: infoLinea[FECHA], Codigo: Codigo(infoLinea[CODIGO])}, vuelo)
	}

}

func (tablero *tablero) guardar(clave Claves, datos Vuelo) {
	tablero.vuelosFechaAsc.Guardar(clave, datos)
	tablero.vuelosFechaDesc.Guardar(clave, datos)
	tablero.vuelos.Guardar(clave.Codigo, datos)
}

func (tablero *tablero) Borrar(desde, hasta Claves) ([]Vuelo, error) {
	if hasta.Fecha < desde.Fecha {
		err := e.ErrorComando{}
		return nil, err
	}
	var claves []Claves
	var vuelos []Vuelo
	iter := tablero.vuelosFechaAsc.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	for i := 0; i < len(vuelos); i++ {
		tablero.vuelosFechaAsc.Borrar(claves[i])
		tablero.vuelosFechaDesc.Borrar(claves[i])
		vuelos = append(vuelos, tablero.vuelos.Borrar(claves[i].Codigo))
	}
	return vuelos, nil
}
