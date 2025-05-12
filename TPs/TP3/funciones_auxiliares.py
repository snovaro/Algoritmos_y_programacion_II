import csv
import grafo

def obtener_info_str(entrada):
    string = " ".join(entrada)
    return string.split(",")

def obtener_dicc_pagina_links(nombreArchivo:str) -> dict:
    dicc: dict = {}
    with open(nombreArchivo, newline='') as usuariosCsv:
        csvReader = csv.reader(usuariosCsv, delimiter = "\t")
        for row in csvReader:
            valores = row[1:]
            dicc[row[0]] = valores
    return dicc

def crear_grafo_de_paginas(nombreArchivo):
    dicc_paginas = obtener_dicc_pagina_links(nombreArchivo)
    g:grafo = grafo.Grafo(True, dicc_paginas.keys())
    for pagina in dicc_paginas:
        for link in dicc_paginas[pagina]:
            g.agregar_arista(pagina, link, 1)
    return g

def impresion_camino(camino, costo = -1):
    if len(camino) == 0:
        print("No se encontro recorrido")
        return
    impresion = ""
    for vertice in camino:
        impresion += vertice + " -> "
    impresion = impresion[:len(impresion)-4] 
    print(impresion)
    if costo != -1:
        print("Costo:", costo)
    return impresion + "\nCosto: " + str(costo)
