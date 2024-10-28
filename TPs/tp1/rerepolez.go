package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	TDAvotos "rerepolez/diseno_alumnos/votos"
	"strings"
	TDAlista "tdas/lista"
)

const (
	SELECCION1 = "ingresar"
	SELECCION2 = "votar"
	SELECCION3 = "deshacer"
	SELECCION4 = "fin-votar"
)

func main() {
	var partidos []TDAvotos.Partido
	var padron []int
	var votante TDAvotos.Votante
	var finalizo bool
	impugnados := 0
	sliceVotantes := make([]int, 0)
	if TDAvotos.LecturaArchivos(&partidos, &padron) {
		return
	}

	listaVotantes := TDAlista.CrearListaEnlazada[TDAvotos.Votante]()
	iteradorVotantes := listaVotantes.Iterador()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && scanner.Text() != "" {
		input := scanner.Text()
		vectorIngreso := strings.Split(input, " ")

		if iteradorVotantes.HaySiguiente() {
			votante = iteradorVotantes.VerActual()
			finalizo = false
		} else if vectorIngreso[0] != SELECCION1 {
			fmt.Fprintln(os.Stdout, new(errores.FilaVacia).Error())
			continue
		}

		if vectorIngreso[0] == SELECCION1 {
			TDAvotos.Ingreso(vectorIngreso, padron, &iteradorVotantes, listaVotantes, sliceVotantes)
		} else if vectorIngreso[0] == SELECCION2 {
			TDAvotos.Votacion(vectorIngreso, partidos, &iteradorVotantes, &votante, sliceVotantes)
		} else if vectorIngreso[0] == SELECCION3 {
			TDAvotos.DeshizoVoto(&votante, sliceVotantes, &iteradorVotantes)
		} else if vectorIngreso[0] == SELECCION4 {
			TDAvotos.Finalizarvoto(&iteradorVotantes, &votante, &sliceVotantes, &finalizo)
		}
	}

	CiudadanosSinVotar(iteradorVotantes)
	RecuentoDeVotos(votante, listaVotantes, &partidos, finalizo, &impugnados)
	ImprimirResultados(&partidos, impugnados)
}

func RecuentoDeVotos(votante TDAvotos.Votante, listaVotantes TDAlista.Lista[TDAvotos.Votante], partidos *[]TDAvotos.Partido, finalizo bool, impugnados *int) {
	var votanteDeLista TDAvotos.Votante

	for !listaVotantes.EstaVacia() && listaVotantes.VerPrimero() != votante {
		votanteDeLista = listaVotantes.BorrarPrimero()
		SumaVotoOImpugnado(votanteDeLista, partidos, impugnados)
	}
	if finalizo {
		SumaVotoOImpugnado(votante, partidos, impugnados)
	}
}

func SumaVotoOImpugnado(votante TDAvotos.Votante, partidos *[]TDAvotos.Partido, impugnados *int) {
	if !votante.FinVoto().Impugnado {
		ContarVotos(partidos, votante)
	} else {
		*impugnados += 1
	}
}

func CiudadanosSinVotar(iter TDAlista.IteradorLista[TDAvotos.Votante]) {
	if iter.HaySiguiente() {
		fmt.Fprintln(os.Stdout, new(errores.ErrorCiudadanosSinVotar).Error())
	}
}

func ContarVotos(listaPartidos *[]TDAvotos.Partido, votante TDAvotos.Votante) {
	votos := votante.FinVoto().VotoPorTipo
	cargos := []TDAvotos.TipoVoto{TDAvotos.PRESIDENTE, TDAvotos.GOBERNADOR, TDAvotos.INTENDENTE}
	for indice, partido := range votos {
		(*listaPartidos)[partido].VotadoPara(cargos[indice])
	}
}

func ImprimirResultados(listaPartidos *[]TDAvotos.Partido, impugnados int) {
	cargos := []TDAvotos.TipoVoto{TDAvotos.PRESIDENTE, TDAvotos.GOBERNADOR, TDAvotos.INTENDENTE}
	cargosString := []string{"Presidente:", "Gobernador:", "Intendente:"}

	for indice, cargoTitulo := range cargosString {
		fmt.Fprintln(os.Stdout, cargoTitulo)
		for _, partido := range *listaPartidos {
			fmt.Fprintln(os.Stdout, partido.ObtenerResultado(cargos[indice]))
		}
		fmt.Fprintln(os.Stdout, "")
	}
	impresionImpugnados := "Votos Impugnados: " + fmt.Sprint(impugnados)
	TDAvotos.ImprimirVotoOVotos(&impresionImpugnados, impugnados)
	fmt.Fprintln(os.Stdout, impresionImpugnados)
}
