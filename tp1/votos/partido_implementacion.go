package votos

import "fmt"

type partidoImplementacion struct {
	nombrePartido        string
	candidatosDelPartido [3]string
	votosActuales        [3]int
}

type partidoEnBlanco struct {
	nombrePartido string
	votosActuales [3]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)

	partido.nombrePartido = nombre

	partido.candidatosDelPartido = candidatos

	var votos [3]int
	partido.votosActuales = votos

	return partido
}

func CrearVotosEnBlanco() Partido {
	partido := new(partidoEnBlanco)

	partido.nombrePartido = "Votos en Blanco"

	var votos [3]int
	partido.votosActuales = votos

	return partido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	(*partido).votosActuales[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.votosActuales[tipo] == 1 {
		return fmt.Sprintf("%s - %s: %d voto", partido.nombrePartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo])
	}
	return fmt.Sprintf("%s - %s: %d votos", partido.nombrePartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	(*blanco).votosActuales[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.votosActuales[tipo] == 1 {
		return fmt.Sprintf("%s: %d voto", blanco.nombrePartido, blanco.votosActuales[tipo])
	}
	return fmt.Sprintf("%s: %d votos", blanco.nombrePartido, blanco.votosActuales[tipo])
}
