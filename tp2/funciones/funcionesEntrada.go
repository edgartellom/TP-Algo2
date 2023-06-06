package funciones

import (
	"strings"
)

func SepararEntrada(entrada string, separador string) []string {
	return strings.Split(entrada, separador)
}


