package votos

import (
	errores "rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	votosRealizados    TDAPila.Pila[*Voto]
	votoInicial        *Voto
	votoActual         *Voto
	votanteFraudulento bool
	dni                int
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)

	votante.dni = dni
	votante.votosRealizados = TDAPila.CrearPilaDinamica[*Voto]()
	votante.votoInicial = new(Voto)
	votante.votoActual = new(Voto)

	votante.votosRealizados.Apilar(votante.votoInicial)
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	copiaDelVotoActual := copiarVotoActual(*votante.votoActual)
	(*votante).votosRealizados.Apilar(copiaDelVotoActual)

	if alternativa == LISTA_IMPUGNA {
		(*votante).votoActual.Impugnado = true
	}
	(*votante).votoActual.VotoPorTipo[tipo] = alternativa

	return votante.comprobarVotanteFraudulento()
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votosRealizados.VerTope() == votante.votoInicial {
		err := new(errores.ErrorNoHayVotosAnteriores)
		return err
	}
	(*votante).votoActual = (*votante).votosRealizados.Desapilar()

	return votante.comprobarVotanteFraudulento()
}

func copiarVotoActual(votoDelVotante Voto) *Voto {
	copia := new(Voto)
	copia.Impugnado = votoDelVotante.Impugnado
	copia.VotoPorTipo = votoDelVotante.VotoPorTipo
	return copia
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	err := votante.comprobarVotanteFraudulento()
	if err == nil {
		votante.votanteFraudulento = true
	}
	return *votante.votoActual, err
}

func (votante votanteImplementacion) comprobarVotanteFraudulento() error {
	if votante.votanteFraudulento {
		err := new(errores.ErrorVotanteFraudulento)
		err.Dni = votante.LeerDNI()
		return err
	}
	return nil
}
