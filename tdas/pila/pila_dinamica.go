package pila

const TAMANIO_PILA = 10
const CRITERIO_ACHICO = 4         //Cuantas veces mas grande debe ser la capacidad respecto a la cantidad
const MULTIPLICADOR_DIMENSION = 2 //Multiplos de que numero van a ser los tamanios posibles de la pila al redimensionar

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, TAMANIO_PILA)
	pila.cantidad = 0
	return pila
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	pila.validarEstado()
	return pila.datos[pila.cantidad-1]
}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimension()
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {
	pila.validarEstado()
	pila.cantidad--
	if pila.cantidad*CRITERIO_ACHICO <= cap(pila.datos) && pila.cantidad > 0 {
		pila.redimension()
	}
	return pila.datos[pila.cantidad]
}

func (pila *pilaDinamica[T]) redimension() {
	copia := pila.datos
	pila.datos = make([]T, pila.cantidad*MULTIPLICADOR_DIMENSION)
	copy(pila.datos, copia)
}

func (pila *pilaDinamica[T]) validarEstado() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}
