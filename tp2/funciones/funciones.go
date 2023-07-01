package funciones

import (
	"algueiza/errores"
	"algueiza/vuelos"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	TDAPila "tdas/pila"
)

const (
	CERO = iota
	UN_PARAMETRO
	DOS_PARAMETROS
	TRES_PARAMETROS
	CUATRO_PARAMETROS
)

const (
	AGREGAR_ARCHIVO = iota
	VER_TABLERO
	INFO_VUELO
	PRIORIDAD_VUELOS
	SIGUIENTE_VUELO
	BORRAR

	CANT_COMANDOS = BORRAR + 1

	MODO_ASCENDETE   = "asc"
	MODO_DESCENDENTE = "desc"
)

var COMANDOS = [CANT_COMANDOS]string{"agregar_archivo", "ver_tablero", "info_vuelo", "prioridad_vuelos", "siguiente_vuelo", "borrar"}

/* ----------------------------------------------------- FUNCIONES AUXILIARES ----------------------------------------------------- */

func abrirArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		err = errores.ErrorComando{Comando: COMANDOS[AGREGAR_ARCHIVO]}
	}
	return archivo, err
}

func cerrarArchivo(archivo *os.File) {
	archivo.Close()
}

func ExtraerInformacion(ruta string) ([]vuelos.Vuelo, error) {
	var vuelosTotales []vuelos.Vuelo

	archivo, err := abrirArchivo(ruta)
	defer cerrarArchivo(archivo)

	if err == nil {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			informacionDeVuelo := scanner.Text()
			vuelo := vuelos.CrearVuelo(informacionDeVuelo)
			vuelosTotales = append(vuelosTotales, vuelo)
		}
	}
	return vuelosTotales, err
}

func InvertirOrden(arreglo []vuelos.Vuelo) {
	pilaAux := TDAPila.CrearPilaDinamica[vuelos.Vuelo]()
	for _, e := range arreglo {
		pilaAux.Apilar(e)
	}
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = pilaAux.Desapilar()
	}
}

func ConvertirAInt(cifra string) int {
	numero, _ := strconv.Atoi(cifra)
	return numero
}

func SepararEntrada(entrada string, separador string) []string {
	return strings.Split(entrada, separador)
}

func CrearMensaje(a, b any) string {
	return fmt.Sprintf("%v - %v", a, b)
}

/* ------------------------------------------------ CON HASH DE COMANDOS Y FUNCIONES ---------------------------------------- */

// func CompletarEntrada(entrada []string) []string {
// 	nuevaEntrada := make([]string, 4)
// 	copy(nuevaEntrada, entrada)
// 	return nuevaEntrada
// }

/* ----------------------------------------------------- FUNCIONES DE SALIDA ----------------------------------------------------- */

func MostrarMensaje(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func MostrarError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

/* --------------------------------------------------- FUNCIONES DE COMPROBACION --------------------------------------------------- */

func ComprobarEntradaDeNumero(cifra string) (int, error) {
	numero, err := strconv.Atoi(cifra)
	if err != nil || numero < CERO {
		err = errores.ErrorComando{Comando: COMANDOS[PRIORIDAD_VUELOS]}
	}
	return numero, err
}

func ComprobarEntradaVerTablero(cantidad, modo, desde, hasta string) (int, error) {
	cant, err := ComprobarEntradaDeNumero(cantidad)
	if (err != nil) || (modo != MODO_ASCENDETE && modo != MODO_DESCENDENTE) {
		err = errores.ErrorComando{Comando: COMANDOS[VER_TABLERO]}
	}
	return cant, err
}

func ComprobarEntradaInfoVuelo(tablero vuelos.SistemaDeVuelos, codigo string) error {
	var err error
	if !tablero.Pertenece(vuelos.Codigo(codigo)) {
		err = errores.ErrorComando{Comando: COMANDOS[INFO_VUELO]}
	}
	return err
}

func ComprobarVuelo(vueloEncontrado *vuelos.Vuelo, origen, destino, fecha string) string {
	if vueloEncontrado == nil {
		return fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", origen, destino, fecha)
	}
	return (*vueloEncontrado).InformacionCompleta
}

func ComprobarEntradaComando(comando string, parametros []string) error {
	var err error
	if comprobarParametros(comando, parametros) {
		err = errores.ErrorComando{Comando: comando}
	}
	return err
}

func comprobarParametros(comando string, parametros []string) bool {
	cantidadParametros := len(parametros)
	if comando == COMANDOS[VER_TABLERO] {
		return cantidadParametros != CUATRO_PARAMETROS
	} else if comando == COMANDOS[SIGUIENTE_VUELO] {
		return cantidadParametros != TRES_PARAMETROS
	} else if comando == COMANDOS[BORRAR] {
		return cantidadParametros != DOS_PARAMETROS
	} else if comando == COMANDOS[AGREGAR_ARCHIVO] || comando == COMANDOS[INFO_VUELO] || comando == COMANDOS[PRIORIDAD_VUELOS] {
		return cantidadParametros != UN_PARAMETRO
	}
	return false
}
