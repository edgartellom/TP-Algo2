package main

import (
	"algueiza/acciones"
	funciones "algueiza/validaciones_y_auxiliares"
	"bufio"
	"os"
)

type indice int

const COMANDO indice = 0

const (
	PARAMETRO_0 = iota
	PARAMETRO_1
	PARAMETRO_2
	PARAMETRO_3

	SEPARADOR = " "
)

func main() {
	sistema := acciones.CrearBaseDeDatos()
	opciones := acciones.CrearOpciones()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := s.Text()
		entradaSeparada := funciones.SepararEntrada(entrada, SEPARADOR)
		comando := entradaSeparada[COMANDO]
		err := funciones.ComprobarEntradaComando(comando, entradaSeparada[PARAMETRO_1:])

		if err != nil {
			funciones.MostrarError(err)
			continue
		}
		accion := opciones.Obtener(comando)
		nuevaEntrada := funciones.CompletarEntrada(entradaSeparada[PARAMETRO_1:])
		accion(&sistema, nuevaEntrada[PARAMETRO_0], nuevaEntrada[PARAMETRO_1], nuevaEntrada[PARAMETRO_2], nuevaEntrada[PARAMETRO_3])
	}
}
