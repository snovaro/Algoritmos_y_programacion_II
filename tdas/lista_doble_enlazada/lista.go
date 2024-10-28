package lista2

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos enlistados, false en caso contrario.
	EstaVacia() bool
	// InsertarPrimero agrega un nuevo elemento a la lista, al principio de la misma.
	InsertarPrimero(T)
	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)
	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
	BorrarPrimero() T
	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T
	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T
	// Largo devuelve la cantidad de elementos que hay en la lista.
	Largo() int
	// Iterar envia cada elemento de la lista a la funcion que recibe por parametro mientras esta devuelva true.
	Iterar(visitar func(T) bool)
	// Iterador crea una nueva variable que depende de la lista a la cual se le pueden realizar
	// todas las primitivas de IteradorLista
	Iterador() IteradorLista[T]

	InvertirLista()
}

type IteradorLista[T any] interface {
	// VerActual devuelve el dato del elemento al que está apuntando el iterandor
	VerActual() T
	// HaySiguiente devuelve true si puedo avanzar a un siguiente elemento de la lista
	// Y devuelve false si no hay siguiente elemento
	HaySiguiente() bool
	// Siguiente itera al siguiente elemento de la lista
	Siguiente()
	// Insertar permite meter un elemento entre el anterior y el actual donhde se esté apuntando
	Insertar(T)
	// Borrar elimina el elemento actualmente apuntado y lo devuelve
	Borrar() T
}
