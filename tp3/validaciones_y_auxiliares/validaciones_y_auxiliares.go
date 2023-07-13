package validaciones_y_auxiliares

import (
	"bufio"
	"flycombi/errores"
	aerolineas "flycombi/sistema_aerolineas"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
)

const (
	COMANDO = iota
	UN_PARAMETRO
	DOS_PARAMETROS
	TRES_PARAMETROS
	CUATRO_PARAMETROS
)

const (
	CAMINO_MAS = iota
	CAMINO_ESCALAS
	CENTRALIDAD
	NUEVA_AEROLINEA
	ITINERARIO
	EXPORTAR_KML
)

const (
	LONGITUD_ENTRADA_COMANDO  = 2
	LONGITUD_ENTRADA_COMPLETA = 4
	CANT_COMANDOS             = 6

	SEPARADOR_1 = ","
	SEPARADOR_2 = " "
	TIPO_BARATO = "barato"
	TIPO_RAPIDO = "rapido"
)

const (
	PRIMER_ELEMENTO = iota
	SEGUNDO_ELEMENTO
)

var COMANDOS = [CANT_COMANDOS]string{"camino_mas", "camino_escalas", "centralidad", "nueva_aerolinea", "itinerario", "exportar_kml"}

/* ---------------------------------------------------------- FUNCIONES DE SALIDA ---------------------------------------------------------- */

func MostrarMensaje(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func MostrarError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

func MostrarSalida(camino []aerolineas.Aeropuerto, separador string) {
	salida := make([]string, len(camino))
	for i, aeropuerto := range camino {
		salida[i] = string(aeropuerto.Codigo)
	}
	MostrarMensaje(strings.Join(salida, separador))
}

func CrearMensaje(arreglo []string, separador string) string {
	return strings.Join(arreglo, separador)
}

/* --------------------------------------------------------- FUNCIONES AUXILIARES ---------------------------------------------------------- */

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		MostrarError(new(errores.ErrorLeerArchivo))
	}
	return archivo
}

func crearEntrada(entradaCompleta string) []string {
	var entrada []string
	entradaSeparada1 := strings.Split(entradaCompleta, SEPARADOR_1)
	entradaSeparada2 := strings.SplitN(entradaSeparada1[PRIMER_ELEMENTO], SEPARADOR_2, LONGITUD_ENTRADA_COMANDO)
	entrada = append(entrada, entradaSeparada2...)
	entrada = append(entrada, entradaSeparada1[SEGUNDO_ELEMENTO:]...)
	return entrada
}

func CompletarYValidarEntrada(entradaReal string) ([]string, error) {
	entradaCompleta := make([]string, LONGITUD_ENTRADA_COMPLETA)
	entrada := crearEntrada(entradaReal)

	err := comprobarParametrosDeComando(entrada[COMANDO], entrada[SEGUNDO_ELEMENTO:])

	copy(entradaCompleta, entrada)
	return entradaCompleta, err
}

type aeropuertoCentralidad struct {
	aeropuerto  aerolineas.Aeropuerto
	centralidad float64
}

func ObtenerMasCentrales(cantidad int, diccDeCentralidades TDADicc.Diccionario[aerolineas.Aeropuerto, float64]) []aerolineas.Aeropuerto {
	aeropuertosOrdenados := ordenarPorCentralidad(diccDeCentralidades)
	var aeropuertosMasCentrales []aerolineas.Aeropuerto
	for i := 0; i < len(aeropuertosOrdenados) && i < cantidad; i++ {
		aeropuertosMasCentrales = append(aeropuertosMasCentrales, aeropuertosOrdenados[i].aeropuerto)
	}
	return aeropuertosMasCentrales
}

func ordenarPorCentralidad(centralidades TDADicc.Diccionario[aerolineas.Aeropuerto, float64]) []aeropuertoCentralidad {
	ordenados := make([]aeropuertoCentralidad, centralidades.Cantidad())
	for iter, i := centralidades.Iterador(), 0; iter.HaySiguiente(); iter.Siguiente() {
		v, cent := iter.VerActual()
		ordenados[i] = aeropuertoCentralidad{v, cent}
		i++
	}
	cmp := func(v1, v2 aeropuertoCentralidad) int { return int(v2.centralidad - v1.centralidad) }
	TDAHeap.HeapSort[aeropuertoCentralidad](ordenados, cmp)
	return ordenados
}

