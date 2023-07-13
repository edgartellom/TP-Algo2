package validaciones_y_auxiliares

import (
	"bufio"
	"flycombi/errores"
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
)

const (
	COMANDO = iota
	UN_PARAMETRO
	DOS_PARAMETROS
	TRES_PARAMETROS
	CUATRO_PARAMETROS
)

const (
	CAMINO_MAS = iota
	CAMINO_ESCALAS
	CENTRALIDAD
	NUEVA_AEROLINEA
	ITINERARIO
	EXPORTAR_KML
)

const (
	LONGITUD_ENTRADA_COMPLETA = 4
	CANT_COMANDOS             = 6

	SEPARADOR_COMA    = ","
	SEPARADOR_ESPACIO = " "
	SEPARADOR_FLECHA  = " -> "

	SALTO_DE_LINEA           = "\n"
	SANGRIA_DE_LINEA         = "	"
	TITULO_KML               = "Camino desde "
	SEPARADOR_ORIGEN_DESTINO = " hasta "
	DESCRIPCION_KML          = "Exporta a un archivo kml el ultimo camino que fue solicitado"

	ENCABEZADO_KML         = `<?xml version="1.0" encoding="UTF-8"?>` + SALTO_DE_LINEA
	DECLARACION_INICIO_KML = `<kml xmlns="http://earth.google.com/kml/2.1">` + SALTO_DE_LINEA
	DECLARACION_CIERRE_KML = `</kml>`
	INICIO_DOCUMENTO       = SANGRIA_DE_LINEA + `<Document>` + SALTO_DE_LINEA
	CIERRE_DOCUMENTO       = SALTO_DE_LINEA + SANGRIA_DE_LINEA + `</Document>` + SALTO_DE_LINEA
	INICIO_PLACEMARK       = SALTO_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<Placemark>` + SALTO_DE_LINEA
	CIERRE_PLACEMARK       = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `</Placemark>` + SALTO_DE_LINEA
	INICIO_PUNTO           = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<Point>` + SALTO_DE_LINEA
	CIERRE_PUNTO           = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `</Point>` + SALTO_DE_LINEA
	INICIO_LINEA           = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<LineString>` + SALTO_DE_LINEA
	CIERRE_LINEA           = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `</LineString>` + SALTO_DE_LINEA
	INICIO_NOMBRE          = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<name>`
	CIERRE_NOMBRE          = `</name>` + SALTO_DE_LINEA
	INICIO_DESCRIPCION     = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<description>`
	CIERRE_DESCRIPCION     = `</description>` + SALTO_DE_LINEA
	INICIO_COORDENADAS     = SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + SANGRIA_DE_LINEA + `<coordinates>`
	CIERRE_COORDENADAS     = `</coordinates>` + SALTO_DE_LINEA
)

var COMANDOS = [CANT_COMANDOS]string{"camino_mas", "camino_escalas", "centralidad", "nueva_aerolinea", "itinerario", "exportar_kml"}

/* ---------------------------------------------------------- FUNCIONES DE SALIDA ---------------------------------------------------------- */

func MostrarMensaje(mensaje string) {
	fmt.Fprintln(os.Stdout, mensaje)
}

func MostrarError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

/* --------------------------------------------------------- FUNCIONES AUXILIARES ---------------------------------------------------------- */

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		MostrarError(new(errores.ErrorLeerArchivo))
	}
	return archivo
}

func MostrarCamino(camino []aerolineas.Aeropuerto) {
	salida := make([]string, len(camino))
	for i, aeropuerto := range camino {
		salida[i] = string(aeropuerto.Codigo)
	}
	MostrarMensaje(strings.Join(salida, SEPARADOR_FLECHA))
}

func crearEntrada(entradaCompleta string) []string {
	var entrada []string
	entradaSeparada1 := strings.Split(entradaCompleta, SEPARADOR_COMA)
	entradaSeparada2 := strings.SplitN(entradaSeparada1[0], SEPARADOR_ESPACIO, 2)
	entrada = append(entrada, entradaSeparada2...)
	entrada = append(entrada, entradaSeparada1[1:]...)
	return entrada
}

func CompletarYValidarEntrada(entradaReal string) ([]string, error) {
	entradaCompleta := make([]string, LONGITUD_ENTRADA_COMPLETA)
	entrada := crearEntrada(entradaReal)

	err := comprobarParametrosDeComando(entrada[COMANDO], entrada[1:])

	copy(entradaCompleta, entrada)
	return entradaCompleta, err
}

/* ------------------------------------------------- FUNCIONES DE EXTRACCION DE INFORMACION ------------------------------------------------ */

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

func ObtenerCiudadesYRutas(ruta string) ([]aerolineas.Ciudad, []aerolineas.Ruta) {
	archivo := abrirArchivo(ruta)
	defer archivo.Close()
	var ciudadesStr []string
	var ciudades []aerolineas.Ciudad
	var rutas []aerolineas.Ruta
	scanner := bufio.NewScanner(archivo)
	primeraLinea := true
	for scanner.Scan() {
		linea := scanner.Text()
		if primeraLinea {
			ciudadesStr = strings.Split(linea, SEPARADOR_COMA)
			for _, ciudadStr := range ciudadesStr {
				ciudades = append(ciudades, aerolineas.Ciudad(ciudadStr))
			}
			primeraLinea = false
		}
		rutaStr := strings.Split(linea, SEPARADOR_COMA)
		ruta := aerolineas.Ruta{
			CiudadOrigen:  aerolineas.Ciudad(rutaStr[0]),
			CiudadDestino: aerolineas.Ciudad(rutaStr[1]),
		}
		rutas = append(rutas, ruta)
	}
	return ciudades, rutas
}

/* -------------------------------------------------------- FUNCIONES DE VALIDACION -------------------------------------------------------- */

func comprobarParametrosDeComando(comando string, parametros []string) error {
	var err error
	if (comando == COMANDOS[CAMINO_MAS] && len(parametros) != TRES_PARAMETROS) || (comando == COMANDOS[CAMINO_ESCALAS] && len(parametros) != DOS_PARAMETROS) ||
		(comando == COMANDOS[CENTRALIDAD] || comando == COMANDOS[NUEVA_AEROLINEA] || comando == COMANDOS[ITINERARIO] || comando == COMANDOS[EXPORTAR_KML]) && len(parametros) != UN_PARAMETRO {
		err = errores.ErrorComando{Comando: comando}
	}
	return err
}

func comprobarPertenencia(sistema aerolineas.SistemaDeAerolineas, entradaOrigen, entradaDestino string) bool {
	return sistema.Pertenece(aerolineas.Ciudad(entradaOrigen)) && sistema.Pertenece(aerolineas.Ciudad(entradaDestino))
}

func ComprobarEntradaCaminoEscalas(sistema aerolineas.SistemaDeAerolineas, entradaOrigen, entradaDestino string) error {
	var err error
	if !comprobarPertenencia(sistema, entradaOrigen, entradaDestino) {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_ESCALAS]}
	}
	return err
}

func ComprobarEntradaCaminoMas(sistema aerolineas.SistemaDeAerolineas, tipo, entradaOrigen, entradaDestino string) error {
	var err error
	if tipo != "barato" && tipo != "rapido" {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	} else if !comprobarPertenencia(sistema, entradaOrigen, entradaDestino) {
		err = errores.ErrorComando{Comando: COMANDOS[CAMINO_MAS]}
	}
	return err
}

func ExportarVuelos(vuelos []aerolineas.Vuelo, ruta string) {
	archivo := crearArchivo(ruta)
	totalVuelos := len(vuelos)
	for i, vuelo := range vuelos {
		archivo.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v", vuelo.AeropuertoOrigen, vuelo.AeropuertoDestino, vuelo.Tiempo, vuelo.Precio, vuelo.Cant_vuelos))
		if i < totalVuelos-1 {
			archivo.WriteString(SALTO_DE_LINEA)
		}
	}

}

func ExportarUltimoCamino(ultimoCamino []aerolineas.Aeropuerto, ruta string) {
	archivo := crearArchivo(ruta)
	archivo.WriteString(ENCABEZADO_KML)

	archivo.WriteString(DECLARACION_INICIO_KML)
	archivo.WriteString(INICIO_DOCUMENTO)
	archivo.WriteString(INICIO_NOMBRE + fmt.Sprintf(TITULO_KML+"%v"+SEPARADOR_ORIGEN_DESTINO+"%v", ultimoCamino[0].Ciudad, ultimoCamino[len(ultimoCamino)-1].Ciudad) + CIERRE_NOMBRE)
	archivo.WriteString(INICIO_DESCRIPCION + DESCRIPCION_KML + CIERRE_DESCRIPCION)

	for _, aeropuerto := range ultimoCamino {
		archivo.WriteString(INICIO_PLACEMARK)
		archivo.WriteString(fmt.Sprintf(SANGRIA_DE_LINEA+INICIO_NOMBRE+"%v"+CIERRE_NOMBRE, aeropuerto.Ciudad))
		archivo.WriteString(INICIO_PUNTO)
		archivo.WriteString(fmt.Sprintf(INICIO_COORDENADAS+"%v, %v"+CIERRE_COORDENADAS, aeropuerto.Longitud, aeropuerto.Latitud))
		archivo.WriteString(CIERRE_PUNTO)
		archivo.WriteString(CIERRE_PLACEMARK)
	}

	for i := 1; i < len(ultimoCamino); i++ {
		archivo.WriteString(INICIO_PLACEMARK)
		archivo.WriteString(INICIO_LINEA)
		archivo.WriteString(fmt.Sprintf(INICIO_COORDENADAS+"%v, %v %v, %v"+CIERRE_COORDENADAS, ultimoCamino[i-1].Longitud, ultimoCamino[i-1].Latitud, ultimoCamino[i].Longitud, ultimoCamino[i].Latitud))
		archivo.WriteString(CIERRE_LINEA)
		archivo.WriteString(CIERRE_PLACEMARK)
	}

	archivo.WriteString(CIERRE_DOCUMENTO)
	archivo.WriteString(DECLARACION_CIERRE_KML)
}

func crearArchivo(ruta string) *os.File {
	archivo, _ := os.Create(ruta) // Será necesario manejar el error que devuelve esta funcion?
	return archivo
}
