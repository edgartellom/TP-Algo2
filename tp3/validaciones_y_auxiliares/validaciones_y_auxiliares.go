package validaciones_y_auxiliares

import (
	"bufio"
	aerolineas "flycombi/sistema_aerolineas"
	"fmt"
	"os"
	"strings"
)

type indice int

const (
	ORIGEN indice = iota
	DESTINO
	TIEMPO_PROMEDIO
	PRECIO
	ESCALAS_ENTRE_AMBOS

	SEPARADOR_1 string = ","
	SEPARADOR_2 string = " "
)

const (
	LONGITUD_ENTRADA_COMPLETA = 4
)

func MostrarError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

func MostrarSalida(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Print("asd")
	}
	return archivo
}

func SepararEntrada(entrada string, separador string) []string {
	return strings.Split(entrada, separador)
}

func ObtenerAeropuertos(ruta string) []aerolineas.Aeropuerto {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()

	var aeropuertos []aerolineas.Aeropuerto

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		aeropuerto := aerolineas.CrearAeropuerto(linea)
		aeropuertos = append(aeropuertos, aeropuerto)
	}
	return aeropuertos
}

func ObtenerVuelos(ruta string) []aerolineas.Vuelo {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()

	var vuelos []aerolineas.Vuelo

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		vuelo := aerolineas.CrearVuelo(linea)
		vuelos = append(vuelos, vuelo)
	}
	return vuelos
}

func armarEntrada(entradaCompleta string) []string {
	var entrada []string
	entradaSeparada1 := strings.Split(entradaCompleta, SEPARADOR_1)
	entradaSeparada2 := strings.SplitN(entradaSeparada1[0], SEPARADOR_2, 2)
	entrada = append(entrada, entradaSeparada2...)
	entrada = append(entrada, entradaSeparada1[1:]...)
	return entrada
}

/* func comprobarEntrada(entrada []string) error {
	comando, parametros := entrada[0], entrada[1:]
	var err error
	if comando == "camino_mas" && len(parametros) != 3 {
		err.Error() // error en comando "camino_mas"
	} else if comando == "camino_escalas" && len(parametros) != 2 {
		err.Error() // error en comando "camino_escalas"
	} else if (comando == "centralidad" || comando == "nueva_aerolinea" || comando == "itinerario" || comando == "exportar_kml") && len(parametros) != 1 {
		err.Error() // error en comando "centralidad"/"nueva_aerolinea"/"itinerario"/"exportar_kml"
	}
	return err
} */

func CompletarEntrada(entrada string) ([]string, error) {
	entradaCompleta := make([]string, LONGITUD_ENTRADA_COMPLETA)
	entradaSeparada := armarEntrada(entrada)

	// err := comprobarEntrada(entradaSeparada)

	copy(entradaCompleta, entradaSeparada)
	return entradaCompleta, nil
}

func ImprimirCamino(aeropuertos []aerolineas.Aeropuerto) {
	result := []string{}
	for _, aeropuerto := range aeropuertos {
		result = append(result, string(aeropuerto.Codigo))
	}

	mensaje := strings.Join(result, " -> ")
	MostrarSalida(mensaje)
}
