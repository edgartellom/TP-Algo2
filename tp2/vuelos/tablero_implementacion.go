package vuelos

import (
	"strings"

	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

const COMPARADOR = 0

type tableroDeControl struct {
	arbolDeVuelos TDADicc.DiccionarioOrdenado[CamposComparables, Vuelo]
	diccDeVuelos  TDADicc.Diccionario[Codigo, Vuelo]
	codigoMayor   Codigo
}

func cmpPrioridad(a, b Vuelo) int {
	superior := b.Prioridad - a.Prioridad
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
	if (*tablero).diccDeVuelos.Pertenece(vuelo.InfoComparable.Codigo) {
		vueloActual := tablero.diccDeVuelos.Obtener(vuelo.InfoComparable.Codigo)
		(*tablero).arbolDeVuelos.Borrar(vueloActual.InfoComparable)
	}
	if vuelo.InfoComparable.Codigo > tablero.codigoMayor {
		(*tablero).codigoMayor = vuelo.InfoComparable.Codigo
	}
	(*tablero).diccDeVuelos.Guardar(vuelo.InfoComparable.Codigo, vuelo)
	(*tablero).arbolDeVuelos.Guardar(vuelo.InfoComparable, vuelo)
}

func (tablero *tableroDeControl) ObtenerVuelosEntreRango(desde, hasta string) []Vuelo {
	var vuelos []Vuelo
	fechaDeSalida, fechaDeLlegada := CamposComparables{Fecha: desde}, CamposComparables{Fecha: hasta, Codigo: tablero.codigoMayor}
	tablero.arbolDeVuelos.IterarRango(&fechaDeSalida, &fechaDeLlegada, func(_ CamposComparables, v Vuelo) bool {
		vuelos = append(vuelos, v)
		return true
	})
	return vuelos
}

func (tablero *tableroDeControl) ObtenerVuelo(codigo string) Vuelo {
	return tablero.diccDeVuelos.Obtener(Codigo(codigo))
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
	if k < len(vuelosPrioritarios) {
		return vuelosPrioritarios[:k]
	}
	return vuelosPrioritarios
}

func (tablero *tableroDeControl) Borrar(desde, hasta string) []Vuelo {
	vuelos := (*tablero).ObtenerVuelosEntreRango(desde, hasta)
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
