package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// e "rerepolez/errores"
	f "rerepolez/funciones"
	"rerepolez/votos"
	TDACola "tdas/cola"
)

func main() {
	// pilaVotos := TDAPila.CrearPilaDinamica[votos.Voto]()
	colaVotantes := TDACola.CrearColaEnlazada[votos.Votante]()

	scanner := bufio.NewScanner(os.Stdin)
	var args = f.ObtenerParametrosEjecucion() // [lista_candidatos.csv, lista_padrones.txt]

	file_candidatos := f.AbrirArchivo(args[0])
	file_padrones := f.AbrirArchivo(args[1])

	votantes := f.ObtenerVotantes(file_padrones)
	partidos := f.ObtenerPartidos(file_candidatos)

	defer file_candidatos.Close()
	defer file_padrones.Close()

	for _, v := range votantes {
		fmt.Printf("%+v\n", v)
	}
	for _, p := range partidos {
		fmt.Printf("%+v\n", p)
	}

	for scanner.Scan() {
		texto := scanner.Text()
		palabras := strings.Split(texto, " ")
		cmd := palabras[0]
		params := palabras[1:]

		switch cmd {
		case "ingresar":
			if len(params) == 1 && params[0] != "" {
				numeroDni := params[0]
				dni := f.ObtenerDniEntero(numeroDni)
				if f.DniValido(dni) {
					votante := f.BuscarVotante(votantes, dni)
					if votante != nil {
						colaVotantes.Encolar(votante)
						f.MostrarSalida("OK")
					}
				}
			}

		case "votar":
			if len(params) == 2 && params[0] != "" && params[1] != "" {
				tipoVoto := params[0]
				// numeroLista := params[1]
				switch tipoVoto {
				case "Presidente":
				case "Gobernador":
				case "Intendente":

				}
			}

		case "deshacer":
			if len(params) == 0 {

			}

		case "fin-votar":
			if len(params) == 0 {
				if colaVotantes.EstaVacia() {

				}
			}

		}

	}
}
