package votos

import (
	"rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni             int
	votoActual      *Voto
	votoInicial     *Voto
	votosRealizados TDAPila.Pila[*Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.votosRealizados = TDAPila.CrearPilaDinamica[*Voto]()

	votoInicial := new(Voto)
	votante.votoInicial = votoInicial
	votante.votosRealizados.Apilar(votoInicial)
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	votoDelVotante := new(Voto)
	if alternativa == LISTA_IMPUGNA {
		votoDelVotante.Impugnado = true
	}
	votoDelVotante.VotoPorTipo[tipo] = alternativa
	(*votante).votoActual = votoDelVotante
	(*votante).votosRealizados.Apilar(votoDelVotante)

	return nil // PENSAR EL ERROR!!!
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votosRealizados.VerTope() == votante.votoInicial {
		errror := new(errores.ErrorNoHayVotosAnteriores)
		return errror
	}
	(*votante).votosRealizados.Desapilar()
	(*votante).votoActual = votante.votosRealizados.VerTope()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	return *votante.votoActual, nil // PENSAR EL ERROR!!!
}
