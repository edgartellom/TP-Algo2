package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	TDACola "tdas/cola"
	"tp1/errores"
	. "tp1/votos"
)

func mostrarError(err string) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
}

func mostrarSalida(mensaje string) {
	fmt.Fprintf(os.Stdout, "%s\n", mensaje)
}

func abrirArchivo(archivo string) *os.File {
	file, err := os.Open(archivo)
	if err != nil {
		mostrarError(new(errores.ErrorLeerArchivo).Error())
		os.Exit(1)
	}
	return file
}

func convertirValor(valor int) TipoVoto {
	switch valor {
	case 0:
		return PRESIDENTE
	case 1:
		return GOBERNADOR
	case 2:
		return INTENDENTE
	}
	return LISTA_IMPUGNA
}

func buscarDni(archivo *os.File, dni string) bool {
	s := bufio.NewScanner(archivo)
	var found bool = false
	for s.Scan() {
		linea := s.Text()
		if strings.TrimSpace(linea) == dni {
			found = true
		}
	}
	return found
}

func obtenerPadrones(ruta string) []int {
	var padrones []int
	archivoPadrones := abrirArchivo(ruta)
	s := bufio.NewScanner(archivoPadrones)
	for s.Scan() {
		dni := s.Text()
		numeroDNI, _ := strconv.Atoi(dni) // Pensar: Si el archivo tiene un documento "1ac434"
		padrones = append(padrones, numeroDNI)
	}
	return padrones
}

func obtenerPartidos(ruta string) []Partido {
	var partidos []Partido
	archivoPartidos := abrirArchivo(ruta)
	s := bufio.NewScanner(archivoPartidos)
	for s.Scan() {
		lineaDePartido := s.Text() // split de las lineas, creo los partidos con las lineas, appendeo los partidos creados
		partidoEnFormaDeLista := strings.Split(lineaDePartido, ",")
		nombre, candidatos := partidoEnFormaDeLista[0], [3]string(partidoEnFormaDeLista[1:])
		partido := CrearPartido(nombre, candidatos)
		partidos = append(partidos, partido)
	}
	return partidos
}

func main() {
	TDACola.CrearColaEnlazada[int]()

	scanner := bufio.NewScanner(os.Stdin)
	var args = os.Args[1:]
	if len(args) < 2 {
		mostrarError(new(errores.ErrorParametros).Error())
		os.Exit(1)
	}

	ruta_partidos := args[0]
	ruta_padrones := args[1]

	lista_partidos := obtenerPartidos(ruta_partidos)
	lista_padrones := obtenerPadrones(ruta_padrones)

	/* ESTOS PRINT SON PARA QUE NO CHILLE EL VSCODE*/

	fmt.Println(lista_partidos) // ESTE SE MUESTRA RARO PORQUE SON STRUCTS
	fmt.Println(lista_padrones) // ESTE SE MUESTRA BIEN

	for scanner.Scan() {
		texto := scanner.Text()
		palabras := strings.Split(texto, " ")
		cmd := palabras[0]
		params := palabras[1:]

		switch cmd {
		case "ingresar":
			if len(params) < 1 || params[0] == "" {
				mostrarError(new(errores.ErrorParametros).Error())
			} else {
				dni := params[0]
				dniEntero, err := strconv.Atoi(dni)
				if err != nil || dniEntero <= 0 {
					mostrarError(new(errores.DNIError).Error())
					// } else {
					// 	if buscarDni(file_padrones, dni) {
					// 		mostrarSalida("OK")
					// 		votos.CrearVotante(dniEntero)
					// 	} else {
					// 		mostrarError(new(errores.DNIFueraPadron).Error())
					// 	}
				}
			}
			break
			/*
				CUANDO SE HAGA EL CASE VOTAR: EL TERCER PARAMETRO ES EL NUMERO DE PARTIDO
				ENTONCES USAS LA FUNCION "convertirValor(tercerParametro)"
			*/
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
