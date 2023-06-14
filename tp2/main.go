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

	SEPARADOR = " "
)

func main() {
	tablero := acciones.CrearBaseDeDatos()
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := s.Text()
		entradaSeparada := funciones.SepararEntrada(entrada, SEPARADOR)
		comando := entradaSeparada[COMANDO]
		err := funciones.ComprobarEntradaComando(comando, entradaSeparada[PARAMETRO_1:])

		switch {
		case err != nil:
			funciones.MostrarSalida(err)

		case comando == funciones.COMANDOS[funciones.AGREGAR_ARCHIVO]:
			acciones.AgregarArchivo(&tablero, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.VER_TABLERO]:
			acciones.VerTablero(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])

		case comando == funciones.COMANDOS[funciones.INFO_VUELO]:
			acciones.InfoVuelo(&tablero, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.PRIORIDAD_VUELOS]:
			acciones.PrioridadVuelos(&tablero, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.SIGUIENTE_VUELO]:
			acciones.ProximoVuelo(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3])

		case comando == funciones.COMANDOS[funciones.BORRAR]:
			acciones.BorrarVuelos(&tablero, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2])
		}
	}

}
