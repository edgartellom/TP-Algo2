package acciones

import (
	"algueiza/funciones"
	"algueiza/vuelos"
)

func CrearBaseDeDatos() vuelos.Sistema {
	return vuelos.CrearSistema()
}

func AgregarArchivo(base *vuelos.Sistema, ruta string) {
	vuelosEnArchivo, err := funciones.ExtraerInformacion(ruta)
	if err == nil {
		for _, vuelo := range vuelosEnArchivo {
			(*base).GuardarVuelo(vuelo)
		}
	}
	funciones.MostrarSalida(err)
}

func VerTablero(base *vuelos.Sistema, k, modo, desde, hasta string) {
	cantidad, err := funciones.ComprobarEntradaVerTablero(k, modo, desde, hasta)
	var vuelosEnRango []vuelos.Vuelo
	if err == nil {
		vuelosEnRango = (*base).ObtenerVuelosEntreRango(desde, hasta)
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

func InfoVuelo(base *vuelos.Sistema, numeroDeVuelo string) {
	err := funciones.ComprobarEntradaInfoVuelo(*base, numeroDeVuelo)
	if err == nil {
		vuelo := (*base).ObtenerVuelo(numeroDeVuelo)
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(err)
}

func PrioridadVuelos(base *vuelos.Sistema, k string) {
	cantidad, err := funciones.ComprobarEntradaDeNumero(k)
	if err == nil {
		vuelosPrioritarios := (*base).ObtenerVuelosPrioritarios(cantidad)
		for _, vuelo := range vuelosPrioritarios {
			mensaje := funciones.CrearMensaje(vuelo.InfoComparable.Prioridad, vuelo.InfoComparable.Codigo)
			funciones.MostrarMensaje(mensaje)
		}
	}
	funciones.MostrarSalida(err)
}

func ProximoVuelo(base *vuelos.Sistema, origen, destino, fecha string) {
	vuelo := (*base).ObtenerSiguienteVuelo(origen, destino, fecha)
	mensaje := funciones.ComprobarVuelo(vuelo, origen, destino, fecha)
	funciones.MostrarMensaje(mensaje)
	funciones.MostrarSalida(nil)
}

func BorrarVuelos(base *vuelos.Sistema, desde, hasta string) {
	vuelosBorrados := (*base).BorrarVuelos(desde, hasta)
	for _, vuelo := range vuelosBorrados {
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(nil)
}
