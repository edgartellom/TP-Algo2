PANIC_NO_PERTENECE = "El vertice no pertenece al Grafo"

"""Clase Grafo(no compuesto)"""
class Grafo:

    """Constructor: Por ahora diccionario principal unicamente."""
    def __init__(self, es_dirigido = False):
        self.diccionario_principal = {}
        self.es_dirigido = es_dirigido

    """Cantidad de vertices totales en el Grafo."""
    def __len__(self):
        return len(self.diccionario_principal)

    """Devuelve si el vertice Pertenece al diccionario principal."""
    def pertenece(self, vertice):
        return vertice in self.diccionario_principal

    """Agrega un vertice al Grafo."""
    def agregar_vertice(self, vertice):
        self.diccionario_principal[vertice] = {}

    """Borrar_vertice, si pertenece, borra el vertice y lo devuelve, sino provoca una excepción."""
    def borrar_vertice(self, vertice):
        if not self.pertenece(vertice):
            raise Exception(PANIC_NO_PERTENECE)
        return self.diccionario_principal.pop(vertice)

    """Agrega una arista entre los vertices indicados por parámetro."""
    def agregar_arista(self, vertice1, vertice2, peso = None):
        self.diccionario_principal[vertice1][vertice2] = peso

        if not self.es_dirigido:
            self.diccionario_principal[vertice2][vertice1] = peso
    
    """Si la arista existe, la borra, en caso contrario provoca una excepción."""
    def borrar_arista(self, vertice1, vertice2):
        if (vertice1 not in self.diccionario_principal) or (vertice2 not in self.diccionario_principal[vertice1]):
            raise Exception(f"Los vertices, {vertice1} y {vertice2}, no están conectados.")
        
        self.diccionario_principal[vertice1].pop(vertice2)
        
        if not self.es_dirigido:
            self.diccionario_principal[vertice2].pop(vertice1)

    """Devuelve una lista de todos los vertices."""
    def obtener_vertices(self):
        return list(self.diccionario_principal.keys())
    
    """Devuelve una lista de todos los vertices adyacentes al vertice indicado por parámetro."""
    def obtener_adyacentes(self, vertice):
        return list(self.diccionario_principal[vertice].keys())

    """Permite recorrer el Grafo con un "for" e ir de vertice en vertice."""
    def __iter__(self):
        return iter(self.diccionario_principal)
    
    """Permite que la funcion "print" muestre el diccionario interno del Grafo."""
    def __str__(self):
        return (f"{self.diccionario_principal}")

    """Devuelve si existe una arista entre los vertices, y en caso de que exista devuelve el peso de esta."""
    def hay_arista(self, vertice1, vertice2):
        if self.pertenece(vertice2):
            return True, self.diccionario_principal[vertice1][vertice2]
        return False, None