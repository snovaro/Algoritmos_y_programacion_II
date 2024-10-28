package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato    T
	proximo *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.panicoEstaVacia()
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) crearNodo(elemento T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = elemento
	nodo.proximo = nil
	return nodo
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nodo := cola.crearNodo(elemento)
	if cola.EstaVacia() {
		cola.primero = nodo
	} else {
		cola.ultimo.proximo = nodo
	}
	cola.ultimo = nodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) Desencolar() T {
	cola.panicoEstaVacia()
	valor := cola.primero.dato
	cola.primero = cola.primero.proximo
	if cola.EstaVacia() {
		cola.ultimo = nil
	}
	return valor
}

func (cola colaEnlazada[T]) panicoEstaVacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}
