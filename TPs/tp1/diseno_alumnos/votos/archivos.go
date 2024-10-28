package votos

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/diseno_alumnos/errores"
	"strconv"
	"strings"
)

func LecturaArchivos(partidos *[]Partido, padron *[]int) bool {
	if len(os.Args) == 3 {
		archPartidos := os.Args[1]
		archPadron := os.Args[2]
		if LeerPadron(archPadron, padron) || LeerPartidos(archPartidos, partidos) {
			return true
		}
	} else {
		fmt.Fprintln(os.Stdout, new(errores.ErrorParametros).Error())
		return true
	}
	return false
}

func LeerPartidos(nombreArchivo string, listaArchivo *[]Partido) bool {
	archivo, err := os.Open(nombreArchivo)
	*listaArchivo = append(*listaArchivo, CrearVotosEnBlanco())
	if err != nil {
		fmt.Fprintln(os.Stdout, new(errores.ErrorLeerArchivo).Error())
		return true
	}
	lineaArchivo := bufio.NewScanner(archivo)
	for lineaArchivo.Scan() {
		lineas := lineaArchivo.Text()
		linea := strings.Split(lineas, ",")
		if len(linea) < 3 {
			fmt.Fprintln(os.Stdout, new(errores.ErrorLeerArchivo).Error())
			return true
		}
		partido := CrearPartido(linea[0], [CANT_VOTACION]string{linea[1], linea[2], linea[3]})
		*listaArchivo = append(*listaArchivo, partido)
	}
	return false
}

func LeerPadron(nombreArchivo string, padron *[]int) bool {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Fprintln(os.Stdout, new(errores.ErrorLeerArchivo).Error())
		return true
	}
	linea := bufio.NewScanner(archivo)
	for linea.Scan() {
		str, err := strconv.Atoi(linea.Text())
		if err != nil {
			fmt.Fprintln(os.Stdout, new(errores.ErrorLeerArchivo).Error())
			return true
		}
		*padron = append(*padron, str)
	}
	*padron = OrdenarPadron(*padron, len(*padron))
	return false
}

func OrdenarPadron(padron []int, largo int) []int {
	if largo == 1 {
		return padron
	}
	medio := largo / 2
	primeraOrdenada := OrdenarPadron(padron[0:medio], medio)
	segundaOrdenada := OrdenarPadron(padron[medio:largo], largo-medio)

	return InsertarOrdenado(primeraOrdenada, segundaOrdenada)
}

func InsertarOrdenado(primeraOrdenada []int, segundaOrdenada []int) []int {
	sliceOrdenado := make([]int, 0)

	for len(primeraOrdenada) > 0 && len(segundaOrdenada) > 0 {
		if primeraOrdenada[0] > segundaOrdenada[0] {
			sliceOrdenado = append(sliceOrdenado, segundaOrdenada[0])
			segundaOrdenada = segundaOrdenada[1:]
		} else {
			sliceOrdenado = append(sliceOrdenado, primeraOrdenada[0])
			primeraOrdenada = primeraOrdenada[1:]
		}
	}
	for len(primeraOrdenada) > 0 {
		sliceOrdenado = append(sliceOrdenado, primeraOrdenada[0])
		primeraOrdenada = primeraOrdenada[1:]
	}
	for len(segundaOrdenada) > 0 {
		sliceOrdenado = append(sliceOrdenado, segundaOrdenada[0])
		segundaOrdenada = segundaOrdenada[1:]
	}
	return sliceOrdenado
}
