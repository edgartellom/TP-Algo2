package main

import (
	"bufio"
	"os"

	"rerepolez/acciones"
	"rerepolez/errores"
	"rerepolez/funciones"
)

const (
	CANTIDAD_PARAMETROS_NECESARIOS = 3

	INDICE_RUTA_PARTIDOS = 1
	INDICE_RUTA_PADRONES = 2

	INDICE_COMANDO = 0

	INDICE_DNI_INGRESADO       = 1
	INDICE_TIPO_INGRESADO      = 1
	INDICE_NUM_LISTA_INGRESADO = 2

	SEPARADOR_ENTRADA      = " "
	COMANDO_INGRESAR       = "ingresar"
	COMANDO_VOTAR          = "votar"
	COMANDO_DESHACER       = "deshacer"
	COMANDO_FINALIZAR_VOTO = "fin-votar"
)

func main() {
	filaDeVotacion := funciones.CrearFilaDeVotacion()
	var documentosQueVotaron []int
	var contadorDeVotosImpugnados int

	scanner := bufio.NewScanner(os.Stdin)
	var parametros []string = os.Args
	if len(parametros) < CANTIDAD_PARAMETROS_NECESARIOS {
		funciones.MostrarError(new(errores.ErrorParametros))
	}

	rutaPartidos := parametros[INDICE_RUTA_PARTIDOS]
	rutaPadrones := parametros[INDICE_RUTA_PADRONES]

	listaDePartidos := funciones.ObtenerPartidos(rutaPartidos)
	listaDeVotantes := funciones.ObtenerVotantes(rutaPadrones)

	for scanner.Scan() {
		entrada := scanner.Text()
		entradaSeparada := funciones.SepararEntrada(entrada, SEPARADOR_ENTRADA)
		comando := entradaSeparada[INDICE_COMANDO]

		switch comando {
		case COMANDO_INGRESAR:
			acciones.AccionIngresar(entradaSeparada[INDICE_DNI_INGRESADO], &filaDeVotacion, listaDeVotantes)

		case COMANDO_VOTAR:
			acciones.AccionVotar(entradaSeparada[INDICE_TIPO_INGRESADO], entradaSeparada[INDICE_NUM_LISTA_INGRESADO], &filaDeVotacion, documentosQueVotaron, listaDePartidos)

		case COMANDO_DESHACER:
			acciones.AccionDeshacer(&filaDeVotacion, documentosQueVotaron)

		case COMANDO_FINALIZAR_VOTO:
			acciones.AccionFinVotar(&filaDeVotacion, &documentosQueVotaron, &listaDePartidos, &contadorDeVotosImpugnados)
		}
	}
	acciones.MostrarResultadosVotaciones(&filaDeVotacion, &listaDePartidos, &contadorDeVotosImpugnados)
}
