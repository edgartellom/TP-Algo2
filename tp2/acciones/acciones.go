package acciones

import (
	"algueiza/funciones"
	"algueiza/vuelos"
)

func CrearBaseDeDatos() vuelos.Tablero {
	return vuelos.CrearTablero()
}

func AgregarArchivo(base *vuelos.Tablero, ruta string) {
	vuelosEnArchivo, err := funciones.ExtraerInformacion(ruta)
	if err == nil {
		for _, vuelo := range vuelosEnArchivo {
			(*base).GuardarVuelo(vuelo)
		}
	}
	funciones.MostrarSalida(err)
}

func InfoVuelo(base *vuelos.Tablero, numeroDeVuelo string) {
	validezDeEntrada, err := funciones.ComprobarEntradaInfoVuelo(*base, numeroDeVuelo)
	if validezDeEntrada {
		vuelo := (*base).ObtenerVuelo(numeroDeVuelo)
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(err)
}

func PrioridadVuelos(base *vuelos.Tablero, k string) {
	cantidad, err := funciones.ComprobarEntradaDeNumero(k)
	if err == nil {
		vuelos := (*base).ObtenerVuelosPrioritarios(cantidad)
		for _, vuelo := range vuelos {
			mensaje := funciones.CrearMensaje(vuelo.Prioridad, vuelo.InfoComparable.Codigo)
			funciones.MostrarMensaje(mensaje)
		}
	}
	funciones.MostrarSalida(err)
}

func VerTablero(base *vuelos.Tablero, k, modo, desde, hasta string) {
	cantidad, err := funciones.ComprobarEntradaVerTablero(k, modo, desde, hasta)
	var vuelos []vuelos.Vuelo
	if err == nil {
		vuelos = (*base).ObtenerVuelosEntreRango(desde, hasta)
		if modo == funciones.MODO_DESCENDENTE {
			funciones.InvertirOrden(vuelos)
		}
		if cantidad < len(vuelos) {
			vuelos = vuelos[:cantidad]
		}
		for _, vuelo := range vuelos {
			mensaje := funciones.CrearMensaje(vuelo.InfoComparable.Fecha, vuelo.InfoComparable.Codigo)
			funciones.MostrarMensaje(mensaje)
		}
	}
	funciones.MostrarSalida(err)
}

func BorrarVuelos(base *vuelos.Tablero, desde, hasta string) {
	vuelosBorrados := (*base).Borrar(desde, hasta)
	for _, vuelo := range vuelosBorrados {
		funciones.MostrarMensaje(vuelo.InformacionCompleta)
	}
	funciones.MostrarSalida(nil)
}

func ProximoVuelo(base *vuelos.Tablero, origen, destino, fecha string) {
	vuelo := (*base).ObtenerSiguienteVuelo(origen, destino, fecha)
	mensaje := funciones.ComprobarVuelo(vuelo, origen, destino, fecha)
	funciones.MostrarMensaje(mensaje)
	funciones.MostrarSalida(nil)
}
