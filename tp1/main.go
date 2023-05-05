package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	errores "rerepolez/errores"
	funcAux "rerepolez/funciones"
	votos "rerepolez/votos"
	TDACola "tdas/cola"
)

func main() {
	colaDeVotacion := TDACola.CrearColaEnlazada[votos.Votante]()
	var votantesQueYaVotaron []int

	scanner := bufio.NewScanner(os.Stdin)
	var parametros []string = os.Args[1:]
	if len(parametros) < 2 {
		funcAux.MostrarError(new(errores.ErrorParametros))
	}

	rutaPartidos := parametros[0]
	rutaPadrones := parametros[1]

	listaDePartidos := funcAux.ObtenerPartidos(rutaPartidos)
	listaDePadrones := funcAux.ObtenerPadrones(rutaPadrones)

	for scanner.Scan() {
		entrada := scanner.Text()
		entradaSeparada := strings.Split(entrada, " ")
		cmd := entradaSeparada[0]

		switch cmd {
		case "ingresar":
			documento := entradaSeparada[1]
			validez, salida := funcAux.VerificarDNI(documento, listaDePadrones)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			numeroDeDocumento, _ := strconv.Atoi(documento)
			votantesQueYaVotaron = append(votantesQueYaVotaron, numeroDeDocumento)
			votante := votos.CrearVotante(numeroDeDocumento)
			colaDeVotacion.Encolar(votante)
			funcAux.MostrarSalida(salida)
		case "votar":
			tipoVoto := entradaSeparada[1]
			numeroDeLista := entradaSeparada[2]
			validez, salida := funcAux.VerificarVoto(tipoVoto, numeroDeLista, colaDeVotacion, votantesQueYaVotaron, listaDePartidos)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			tipo := funcAux.ConvertirEntradaATipoVoto(tipoVoto)
			alternativa, _ := strconv.Atoi(numeroDeLista)
			votante := colaDeVotacion.VerPrimero()
			votante.Votar(tipo, alternativa)
			funcAux.MostrarSalida(salida)
		case "deshacer":
			validez, salida := funcAux.VerificarColaYVotante(colaDeVotacion, votantesQueYaVotaron)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			votante := colaDeVotacion.VerPrimero()
			err := votante.Deshacer()
			if err != nil {
				funcAux.MostrarSalida(err.Error())
			}
			funcAux.MostrarSalida(salida)
		case "fin-votar":

		}
	}
}

/* --------------------------------------------------------------------------------------------*/

/* MAIN POR SI QUIERES VER QUE FUNCIONA TODO BIEN EL CONSEGUIR LISTA DE PARTIDOS Y DE PADRONES */

/* --------------------------------------------------------------------------------------------*/

// func main() {
// 	TDACola.CrearColaEnlazada[int]()

// 	bufio.NewScanner(os.Stdin)
// 	var args = os.Args[1:]
// 	if len(args) < 2 {
// 		mostrarError(new(errores.ErrorParametros).Error())
// 		os.Exit(1)
// 	}

// 	ruta_partidos := args[0]
// 	ruta_padrones := args[1]

// 	// fmt.Println(scanner)
// 	fmt.Println(ruta_partidos)
// 	lista_partidos := obtenerPartidos(ruta_partidos)
// 	for _, partido := range lista_partidos {
// 		tipo1, tipo2, tipo3 := convertirValor(0), convertirValor(1), convertirValor(2)
// 		fmt.Println(partido.ObtenerResultado(tipo1))
// 		fmt.Println(partido.ObtenerResultado(tipo2))
// 		fmt.Println(partido.ObtenerResultado(tipo3))
// 	}
// 	fmt.Println()
// 	fmt.Println(ruta_padrones)
// 	lista_padrones := obtenerPadrones(ruta_padrones)
// 	fmt.Println(lista_padrones)
// }
