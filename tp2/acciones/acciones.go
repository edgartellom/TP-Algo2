package acciones

import (
	funciones "algueiza/validaciones_y_auxiliares"
	"algueiza/vuelos"
	TDADicc "tdas/diccionario"
)

type Comandos func(*vuelos.SistemaDeVuelos, string, string, string, string)

const SALIDA_EXITOSA = "OK"

var acciones = [funciones.CANT_COMANDOS]Comandos{AgregarArchivo, VerTablero, InfoVuelo, PrioridadVuelos, ProximoVuelo, BorrarVuelos}

func CrearOpciones() TDADicc.Diccionario[string, Comandos] {
	opciones := TDADicc.CrearHash[string, Comandos]()
	for i := 0; i < funciones.CANT_COMANDOS; i++ {
		opciones.Guardar(funciones.COMANDOS[i], acciones[i])
	}
	return opciones
}

func CrearBaseDeDatos() vuelos.SistemaDeVuelos {
	return vuelos.CrearSistema()
}

func AgregarArchivo(sistema *vuelos.SistemaDeVuelos, ruta, _, _, _ string) {
	vuelosEnArchivo, err := funciones.ExtraerInformacion(ruta)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	for _, vuelo := range vuelosEnArchivo {
		(*sistema).GuardarVuelo(vuelo)
	}
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func VerTablero(sistema *vuelos.SistemaDeVuelos, k, modo, desde, hasta string) {
	cantidad, err := funciones.ComprobarEntradaVerTablero(k, modo, desde, hasta)
	var vuelosEnRango []vuelos.Vuelo
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	vuelosEnRango = (*sistema).ObtenerVuelosEntreRango(desde, hasta)
	if modo == funciones.MODO_DESCENDENTE {
		funciones.InvertirOrden(vuelosEnRango)
	}
	if cantidad < len(vuelosEnRango) {
		vuelosEnRango = vuelosEnRango[:cantidad]
	}
	for _, vuelo := range vuelosEnRango {
		mensaje := funciones.CrearMensaje(vuelo.InfoComparable.Fecha, vuelo.InfoComparable.Codigo)
		funciones.MostrarMensaje(mensaje)
	}
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func InfoVuelo(sistema *vuelos.SistemaDeVuelos, numeroDeVuelo, _, _, _ string) {
	err := funciones.ComprobarEntradaInfoVuelo(*sistema, numeroDeVuelo)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	vuelo := (*sistema).ObtenerVuelo(numeroDeVuelo)
	funciones.MostrarMensaje(vuelo.InformacionCompleta)
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func PrioridadVuelos(sistema *vuelos.SistemaDeVuelos, k, _, _, _ string) {
	cantidad, err := funciones.ComprobarEntradaDeNumero(k)
	if err != nil {
		funciones.MostrarError(err)
		return
	}
	vuelosPrioritarios := (*sistema).ObtenerVuelosPrioritarios(cantidad)
	for _, vuelo := range vuelosPrioritarios {
		mensaje := funciones.CrearMensaje(vuelo.InfoComparable.Prioridad, vuelo.InfoComparable.Codigo)
		funciones.MostrarMensaje(mensaje)
	}
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func ProximoVuelo(sistema *vuelos.SistemaDeVuelos, origen, destino, fecha, _ string) {
	vuelo := (*sistema).ObtenerSiguienteVuelo(origen, destino, fecha)
	mensaje := funciones.ComprobarVuelo(vuelo, origen, destino, fecha)
	funciones.MostrarMensaje(mensaje)
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}

func BorrarVuelos(sistema *vuelos.SistemaDeVuelos, desde, hasta, _, _ string) {
	vuelosBorrados := (*sistema).BorrarVuelos(desde, hasta)
	for _, vuelo := range vuelosBorrados {
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarMensaje(SALIDA_EXITOSA)
}
