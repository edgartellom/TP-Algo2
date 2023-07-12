package validaciones_y_auxiliares

import (
	"bufio"
	"flycombi/errores"
	aerolineas "flycombi/sistema_aerolineas"
	"fmt"
	"os"
	"strings"
)

type indice int

const (
	ORIGEN indice = iota
	DESTINO
	TIEMPO_PROMEDIO
	PRECIO
	ESCALAS_ENTRE_AMBOS
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
	LONGITUD_ENTRADA_COMPLETA = 4
	CANT_COMANDOS             = 6

	SEPARADOR_1 string = ","
	SEPARADOR_2 string = " "
)

var COMANDOS = [CANT_COMANDOS]string{"camino_mas", "camino_escalas", "centralidad", "nueva_aerolinea", "itinerario", "exportar_kml"}

/* ---------------------------------------------------------- FUNCIONES DE SALIDA ---------------------------------------------------------- */

func MostrarMensaje(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func MostrarError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

/* --------------------------------------------------------- FUNCIONES AUXILIARES ---------------------------------------------------------- */

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		MostrarError(new(errores.ErrorLeerArchivo))
	}
	return archivo
}

func MostrarCamino(camino []aerolineas.Aeropuerto) {
	salida := make([]string, len(camino))
	for i, aeropuerto := range camino {
		salida[i] = string(aeropuerto.Codigo)
	}
	MostrarMensaje(strings.Join(salida, " -> "))
}

func crearEntrada(entradaCompleta string) []string {
	var entrada []string
	entradaSeparada1 := strings.Split(entradaCompleta, SEPARADOR_1)
	entradaSeparada2 := strings.SplitN(entradaSeparada1[0], SEPARADOR_2, 2)
	entrada = append(entrada, entradaSeparada2...)
	entrada = append(entrada, entradaSeparada1[1:]...)
	return entrada
}

func CompletarYValidarEntrada(entradaReal string) ([]string, error) {
	entradaCompleta := make([]string, LONGITUD_ENTRADA_COMPLETA)
	entrada := crearEntrada(entradaReal)

	err := comprobarParametrosDeComando(entrada[COMANDO], entrada[1:])

	copy(entradaCompleta, entrada)
	return entradaCompleta, err
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

/* -------------------------------------------------------- FUNCIONES DE VALIDACION -------------------------------------------------------- */

func comprobarParametrosDeComando(comando string, parametros []string) error {
	var err error
	if (comando == COMANDOS[CAMINO_MAS] && len(parametros) != TRES_PARAMETROS) || (comando == COMANDOS[CAMINO_ESCALAS] && len(parametros) != DOS_PARAMETROS) ||
		(comando == COMANDOS[CENTRALIDAD] || comando == COMANDOS[NUEVA_AEROLINEA] || comando == COMANDOS[ITINERARIO] || comando == COMANDOS[EXPORTAR_KML]) && len(parametros) != UN_PARAMETRO {
		err = errores.ErrorComando{Comando: comando}
	}
	return err
}

func ComprobarEntradaCaminoMas(sistema aerolineas.SistemaDeAerolineas, tipo, entradaOrigen, entradaDestino string) error {
	var err error
	if tipo != "barato" && tipo != "rapido" {
		MostrarMensaje("error en el tipo")
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	} else if !sistema.Pertenece(aerolineas.Ciudad(entradaOrigen)) || !sistema.Pertenece(aerolineas.Ciudad(entradaDestino)) {
		MostrarMensaje("error que no pertenece")
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	}
	return err
}
