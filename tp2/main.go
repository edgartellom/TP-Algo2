package main

import (
	"algueiza/acciones"
	"algueiza/funciones"
	"bufio"
	"os"
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
		entradaSeparada := funciones.SepararEntrada(entrada, " ")
		comando := entradaSeparada[COMANDO]
		switch comando {
		case "agregar_archivo":
			acciones.AgregarArchivo(&tablero, entradaSeparada[PARAMETRO_1])
		case "ver_tablero":
			acciones.VerTablero(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])
		case "info_vuelo":
			acciones.InfoVuelo(&tablero, entradaSeparada[PARAMETRO_1])
		case "prioridad_vuelos":
			acciones.PrioridadVuelos(&tablero, entradaSeparada[PARAMETRO_1])
		case "siguiente_vuelo":
			acciones.ProximoVuelo(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3])
		case "borrar":
			acciones.BorrarVuelos(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2])
		}
	}

}
