package lista2

type listaDobleEnlazada[T any] struct {
	longitud int
	primero  *nodoLista[T]
	ultimo   *nodoLista[T]
}

type nodoLista[T any] struct {
	dato    T
	proximo *nodoLista[T]
	anterior *nodoLista[T]
}

type iteradorListaDobleEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	iterado  int
	lista    *listaDobleEnlazada[T]
}

func CrearListaDobleEnlazada[T any]() Lista[T] {
	lista := new(listaDobleEnlazada[T])
	lista.longitud = 0
	lista.ultimo = nil
	lista.primero = nil
	return lista
}

func (lista *listaDobleEnlazada[T]) crearNodo(elemento T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = elemento
	return nodo
}

func (lista listaDobleEnlazada[T]) EstaVacia() bool {
	return lista.longitud == 0
}

func (lista *listaDobleEnlazada[T]) InsertarPrimero(elem T) {
	nodo := lista.crearNodo(elem)
	if !lista.insertarVacia(nodo) {
		nodo.proximo = lista.primero
		lista.primero = nodo
	}
	lista.longitud++
}

func (lista *listaDobleEnlazada[T]) insertarVacia(nodo *nodoLista[T]) bool {
	if lista.EstaVacia() {
		lista.primero = nodo
		lista.ultimo = nodo
		return true
	}
	return false
}

func (lista *listaDobleEnlazada[T]) InsertarUltimo(elem T) {
	nodo := lista.crearNodo(elem)
	if !lista.insertarVacia(nodo) {
		nodo.anterior = lista.ultimo
		lista.ultimo.proximo = nodo
		lista.ultimo = nodo
		if lista.longitud == 1 {
			lista.primero.proximo = lista.ultimo
		}
	}
	lista.longitud++
}

func (lista *listaDobleEnlazada[T]) BorrarPrimero() T {
	lista.panicEstaVacia()
	dato := lista.primero.dato

	if lista.longitud == 1 {
		lista.primero = nil
		lista.ultimo = nil
	} else {
		*lista.primero = *lista.primero.proximo
		lista.primero.anterior = nil
	}
	lista.longitud--
	return dato
}

func (lista listaDobleEnlazada[T]) VerPrimero() T {
	lista.panicEstaVacia()
	return lista.primero.dato
}

func (lista listaDobleEnlazada[T]) VerUltimo() T {
	lista.panicEstaVacia()
	return lista.ultimo.dato
}

func (lista listaDobleEnlazada[T]) panicEstaVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista listaDobleEnlazada[T]) Largo() int {
	return lista.longitud
}

func (lista listaDobleEnlazada[T]) Iterar(visitar func(elem T) bool) {
	iterado := lista.primero
	for iterado != nil && visitar(iterado.dato) {
		iterado = iterado.proximo
	}
}

func (lista *listaDobleEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorListaDobleEnlazada[T])
	iter.actual = lista.primero
	iter.iterado = 0
	iter.lista = lista
	return iter
}

func (iter iteradorListaDobleEnlazada[T]) VerActual() T {
	iter.panicIteracionTerminada()
	return iter.actual.dato
}

func (iterador iteradorListaDobleEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaDobleEnlazada[T]) Siguiente() {
	iterador.panicIteracionTerminada()
	iterador.anterior, iterador.actual = iterador.actual, iterador.actual.proximo
	iterador.iterado++
}

func (iter iteradorListaDobleEnlazada[T]) panicIteracionTerminada() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iteradorListaDobleEnlazada[T]) Borrar() T {
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

func (iterador *iteradorListaDobleEnlazada[T]) Insertar(elem T) {
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

func (lista *listaDobleEnlazada[T]) InvertirLista(){
	actual := lista.primero
	var anterior *nodoLista[T]
	for actual != nil{
		siguiente := actual.proximo
		actual.proximo = anterior
		actual.anterior = siguiente
		anterior = actual
		actual = siguiente
	}
	lista.primero, lista.ultimo = lista.ultimo, lista.primero
}