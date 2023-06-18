package vuelos

import (
	"strconv"
	"strings"

	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

const (
	COMPARADOR        = 0
	SEPARADOR_ENTRADA = ","
	SEPARADOR         = " "
)

type sistemaDeVuelos struct {
	vuelosOrdenados   TDADicc.DiccionarioOrdenado[CamposComparables, Vuelo]
	vuelosAlmacenados TDADicc.Diccionario[Codigo, Vuelo]
	codigoMayorActual Codigo
}

/* -------------------------------------------------- FUNCIONES DE COMPARACION -------------------------------------------------- */

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

/* ---------------------------------------------------- FUNCIONES DE CREACION ---------------------------------------------------- */

func CrearVuelo(infoDeVuelo string) Vuelo {
	informacion := strings.Split(infoDeVuelo, SEPARADOR_ENTRADA)

	prioridad, demora := convertirAInt(informacion[PRIORIDAD]), convertirAInt(informacion[DEMORA])
	informacion[PRIORIDAD], informacion[DEMORA] = strconv.Itoa(prioridad), strconv.Itoa(demora)

	campos := CamposComparables{Prioridad: prioridad, Codigo: Codigo(informacion[CODIGO]), Fecha: informacion[FECHA]}

	return Vuelo{InfoComparable: campos, Origen: informacion[ORIGEN], Destino: informacion[DESTINO], InformacionCompleta: strings.Join(informacion, SEPARADOR)}
}

func CrearSistema() SistemaDeVuelos {
	arbolDeVuelos := TDADicc.CrearABB[CamposComparables, Vuelo](cmpTablero)
	diccDeVuelos := TDADicc.CrearHash[Codigo, Vuelo]()

	return &sistemaDeVuelos{vuelosOrdenados: arbolDeVuelos, vuelosAlmacenados: diccDeVuelos}
}

/* ---------------------------------------------------- PRIMITIVAS DE SISTEMA ---------------------------------------------------- */

func (sistema *sistemaDeVuelos) GuardarVuelo(vuelo Vuelo) {
	if (*sistema).vuelosAlmacenados.Pertenece(vuelo.InfoComparable.Codigo) {
		vueloActual := sistema.vuelosAlmacenados.Obtener(vuelo.InfoComparable.Codigo)
		(*sistema).vuelosOrdenados.Borrar(vueloActual.InfoComparable)
	}
	if vuelo.InfoComparable.Codigo > sistema.codigoMayorActual {
		(*sistema).codigoMayorActual = vuelo.InfoComparable.Codigo
	}
	(*sistema).vuelosAlmacenados.Guardar(vuelo.InfoComparable.Codigo, vuelo)
	(*sistema).vuelosOrdenados.Guardar(vuelo.InfoComparable, vuelo)
}

func (sistema *sistemaDeVuelos) ObtenerVuelosEntreRango(desde, hasta string) []Vuelo {
	var vuelos []Vuelo
	fechaDeSalida, fechaDeLlegada := CamposComparables{Fecha: desde}, CamposComparables{Fecha: hasta, Codigo: sistema.codigoMayorActual}
	sistema.vuelosOrdenados.IterarRango(&fechaDeSalida, &fechaDeLlegada, func(_ CamposComparables, v Vuelo) bool {
		vuelos = append(vuelos, v)
		return true
	})
	return vuelos
}

func (sistema *sistemaDeVuelos) ObtenerVuelo(codigo string) Vuelo {
	return sistema.vuelosAlmacenados.Obtener(Codigo(codigo))
}

func (sistema *sistemaDeVuelos) ObtenerVuelosPrioritarios(k int) []Vuelo {
	var vuelosPrioritarios []Vuelo
	for iter := sistema.vuelosAlmacenados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		vuelosPrioritarios = append(vuelosPrioritarios, vuelo)
	}
	TDAHeap.HeapSort(vuelosPrioritarios, cmpPrioridad)
	if k < len(vuelosPrioritarios) {
		return vuelosPrioritarios[:k]
	}
	return vuelosPrioritarios
}

func (sistema *sistemaDeVuelos) ObtenerSiguienteVuelo(origen, destino, fecha string) *Vuelo {
	var vuelo *Vuelo
	fechaDesde := CamposComparables{Fecha: fecha}
	sistema.vuelosOrdenados.IterarRango(&fechaDesde, nil, func(_ CamposComparables, v Vuelo) bool {
		if v.Origen == origen && v.Destino == destino {
			vuelo = &v
			return false
		}
		return true
	})
	return vuelo
}

func (sistema *sistemaDeVuelos) BorrarVuelos(desde, hasta string) []Vuelo {
	vuelos := (*sistema).ObtenerVuelosEntreRango(desde, hasta)
	for _, vuelo := range vuelos {
		(*sistema).vuelosOrdenados.Borrar(vuelo.InfoComparable)
		(*sistema).vuelosAlmacenados.Borrar(vuelo.InfoComparable.Codigo)
	}
	return vuelos
}

func (sistema *sistemaDeVuelos) Pertenece(numeroDeVuelo Codigo) bool {
	return sistema.vuelosAlmacenados.Pertenece(numeroDeVuelo)
}

/* ---------------------------------------------------- FUNCION AUXILIAR ---------------------------------------------------- */

func convertirAInt(cadena string) int {
	numero, _ := strconv.Atoi(cadena)
	return numero
}
