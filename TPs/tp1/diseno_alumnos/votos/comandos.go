package votos

import (
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	"strconv"
	TDAlista "tdas/lista"
)

func Ingreso(vectorIngreso []string, padron []int, iteradorVotantes *TDAlista.IteradorLista[Votante], listaVotantes TDAlista.Lista[Votante], sliceVotantes []int) {
	dni, err := strconv.Atoi(vectorIngreso[1])
	if err == nil && BuscarPadron(0, len(padron)-1, dni, padron) {
		if !(*iteradorVotantes).HaySiguiente() {
			(*iteradorVotantes).Insertar(CrearVotante(dni))
		} else {
			listaVotantes.InsertarUltimo(CrearVotante(dni))
		}

		fmt.Fprintln(os.Stdout, "OK")

	} else if err != nil {
		fmt.Fprintln(os.Stdout, new(errores.DNIError).Error())

	} else {
		fmt.Fprintln(os.Stdout, new(errores.DNIFueraPadron).Error())

	}
}

func Votacion(vectorIngreso []string, partidos []Partido, iteradorVotantes *TDAlista.IteradorLista[Votante], votante *Votante, sliceVotantes []int) {
	tipoCargo := asignacionCargo(vectorIngreso[1])
	if tipoCargo != -1 {
		alternativa, errAlternativa := strconv.Atoi(vectorIngreso[2])
		if errAlternativa == nil && alternativa < len(partidos) && alternativa >= 0 {
			if VerificarFraude(sliceVotantes, votante) {
				ImprimirErrorFraude(iteradorVotantes, votante)
			} else {
				(*votante).Votar(tipoCargo, alternativa)
			}
		} else {
			fmt.Fprintln(os.Stdout, new(errores.ErrorAlternativaInvalida).Error())

		}
	} else {
		fmt.Fprintln(os.Stdout, new(errores.ErrorTipoVoto).Error())

	}
}

func DeshizoVoto(votante *Votante, sliceVotantes []int, iteradorVotantes *TDAlista.IteradorLista[Votante]) {
	if VerificarFraude(sliceVotantes, votante) {
		ImprimirErrorFraude(iteradorVotantes, votante)
	} else {
		errDeshacer := (*votante).Deshacer()
		if errDeshacer != nil {
			fmt.Fprintln(os.Stdout, errDeshacer)
		}
	}
}

func Finalizarvoto(iteradorVotantes *TDAlista.IteradorLista[Votante], votante *Votante, sliceVotantes *[]int, finalizo *bool) {
	*finalizo = true
	if VerificarFraude(*sliceVotantes, votante) {
		ImprimirErrorFraude(iteradorVotantes, votante)
	} else {
		fmt.Fprintln(os.Stdout, "OK")
		(*iteradorVotantes).Siguiente()
	}
	if !BuscarPadron(0, len(*sliceVotantes), (*votante).LeerDNI(), *sliceVotantes) {
		*sliceVotantes = InsertarOrdenado(*sliceVotantes, []int{(*votante).LeerDNI()})
	}
}
