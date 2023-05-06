package votos

import (
	"rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni             int
	yaVoto          bool
	votoActual      *Voto
	votoInicial     *Voto
	votosRealizados TDAPila.Pila[*Voto]
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
	copiaDelVoto := copiarVoto(*votante.votoActual)
	(*votante).votosRealizados.Apilar(copiaDelVoto)
	if alternativa == LISTA_IMPUGNA {
		(*votante).votoActual.Impugnado = true
	}
	(*votante).votoActual.VotoPorTipo[tipo] = alternativa

	return votante.comprobarSiYaVoto()
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.votosRealizados.VerTope() == votante.votoInicial {
		errror := new(errores.ErrorNoHayVotosAnteriores)
		return errror
	}
	(*votante).votoActual = (*votante).votosRealizados.Desapilar()

	return votante.comprobarSiYaVoto()
}

func copiarVoto(votoDelVotante Voto) *Voto {
	copia := new(Voto)
	copia.Impugnado = votoDelVotante.Impugnado
	copia.VotoPorTipo = votoDelVotante.VotoPorTipo
	return copia
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	errror := votante.comprobarSiYaVoto()
	if errror == nil {
		votante.yaVoto = true
	}
	return *votante.votoActual, errror
}

func (votante votanteImplementacion) comprobarSiYaVoto() error {
	if votante.yaVoto {
		errror := new(errores.ErrorVotanteFraudulento)
		errror.Dni = votante.LeerDNI()
		return errror
	}
	return nil
}
