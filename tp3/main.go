package main

import (
	"bufio"
	"flycombi/acciones"
	funciones "flycombi/validaciones_y_auxiliares"
	"fmt"
	"os"
)

type indice int

const (
	PARAMETRO_1 indice = iota + 1
	PARAMETRO_2
	PARAMETRO_3
)

const (
	SEPARADOR_1 = ","
	SEPARADOR_2 = " "
)

const (
	INDICE_COMANDO = iota
	INDICE_RUTA_AEROPUERTOS
	INDICE_RUTA_VUELOS

	CANT_PARAMETROS_INICIALES
)

func main() {
	sistema := acciones.CrearBaseDeDatos()
	opciones := acciones.CrearOpciones()

	fmt.Println(sistema, opciones) // BORRAR

	scanner := bufio.NewScanner(os.Stdin)
	parametros := os.Args
	if len(parametros) != CANT_PARAMETROS_INICIALES {
		fmt.Println("Error") // CAMBIAR POR ERROR APROPIADO, PERO NO DICE NADA DE ERRORES.
	}

	rutaDeAeropuertos, rutaDeVuelos := parametros[INDICE_RUTA_AEROPUERTOS], parametros[INDICE_RUTA_VUELOS]
	acciones.GuardarInformacion(sistema, rutaDeAeropuertos, rutaDeVuelos)

	for scanner.Scan() {
		entrada := scanner.Text()
		entradaCompleta, err := funciones.CompletarEntrada(entrada)

		if err != nil {
			fmt.Println("Error") // CAMBIAR POR ERROR APROPIADO...
			continue
		}

		accion := opciones.Obtener(entradaCompleta[INDICE_COMANDO])
		accion(entradaCompleta[PARAMETRO_1], entradaCompleta[PARAMETRO_2], entradaCompleta[PARAMETRO_3])
	}
}
