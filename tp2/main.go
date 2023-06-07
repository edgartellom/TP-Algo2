package main

import (
	a "algueiza/acciones"
	f "algueiza/funciones"
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
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := s.Text()
		entradaSeparada := f.SepararEntrada(entrada, " ")
		comando := entradaSeparada[COMANDO]
		switch comando {
		case a.LISTA_COMANDOS[a.AGREGAR_ARCHIVO]:
			a.AgregarArchivo(entradaSeparada[PARAMETRO_1])
		case a.LISTA_COMANDOS[a.VER_TABLERO]:
			a.VerTablero(entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])
		case a.LISTA_COMANDOS[a.INFO_VUELO]:
			a.InfoVuelo(entradaSeparada[PARAMETRO_1])
		case a.LISTA_COMANDOS[a.PRIORIDAD_VUELOS]:
			a.PrioridadVuelos(entradaSeparada[PARAMETRO_1])
		case a.LISTA_COMANDOS[a.SIGUIENTE_VUELO]:
		case a.LISTA_COMANDOS[a.BORRAR]:
		}
	}

}
