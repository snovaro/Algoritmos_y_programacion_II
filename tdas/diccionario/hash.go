package diccionario

import (
	"fmt"
)

type estados int

const (
	VACIO estados = iota
	OCUPADO
	BORRADO
	TAMANIO_INICIAL = 4
)

func funcion_hash[K comparable, V any](h *hash[K, V], clave K, largo int) int {
	var hashh uint32 = 2166136261
	claveStr := convertirABytes[K](clave)
	for _, c := range claveStr {
		hashh ^= uint32(c)
		hashh *= 16777619
	}
	return int(hashh) % largo
}

type hash[K comparable, V any] struct {
	lista            []nodoLista[K, V]
	cantidad         int
	cantidadBorrados int
}

type nodoLista[K comparable, V any] struct {
	elemento claveValor[K, V]
	estado   estados
}

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}
type iterDicccionario[K comparable, V any] struct {
	lista           []nodoLista[K, V]
	iteracionActual int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hash[K, V])
	slice := make([]nodoLista[K, V], TAMANIO_INICIAL)
	hash.lista = slice
	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hash[K, V]) Guardar(clave K, dato V) {
	posicion := funcion_hash(hash, clave, len(hash.lista))
	largoLista := len(hash.lista)
	for hash.lista[posicion].estado == OCUPADO {
		if hash.lista[posicion].elemento.clave == clave {
			(*hash).lista[posicion].elemento.valor = dato
			return
		}
		if posicion == largoLista-1 {
			posicion = 0
		} else {
			posicion++
		}
	}
	nuevoElemento := new(claveValor[K, V])
	nuevoElemento.clave = clave
	nuevoElemento.valor = dato
	(*hash).lista[posicion].elemento = *nuevoElemento
	(*hash).lista[posicion].estado = OCUPADO
	hash.cantidad++
	if 3*(hash.cantidad+hash.cantidadBorrados) > 2*len(hash.lista) {
		hash.redimensionar()
	}
}

func (hash hash[K, V]) Pertenece(clave K) bool {
	posicion := hash.buscarPosicion(clave)
	return posicion != -1
}
func (hash hash[K, V]) buscarPosicion(clave K) int {
	posicion := funcion_hash(&hash, clave, len(hash.lista))
	for hash.lista[posicion].estado != VACIO {
		if hash.lista[posicion].elemento.clave == clave && hash.lista[posicion].estado == OCUPADO {
			return posicion
		}
		posicion++
		if posicion == len(hash.lista) {
			posicion = 0
		}
	}
	return -1
}
func (hash hash[K, V]) Obtener(clave K) V {
	posicion := hash.buscarPosicion(clave)
	if posicion == -1 {
		panic("La clave no pertenece al diccionario")
	}
	return hash.lista[posicion].elemento.valor
}
func (hash *hash[K, V]) Borrar(clave K) V {
	if 5*hash.cantidad < len(hash.lista) {
		hash.redimensionar()
	}
	posicion := funcion_hash(hash, clave, len(hash.lista))
	for hash.lista[posicion].estado != VACIO {
		if hash.lista[posicion].elemento.clave == clave && hash.lista[posicion].estado == OCUPADO {
			hash.lista[posicion].estado = BORRADO
			hash.cantidad--
			hash.cantidadBorrados++
			return hash.lista[posicion].elemento.valor
		}
		posicion++
		if posicion == len(hash.lista) {
			posicion = 0
		}
	}
	panic("La clave no pertenece al diccionario")
}
func (hash hash[K, V]) Cantidad() int {
	return hash.cantidad
}
func (hash hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	posicion := 0
	for hash.lista[posicion].estado != OCUPADO {
		posicion++
	}
	claveValor := hash.lista[posicion]
	for i := 0; visitar(claveValor.elemento.clave, claveValor.elemento.valor) && posicion < len(hash.lista)-1; i++ {
		posicion++
		for posicion < len(hash.lista) && hash.lista[posicion].estado != OCUPADO {
			posicion++
		}
		if posicion == len(hash.lista) || hash.lista[posicion].estado != OCUPADO {
			return
		}
		claveValor = hash.lista[posicion]

	}
}
func (hash hash[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDicccionario[K, V])
	iter.lista = hash.lista
	pos := 0
	for pos < len(iter.lista) && iter.lista[pos].estado != OCUPADO {
		pos++
	}
	iter.iteracionActual = pos
	return iter
}

func (iter iterDicccionario[K, V]) HaySiguiente() bool {
	return iter.iteracionActual < len(iter.lista)
}
func (iter iterDicccionario[K, V]) VerActual() (K, V) {
	iter.panicIteracion()
	clave := iter.lista[iter.iteracionActual].elemento.clave
	valor := iter.lista[iter.iteracionActual].elemento.valor
	return clave, valor
}
func (iter *iterDicccionario[K, V]) Siguiente() {
	iter.panicIteracion()
	pos := iter.iteracionActual + 1
	for pos < len(iter.lista) && iter.lista[pos].estado != OCUPADO {
		pos++
	}
	iter.iteracionActual = pos
}

func (iter iterDicccionario[K, V]) panicIteracion() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (hash *hash[K, V]) redimensionar() {
	if 3*(hash.cantidad+hash.cantidadBorrados) > 2*len(hash.lista) {
		listaVieja := hash.lista
		slice := make([]nodoLista[K, V], 2*len(hash.lista))
		hash.lista = slice
		hash.copiarRedimension(listaVieja)
	} else if 5*hash.cantidad < len(hash.lista) {
		listaVieja := hash.lista
		slice := make([]nodoLista[K, V], len(hash.lista)/2)
		hash.lista = slice
		hash.copiarRedimension(listaVieja)
	}
}

func (hash *hash[K, V]) copiarRedimension(listaVieja []nodoLista[K, V]) {
	hash.cantidad = 0
	hash.cantidadBorrados = 0
	for _, campo := range listaVieja {
		if campo.estado == OCUPADO {
			clave := campo.elemento.clave
			valor := campo.elemento.valor
			hash.Guardar(clave, valor)
		}
	}
}
