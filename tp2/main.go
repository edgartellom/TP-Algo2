package main

import (
	"algueiza/acciones"
	"bufio"
	"os"
	"strings"
)

type indice int

const (
	COMANDO indice = iota
	PARAMETRO_1
	PARAMETRO_2
	PARAMETRO_3
	PARAMETRO_4
)

func main() {
	tablero := acciones.CrearBaseDeDatos()
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := s.Text()
		entradaSeparada := strings.Split(entrada, " ")
		comando := entradaSeparada[COMANDO]
		switch comando {
		case "agregar_archivo":
			acciones.AgregarArchivo(&tablero, entradaSeparada[PARAMETRO_1])
		case "ver_tablero":
			acciones.VerTablero(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])
		case "info_vuelo":
			acciones.InfoVuelo(&tablero, entradaSeparada[PARAMETRO_1])
			// case a.LISTA_COMANDOS[a.AGREGAR_ARCHIVO]:
			// 	a.AgregarArchivo(entradaSeparada[PARAMETRO_1])
			// case a.LISTA_COMANDOS[a.VER_TABLERO]:
			// case a.LISTA_COMANDOS[a.INFO_VUELO]:
			// 	a.InfoVuelo(entradaSeparada[PARAMETRO_1])
			// case a.LISTA_COMANDOS[a.PRIORIDAD_VUELOS]:
			// 	parametro, _ := strconv.Atoi(entradaSeparada[PARAMETRO_1])
			// 	a.PrioridadVuelos(parametro)
			// case a.LISTA_COMANDOS[a.SIGUIENTE_VUELO]:
			// case a.LISTA_COMANDOS[a.BORRAR]:
		}
	}

}
