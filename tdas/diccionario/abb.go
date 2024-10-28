package diccionario

import (
	TDACola "tdas/cola"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
)

type arbolBinarioDeBusqueda[K comparable, V any] struct {
	raiz        *nodoAbb[K, V]
	cantidad    int
	funcion_cmp func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	clave    K
	valor    V
	hijo_izq *nodoAbb[K, V]
	hijo_der *nodoAbb[K, V]
}

type iteradorDiccionarioOrdenado[K comparable, V any] struct {
	cola            TDACola.Cola[*nodoAbb[K, V]]
	desde           *K
	hasta           *K
	funcion_cmp     func(K, V) bool
	funcion_cmp_abb func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	dicc := new(arbolBinarioDeBusqueda[K, V])
	dicc.funcion_cmp = funcion_cmp
	return dicc
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nuevoNodo := new(nodoAbb[K, V])
	nuevoNodo.valor = valor
	nuevoNodo.clave = clave
	return nuevoNodo
}

func (abb *arbolBinarioDeBusqueda[K, V]) buscarClave(clave K, nodo *nodoAbb[K, V], padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nodo, padre
	}
	if abb.funcion_cmp(nodo.clave, clave) == 0 {
		return nodo, padre
	} else if abb.funcion_cmp(nodo.clave, clave) < 0 {
		return abb.buscarClave(clave, nodo.hijo_der, nodo)
	}
	return abb.buscarClave(clave, nodo.hijo_izq, nodo)
}

func (abb *arbolBinarioDeBusqueda[K, V]) Guardar(clave K, valor V) {
	nodo, padre := abb.buscarClave(clave, abb.raiz, nil)
	if nodo != nil {
		nodo.valor = valor
		return
	}
	abb.cantidad++
	nuevoNodo := crearNodo[K, V](clave, valor)
	if padre == nil {
		abb.raiz = nuevoNodo
		return
	}
	if abb.funcion_cmp(padre.clave, clave) < 0 {
		padre.hijo_der = nuevoNodo
		return
	}
	padre.hijo_izq = nuevoNodo
}

