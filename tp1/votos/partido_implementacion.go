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

/*
PARTIDO : NUMERO DE LISTA: 1
		  NOMBRE: LOS GATOS
		  CANDIDATOS: [Mondi, Viena, Mrużka] ===> SIENDO: Mondi  PRESIDENTE (TipoVoto 0),
															Viena  GOBERNADOR (TipoVoto 1),
															Mruźka INTENDENTE (TipoVoto 2).
			VOTOS
			  A     : [  0  ,   0  ,   0   ]
		  CANDIDATO
*/

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
	return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	(*blanco).votosActuales[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("%s: %d votos\n", blanco.nombrePartido, blanco.votosActuales[tipo])
}
