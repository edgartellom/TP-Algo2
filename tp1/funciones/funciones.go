package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	errores "rerepolez/errores"
	votos "rerepolez/votos"
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

func ObtenerVotantes(rutaPadrones string) []votos.Votante {
	padrones := ObtenerPadrones(rutaPadrones)
	listaDeVotantes := make([]votos.Votante, len(padrones))

	for i, dni := range padrones {
		listaDeVotantes[i] = votos.CrearVotante(dni)
	}
	return listaDeVotantes
}

func ObtenerPadrones(ruta string) []int {
	var padrones []int
	archivoPadrones := AbrirArchivo(ruta)
	defer archivoPadrones.Close()

	s := bufio.NewScanner(archivoPadrones)
	for s.Scan() {
		dni := s.Text()
		numeroDNI, _ := strconv.Atoi(dni)
		padrones = append(padrones, numeroDNI)
	}

	padrones = ordenarPadronesMergeSort(padrones)
	return padrones
}

func merge(izquierda, derecha []int) []int {
	i, j := 0, 0
	result := make([]int, 0)
	for i < len(izquierda) && j < len(derecha) {
		if izquierda[i] < derecha[j] {
			result = append(result, izquierda[i])
			i++
		} else {
			result = append(result, derecha[j])
			j++
		}
	}
	result = append(result, izquierda[i:]...)
	result = append(result, derecha[j:]...)
	return result
}

func ordenarPadronesMergeSort(padrones []int) []int {
	if len(padrones) < 2 {
		return padrones
	}
	medio := len(padrones) / 2
	izquierda := ordenarPadronesMergeSort(padrones[:medio])
	derecha := ordenarPadronesMergeSort(padrones[medio:])
	return merge(izquierda, derecha)
}

func documentoEnVotantes(dni int, votantes []votos.Votante) bool {
	if len(votantes) == 0 {
		return false
	}
	medio := len(votantes) / 2
	if dni == votantes[medio].LeerDNI() {
		return true
	}
	if dni < votantes[medio].LeerDNI() {
		return documentoEnVotantes(dni, votantes[:medio])
	}
	return documentoEnVotantes(dni, votantes[medio+1:])
}

func VerificarDNI(dni string, votantes []votos.Votante) (bool, string) {
	numeroDNI, err := strconv.Atoi(dni)
	if err != nil || numeroDNI <= 0 {
		errror := new(errores.DNIError)
		return false, errror.Error()
	} else if !documentoEnVotantes(numeroDNI, votantes) {
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
	if verificarVotante(votante, votantesQueVotaron) {
		errror := new(errores.ErrorVotanteFraudulento)
		errror.Dni = votante.LeerDNI()
		colaDeVotantes.Desencolar()
		return false, errror.Error()
	}
	return true, "OK"
}

func VerificarVoto(tipoDeVoto string, numeroDeLista string, colaDeVotantes TDACola.Cola[votos.Votante], votantesQueVotaron []int, listaDePartidos []votos.Partido) (bool, string) {
	validez, salida := VerificarColaYVotante(colaDeVotantes, votantesQueVotaron)
	if !validez {
		return validez, salida
	}
	if !tipoValido(tipoDeVoto) {
		errror := new(errores.ErrorTipoVoto)
		return false, errror.Error()
	}
	if !numeroDeListaValido(numeroDeLista, len(listaDePartidos)-1) {
		errror := new(errores.ErrorAlternativaInvalida)
		return false, errror.Error()
	}
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

func ImprimirTipoCompleto(tipo votos.TipoVoto, partido []votos.Partido) {
	switch tipo {
	case 0:
		fmt.Println("Presidente:")
		for i := 0; i < len(partido); i++ {
			fmt.Println(partido[i].ObtenerResultado(tipo))
		}
		fmt.Println()
	case 1:
		fmt.Println("Gobernador:")
		for i := 0; i < len(partido); i++ {
			fmt.Println(partido[i].ObtenerResultado(tipo))
		}
		fmt.Println()
	case 2:
		fmt.Println("Intendente:")
		for i := 0; i < len(partido); i++ {
			fmt.Println(partido[i].ObtenerResultado(tipo))
		}
		fmt.Println()
	}
}

func ImprimirImpugnadosSegunCantidad(cantidad int) {
	if cantidad == 1 {
		fmt.Printf("Votos Impugnados: %d voto\n", cantidad)
	} else {
		fmt.Printf("Votos Impugnados: %d votos\n", cantidad)
	}
}
