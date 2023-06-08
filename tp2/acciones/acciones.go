package acciones

import (
	"algueiza/funciones"
	"algueiza/vuelos"
	"fmt"
)

const (
	SALIDA_EXITOSA = "OK"
)

func CrearBaseDeDatos() vuelos.Tablero {
	return vuelos.CrearTablero()
}

func AgregarArchivo(base *vuelos.Tablero, ruta string) {
	(*base).CargarInformacion(ruta)
	funciones.MostrarSalida(SALIDA_EXITOSA)
}

func InfoVuelo(base *vuelos.Tablero, numeroDeVuelo string) {
	vuelo, err := (*base).ObtenerVuelo(numeroDeVuelo)
	if err != nil {
		funciones.MostrarSalida(err.Error())
	} else {
		mensaje := fmt.Sprintf("%s\n", (*vuelo).ObtenerInformacionDeVuelo())
		funciones.MostrarSalida(mensaje)
		funciones.MostrarSalida(SALIDA_EXITOSA)
	}
}

func VerTablero(base *vuelos.Tablero, k, modo, desde, hasta string) {
	cantidadAMostrar, err := funciones.ComprobarEntradaVerTablero(k, modo, desde, hasta)
	var vuelos []vuelos.Vuelo
	if err == nil {
		vuelos = (*base).ObtenerVuelosEntreRango(cantidadAMostrar, desde, hasta)
		if modo == funciones.MODO_DESCENDENTE {
			funciones.InvertirOrden(vuelos)
		}
		for _, vuelo := range vuelos {
			mensaje := fmt.Sprintf("%s - %s", vuelo.VerFecha(), vuelo.VerCodigo())
			funciones.MostrarSalida(mensaje)
		}
		funciones.MostrarSalida(SALIDA_EXITOSA)
	}
	funciones.MostrarSalida(err.Error())
}
