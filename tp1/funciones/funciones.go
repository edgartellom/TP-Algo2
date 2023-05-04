package funciones

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	e "rerepolez/errores"
	"rerepolez/votos"
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

func ObtenerVotantes(archivo *os.File) []votos.Votante {
	var votantes []votos.Votante
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		padron := s.Text()
		dni := ObtenerDniEntero(padron)
		votante := votos.CrearVotante(dni)
		votantes = append(votantes, votante)
	}

	sort.Slice(votantes, func(i, j int) bool {
		return votantes[i].LeerDNI() < votantes[j].LeerDNI()
	})

	return votantes
}

func ObtenerPartidos(archivo *os.File) []votos.Partido {
	var listaPartidos []votos.Partido
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		lineaPartido := s.Text()
		listaPartido := strings.Split(lineaPartido, ",")
		nombre, candidatos := listaPartido[0], [votos.CANT_VOTACION]string(listaPartido[1:])
		partido := votos.CrearPartido(nombre, candidatos)
		listaPartidos = append(listaPartidos, partido)
	}

	return listaPartidos
}

func DniValido(dni int) bool {
	if dni != 0 {
		return true
	} else {
		MostrarSalida(new(e.DNIError).Error())
		return false
	}
}

func BuscarVotante(votantes []votos.Votante, dni int) votos.Votante {
	if len(votantes) == 0 {
		MostrarSalida(new(e.DNIFueraPadron).Error())
		return nil
	}

	medio := len(votantes) / 2

	if votantes[medio].LeerDNI() == dni {
		return votantes[medio]
	} else if votantes[medio].LeerDNI() < dni {
		return BuscarVotante(votantes[medio+1:], dni)
	} else {
		return BuscarVotante(votantes[:medio], dni)
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
