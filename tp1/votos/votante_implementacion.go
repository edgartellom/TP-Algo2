package votos

import (
	e "rerepolez/errores"
	"tdas/pila"
)

type votanteImplementacion struct {
	dni   int
	voto  *Voto
	votos pila.Pila[int]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	tipo = TipoVoto(alternativa)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votos.EstaVacia() {
		return e.ErrorNoHayVotosAnteriores{}
	} else {
		// votoAnterior := votante.votos.Desapilar()

		// votante.voto = TipoVoto(votoAnterior)
		return nil
	}
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return Voto{}, nil
}
