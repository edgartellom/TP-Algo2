import random

"""Clase Grafo (no compuesto)"""
class Grafo:
# Grafo() para crear un grafo no dirigido
# Grafo(dirigido=True) para crear un grafo dirigido

    """Constructor: Por ahora diccionario principal unicamente."""
    def __init__(self, dirigido = False, vertices = {}):
        self.dirigido = dirigido
        self.vertices = vertices


    """Agrega un vertice al Grafo."""
    def agregar_vertice(self, v):
        self.vertices[v] = {}


    """Borrar_vertice, borra el vertice y lo devuelve, sino provoca una excepción."""
    def borrar_vertice(self, v):
        self.comprobar_pertenece(v)

        if v in self.vertices:
            for w in self.vertices[v]:
                del self.vertices[w][v]
            del self.vertices[v]


    """Agrega una arista entre los vertices indicados por parámetro."""
    def agregar_arista(self, v, w, peso = 1):
        self.comprobar_pertenece(v , w)
        
        self.vertices[v][w] = peso

        if not self.dirigido:
            self.vertices[w][v] = peso


    """Si la arista existe, la borra, en caso contrario provoca una excepción."""
    def borrar_arista(self, v, w):
        self.comprobar_pertenece(v, w)
        
        self.comprobar_hay_arista(v, w)
        
        del self.vertices[v][w]
        
        if not self.dirigido:
            del self.vertices[w][v]


    """Devuelve si existe una arista entre los vertices."""
    def hay_arista(self, v, w):
        return w in self.vertices[v] or v in self.vertices[w]


    """Devuelve el peso de una arista."""
    def peso_arista(self, v, w):
        return self.vertices[v][w]


    """Devuelve una lista con todos los vertices del grafo."""
    def obtener_vertices(self):
        return list(self.vertices)


    """Devuelve un vertice aleatorio del grafo."""
    def vertice_aleatorio(self):
        if not self.vertices:
            return None
        return random.choice(list(self.vertices))


    """Devuelve una lista de todos los vertices adyacentes al vertice indicado por parámetro."""
    def adyacentes(self, v):
        return list(self.vertices[v])


    """Devuelve si el(los) vertices pertenece(n) al grafo."""    
    def pertenece(self, v, w=None):
        if w != None:
            return v in self.vertices and w in self.vertices
        return v in self.vertices


    """Permite recorrer el Grafo con un "for" e ir de vertice en vertice."""
    def __iter__(self):
        return iter(self.vertices)


    """Permite que la funcion "print" muestre el diccionario interno del Grafo."""
    def __str__(self):
        return f"Grafo con {len(self.vertices)} vértices: {', '.join(self.vertices)}" 


    def comprobar_pertenece(self, v, w=None):
        if not self.pertenece(v, w):
            if w != None:
                raise Exception(f"El vértice {v} o {w} no pertenece al grafo")
            raise Exception(f"El vértice {v} no pertenece al grafo")


    def comprobar_hay_arista(self, v, w):
        if not self.hay_arista(v, w):
            raise Exception(f"Los vertices, {v} y {w} no están conectados.")