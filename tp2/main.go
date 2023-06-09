package main

import (
	a "algueiza/acciones"
	e "algueiza/errores"
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
		parametros := entradaSeparada[PARAMETRO_1:]
		switch comando {
		case a.LISTA_COMANDOS[a.AGREGAR_ARCHIVO]:
			if len(parametros) != 1 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.AgregarArchivo(entradaSeparada[PARAMETRO_1])
			}
		case a.LISTA_COMANDOS[a.VER_TABLERO]:
			if len(parametros) != 4 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.VerTablero(entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])
			}
		case a.LISTA_COMANDOS[a.INFO_VUELO]:
			if len(parametros) != 1 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.InfoVuelo(entradaSeparada[PARAMETRO_1])
			}
		case a.LISTA_COMANDOS[a.PRIORIDAD_VUELOS]:
			if len(parametros) != 1 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.PrioridadVuelos(entradaSeparada[PARAMETRO_1])
			}
		case a.LISTA_COMANDOS[a.SIGUIENTE_VUELO]:
			if len(parametros) != 3 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.SiguienteVuelo(entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3])
			}
		case a.LISTA_COMANDOS[a.BORRAR]:
			if len(parametros) != 2 {
				f.MostrarError(e.ErrorComando{Comando: comando})
			} else {
				a.Borrar(entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2])
			}
		}
	}

}
