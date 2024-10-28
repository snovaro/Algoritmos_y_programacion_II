package votos

import (
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	TDAlista "tdas/lista"
)

func asignacionCargo(tipo string) TipoVoto {
	if tipo == "Presidente" {
		return PRESIDENTE
	} else if tipo == "Gobernador" {
		return GOBERNADOR
	} else if tipo == "Intendente" {
		return INTENDENTE
	}
	return -1
}

func BuscarPadron(ini, fin, dniIngreso int, padron []int) bool {
	if len(padron) <= ini {
		return false
	}
	medio := (ini + fin) / 2
	if ini == fin {
		return padron[ini] == dniIngreso
	}
	if padron[medio] == dniIngreso {
		return true
	}
	if padron[medio] < dniIngreso {
		return BuscarPadron(medio+1, fin, dniIngreso, padron)
	}
	return BuscarPadron(ini, medio, dniIngreso, padron)
}

func VerificarFraude(sliceVotantes []int, votante *Votante) bool {
	return BuscarPadron(0, len(sliceVotantes), (*votante).LeerDNI(), sliceVotantes)
}

func ImprimirErrorFraude(iteradorVotantes *TDAlista.IteradorLista[Votante], votante *Votante) {
	(*iteradorVotantes).Borrar()
	errFraude := new(errores.ErrorVotanteFraudulento)
	errFraude.Dni = (*votante).LeerDNI()
	fmt.Fprintln(os.Stdout, errFraude)
}
