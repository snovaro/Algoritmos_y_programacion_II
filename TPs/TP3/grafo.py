from collections import deque
import random
#Grafo:
#es_dirigido bool
#diccionario {{}}
class Grafo:
    def __init__(self, dirigido:bool = False, vertices_iniciales = []) -> None:
        self.dirigido:bool = dirigido
        self.diccionario:dict = {}
        for v in vertices_iniciales:
            self.diccionario[v] = {}
    def __str__(self) -> str:
        cadenaPrint = "{"
        contadorSeparador1 = 0
        for v in self.diccionario:
            if contadorSeparador1 >0:
                cadenaPrint += ","
            cadenaPrint += "\n" + str(v) + ":{"
            contadorSeparador2 = 0
            for w in self.diccionario[v]:
                if contadorSeparador2 > 0:
                    cadenaPrint += ", "
                contadorSeparador2 += 1
                cadenaPrint += str(w)+ ":" + str(self.diccionario[v][w])
            cadenaPrint += "}"
        cadenaPrint += "\n}"
        return cadenaPrint
    def agregar_vertice(self, v):
        if v in self.diccionario:
            print("El vertice ya esta cargado.")
            return #error
        self.diccionario[v] = {}
    def borrar_vertice(self, v):
        del self.diccionario[v]
        for w in self.diccionario:
            if v in self.diccionario[w]:
                del self.diccionario[w][v]
    def agregar_arista(self, v, w, peso = 0):
        self.diccionario[v][w] = peso
        if not self.dirigido:
            self.diccionario[w][v] = peso
    def borrar_arista(self, v, w):
        if v not in self.diccionario:
            print(f"El vertice {v} no esta en el diccionario")
            return
        if not self.estan_unidos(v,w):
            print(f"El vertice {v} no se encuentra unido al vertice {w}")
            return
        del self.diccionario[v][w]
        if not self.dirigido:
            del self.diccionario[w][v]
    def estan_unidos(self, v, w):
        return w in self.diccionario[v]
    def peso_arista(self, v, w):
        if not self.estan_unidos(v,w):
            print(f"No hay arista entres estos dos vertices o no hay arista desde {v} hacia {w}")
            return #error
        return self.diccionario[v][w]
    def obtener_vertices(self):
        return list(self.diccionario.keys())
    def vertice_aleatorio(self):
        return random.choice(list(self.diccionario.keys()))
    def adyacentes(self, v):
        if v in self.diccionario:
            return list(self.diccionario[v].keys())
        print(f"El vertice {v} no esta cargado en el grafo")