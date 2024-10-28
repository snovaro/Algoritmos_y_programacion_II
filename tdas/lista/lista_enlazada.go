package lista

type listaEnlazada[T any] struct {
	longitud int
	primero  *nodoLista[T]
	ultimo   *nodoLista[T]
}

type nodoLista[T any] struct {
	dato    T
	proximo *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	iterado  int
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.longitud = 0
	lista.ultimo = nil
	lista.primero = nil
	return lista
}

func (lista *listaEnlazada[T]) crearNodo(elemento T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = elemento
	return nodo
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.longitud == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nodo := lista.crearNodo(elem)
	if !lista.insertarVacia(nodo) {
		nodo.proximo = lista.primero
		lista.primero = nodo
	}
	lista.longitud++
}

func (lista *listaEnlazada[T]) insertarVacia(nodo *nodoLista[T]) bool {
	if lista.EstaVacia() {
		lista.primero = nodo
		lista.ultimo = nodo
		return true
	}
	return false
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nodo := lista.crearNodo(elem)

	if !lista.insertarVacia(nodo) {
		lista.ultimo.proximo = nodo
		lista.ultimo = nodo
		if lista.longitud == 1 {
			lista.primero.proximo = lista.ultimo
		}
	}
	lista.longitud++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.panicEstaVacia()
	dato := lista.primero.dato

	if lista.longitud == 1 {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		*lista.primero = *lista.primero.proximo
	}
	lista.longitud--
	return dato
}

func (lista listaEnlazada[T]) VerPrimero() T {
	lista.panicEstaVacia()
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	lista.panicEstaVacia()
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) panicEstaVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.longitud
}

func (lista listaEnlazada[T]) Iterar(visitar func(elem T) bool) {
	iterado := lista.primero
	for iterado != nil && visitar(iterado.dato) {
		iterado = iterado.proximo
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorListaEnlazada[T])
	iter.actual = lista.primero
	iter.iterado = 0
	iter.lista = lista
	return iter
}

func (iter iteradorListaEnlazada[T]) VerActual() T {
	iter.panicIteracionTerminada()
	return iter.actual.dato
}

func (iterador iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	iterador.panicIteracionTerminada()
	iterador.anterior, iterador.actual = iterador.actual, iterador.actual.proximo
	iterador.iterado++
}

func (iter iteradorListaEnlazada[T]) panicIteracionTerminada() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	iterador.panicIteracionTerminada()
	dato := iterador.actual.dato
	iterador.actual = iterador.actual.proximo
	iterador.lista.longitud--
	if iterador.iterado == 0 {
		iterador.lista.primero = iterador.actual
	}
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = iterador.anterior
	}
	if iterador.anterior != nil {
		iterador.anterior.proximo = iterador.actual
	}

	return dato
}

func (iterador *iteradorListaEnlazada[T]) Insertar(elem T) {
	nodo := iterador.lista.crearNodo(elem)
	nodo.proximo = iterador.actual
	iterador.actual = nodo
	if iterador.iterado == 0 {
		iterador.lista.primero = nodo
	}
	if iterador.iterado == iterador.lista.longitud {
		iterador.lista.ultimo = nodo
	}
	if iterador.anterior != nil {
		iterador.anterior.proximo = nodo
	}
	iterador.lista.longitud++
}

func (lista *listaEnlazada[T]) InvertirLista() {
	actual := lista.primero
	var anterior *nodoLista[T]
	for actual != nil {
		siguiente := actual.proximo
		actual.proximo = anterior
		anterior = actual
		actual = siguiente
	}
	lista.primero, lista.ultimo = lista.ultimo, lista.primero
}

func (lista *listaEnlazada[T]) EliminarPosicionesPares() {
	actual := lista.primero
	siguiente := actual.proximo

	for siguiente != nil {
		lista.longitud--
		actual.proximo = siguiente.proximo
		actual = actual.proximo
		if actual != nil {
			siguiente = actual.proximo
		} else {
			siguiente = nil
		}
	}
}
