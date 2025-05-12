#!/usr/bin/env python3

import grafo
import sys
import funciones_auxiliares as aux
import comandos

LISTAR = "listar_operaciones"
OPCIONES = (
"camino",
"lectura",
"diametro",
"rango",
"navegacion",
"ciclo",
"clustering"
)

def main():
    nombreArchivo:str = sys.argv[1]
    g:grafo.Grafo = aux.crear_grafo_de_paginas(nombreArchivo)
    sacamos_diametro = False
    impresion_diametro = ""
    ingreso = input()
    while ingreso != "":
        entrada = ingreso.split()
        if entrada[0] == LISTAR:
            comandos.listar_operaciones(OPCIONES)
        elif entrada[0] == OPCIONES[0]:
            comandos.camino(entrada[1:], g)
        elif entrada[0] == OPCIONES[1]:
            comandos.lectura_antitopologica(entrada[1:], g)
        elif entrada[0] == OPCIONES[2]:
            if not sacamos_diametro:
                impresion_diametro = comandos.diametro(g)
                sacamos_diametro = True
            else:
                print(impresion_diametro)
        elif entrada[0] == OPCIONES[3]:
            comandos.a_rango_N(entrada[1:], g)
        elif entrada[0] == OPCIONES[4]:
            comandos.navegar_primer_link(entrada, g)
        elif entrada[0] == OPCIONES[5]:
            comandos.ciclo_N(entrada[1:], g)
        elif entrada[0] == OPCIONES[6]:
            comandos.clustering(entrada, g)
        ingreso = input()

main()