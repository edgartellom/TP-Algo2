package main

import (
	"bufio"
	acciones "flycombi/acciones"
	errores "flycombi/errores"
	funciones "flycombi/validaciones_y_auxiliares"
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

	scanner := bufio.NewScanner(os.Stdin)
	parametros := os.Args
	if len(parametros) != CANT_PARAMETROS_INICIALES {
		funciones.MostrarError(new(errores.ErrorParametros))
	}

	rutaDeAeropuertos, rutaDeVuelos := parametros[INDICE_RUTA_AEROPUERTOS], parametros[INDICE_RUTA_VUELOS]
	acciones.GuardarInformacion(sistema, rutaDeAeropuertos, rutaDeVuelos)

	for scanner.Scan() {
		entrada := scanner.Text()
		entradaCompleta, err := funciones.CompletarYValidarEntrada(entrada)

		if err != nil {
			funciones.MostrarError(err)
			continue
		}

		accion := opciones.Obtener(entradaCompleta[INDICE_COMANDO])
		accion(sistema, entradaCompleta[PARAMETRO_1], entradaCompleta[PARAMETRO_2], entradaCompleta[PARAMETRO_3])
	}
}
