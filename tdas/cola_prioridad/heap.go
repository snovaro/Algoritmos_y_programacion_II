package cola_prioridad

const (
	TAMANIO_INICIAL = 7
)

type heap[T any] struct {
	lista       *[]T
	cantidad    int
	funcion_cmp func(T, T) int
}

func iniciarHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	slice := make([]T, TAMANIO_INICIAL)
	heap.lista = &slice
	heap.funcion_cmp = funcion_cmp
	return heap
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := iniciarHeap[T](funcion_cmp)
	return heap
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heapify[T](&arreglo, funcion_cmp)
	heap := new(heap[T])
	heap.lista = &arreglo
	heap.funcion_cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	return heap
}

func heapify[T any](lista *[]T, funcion_cmp func(T, T) int) {
	ultimaPos := len(*lista) - 1
	for i := range *lista { //for i:=0; i<len(lista);i++{
		downHeap(lista, ultimaPos-i, len(*lista), funcion_cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	if len(elementos) <= 1 {
		return
	}
	heapify[T](&elementos, funcion_cmp)
	heapSortAux[T](&elementos, len(elementos), funcion_cmp)
}

func heapSortAux[T any](elementos *[]T, largo int, funcion_cmp func(T, T) int) {
	if largo == 1 {
		return
	}
	swap(elementos, 0, largo-1)
	downHeap(elementos, 0, largo-1, funcion_cmp)
	heapSortAux(elementos, largo-1, funcion_cmp)
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elemento T) {
	heap.redimensionar()
	(*heap.lista)[heap.cantidad] = elemento
	upheap(heap.lista, heap.cantidad, heap.funcion_cmp)
	heap.cantidad++
}

func (heap heap[T]) VerMax() T {
	heap.panicVacia()
	return (*heap.lista)[0]
}

func (heap *heap[T]) Desencolar() T {
	heap.panicVacia()
	heap.redimensionar()
	swap(heap.lista, 0, heap.cantidad-1)
	heap.cantidad--
	downHeap(heap.lista, 0, heap.cantidad, heap.funcion_cmp)
	return (*heap.lista)[heap.cantidad]
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap heap[T]) panicVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (heap *heap[T]) redimensionar() {
	if len(*heap.lista) >= 4*heap.cantidad || 2*len(*heap.lista) < 3*heap.cantidad {
		tamanio := TAMANIO_INICIAL
		if heap.cantidad != 0 {
			tamanio = 2 * heap.cantidad
		}
		newLista := make([]T, tamanio)
		copy(newLista, *heap.lista)
		heap.lista = &newLista
	}
}

func upheap[T any](lista *[]T, posicion int, funcion_cmp func(T, T) int) {
	if posicion == 0 {
		return
	}
	padre := (posicion - 1) / 2
	if funcion_cmp((*lista)[posicion], (*lista)[padre]) == 1 {
		swap(lista, posicion, padre)
		upheap(lista, padre, funcion_cmp)
	}
}

func swap[T any](lista *[]T, posicion1 int, posicion2 int) {
	(*lista)[posicion1], (*lista)[posicion2] = (*lista)[posicion2], (*lista)[posicion1]
}

func downHeap[T any](lista *[]T, posicion int, cantidad int, funcion_cmp func(T, T) int) {
	hijoMayor := 2*posicion + 1
	hijoDer := 2*posicion + 2
	if hijoMayor >= cantidad {
		return
	}
	if hijoDer < cantidad && funcion_cmp((*lista)[hijoDer], (*lista)[hijoMayor]) == 1 {
		hijoMayor = hijoDer
	}
	actual := (*lista)[posicion]
	hijo := (*lista)[hijoMayor]
	if funcion_cmp(actual, hijo) == -1 {
		swap(lista, posicion, hijoMayor)
		downHeap(lista, hijoMayor, cantidad, funcion_cmp)
	}
}
