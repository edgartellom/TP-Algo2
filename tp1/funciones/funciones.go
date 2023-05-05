package funciones

import (
	"bufio"
	"fmt"
	"os"
	errores "rerepolez/errores"
	votos "rerepolez/votos"
	"strconv"
	"strings"
	TDACola "tdas/cola"
)

func MostrarError(err error) {
	fmt.Fprintln(os.Stdout, err)
	os.Exit(1)
}

func MostrarSalida(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func AbrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		MostrarError(new(errores.ErrorLeerArchivo))
	}
	return archivo
}

func ObtenerPartidos(ruta string) []votos.Partido {
	var partidos []votos.Partido
	archivoPartidos := AbrirArchivo(ruta)
	defer archivoPartidos.Close()

	partidoEnBlanco := votos.CrearVotosEnBlanco()
	partidos = append(partidos, partidoEnBlanco)

	s := bufio.NewScanner(archivoPartidos)
	for s.Scan() {
		lineaDePartido := s.Text()
		partidoEnFormaDeLista := strings.Split(lineaDePartido, ",")
		nombre, candidatos := partidoEnFormaDeLista[0], [3]string(partidoEnFormaDeLista[1:])
		partido := votos.CrearPartido(nombre, candidatos)
		partidos = append(partidos, partido)
	}
	return partidos
}

func ObtenerPadrones(ruta string) []int {
	var padrones []int
	archivoPadrones := AbrirArchivo(ruta)
	defer archivoPadrones.Close()

	s := bufio.NewScanner(archivoPadrones)
	for s.Scan() {
		dni := s.Text()
		numeroDNI, _ := strconv.Atoi(dni) // Pensar: Si el archivo tiene un DNI "1ac434"
		padrones = append(padrones, numeroDNI)
	}
	return padrones
}

func documentoEnPadron(dni int, padrones []int) bool {
	if len(padrones) == 0 {
		return false
	}
	medio := len(padrones) / 2
	if dni == padrones[medio] {
		return true
	}
	if dni < padrones[medio] {
		return documentoEnPadron(dni, padrones[:medio])
	}
	return documentoEnPadron(dni, padrones[medio:])
}

func VerificarDNI(dni string, padrones []int) (bool, string) {
	numeroDNI, err := strconv.Atoi(dni)
	if err != nil || numeroDNI <= 0 {
		errror := new(errores.DNIError)
		return false, errror.Error()
	} else if !documentoEnPadron(numeroDNI, padrones) {
		errror := new(errores.DNIFueraPadron)
		return false, errror.Error()
	}
	return true, "OK"
}

func tipoValido(tipoIngresado string) bool {
	if tipoIngresado == "Presidente" || tipoIngresado == "Gobernador" || tipoIngresado == "Intendente" {
		return true
	}
	return false
}

func numeroDeListaValido(numeroDeLista string, cantidadDePartidos int) bool {
	listaNumero, err := strconv.Atoi(numeroDeLista)
	if err != nil || listaNumero > cantidadDePartidos {
		return false
	}
	return true
}

func verificarVotante(votante votos.Votante, votantesQueVotaron []int) bool {
	documentoDelVotante := votante.LeerDNI()
	for _, documento := range votantesQueVotaron {
		if documentoDelVotante == documento {
			return true
		}
	}
	return false
}

func VerificarColaYVotante(colaDeVotantes TDACola.Cola[votos.Votante], votantesQueVotaron []int) (bool, string) {
	if colaDeVotantes.EstaVacia() {
		errror := new(errores.FilaVacia)
		return false, errror.Error()
	}
	votante := colaDeVotantes.VerPrimero()
	if !verificarVotante(votante, votantesQueVotaron) {
		errror := new(errores.ErrorVotanteFraudulento)
		errror.Dni = votante.LeerDNI()
		colaDeVotantes.Desencolar()
		return false, fmt.Sprintf("%s", errror)
	}
	return true, "OK"
}

func VerificarVoto(tipoDeVoto string, numeroDeLista string, colaDeVotantes TDACola.Cola[votos.Votante], votantesQueVotaron []int, listaDePartidos []votos.Partido) (bool, string) {
	// if colaDeVotantes.EstaVacia() {
	// 	errror := new(errores.FilaVacia)
	// 	return false, errror.Error()
	// }
	validez, salida := VerificarColaYVotante(colaDeVotantes, votantesQueVotaron)
	if !validez {
		return validez, salida
	}
	if !tipoValido(tipoDeVoto) {
		errror := new(errores.ErrorTipoVoto)
		return false, errror.Error()
	}
	if !numeroDeListaValido(numeroDeLista, len(listaDePartidos)) {
		errror := new(errores.ErrorAlternativaInvalida)
		return false, errror.Error()
	}
	// votante := colaDeVotantes.VerPrimero()
	// if !verificarVotante(votante, votantesQueVotaron) {
	// 	errror := new(errores.ErrorVotanteFraudulento)
	// 	errror.Dni = votante.LeerDNI()
	// 	colaDeVotantes.Desencolar()
	// 	return false, fmt.Sprintf("%s", errror)
	// }
	return true, "OK"
}

func ConvertirEntradaATipoVoto(tipoVotoIngresado string) votos.TipoVoto {
	if tipoVotoIngresado == "Presidente" {
		return votos.PRESIDENTE
	}
	if tipoVotoIngresado == "Gobernador" {
		return votos.GOBERNADOR
	}
	return votos.INTENDENTE
}
