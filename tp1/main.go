package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	errores "rerepolez/errores"
	funcAux "rerepolez/funciones"
	votos "rerepolez/votos"
	TDACola "tdas/cola"
)

const (
	INGRESAR       = "ingresar"
	VOTAR          = "votar"
	DESHACER       = "deshacer"
	FINALIZAR_VOTO = "fin-votar"
)

func main() {
	colaDeVotacion := TDACola.CrearColaEnlazada[votos.Votante]()
	var votantesQueYaVotaron []int
	var contadorImpugnados int

	scanner := bufio.NewScanner(os.Stdin)
	var parametros []string = os.Args[1:]
	if len(parametros) < 2 {
		funcAux.MostrarError(new(errores.ErrorParametros))
	}

	rutaPartidos := parametros[0]
	rutaPadrones := parametros[1]

	listaDePartidos := funcAux.ObtenerPartidos(rutaPartidos)
	listaDeVotantes := funcAux.ObtenerVotantes(rutaPadrones)

	for scanner.Scan() {
		entrada := scanner.Text()
		entradaSeparada := strings.Split(entrada, " ")
		cmd := entradaSeparada[0]

		switch cmd {
		case INGRESAR:
			documento := entradaSeparada[1]
			validez, salida := funcAux.VerificarDNI(documento, listaDeVotantes)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			numeroDeDocumento, _ := strconv.Atoi(documento)
			votante := votos.CrearVotante(numeroDeDocumento)
			colaDeVotacion.Encolar(votante)
			funcAux.MostrarSalida(salida)

		case VOTAR:
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

		case DESHACER:
			validez, salida := funcAux.VerificarColaYVotante(colaDeVotacion, votantesQueYaVotaron)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			votante := colaDeVotacion.VerPrimero()
			err := votante.Deshacer()
			if err != nil {
				funcAux.MostrarSalida(err.Error())
				continue
			}
			funcAux.MostrarSalida(salida)

		case FINALIZAR_VOTO:
			validez, salida := funcAux.VerificarColaYVotante(colaDeVotacion, votantesQueYaVotaron)
			if !validez {
				funcAux.MostrarSalida(salida)
				continue
			}
			votante := colaDeVotacion.Desencolar()
			voto, err := votante.FinVoto()
			if err != nil {
				funcAux.MostrarSalida(err.Error())
				continue
			}

			votantesQueYaVotaron = append(votantesQueYaVotaron, votante.LeerDNI())
			funcAux.MostrarSalida(salida)

			if voto.Impugnado {
				contadorImpugnados++
				continue
			}
			var tipo votos.TipoVoto
			for i := 0; i < len(voto.VotoPorTipo); i++ {
				numeroDeLista := voto.VotoPorTipo[i]
				listaDePartidos[numeroDeLista].VotadoPara(tipo)
				tipo++
			}
		}
	}
	if !colaDeVotacion.EstaVacia() {
		errror := new(errores.ErrorCiudadanosSinVotar)
		for !colaDeVotacion.EstaVacia() {
			colaDeVotacion.Desencolar()
		}
		fmt.Println(errror)
	}
	var tipo votos.TipoVoto
	for i := tipo; i < 3; i++ {
		funcAux.ImprimirTipoCompleto(tipo, listaDePartidos)
		tipo++
	}
	funcAux.ImprimirImpugnadosSegunCantidad(contadorImpugnados)
}
