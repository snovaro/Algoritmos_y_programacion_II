#Entiendo que aca hay que agregar 
#funciones como puede ser buscar camino minimo, o buscar tal cosa, o nose la verdad
import grafo
from collections import deque

def bfs_desde_v(g:grafo.Grafo, origen):
    cola = deque()
    padres = {}
    orden = {}
    cola.append(origen)
    orden[origen] = 0
    padres[origen] = None
    while len(cola) > 0:
        v = cola.popleft()
        for w in g.adyacentes(v):
            if w not in orden:
                orden[w] = orden[v] + 1
                cola.append(w)
                padres[w] = v
    return padres, orden

def camino_minimo(g:grafo.Grafo, origen, destino):
    padres,dist = bfs_desde_v(g, origen)
    if destino in padres:
        return reconstruir_camino(padres, destino), dist[destino]
    return []   ##NO HAY CAMINO DESDE ORIGEN A DESTINO

def grafo_esta_vacio(g) -> bool:
    vertices = g.obtener_vertices()
    if len(vertices) == 0:
        return True
    return False

def diametro(g:grafo.Grafo):
    vertices = g.obtener_vertices()
    if len(vertices) == 0:
        return 0
    max = 0
    mas_lejano = g.vertice_aleatorio()
    for v in vertices:
        padres,distancias = bfs_desde_v(g, v)
        for w,distancia in distancias.items():
            if distancia > max:
                max = distancia
                mas_lejano = w
                padres_max = padres
    camino_al_mas_lejos = reconstruir_camino(padres_max, mas_lejano)
    return max, camino_al_mas_lejos

def reconstruir_camino(padres, destino):
    camino = []
    v = destino
    while v != None:
        camino.append(v)
        v = padres[v]
    camino = camino[::-1]
    return camino

def pila_a_lista(pila):
    lista = []
    while len(pila) != 0:
        lista.append(pila.pop())
    return lista

def ciclo_de_n(g,v,n):
    pila = deque()
    visitados = []
    ciclo_de_n_recur(g,v,v,n,pila,visitados)
    return pila_a_lista(pila)


def ciclo_de_n_recur(g,v_ini,v,n,pila,visitados):
    if v != v_ini:
        visitados.append(v)
    if n == 0:
        if v == v_ini:
            pila.append(v)
        return
    for w in g.adyacentes(v):
        if len(pila) != 0:
            break
        if w not in visitados:
            ciclo_de_n_recur(g,v_ini,w,n-1,pila,visitados)
    if len(pila) != 0:
        pila.append(v)
    return

def todos_en_Rango_N(g, n, v):
    _, dist = bfs_desde_v(g, v)
    contador = 0
    for vertice in dist:
        if dist[vertice] == n:
            contador += 1
    return contador

def orden_inverso(g:grafo.Grafo):
    vertices = g.obtener_vertices()
    orden = {}
    hijos = {}
    cola = deque()
    res = []
    for v in vertices:
        orden[v] = 0
        hijos[v] = []
    for v in vertices:
        for w in g.adyacentes(v):
            if w in hijos[v]:
                return []#no hay orden posible
            orden[v] -=1
            hijos[w].append(v)
    for v in vertices:
        if orden[v] == 0:
            cola.append(v)
    while len(cola) > 0:
        v = cola.popleft()
        res.append(v)
        for w in hijos[v]:
            orden[w] +=1
            if orden[w] == 0:
                cola.append(w)
    return res


def clustering_desde_v(g,v):
    cantidad_adyacentes = 0
    adyacentes_v = g.adyacentes(v)
    g_salida = len(adyacentes_v)
    if v in adyacentes_v:
        g_salida -= 1
    if g_salida < 2 :
        return 0.000
    cantidad_maxima_adyacentes = (g_salida*(g_salida-1))
    for w in adyacentes_v:
        if w != v:    
            for x in g.adyacentes(w):
                if x != v and x != w and x in adyacentes_v:
                    cantidad_adyacentes += 1
    coeficiente = cantidad_adyacentes / cantidad_maxima_adyacentes
    return coeficiente

def clustering(g,v):
    cantidad_coeficientes = 0
    coeficiente = 0.000
    if v != None:
        coeficiente = clustering_desde_v(g,v)
    else:
        for w in g.obtener_vertices():
            coeficiente += clustering_desde_v(g,w)
            cantidad_coeficientes += 1
        coeficiente /= cantidad_coeficientes
    return round(coeficiente,3)