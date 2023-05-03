package main

import (
	"bufio"
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

	padrones := f.GuardarArchivoPadrones(file_padrones)
	candidatos := f.GuardarArchivoCandidatos(file_candidatos)

	defer file_candidatos.Close()
	defer file_padrones.Close()

	for _, lista_partido := range candidatos {
		partido := lista_partido[0]
		candidatos := [votos.CANT_VOTACION]string{lista_partido[1], lista_partido[2], lista_partido[3]}
		votos.CrearPartido(partido, candidatos)
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
					if f.ExisteDni(padrones, dni) {
						votante := votos.CrearVotante(dni)
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
