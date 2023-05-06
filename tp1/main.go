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
	ACCION_INGRESAR       = "ingresar"
	ACCION_VOTAR          = "votar"
	ACCION_DESHACER       = "deshacer"
	ACCION_FINALIZAR_VOTO = "fin-votar"
)

func main() {
	filaDeVotacion := TDACola.CrearColaEnlazada[votos.Votante]()
	var documentosQueVotaron []int
	var contadorDeVotosImpugnados int

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
		comando := entradaSeparada[0]

		switch comando {
		case ACCION_INGRESAR:
			documentoIngresado := entradaSeparada[1]
			validezDeEntrada, salida := funcAux.VerificarDNI(documentoIngresado, listaDeVotantes)
			if !validezDeEntrada {
				funcAux.MostrarSalida(salida)
				continue
			}
			numeroDeDNI, _ := strconv.Atoi(documentoIngresado)
			votante := votos.CrearVotante(numeroDeDNI)
			filaDeVotacion.Encolar(votante)
			funcAux.MostrarSalida(salida)

		case ACCION_VOTAR:
			tipoIngresado := entradaSeparada[1]
			numDeListaIngresado := entradaSeparada[2]
			validezDeEntrada, salida := funcAux.VerificarVoto(tipoIngresado, numDeListaIngresado, filaDeVotacion, documentosQueVotaron, listaDePartidos)
			if !validezDeEntrada {
				funcAux.MostrarSalida(salida)
				continue
			}
			tipoDeVoto := funcAux.ConvertirEntradaATipoVoto(tipoIngresado)
			numeroDeLista, _ := strconv.Atoi(numDeListaIngresado)
			votante := filaDeVotacion.VerPrimero()
			votante.Votar(tipoDeVoto, numeroDeLista)
			funcAux.MostrarSalida(salida)

		case ACCION_DESHACER:
			validezDeAccion, salida := funcAux.VerificarColaYVotante(filaDeVotacion, documentosQueVotaron)
			if !validezDeAccion {
				funcAux.MostrarSalida(salida)
				continue
			}
			votante := filaDeVotacion.VerPrimero()
			err := votante.Deshacer()
			if err != nil {
				funcAux.MostrarSalida(err.Error())
				continue
			}
			funcAux.MostrarSalida(salida)

		case ACCION_FINALIZAR_VOTO:
			validezDeAccion, salida := funcAux.VerificarColaYVotante(filaDeVotacion, documentosQueVotaron)
			if !validezDeAccion {
				funcAux.MostrarSalida(salida)
				continue
			}
			votante := filaDeVotacion.Desencolar()
			voto, err := votante.FinVoto()
			if err != nil {
				funcAux.MostrarSalida(err.Error())
				continue
			}

			documentosQueVotaron = append(documentosQueVotaron, votante.LeerDNI())
			funcAux.MostrarSalida(salida)

			if voto.Impugnado {
				contadorDeVotosImpugnados++
				continue
			}
			var tipoDeVoto votos.TipoVoto
			for i := 0; i < len(voto.VotoPorTipo); i++ {
				numeroDeLista := voto.VotoPorTipo[i]
				listaDePartidos[numeroDeLista].VotadoPara(tipoDeVoto)
				tipoDeVoto++
			}
		}
	}
	if !filaDeVotacion.EstaVacia() {
		err := new(errores.ErrorCiudadanosSinVotar)
		for !filaDeVotacion.EstaVacia() {
			filaDeVotacion.Desencolar()
		}
		fmt.Println(err)
	}
	var tipoDeVoto votos.TipoVoto
	for tipoDeVoto = 0; tipoDeVoto < 3; tipoDeVoto++ {
		funcAux.ImprimirTipoCompleto(tipoDeVoto, listaDePartidos)
	}
	funcAux.ImprimirImpugnados(contadorDeVotosImpugnados)
}
