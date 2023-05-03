package funciones

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	e "rerepolez/errores"
)

func ObtenerParametrosEjecucion() []string {
	var args = os.Args[1:]
	if len(args) < 2 {
		MostrarSalida(new(e.ErrorParametros).Error())
		os.Exit(1)
	}
	return args
}

// func MostrarError(err string) {
// 	fmt.Fprintf(os.Stderr, "%s\n", err)
// }

func MostrarSalida(salida string) {
	fmt.Fprintf(os.Stdout, "%s\n", salida)
}

func AbrirArchivo(archivo string) *os.File {
	file, err := os.Open(archivo)
	if err != nil {
		MostrarSalida(new(e.ErrorLeerArchivo).Error())
		os.Exit(1)
	}
	return file
}

func GuardarArchivoPadrones(archivo *os.File) []int {
	var padrones []int
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		padron := s.Text()
		dni := ObtenerDniEntero(padron)
		padrones = append(padrones, dni)
	}
	sort.Slice(padrones, func(i, j int) bool {
		return padrones[i] < padrones[j]
	})

	fmt.Println(padrones)
	return padrones
}

func GuardarArchivoCandidatos(archivo *os.File) [][]string {
	var lista_partidos [][]string
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		lista_partido := strings.Split(linea, ",")
		lista_partidos = append(lista_partidos, lista_partido)
	}
	fmt.Println(lista_partidos)
	return lista_partidos
}

func DniValido(dni int) bool {
	if dni != 0 {
		return true
	} else {
		MostrarSalida(new(e.DNIError).Error())
		return false
	}
}

func ExisteDni(padrones []int, dni int) bool {
	if len(padrones) == 0 {
		MostrarSalida(new(e.DNIFueraPadron).Error())
		return false
	}

	medio := len(padrones) / 2

	if padrones[medio] == dni {
		return true
	} else if padrones[medio] < dni {
		return ExisteDni(padrones[medio+1:], dni)
	} else {
		return ExisteDni(padrones[:medio], dni)
	}
}

func ObtenerDniEntero(dni string) int {
	var dniEntero int
	dniEntero, err := strconv.Atoi(dni)
	if err != nil || dniEntero <= 0 {
		dniEntero = 0
	}
	return dniEntero
}