/* ------------------------------------------------- FUNCIONES DE EXTRACCION DE INFORMACION ------------------------------------------------ */

func ObtenerAeropuertos(ruta string) []aerolineas.Aeropuerto {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()

	var aeropuertos []aerolineas.Aeropuerto

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		aeropuerto := aerolineas.CrearAeropuerto(linea)
		aeropuertos = append(aeropuertos, aeropuerto)
	}
	return aeropuertos
}

func ObtenerVuelos(ruta string) []aerolineas.Vuelo {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()

	var vuelos []aerolineas.Vuelo

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		vuelo := aerolineas.CrearVuelo(linea)
		vuelos = append(vuelos, vuelo)
	}
	return vuelos
}

func ObtenerCiudadesYRutas(ruta string) ([]aerolineas.Ciudad, []aerolineas.Ruta) {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()

	var ciudadesStr []string
	var rutas []aerolineas.Ruta
	primeraLinea := true

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		linea := scanner.Text()
		if primeraLinea {
			ciudadesStr = strings.Split(linea, SEPARADOR_1)
			primeraLinea = false
		} else {
			rutaStr := strings.Split(linea, SEPARADOR_1)
			ruta := aerolineas.Ruta{
				CiudadOrigen:  aerolineas.Ciudad(rutaStr[PRIMER_ELEMENTO]),
				CiudadDestino: aerolineas.Ciudad(rutaStr[SEGUNDO_ELEMENTO]),
			}
			rutas = append(rutas, ruta)
		}
	}

	ciudades := make([]aerolineas.Ciudad, len(ciudadesStr))
	for i, ciudad := range ciudadesStr {
		ciudades[i] = aerolineas.Ciudad(ciudad)
	}
	return ciudades, rutas
}

/* -------------------------------------------------------- FUNCIONES DE VALIDACION -------------------------------------------------------- */

func comprobarParametrosDeComando(comando string, parametros []string) error {
	var err error
	if (comando == COMANDOS[CAMINO_MAS] && len(parametros) != TRES_PARAMETROS) ||
		(comando == COMANDOS[CAMINO_ESCALAS] && len(parametros) != DOS_PARAMETROS) ||
		(comando == COMANDOS[CENTRALIDAD] ||
			comando == COMANDOS[NUEVA_AEROLINEA] ||
			comando == COMANDOS[ITINERARIO] ||
			comando == COMANDOS[EXPORTAR_KML]) && len(parametros) != UN_PARAMETRO {
		err = errores.ErrorComando{Comando: comando}
	}
	return err
}

func comprobarPertenencia(sistema aerolineas.SistemaDeAerolineas, entradaOrigen, entradaDestino string) bool {
	return sistema.Pertenece(aerolineas.Ciudad(entradaOrigen)) && sistema.Pertenece(aerolineas.Ciudad(entradaDestino))
}

func ComprobarEntradaCaminoEscalas(sistema aerolineas.SistemaDeAerolineas, entradaOrigen, entradaDestino string) error {
	var err error
	if !comprobarPertenencia(sistema, entradaOrigen, entradaDestino) || entradaOrigen == entradaDestino {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_ESCALAS]}
	}
	return err
}

func ComprobarEntradaCaminoMas(sistema aerolineas.SistemaDeAerolineas, tipo, entradaOrigen, entradaDestino string) error {
	var err error
	if tipo != TIPO_BARATO && tipo != TIPO_RAPIDO {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	} else if !comprobarPertenencia(sistema, entradaOrigen, entradaDestino) || entradaOrigen == entradaDestino {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	}
	return err
}

func ComprobarEntradaCentralidad(digito string) (int, error) {
	return strconv.Atoi(digito)
}
