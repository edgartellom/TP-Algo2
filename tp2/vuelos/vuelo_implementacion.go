package vuelos

import (
	"strconv"
	"strings"
)

type indice int

const (
	CODIGO indice = iota
	AEROLINEA
	ORIGEN
	DESTINO
	NUM_COLA
	PRIORIDAD
	FECHA
	DEMORA
	TIEMPO
	CANCELADO
)

type vuelo struct {
	info Claves
	// numeroDeVuelo       Clave
	// origen              Clave
	// destino             Clave
	prioridad int
	// fecha               Clave
	informacionCompleta string
}

func CrearInformacionPrincipal(codigo, fecha, origen, destino string) Claves {
	return Claves{codigo, fecha, origen, destino}
}

func CrearVuelo(informacionDeVuelo string) Vuelo {
	camposDeVuelo := strings.Split(informacionDeVuelo, ",")
	nPrioridad, _ := strconv.Atoi(camposDeVuelo[PRIORIDAD])

	vuelo := new(vuelo)
	vuelo.info = CrearInformacionPrincipal(camposDeVuelo[CODIGO], camposDeVuelo[FECHA], camposDeVuelo[ORIGEN], camposDeVuelo[DESTINO])
	// vuelo.origen, vuelo.destino = camposDeVuelo[ORIGEN], camposDeVuelo[DESTINO]
	vuelo.informacionCompleta = strings.Join(camposDeVuelo, " ")
	// vuelo.numeroDeVuelo = camposDeVuelo[CODIGO]
	// vuelo.fecha = camposDeVuelo[FECHA]
	vuelo.prioridad = nPrioridad
	return vuelo
}

func (vuelo vuelo) VerInformacionPrincipal() Claves {
	return vuelo.info
}

func (vuelo vuelo) VerPrioridad() int {
	return vuelo.prioridad
}

func (vuelo vuelo) VerCodigo() string {
	return vuelo.info.codigo
}

func (vuelo vuelo) VerFecha() string {
	return vuelo.info.fecha
}

func (vuelo vuelo) ObtenerInformacionDeVuelo() string {
	return vuelo.informacionCompleta
}
