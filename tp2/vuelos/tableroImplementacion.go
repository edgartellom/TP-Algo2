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

func cmpPrioridad(a, b Vuelo) int {
	cmp := b.Prioridad - a.Prioridad

	if cmp == COMPARADOR {
		return strings.Compare(string(a.Claves.Codigo), string(b.Claves.Codigo))
	}
	return cmp
}

func cmpOrdenadosAsc(a, b Claves) int {
	if a.Fecha.After(b.Fecha) {
		return 1
	}

	if a.Fecha.Before(b.Fecha) {
		return -1
	}
	return strings.Compare(string(a.Codigo), (string(b.Codigo)))
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

	if K <= 0 || (modo != ASCENDENTE && modo != DESCENDENTE) {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	var contador int
	contPtr := &contador
	if !(cmpOrdenadosAsc(hasta, desde) < COMPARADOR) {
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
	}
	return vuelos, nil
}

func (tablero *tablero) ObtenerVuelo(codigo Codigo) (*Vuelo, error) {
	if !tablero.vuelos.Pertenece(codigo) {
		err := e.ErrorComando{}
		return nil, err
	}
	vuelo := tablero.vuelos.Obtener(codigo)
	return &vuelo, nil
}

func (tablero *tablero) ObtenerVuelosPrioritarios(K int) []Vuelo {
	var vuelosPrioritarios []Vuelo
	// heap := TDAHeap.CrearHeap(cmpPrioridad)
	for iter := tablero.vuelos.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelosPrioritarios = append(vuelosPrioritarios, vuelo)
		// heap.Encolar(&vuelo)
	}
	// var vuelos []Vuelo
	TDAHeap.HeapSort(vuelosPrioritarios, cmpPrioridad)
	if K < len(vuelosPrioritarios) {
		return vuelosPrioritarios[:K]
	}
	return vuelosPrioritarios
}

func (tablero *tablero) SiguienteVuelo(origen, destino string, clavesDesde Claves) (*Vuelo, error) {
	var vuelo *Vuelo

	tablero.vuelosFechaAsc.IterarRango(&clavesDesde, nil, func(c Claves, v Vuelo) bool {
		if v.Origen == origen && v.Destino == destino {
			vuelo = &v
			return false
		}
		return true
	})
	if vuelo == nil {
		cadenaFecha := f.ConvertirFechaACadena(clavesDesde.Fecha)
		err := e.ErrorSiguienteVuelo{Origen: origen, Destino: destino, Fecha: cadenaFecha}
		return nil, err
	}
	return vuelo, nil
}

func (tablero *tablero) ActualizarTablero(archivo *os.File) {
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		infoLinea := f.SepararEntrada(linea, ",")
		fecha := f.ConvertirCadenaAFecha(infoLinea[FECHA])
		claves := Claves{Codigo: Codigo(infoLinea[CODIGO]), Fecha: fecha}
		prioridad, _ := strconv.Atoi(infoLinea[PRIORIDAD])
		demora, _ := strconv.Atoi(infoLinea[DEMORA])
		tiempo, _ := strconv.Atoi(infoLinea[TIEMPO])
		cancelado, _ := strconv.Atoi(infoLinea[CANCELADO])
		vuelo := Vuelo{claves, infoLinea[AEROLINEA], infoLinea[ORIGEN], infoLinea[DESTINO],
			infoLinea[NUM_COLA], prioridad, demora, tiempo, cancelado}
		tablero.guardarVuelos(Claves{Fecha: fecha, Codigo: Codigo(infoLinea[CODIGO])}, vuelo)
	}

}

func (tablero *tablero) guardarVuelos(claves Claves, datos Vuelo) {
	vuelo, _ := tablero.ObtenerVuelo(claves.Codigo)
	if vuelo != nil {
		claves := Claves{Codigo: (*vuelo).Claves.Codigo, Fecha: (*vuelo).Claves.Fecha}
		tablero.vuelosFechaAsc.Borrar(claves)
		tablero.vuelosFechaDesc.Borrar(claves)
	}

	tablero.vuelos.Guardar(claves.Codigo, datos)
	tablero.vuelosFechaAsc.Guardar(claves, datos)
	tablero.vuelosFechaDesc.Guardar(claves, datos)
}

func (tablero *tablero) BorrarVuelos(clavesDesde, clavesHasta Claves) ([]Vuelo, error) {
	// if clavesHasta.Fecha.Before(clavesDesde.Fecha) {
	// 	err := e.ErrorComando{}
	// 	return nil, err
	// }
	var conjClaves []Claves
	var vuelos []Vuelo
	if clavesDesde.Fecha.Before(clavesHasta.Fecha) || clavesDesde.Fecha.Equal(clavesHasta.Fecha) {
		// iter := tablero.vuelosFechaAsc.IteradorRango(&clavesDesde, &clavesHasta)
		// for iter.HaySiguiente() {
		// 	clave, vuelo := iter.VerActual()
		// 	claves = append(claves, clave)
		// 	vuelos = append(vuelos, vuelo)
		// 	iter.Siguiente()
		// }
		tablero.vuelosFechaAsc.IterarRango(&clavesDesde, &clavesHasta, func(c Claves, v Vuelo) bool {
			conjClaves = append(conjClaves, c)
			vuelos = append(vuelos, v)
			return true
		})
		for _, claves := range conjClaves {
			tablero.vuelosFechaAsc.Borrar(claves)
			tablero.vuelosFechaDesc.Borrar(claves)
			tablero.vuelos.Borrar(claves.Codigo)
		}
	}
	return vuelos, nil
}
