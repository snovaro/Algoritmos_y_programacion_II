package votos

import "fmt"

type partidoImplementacion struct {
	nombrePartido string
	votos         [CANT_VOTACION]int
	candidatos    [CANT_VOTACION]string
}

type partidoEnBlanco struct {
	votos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombrePartido = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	partido := new(partidoEnBlanco)
	return partido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo] += 1
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	impresion := ""
	impresion += fmt.Sprint(partido.nombrePartido) + " - " + fmt.Sprint(partido.candidatos[tipo]) + ": " + fmt.Sprint(partido.votos[tipo])

	ImprimirVotoOVotos(&impresion, partido.votos[tipo])
	return impresion
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo] += 1
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	impresion := "Votos en Blanco: " + fmt.Sprint(blanco.votos[tipo])
	ImprimirVotoOVotos(&impresion, blanco.votos[tipo])
	return impresion
}
