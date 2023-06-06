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
	if (*a).Prioridad > (*b).Prioridad {
		return 1
	}
	if (*a).Prioridad == (*b).Prioridad {
		return strings.Compare((*b).Codigo, (*a).Codigo)
	}
	return -1
}

func cmpTablero(a, b *Vuelo) int {
	if (*a).Fecha > (*b).Fecha {
		return 1
	}
	if (*a).Fecha == (*b).Fecha {
		return strings.Compare((*a).Codigo, (*b).Codigo)
	}
	return -1
}

func CrearTablero() Tablero {
	abb := TDADic.CrearABB[*Vuelo, Vuelo](cmpTablero)
	hash := TDADic.CrearHash[string, Vuelo]()
	return &tablero{abb: abb, hash: hash}
}

func (tablero *tablero) ObtenerVuelos(K int, modo string, desde, hasta *Vuelo) ([]Vuelo, error) {
	if K <= 0 || (modo != "asc" && modo != "desc") || hasta.Fecha < desde.Fecha {
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
		codigo := infoLinea[CODIGO]
		prioridad, _ := strconv.Atoi(infoLinea[PRIORIDAD])
		origen := infoLinea[ORIGEN]
		destino := infoLinea[DESTINO]
		fecha := infoLinea[FECHA]
		vuelo := Vuelo{codigo, prioridad, origen, destino, fecha, infoLinea}
		tablero.guardar(vuelo, vuelo)
		tablero.vuelos = append(tablero.vuelos, &vuelo)
	}

}

func (tablero *tablero) guardar(clave Vuelo, datos Vuelo) {
	tablero.abb.Guardar(&clave, datos)
	tablero.hash.Guardar(clave.Codigo, datos)
}

func (tablero *tablero) Borrar(desde, hasta *Vuelo) ([]Vuelo, error) {
	if hasta.Fecha < desde.Fecha {
		err := e.ErrorComando{}
		return nil, err
	}
	var vuelos []Vuelo
	tablero.abb.IterarRango(&desde, &hasta, func(c *Vuelo, d Vuelo) bool {
		vuelos = append(vuelos, d)
		return true
	})
	for i, vuelo := range vuelos {
		tablero.abb.Borrar(&vuelo)
		tablero.hash.Borrar(vuelo.Codigo)
		tablero.vuelos = append(tablero.vuelos[:i], tablero.vuelos[i+1:]...)
	}
	return vuelos, nil
}
