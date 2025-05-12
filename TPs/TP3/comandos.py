import funciones_auxiliares as aux
import grafo
import biblioteca_funciones as biblioteca_funciones

def listar_operaciones(opciones:tuple):
    for opcion in opciones:
        print(opcion)

def camino(entrada:list, g:grafo.Grafo):
    origen, destino = aux.obtener_info_str(entrada)
    camino = biblioteca_funciones.camino_minimo(g, origen, destino)
    if len(camino) == 0:
        print("No se encontro recorrido")
    else:
        aux.impresion_camino(camino[0], camino[1])

def diametro(g:grafo.Grafo):
    diam, camino = biblioteca_funciones.diametro(g)
    return aux.impresion_camino(camino, diam)

def ciclo_N(entrada:list, g):
    pagina, n = aux.obtener_info_str(entrada)
    if pagina not in g.obtener_vertices():
        print("La pagina no se encuentra en este archivo")
        return
    n = int(n)
    camino = biblioteca_funciones.ciclo_de_n(g, pagina, n)
    aux.impresion_camino(camino)

def a_rango_N(entrada:list, g:grafo.Grafo):
    pagina, n = aux.obtener_info_str(entrada)
    n = int(n)
    print(biblioteca_funciones.todos_en_Rango_N(g, n, pagina))

def lectura_antitopologica(entrada, g):
    paginas = aux.obtener_info_str(entrada)
    g2 = grafo.Grafo(True, paginas)
    for pagina in paginas:
        for link in g.adyacentes(pagina):
            if link in paginas:
                g2.agregar_arista(pagina, link)
    orden = biblioteca_funciones.orden_inverso(g2)
    if len(orden) != len(paginas):
        print("No existe forma de leer las paginas en orden")
    else:
        for i in range(len(orden)):
            if i != len(orden)-1:
                print(orden[i], end=", ")
            else:
                print(orden[i])

def navegar_primer_link(entrada, g):
    pagina = entrada[1]
    if len(entrada) > 2:
        pagina = " ".join(entrada[1:])
    camino = []
    if pagina not in g.obtener_vertices():
        return
    while len(camino) <21:
        camino.append(pagina)
        adyacentes = g.adyacentes(pagina)
        if len(adyacentes) == 0:
            break
        pagina = adyacentes[0]
    aux.impresion_camino(camino)

def clustering(entrada, g):
    pagina = None
    if len(entrada) > 1:
        pagina = " ".join(entrada[1:])
    resultado = biblioteca_funciones.clustering(g, pagina)
    print("%.3f" % resultado)