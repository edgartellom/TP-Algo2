package acciones

import (
	"fmt"
	"strconv"

	"rerepolez/errores"
	"rerepolez/funciones"
	"rerepolez/votos"
	TDACola "tdas/cola"
)

func AccionIngresar(dniIngresado string, filaDeVotacion *TDACola.Cola[votos.Votante], listaDeVotantes []votos.Votante) {
	validezDeEntrada, salida := funciones.VerificarDNI(dniIngresado, listaDeVotantes)
	if validezDeEntrada {
		numeroDeDNI, _ := strconv.Atoi(dniIngresado)
		votante := votos.CrearVotante(numeroDeDNI)
		(*filaDeVotacion).Encolar(votante)
	}
	funciones.MostrarSalida(salida)
}

func AccionVotar(tipoIngresado string, numListaIngresado string, filaDeVotacion *TDACola.Cola[votos.Votante], DNIsQueVotaron []int, partidos []votos.Partido) {
	validezDeEntrada, salida := funciones.VerificarVoto(tipoIngresado, numListaIngresado, *filaDeVotacion, DNIsQueVotaron, partidos)
	if validezDeEntrada {
		tipoDeVoto := funciones.ConvertirEntradaATipoVoto(tipoIngresado)
		numeroDeLista, _ := strconv.Atoi(numListaIngresado)
		votante := (*filaDeVotacion).VerPrimero()
		votante.Votar(tipoDeVoto, numeroDeLista)
	}
	funciones.MostrarSalida(salida)
}

func AccionDeshacer(filaDeVotacion *TDACola.Cola[votos.Votante], DNIsQueVotaron []int) {
	validezDeAccion, salida := funciones.VerificarColaYVotante(*filaDeVotacion, DNIsQueVotaron)
	if validezDeAccion {
		votante := (*filaDeVotacion).VerPrimero()
		err := votante.Deshacer()
		if err != nil {
			salida = err.Error()
		}
	}
	funciones.MostrarSalida(salida)
}

func AccionFinVotar(filaDeVotacion *TDACola.Cola[votos.Votante], DNIsQueVotaron *[]int, listaDePartidos *[]votos.Partido, votosImpugnados *int) {
	validezDeAccion, salida := funciones.VerificarColaYVotante(*filaDeVotacion, *DNIsQueVotaron)
	if validezDeAccion {
		votante := (*filaDeVotacion).Desencolar()
		voto, err := votante.FinVoto()
		if err != nil {
			funciones.MostrarSalida(err.Error())
			return
		}
		*DNIsQueVotaron = append(*DNIsQueVotaron, votante.LeerDNI())
		if voto.Impugnado {
			*votosImpugnados++
		} else {
			for i := 0; i < len(voto.VotoPorTipo); i++ {
				numeroDeLista := voto.VotoPorTipo[i]
				(*listaDePartidos)[numeroDeLista].VotadoPara(votos.TipoVoto(i))
			}
		}
	}
	funciones.MostrarSalida(salida)
}

func MostrarResultadosVotaciones(filaDeVotacion *TDACola.Cola[votos.Votante], listaDePartidos *[]votos.Partido, votosImpugnados *int) {
	if !(*filaDeVotacion).EstaVacia() {
		for !(*filaDeVotacion).EstaVacia() {
			(*filaDeVotacion).Desencolar()
		}
		fmt.Println(new(errores.ErrorCiudadanosSinVotar))
	}
	for i := 0; i < int(votos.CANT_VOTACION); i++ {
		tipoDeVoto := votos.TipoVoto(i)
		funciones.ImprimirTipoCompleto(tipoDeVoto, *listaDePartidos)
	}
	cantidad := funciones.PalabraSegunCantidad(*votosImpugnados)
	fmt.Printf("Votos Impugnados: %d %s\n", *votosImpugnados, cantidad)
}
