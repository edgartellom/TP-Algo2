package exportaciones_de_datos

import (
	"fmt"
	"os"

	aerolineas "flycombi/sistema_aerolineas"
)

const (
	ORIGEN     = 0
	NUMERO_UNO = 1
)

const (
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

func EscribirArchivoKML(ruta string, ultimoCamino []aerolineas.Aeropuerto) {
	archivo, _ := os.Create(ruta)
	defer archivo.Close()

	archivo.WriteString(ENCABEZADO_KML)

	archivo.WriteString(DECLARACION_INICIO_KML)
	archivo.WriteString(INICIO_DOCUMENTO)
	archivo.WriteString(INICIO_NOMBRE + fmt.Sprintf(TITULO_KML+"%v"+SEPARADOR_ORIGEN_DESTINO+"%v", ultimoCamino[ORIGEN].Ciudad, ultimoCamino[len(ultimoCamino)-NUMERO_UNO].Ciudad) + CIERRE_NOMBRE)
	archivo.WriteString(INICIO_DESCRIPCION + DESCRIPCION_KML + CIERRE_DESCRIPCION)

	for _, aeropuerto := range ultimoCamino {
		archivo.WriteString(INICIO_PLACEMARK)
		archivo.WriteString(fmt.Sprintf(SANGRIA_DE_LINEA+INICIO_NOMBRE+"%v"+CIERRE_NOMBRE, aeropuerto.Codigo))
		archivo.WriteString(INICIO_PUNTO)
		archivo.WriteString(fmt.Sprintf(INICIO_COORDENADAS+"%v, %v"+CIERRE_COORDENADAS, aeropuerto.Longitud, aeropuerto.Latitud))
		archivo.WriteString(CIERRE_PUNTO)
		archivo.WriteString(CIERRE_PLACEMARK)
	}

	for i := NUMERO_UNO; i < len(ultimoCamino); i++ {
		archivo.WriteString(INICIO_PLACEMARK)
		archivo.WriteString(INICIO_LINEA)
		archivo.WriteString(fmt.Sprintf(INICIO_COORDENADAS+"%v, %v %v, %v"+CIERRE_COORDENADAS, ultimoCamino[i-NUMERO_UNO].Longitud, ultimoCamino[i-NUMERO_UNO].Latitud, ultimoCamino[i].Longitud, ultimoCamino[i].Latitud))
		archivo.WriteString(CIERRE_LINEA)
		archivo.WriteString(CIERRE_PLACEMARK)
	}

	archivo.WriteString(CIERRE_DOCUMENTO)
	archivo.WriteString(DECLARACION_CIERRE_KML)
}

func EscribirArchivoCSV(ruta string, vuelos []aerolineas.Vuelo) {
	archivo, _ := os.Create(ruta)
	defer archivo.Close()

	totalVuelos := len(vuelos)
	for i, vuelo := range vuelos {
		archivo.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v", vuelo.AeropuertoOrigen, vuelo.AeropuertoDestino, vuelo.Tiempo, vuelo.Precio, vuelo.Cant_vuelos))
		if i < totalVuelos-1 {
			archivo.WriteString(SALTO_DE_LINEA)
		}
	}
}
