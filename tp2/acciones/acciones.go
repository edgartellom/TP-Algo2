package acciones

import (
	"algueiza/funciones"
	"algueiza/vuelos"
)

func CrearBaseDeDatos() vuelos.SistemaDeVuelos {
	return vuelos.CrearSistema()
}

func AgregarArchivo(sistema *vuelos.SistemaDeVuelos, ruta string) {
	vuelosEnArchivo, err := funciones.ExtraerInformacion(ruta)
	if err == nil {
		for _, vuelo := range vuelosEnArchivo {
			(*sistema).GuardarVuelo(vuelo)
		}
	}
	funciones.MostrarSalida(err)
}

func VerTablero(sistema *vuelos.SistemaDeVuelos, k, modo, desde, hasta string) {
	cantidad, err := funciones.ComprobarEntradaVerTablero(k, modo, desde, hasta)
	var vuelosEnRango []vuelos.Vuelo
	if err == nil {
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
	}
	funciones.MostrarSalida(err)
}

func InfoVuelo(sistema *vuelos.SistemaDeVuelos, numeroDeVuelo string) {
	err := funciones.ComprobarEntradaInfoVuelo(*sistema, numeroDeVuelo)
	if err == nil {
		vuelo := (*sistema).ObtenerVuelo(numeroDeVuelo)
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(err)
}

func PrioridadVuelos(sistema *vuelos.SistemaDeVuelos, k string) {
	cantidad, err := funciones.ComprobarEntradaDeNumero(k)
	if err == nil {
		vuelosPrioritarios := (*sistema).ObtenerVuelosPrioritarios(cantidad)
		for _, vuelo := range vuelosPrioritarios {
			mensaje := funciones.CrearMensaje(vuelo.InfoComparable.Prioridad, vuelo.InfoComparable.Codigo)
			funciones.MostrarMensaje(mensaje)
		}
	}
	funciones.MostrarSalida(err)
}

func ProximoVuelo(sistema *vuelos.SistemaDeVuelos, origen, destino, fecha string) {
	vuelo := (*sistema).ObtenerSiguienteVuelo(origen, destino, fecha)
	mensaje := funciones.ComprobarVuelo(vuelo, origen, destino, fecha)
	funciones.MostrarMensaje(mensaje)
	funciones.MostrarSalida(nil)
}

func BorrarVuelos(sistema *vuelos.SistemaDeVuelos, desde, hasta string) {
	vuelosBorrados := (*sistema).BorrarVuelos(desde, hasta)
	for _, vuelo := range vuelosBorrados {
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(nil)
}
