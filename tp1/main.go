package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	TDAVotos "tp1/diseno_alumnos/votos"
)

func mostrarError(mensaje string) {
	fmt.Fprintf(os.Stderr, "%s\n", mensaje)
	os.Exit(1)
}

func mostrarSalida(mensaje string) {
	fmt.Fprintf(os.Stdout, "%s\n", mensaje)
}

func abrirArchivo(archivo string) *os.File {
	file, err := os.Open(archivo)
	if err != nil {
		mostrarError("ERROR: Lectura de archivos")
	}
	// defer file.Close()
	return file
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var args = os.Args[1:]
	if len(args) < 2 {
		mostrarError("ERROR: Faltan parámetros")
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
		TDAVotos.CrearPartido(partido, candidatos)
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
				mostrarError("ERROR: Falta parámetro")
			}
			dni := params[0]
			dniEntero, err := strconv.Atoi(dni)
			if err != nil || dniEntero <= 0 {
				mostrarError("ERROR: DNI incorrecto")
			}
			s := bufio.NewScanner(file_padrones)
			for s.Scan() {
				linea := s.Text()
				if strings.Contains(linea, dni) {
					mostrarSalida("OK")
					TDAVotos.CrearVotante(dniEntero)
				} else {
					mostrarError("ERROR: DNI fuera del padrón")
				}

			}

		}
		defer file_candidatos.Close()
		defer file_padrones.Close()
	}
}
