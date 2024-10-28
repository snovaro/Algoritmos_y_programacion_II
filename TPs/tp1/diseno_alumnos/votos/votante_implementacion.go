package votos

import (
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	TDApila "tdas/pila"
)

type votanteImplementacion struct {
	DNI       int
	pilaVotos TDApila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.DNI = dni
	votante.pilaVotos = TDApila.CrearPilaDinamica[Voto]()
	votante.pilaVotos.Apilar(*new(Voto))
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.DNI
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) {
	ViejoVoto := votante.pilaVotos.VerTope()
	NuevoVoto := new(Voto)
	NuevoVoto.VotoPorTipo = ViejoVoto.VotoPorTipo
	(*NuevoVoto).Impugnado = ViejoVoto.Impugnado
	(*NuevoVoto).VotoPorTipo[tipo] = alternativa
	if alternativa == 0 {
		(*NuevoVoto).Impugnado = true
	}
	votante.pilaVotos.Apilar(*NuevoVoto)
	fmt.Fprintln(os.Stdout, "OK")
}

func (votante *votanteImplementacion) Deshacer() error {
	ViejoVoto := votante.pilaVotos.Desapilar()
	if votante.pilaVotos.EstaVacia() {
		votante.pilaVotos.Apilar(ViejoVoto)
		return fmt.Errorf(new(errores.ErrorNoHayVotosAnteriores).Error())
	}
	fmt.Fprintln(os.Stdout, "OK")
	return nil
}

func (votante *votanteImplementacion) FinVoto() Voto {
	voto := votante.pilaVotos.VerTope()
	return voto
}
