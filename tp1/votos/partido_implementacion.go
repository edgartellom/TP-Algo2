package votos

import "fmt"

type partidoImplementacion struct {
	nombre        string
	candidatos    []string
	votosActuales [CANT_VOTACION]int
	numeroLista   int
}

type partidoEnBlanco struct {
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos[:]
	return partido
}

func CrearVotosEnBlanco() Partido {
	return nil
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosActuales[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("%s - %s: %d votos\n", partido.nombre, partido.candidatos[tipo], partido.votosActuales[tipo])
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return ""
}
