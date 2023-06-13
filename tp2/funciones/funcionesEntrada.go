package funciones

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const formatoFecha = "2006-01-02T15:04:05"

func SepararEntrada(entrada string, separador string) []string {
	return strings.Split(entrada, separador)
}

func ConvertirCadenaAFechaEnteros(cadena string) (int, int, int, int, int, int) {
	// Separar los componentes de fecha y hora
	parts := strings.Split(cadena, "T")
	fechaPart := parts[0]
	horaPart := parts[1]

	// Extraer los componentes de fecha
	fechaComponents := strings.Split(fechaPart, "-")
	year, _ := strconv.Atoi(fechaComponents[0])
	month, _ := strconv.Atoi(fechaComponents[1])
	day, _ := strconv.Atoi(fechaComponents[2])

	// Extraer los componentes de hora
	horaComponents := strings.Split(horaPart, ":")
	hour, _ := strconv.Atoi(horaComponents[0])
	minute, _ := strconv.Atoi(horaComponents[1])
	second, _ := strconv.Atoi(horaComponents[2])

	// Devolver los valores en enteros

	return year, month, day, hour, minute, second
}

func ConvertirCadenaAFecha(cadena string) time.Time {
	year, month, day, hour, minute, second := ConvertirCadenaAFechaEnteros(cadena)
	fecha := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	return fecha
}

func ConvertirFechaACadena(fecha time.Time) string {
	cadenaFecha := fecha.Format(formatoFecha)
	return cadenaFecha
}

func ConvertirFechaACadenaEnteros(fecha time.Time) string {
	cadenaFecha := ConvertirFechaACadena(fecha)
	year, month, day, hour, minute, second := ConvertirCadenaAFechaEnteros(cadenaFecha)
	cadena := fmt.Sprintf("%d-%d-%dT%d:%d:%d", year, month, day, hour, minute, second)
	return cadena
}
