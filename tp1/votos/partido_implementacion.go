package votos

import "fmt"

const (
	NOMBRE_PARTIDO_EN_BLANCO = "Votos en Blanco"
	SINGULAR                 = "voto"
	PLURAL                   = "votos"
)

type partido struct {
	nombreDelPartido string
	votosActuales    [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	partido
}

type partidoImplementacion struct {
	partido
	candidatosDelPartido [CANT_VOTACION]string
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

func (partido *partido) palabraSegunCantidad(tipo TipoVoto) string {
	if partido.votosActuales[tipo] == 1 {
		return SINGULAR
	}
	return PLURAL
}

func (partido *partido) VotadoPara(tipo TipoVoto) {
	(*partido).votosActuales[tipo]++
}

func (partido *partido) ObtenerResultado(tipo TipoVoto) string {
	cantidad := partido.palabraSegunCantidad(tipo)
	return fmt.Sprintf("%s: %d %s", partido.nombreDelPartido, partido.votosActuales[tipo], cantidad)
}

func (partido *partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	cantidad := partido.palabraSegunCantidad(tipo)
	return fmt.Sprintf("%s - %s: %d %s", partido.nombreDelPartido, partido.candidatosDelPartido[tipo], partido.votosActuales[tipo], cantidad)
}
