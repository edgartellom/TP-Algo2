package votos

import "fmt"

const NOMBRE_PARTIDO_EN_BLANCO = "Votos en Blanco"

type partidoImplementacion struct {
	nombreDelPartido     string
	votosActuales        [CANT_VOTACION]int
	candidatosDelPartido [CANT_VOTACION]string
}

type partidoEnBlanco struct {
	nombreDelPartido string
	votosActuales    [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)

	partido.nombreDelPartido = nombre
	partido.candidatosDelPartido = candidatos
	partido.votosActuales = [CANT_VOTACION]int{}
	return partido
}

func CrearVotosEnBlanco() Partido {
	partido := new(partidoEnBlanco)

	partido.nombreDelPartido = NOMBRE_PARTIDO_EN_BLANCO
	partido.votosActuales = [CANT_VOTACION]int{}
	return partido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	(*partido).votosActuales[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.votosActuales[tipo] == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido.nombreDelPartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo])
	}
	return fmt.Sprintf("%s - %s: %d votos", partido.nombreDelPartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	(*blanco).votosActuales[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votosActuales[tipo] == 1 {
		return fmt.Sprintf("%s: %d voto", blanco.nombreDelPartido, blanco.votosActuales[tipo])
	}
	return fmt.Sprintf("%s: %d votos", blanco.nombreDelPartido, blanco.votosActuales[tipo])
}