func (abb arbolBinarioDeBusqueda[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarClave(clave, abb.raiz, nil)
	return nodo != nil
}

func (abb *arbolBinarioDeBusqueda[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarClave(clave, abb.raiz, nil)
	abb.panicNoPertenece(nodo)
	return nodo.valor
}

func (abb *arbolBinarioDeBusqueda[K, V]) panicNoPertenece(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

func (abb *arbolBinarioDeBusqueda[K, V]) Borrar(elem K) V {
	nodo, padre := abb.buscarClave(elem, abb.raiz, nil)
	abb.panicNoPertenece(nodo)
	abb.cantidad--
	valor := nodo.valor
	if padre == nil {
		if nodo.hijo_izq == nil && nodo.hijo_der == nil {
			abb.raiz = nil
			return valor
		}
		if abb.raiz.hijo_izq == nil {
			abb.raiz = nodo.hijo_der
			return valor
		}
		if abb.raiz.hijo_der == nil {
			abb.raiz = nodo.hijo_izq
			return valor
		}
		reemplazo, padreDelReemplazo := abb.buscarMayor(abb.raiz.hijo_izq, nil)
		abb.raiz.clave = reemplazo.clave
		abb.raiz.valor = reemplazo.valor
		if padreDelReemplazo != nil {
			padreDelReemplazo.hijo_der = reemplazo.hijo_izq
		} else {
			abb.raiz.hijo_izq = reemplazo.hijo_izq
		}
		return valor
	}
	hijoPadreNodo := &padre.hijo_der
	if abb.funcion_cmp(padre.clave, nodo.clave) > 0 {
		hijoPadreNodo = &padre.hijo_izq
	}
	if nodo.hijo_izq == nil && nodo.hijo_der == nil {
		*hijoPadreNodo = nil
		return valor
	}
	if nodo.hijo_izq == nil {
		*hijoPadreNodo = nodo.hijo_der
		return valor
	}
	if nodo.hijo_der == nil {
		*hijoPadreNodo = nodo.hijo_izq
		return valor
	}
	reemplazo, padreDelReemplazo := abb.buscarMayor(nodo.hijo_izq, nil)
	nodo.clave = reemplazo.clave
	nodo.valor = reemplazo.valor
	if padreDelReemplazo == nil {
		nodo.hijo_izq = reemplazo.hijo_izq
	} else {
		padreDelReemplazo.hijo_der = reemplazo.hijo_izq
	}
	return valor
}
func (abb arbolBinarioDeBusqueda[K, V]) buscarMayor(padre *nodoAbb[K, V], Abuelo *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if padre.hijo_der == nil {
		return padre, Abuelo
	}
	return abb.buscarMayor(padre.hijo_der, padre)
}

func (abb arbolBinarioDeBusqueda[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb arbolBinarioDeBusqueda[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorDiccionarioOrdenado[K, V])
	iter.cola = TDACola.CrearColaEnlazada[*nodoAbb[K, V]]()
	iter.encolarEnRango(abb.raiz)
	return iter
}

func (abb arbolBinarioDeBusqueda[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	cola := TDACola.CrearColaEnlazada[*nodoAbb[K, V]]()
	abb.encolarDesdeElMenor(&cola, abb.raiz)
	for !cola.EstaVacia() && visitar(cola.VerPrimero().clave, cola.VerPrimero().valor) {
		cola.Desencolar()
	}
}

func (abb arbolBinarioDeBusqueda[K, V]) encolarDesdeElMenor(cola *TDACola.Cola[*nodoAbb[K, V]], nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	abb.encolarDesdeElMenor(cola, nodo.hijo_izq)
	(*cola).Encolar(nodo)
	abb.encolarDesdeElMenor(cola, nodo.hijo_der)
}

func (abb *arbolBinarioDeBusqueda[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iter := new(iteradorDiccionarioOrdenado[K, V])
	iter.cola = TDACola.CrearColaEnlazada[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.funcion_cmp = visitar
	iter.funcion_cmp_abb = abb.funcion_cmp
	iter.encolarEnRango(abb.raiz)
	for !iter.cola.EstaVacia() && visitar(iter.VerActual()) {
		iter.cola.Desencolar()
	}
}

func (abb *arbolBinarioDeBusqueda[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorDiccionarioOrdenado[K, V])
	iter.cola = TDACola.CrearColaEnlazada[*nodoAbb[K, V]]()
	iter.desde = desde
	iter.hasta = hasta
	iter.funcion_cmp_abb = abb.funcion_cmp
	iter.encolarEnRango(abb.raiz)
	return iter
}

func (iter *iteradorDiccionarioOrdenado[K, V]) encolarEnRango(raiz *nodoAbb[K, V]) {
	if raiz == nil {
		return
	}
	iter.encolarEnRango(raiz.hijo_izq)

	if iter.desde == nil && iter.hasta == nil {
		(*iter).cola.Encolar(raiz)
	} else if iter.hasta == nil {
		if iter.funcion_cmp_abb(raiz.clave, *iter.desde) >= 0 {
			(*iter).cola.Encolar(raiz)
		}
	} else if iter.desde == nil {
		if iter.funcion_cmp_abb(raiz.clave, *iter.hasta) <= 0 {
			(*iter).cola.Encolar(raiz)
		}
	} else {
		if iter.funcion_cmp_abb(raiz.clave, *iter.hasta) <= 0 {
			if iter.funcion_cmp_abb(raiz.clave, *iter.desde) >= 0 {
				(*iter).cola.Encolar(raiz)
			}
		}
	}
	iter.encolarEnRango(raiz.hijo_der)
}

func (iter iteradorDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !iter.cola.EstaVacia()
}

func (iter *iteradorDiccionarioOrdenado[K, V]) Siguiente() {
	iter.panicIteracion()
	(*iter).cola.Desencolar()
}

func (iter iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	iter.panicIteracion()
	actual := iter.cola.VerPrimero()
	return actual.clave, actual.valor
}

func (iter iteradorDiccionarioOrdenado[K, V]) panicIteracion() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (abb *arbolBinarioDeBusqueda[K, V]) IteradorRangoNiveles(desde, hasta *K, nivelMax int) TDALista.Lista[V] {
	lista := TDALista.CrearListaEnlazada[V]()
	abb.raiz.iteradorRangoNiveles(desde, hasta, 1, nivelMax, abb.funcion_cmp, &lista)
	return lista
}

func (nodo *nodoAbb[K, V]) iteradorRangoNiveles(desde, hasta *K, nivelActual, nivelMax int, cmp func(K, K) int, lista *TDALista.Lista[V]) {
	if nodo == nil || nivelActual > nivelMax {
		return
	}
	if cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0 {
		(*lista).InsertarUltimo(nodo.valor)
	}
	if cmp(nodo.clave, *hasta) <= 0 {
		nodo.hijo_izq.iteradorRangoNiveles(desde, hasta, nivelActual+1, nivelMax, cmp, lista)
	}
	if cmp(nodo.clave, *desde) >= 0 {
		nodo.hijo_der.iteradorRangoNiveles(desde, hasta, nivelActual+1, nivelMax, cmp, lista)
	}
}

type par[K comparable, V any] struct {
	clave K
	valor V
}

func (abb *arbolBinarioDeBusqueda[K, V]) IterNivelesInversos(visitar func(K, V) bool) {
	pila := TDAPila.CrearPilaDinamica[par[K, V]]()
	abb.Iterar(func(clave K, valor V) bool {
		nuevoPar := *new(par[K, V])
		nuevoPar.clave = clave
		nuevoPar.valor = valor
		pila.Apilar(nuevoPar)
		return true
	})
	visitado := true
	for !pila.EstaVacia() && visitado {
		claveValor := pila.Desapilar()
		clave := claveValor.clave
		valor := claveValor.valor
		visitado = visitar(clave, valor)
	}
}
