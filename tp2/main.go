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

/*
CAMBIOS QUE HICE:
	- Los cambios para usar un Hash, todo lo dejo comentado para que lo pruebes en (main.go, acciones.go y funciones.go)

	- Algo que me falto fue no usar numeros magicos en la línea 75 y 76 de este archivo.

	- Separe las funciones y las llamé: "MostrarMensaje(mensaje)" y "MostrarError(error)".

	- La verificación del error antes de la funcionalidad funciona de las 2 formas;
	agregando un return o poniendo la funcionalidad dentro de un else...
	(igual te dejo la opcion del else, y la posicion del "return" comentado, elige la que te parezca mejor)
*/

func main() {
	sistema := acciones.CrearBaseDeDatos()
	// opciones := acciones.CrearOpciones()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		entrada := s.Text()
		entradaSeparada := funciones.SepararEntrada(entrada, SEPARADOR)
		comando := entradaSeparada[COMANDO]
		err := funciones.ComprobarEntradaComando(comando, entradaSeparada[PARAMETRO_1:])

		switch {
		case err != nil:
			funciones.MostrarError(err)

		case comando == funciones.COMANDOS[funciones.AGREGAR_ARCHIVO]:
			acciones.AgregarArchivo(&sistema, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.VER_TABLERO]:
			acciones.VerTablero(&sistema, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3], entradaSeparada[PARAMETRO_4])

		case comando == funciones.COMANDOS[funciones.INFO_VUELO]:
			acciones.InfoVuelo(&sistema, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.PRIORIDAD_VUELOS]:
			acciones.PrioridadVuelos(&sistema, entradaSeparada[PARAMETRO_1])

		case comando == funciones.COMANDOS[funciones.SIGUIENTE_VUELO]:
			acciones.ProximoVuelo(&sistema, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2], entradaSeparada[PARAMETRO_3])

		case comando == funciones.COMANDOS[funciones.BORRAR]:
			acciones.BorrarVuelos(&sistema, entradaSeparada[PARAMETRO_1], entradaSeparada[PARAMETRO_2])
		}

		/* ------------------------------------------------ CON HASH DE COMANDOS Y FUNCIONES ---------------------------------------- */
		// if err != nil {
		// 	funciones.MostrarError(err)
		// } else {
		// 	accion := opciones.Obtener(comando)
		// 	nuevaEntrada := funciones.CompletarEntrada(entradaSeparada[1:])
		// 	accion(&sistema, nuevaEntrada[0], nuevaEntrada[1], nuevaEntrada[2], nuevaEntrada[3])
		// }

	}
}
