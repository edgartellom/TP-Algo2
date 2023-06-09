package vuelos

import (
	"strings"

	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

const COMPARADOR = 0

type tableroDeControl struct {
	arbolDeVuelos   TDADicc.DiccionarioOrdenado[CamposComparables, Vuelo]
	diccDeVuelos    TDADicc.Diccionario[Codigo, Vuelo]
	listadoDeVuelos []Vuelo
}

func cmpPrioridad(a, b Vuelo) int {
	superior := b.InfoComparable.Prioridad - a.InfoComparable.Prioridad
	if superior == COMPARADOR {
		return strings.Compare(string(a.InfoComparable.Codigo), string(b.InfoComparable.Codigo))
	}
	return superior
}

func cmpTablero(a, b CamposComparables) int {
	superior := strings.Compare(a.Fecha, b.Fecha)
	if superior == COMPARADOR {
		return strings.Compare(string(a.Codigo), string(b.Codigo))
	}
	return superior
}

func CrearTablero() Tablero {
	arbolDelTablero := TDADicc.CrearABB[CamposComparables, Vuelo](cmpTablero)
	diccDelTablero := TDADicc.CrearHash[Codigo, Vuelo]()
	return &tableroDeControl{arbolDeVuelos: arbolDelTablero, diccDeVuelos: diccDelTablero}
}

func (tablero *tableroDeControl) GuardarVuelo(vuelo Vuelo) {
	(*tablero).diccDeVuelos.Guardar(vuelo.InfoComparable.Codigo, vuelo)
	(*tablero).arbolDeVuelos.Guardar(vuelo.InfoComparable, vuelo)
	(*tablero).listadoDeVuelos = append((*tablero).listadoDeVuelos, vuelo)
}

func (tablero *tableroDeControl) ObtenerVuelosEntreRango(k int, desde, hasta string) []Vuelo {
	var vuelos []Vuelo
	fechaDeSalida, fechaDeLlegada := CamposComparables{Fecha: desde}, CamposComparables{Fecha: hasta}
	for iter := tablero.arbolDeVuelos.IteradorRango(&fechaDeSalida, &fechaDeLlegada); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelos = append(vuelos, vuelo)
	}
	return vuelos
}

func (tablero *tableroDeControl) ObtenerVuelo(codigo string) Vuelo {
	vuelo := tablero.diccDeVuelos.Obtener(Codigo(codigo))
	return vuelo
}

func (tablero *tableroDeControl) Pertenece(numeroDeVuelo string) bool {
	return tablero.diccDeVuelos.Pertenece(Codigo(numeroDeVuelo))
}

func (tablero *tableroDeControl) ObtenerVuelosPrioritarios(k int) []Vuelo {
	var vuelosPrioritarios []Vuelo
	for iter := tablero.diccDeVuelos.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelosPrioritarios = append(vuelosPrioritarios, vuelo)
	}
	TDAHeap.HeapSort(vuelosPrioritarios, cmpPrioridad)
	return vuelosPrioritarios[:k]
}

func (tablero *tableroDeControl) Borrar(desde, hasta string) []Vuelo {
	var vuelos []Vuelo
	fechaDesde, fechaHasta := CamposComparables{Fecha: desde}, CamposComparables{Fecha: hasta}
	for iter := tablero.arbolDeVuelos.IteradorRango(&fechaDesde, &fechaHasta); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelos = append(vuelos, vuelo)
	}
	for _, vuelo := range vuelos {
		(*tablero).arbolDeVuelos.Borrar(vuelo.InfoComparable)
		(*tablero).diccDeVuelos.Borrar(vuelo.InfoComparable.Codigo)
	}
	return vuelos
}

func (tablero *tableroDeControl) ObtenerSiguienteVuelo(origen, destino, fecha string) *Vuelo {
	var vuelo *Vuelo
	fechaDesde := CamposComparables{Fecha: fecha}
	tablero.arbolDeVuelos.IterarRango(&fechaDesde, nil, func(_ CamposComparables, v Vuelo) bool {
		if v.Origen == origen && v.Destino == destino {
			vuelo = &v
			return false
		}
		return true
	})
	return vuelo
}
