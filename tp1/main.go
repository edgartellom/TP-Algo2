package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	TDACola "tdas/cola"
	"tp1/errores"
	"tp1/votos"
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

func main() {
	// pilaVotos := TDAPila.CrearPilaDinamica[votos.Voto]()
	TDACola.CrearColaEnlazada[votos.Votante]()

	scanner := bufio.NewScanner(os.Stdin)
	var args = os.Args[1:]
	if len(args) < 2 {
		mostrarError(new(errores.ErrorParametros).Error())
		os.Exit(1)
	}

	lista_candidatos := args[0]
	lista_padrones := args[1]

	file_candidatos := abrirArchivo(lista_candidatos)
	s := bufio.NewScanner(file_candidatos)
	for s.Scan() {
		linea := s.Text()
		lista_partido := strings.Split(linea, ",")
		partido := lista_partido[0]
		candidatos := [3]string{lista_partido[1], lista_partido[2], lista_partido[3]}
		votos.CrearPartido(partido, candidatos)
	}
	file_padrones := abrirArchivo(lista_padrones)

	for scanner.Scan() {
		texto := scanner.Text()
		palabras := strings.Split(texto, " ")
		cmd := palabras[0]
		params := palabras[1:]

		switch cmd {
		case "ingresar":
			if len(params) < 1 {
				mostrarError(new(errores.ErrorParametros).Error())
			} else {
				dni := params[0]
				dniEntero, err := strconv.Atoi(dni)
				if err != nil || dniEntero <= 0 {
					mostrarError(new(errores.DNIError).Error())
				}
				if buscarDni(file_padrones, dni) {
					mostrarSalida("OK")
					votos.CrearVotante(dniEntero)
					file_padrones.Seek(0, 0)
				} else {
					mostrarError(new(errores.DNIFueraPadron).Error())
				}

			}
		}
		defer file_candidatos.Close()
		defer file_padrones.Close()
	}
}
