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
	SALIDA_EXITOSA   = "OK"
	MODO_ASCENDETE   = "asc"
	MODO_DESCENDENTE = "desc"
)

const (
	A_ARCHIVO = "agregar_archivo"
	V_VUELOS  = "ver_tablero"
	I_VUELO   = "info_vuelo"
	P_VUELO   = "prioridad_vuelos"
	S_VUELO   = "siguiente_vuelo"
	B_VUELOS  = "borrar"
)

const (
	AGREGAR_ARCHIVO = iota
	VER_TABLERO
	INFO_VUELO
	PRIORIDAD_VUELOS
	SIGUIENTE_VUELO
	BORRAR
)

const CANT_COMANDOS = BORRAR + 1

var LISTA_COMANDOS = [CANT_COMANDOS]string{"agregar_archivo", "ver_tablero", "info_vuelo", "prioridad_vuelos", "siguiente_vuelo", "borrar"}

/* -------------------------------------------------- FUNCION AUX -------------------------------------------------- */

func ComprobarVuelo(vueloEncontrado *vuelos.Vuelo, origen, destino, fecha string) string {
	if vueloEncontrado == nil {
		return fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", origen, destino, fecha)
	}
	return (*vueloEncontrado).InformacionCompleta
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

func ComprobarEntradaComando(comando string, parametros []string) error {
	var err error
	if (comando == V_VUELOS && len(parametros) != 4) || (comando == S_VUELO && len(parametros) != 3) || (comando == B_VUELOS && len(parametros) != 2) || ((comando == A_ARCHIVO || comando == I_VUELO || comando == P_VUELO) && len(parametros) != 1) {
		err = errores.ErrorComando{Comando: comando}
	}
	return err
}

/* -------------------------------------------------- FUNCIONES DE SALIDA -------------------------------------------------- */

func MostrarMensaje(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func MostrarSalida(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		MostrarMensaje(SALIDA_EXITOSA)
	}
}

func abrirArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		err = errores.ErrorComando{Comando: LISTA_COMANDOS[AGREGAR_ARCHIVO]}
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

/* -------------------------------------------------- FUNCIONES DE COMPROBACION -------------------------------------------------- */

func ComprobarEntradaDeNumero(cifra string) (int, error) {
	numero, err := strconv.Atoi(cifra)
	if err != nil || numero < 0 {
		err = errores.ErrorComando{Comando: LISTA_COMANDOS[PRIORIDAD_VUELOS]}
	}
	return numero, err
}

func ComprobarEntradaDeRango(desde, hasta string) error {
	// var err error
	// if hasta < desde {
	// 	err = errores.ErrorComando{Comando: LISTA_COMANDOS[BORRAR]}
	// }
	// return err
	if strings.Compare(hasta, desde) < 0 {
		return errores.ErrorComando{Comando: LISTA_COMANDOS[BORRAR]}
	}
	return nil
}

func ComprobarEntradaVerTablero(cantidad, modo, desde, hasta string) (int, error) {
	cant, err := strconv.Atoi(cantidad)
	if (err != nil) || (modo != MODO_ASCENDETE && modo != MODO_DESCENDENTE) {
		err = errores.ErrorComando{Comando: LISTA_COMANDOS[VER_TABLERO]}
		return -1, err
	} //else if ComprobarEntradaDeRango(desde, hasta) != nil {
	// 	err = errores.ErrorComando{Comando: LISTA_COMANDOS[VER_TABLERO]}
	// }
	return cant, nil
}

func ComprobarEntradaInfoVuelo(tablero vuelos.Tablero, codigo string) (bool, error) {
	var err error
	pertenece := tablero.Pertenece(codigo)
	if !pertenece {
		err = errores.ErrorComando{Comando: LISTA_COMANDOS[INFO_VUELO]}
	}
	return pertenece, err
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
