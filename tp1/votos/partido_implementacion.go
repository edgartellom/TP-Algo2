package votos

import "fmt"

type partidoImplementacion struct {
	numeroDeLista        int
	nombrePartido        string
	candidatosDelPartido []string
	votosActuales        [3]int
}

type partidoEnBlanco struct {
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

	partido.numeroDeLista = 1

	partido.nombrePartido = nombre

	postulantes := make([]string, len(candidatos))
	copy(postulantes, candidatos[:])
	partido.candidatosDelPartido = postulantes

	var votos [3]int
	partido.votosActuales = votos

	return partido
}

func CrearVotosEnBlanco() Partido {
	return nil
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	switch tipo {
	case PRESIDENTE:
		(*partido).votosActuales[0]++
	case GOBERNADOR:
		(*partido).votosActuales[1]++
	case INTENDENTE:
		(*partido).votosActuales[2]++
	}
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	switch tipo {
	case PRESIDENTE:
		return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatosDelPartido[0], partido.votosActuales[tipo])
	case GOBERNADOR:
		return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatosDelPartido[1], partido.votosActuales[tipo])
	case INTENDENTE:
		return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatosDelPartido[2], partido.votosActuales[tipo])
	}
	return ""
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {

}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
