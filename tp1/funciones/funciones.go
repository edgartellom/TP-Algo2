package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"rerepolez/errores"
	"rerepolez/votos"
	TDACola "tdas/cola"
)

const (
	SEPARADOR_ARCHIVO = ","
	SINGULAR          = "voto"
	PLURAL            = "votos"

	SALIDA_EXITOSA = "OK"

	PRESIDENTE = "Presidente"
	GOBERNADOR = "Gobernador"
	INTENDENTE = "Intendente"
)

/* ---------------------------- FUNCIONES DE GENERALES ---------------------------- */

func MostrarError(err error) {
	fmt.Fprintln(os.Stdout, err)
	os.Exit(1)
}

func MostrarSalida(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		MostrarError(new(errores.ErrorLeerArchivo))
	}
	return archivo
}

func CrearFilaDeVotacion() TDACola.Cola[votos.Votante] {
	return TDACola.CrearColaEnlazada[votos.Votante]()
}

func SepararEntrada(entrada string, separador string) []string {
	return strings.Split(entrada, separador)
}

/* ---------------------------- FUNCIONES DE EXTRACCION DE INFORMACION ---------------------------- */

func ObtenerPartidos(ruta string) []votos.Partido {
	archivoDePartidos := abrirArchivo(ruta)
	defer archivoDePartidos.Close()

	var partidos []votos.Partido
	partidoEnBlanco := votos.CrearVotosEnBlanco()
	partidos = append(partidos, partidoEnBlanco)

	scanner := bufio.NewScanner(archivoDePartidos)
	for scanner.Scan() {
		lineaDePartido := scanner.Text()
		partidoEnFormaDeLista := strings.Split(lineaDePartido, SEPARADOR_ARCHIVO)
		nombre := partidoEnFormaDeLista[0]
		candidatos := [votos.CANT_VOTACION]string{partidoEnFormaDeLista[1], partidoEnFormaDeLista[2], partidoEnFormaDeLista[3]}
		partidoNuevo := votos.CrearPartido(nombre, candidatos)
		partidos = append(partidos, partidoNuevo)
	}
	return partidos
}

func ObtenerVotantes(rutaPadrones string) []votos.Votante {
	padrones := obtenerPadrones(rutaPadrones)
	listaDeVotantes := make([]votos.Votante, len(padrones))

	for i, dni := range padrones {
		listaDeVotantes[i] = votos.CrearVotante(dni)
	}
	return listaDeVotantes
}

func obtenerPadrones(ruta string) []int {
	archivoDePadrones := abrirArchivo(ruta)
	defer archivoDePadrones.Close()

	var padrones []int

	scanner := bufio.NewScanner(archivoDePadrones)
	for scanner.Scan() {
		dni := scanner.Text()
		numeroDNI, _ := strconv.Atoi(dni)
		padrones = append(padrones, numeroDNI)
	}

	padrones = ordenarPadronesMergeSort(padrones)
	return padrones
}

func merge(izquierda, derecha []int) []int {
	i, j := 0, 0
	resultante := make([]int, 0)
	for i < len(izquierda) && j < len(derecha) {
		if izquierda[i] < derecha[j] {
			resultante = append(resultante, izquierda[i])
			i++
		} else {
			resultante = append(resultante, derecha[j])
			j++
		}
	}
	resultante = append(resultante, izquierda[i:]...)
	resultante = append(resultante, derecha[j:]...)
	return resultante
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

/* ---------------------------- FUNCIONES DE VERIFICACION ---------------------------- */

func VerificarDNI(dni string, votantes []votos.Votante) (bool, string) {
	numeroDNI, err := strconv.Atoi(dni)
	if err != nil || numeroDNI <= 0 {
		err := new(errores.DNIError)
		return false, err.Error()
	} else if !documentoEnVotantes(numeroDNI, votantes) {
		err := new(errores.DNIFueraPadron)
		return false, err.Error()
	}
	return true, SALIDA_EXITOSA
}

func tipoValido(tipoIngresado string) bool {
	return tipoIngresado == PRESIDENTE || tipoIngresado == GOBERNADOR || tipoIngresado == INTENDENTE
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

func VerificarColaYVotante(filaDeVotantes TDACola.Cola[votos.Votante], votantesQueVotaron []int) (bool, string) {
	if filaDeVotantes.EstaVacia() {
		err := new(errores.FilaVacia)
		return false, err.Error()
	}
	votante := filaDeVotantes.VerPrimero()
	if verificarVotante(votante, votantesQueVotaron) {
		err := new(errores.ErrorVotanteFraudulento)
		err.Dni = votante.LeerDNI()
		filaDeVotantes.Desencolar()
		return false, err.Error()
	}
	return true, SALIDA_EXITOSA
}

func VerificarVoto(tipoDeVoto string, numeroDeLista string, filaDeVotantes TDACola.Cola[votos.Votante], votantesQueVotaron []int, listaDePartidos []votos.Partido) (bool, string) {
	validez, salida := VerificarColaYVotante(filaDeVotantes, votantesQueVotaron)
	if !validez {
		return validez, salida
	}
	if !tipoValido(tipoDeVoto) {
		err := new(errores.ErrorTipoVoto)
		return false, err.Error()
	}
	if !numeroDeListaValido(numeroDeLista, len(listaDePartidos)-1) {
		err := new(errores.ErrorAlternativaInvalida)
		return false, err.Error()
	}
	return true, SALIDA_EXITOSA
}

/* ---------------------------- FUNCION PARA CONVERTIR ENTRADA ---------------------------- */

func ConvertirEntradaATipoVoto(tipoVotoIngresado string) votos.TipoVoto {
	if tipoVotoIngresado == PRESIDENTE {
		return votos.PRESIDENTE
	}
	if tipoVotoIngresado == GOBERNADOR {
		return votos.GOBERNADOR
	}
	return votos.INTENDENTE
}

/* ---------------------------- FUNCIONES DE IMPRESION ---------------------------- */

func imprimirNombreDelTipo(tipo votos.TipoVoto) {
	switch tipo {
	case votos.PRESIDENTE:
		fmt.Printf("%s:\n", PRESIDENTE)
	case votos.GOBERNADOR:
		fmt.Printf("%s:\n", GOBERNADOR)
	case votos.INTENDENTE:
		fmt.Printf("%s:\n", INTENDENTE)
	}
}

func ImprimirTipoCompleto(tipo votos.TipoVoto, partidos []votos.Partido) {
	imprimirNombreDelTipo(tipo)
	for i := 0; i < len(partidos); i++ {
		fmt.Println(partidos[i].ObtenerResultado(tipo))
	}
	fmt.Println()
}

func PalabraSegunCantidad(cantidad int) string {
	if cantidad == 1 {
		return SINGULAR
	}
	return PLURAL
}
