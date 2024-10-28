package cola_prioridad_test

import (
	TDAheap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 1000

func TestCrearArrVacioYEncolar(t *testing.T) {
	heap := TDAheap.CrearHeapArr[int]([]int{}, func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.True(t, heap.EstaVacia())
	heap.Encolar(3)
	heap.Encolar(120)
	require.EqualValues(t, 120, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 120, heap.Desencolar())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())

}

func TestCrearArrEncoloYDesencolo(t *testing.T) {
	arreglo := []int{6, 7, 1, 10, -5, 7, 100, 4}
	heap := TDAheap.CrearHeapArr[int](arreglo, func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
	heap.Encolar(5)
	heap.Encolar(300)
	require.EqualValues(t, 300, heap.VerMax())
	require.EqualValues(t, 300, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 6, heap.VerMax())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 4, heap.VerMax())
	require.EqualValues(t, 4, heap.Desencolar())
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Desencolar())
	require.EqualValues(t, -5, heap.VerMax())
	require.EqualValues(t, -5, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	heap.Encolar(1000)
	heap.Encolar(5)
	heap.Encolar(100)
	require.EqualValues(t, 1000, heap.VerMax())
	require.EqualValues(t, 1000, heap.Desencolar())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
}
func TestCrearHeapArrPocosElems(t *testing.T) {
	arreglo := []int{10, -5, 100}
	heap := TDAheap.CrearHeapArr[int](arreglo, func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, -5, heap.VerMax())
	require.EqualValues(t, -5, heap.Desencolar())
}
func TestCrearHeapArrGeneral(t *testing.T) {
	arreglo := []int{6, 7, 1, 10, -5, 7, 100}
	heap := TDAheap.CrearHeapArr[int](arreglo, func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.False(t, heap.EstaVacia())
	heap.Encolar(100)
	heap.Encolar(5)
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	heap.Encolar(8)
	heap.Encolar(0)
	heap.Encolar(-1000)
	require.EqualValues(t, 8, heap.VerMax())
	require.EqualValues(t, 8, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 6, heap.VerMax())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Desencolar())
	require.EqualValues(t, 0, heap.VerMax())
	require.EqualValues(t, 0, heap.Desencolar())
	require.EqualValues(t, -5, heap.VerMax())
	require.EqualValues(t, -5, heap.Desencolar())
	require.EqualValues(t, -1000, heap.VerMax())
	require.EqualValues(t, -1000, heap.Desencolar())
}
func TestHeapsort(t *testing.T) {
	arreglo := []int{6, 7, 1, 10, -5, 7, 100, 4}
	arregloOrdenado := []int{-5, 1, 4, 6, 7, 7, 10, 100}
	TDAheap.HeapSort[int](arreglo, func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.EqualValues(t, arregloOrdenado, arreglo)
}

func TestHeapVacio(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapEnteroUnico(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	heap.Encolar(5)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEnterosVarios(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	heap.Encolar(-9)
	heap.Encolar(7)
	heap.Encolar(30)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 30, heap.VerMax())
	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, -9, heap.VerMax())
	require.EqualValues(t, -9, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapFloatUnico(t *testing.T) {
	heap := TDAheap.CrearHeap[float32](func(V1, V2 float32) int { return funcionCmpfloat(V1, V2) })
	heap.Encolar(3.1)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 3.1, heap.VerMax())
	require.EqualValues(t, 3.1, heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func TestFloatVarios(t *testing.T) {
	heap := TDAheap.CrearHeap[float32](func(V1, V2 float32) int { return funcionCmpfloat(V1, V2) })
	heap.Encolar(9.9)
	heap.Encolar(0.0)
	heap.Encolar(3.1)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 9.9, heap.VerMax())
	require.EqualValues(t, 9.9, heap.Desencolar())
	require.EqualValues(t, 3.1, heap.VerMax())
	require.EqualValues(t, 3.1, heap.Desencolar())
	require.EqualValues(t, 0.0, heap.VerMax())
	require.EqualValues(t, 0.0, heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func TestHeapStringUnico(t *testing.T) {
	heap := TDAheap.CrearHeap[string](func(V1, V2 string) int { return funcionCmpstring(V1, V2) })
	heap.Encolar("Hola Mundo")
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Hola Mundo", heap.VerMax())
	require.EqualValues(t, "Hola Mundo", heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func TestHeapStringVarios(t *testing.T) {
	heap := TDAheap.CrearHeap[string](func(V1, V2 string) int { return funcionCmpstring(V1, V2) })
	heap.Encolar("Mundo")
	heap.Encolar("Dias")
	heap.Encolar("Buenos")
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, "Buenos", heap.VerMax())
	require.EqualValues(t, "Buenos", heap.Desencolar())
	require.EqualValues(t, "Mundo", heap.VerMax())
	require.EqualValues(t, "Mundo", heap.Desencolar())
	require.EqualValues(t, "Dias", heap.VerMax())
	require.EqualValues(t, "Dias", heap.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func TestVariosEncoloDesencoloDesordenados(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	heap.Encolar(30)
	heap.Encolar(7)
	heap.Encolar(-9)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 30, heap.VerMax())
	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 7, heap.Desencolar())
	heap.Encolar(12)
	require.EqualValues(t, 12, heap.VerMax())
	require.EqualValues(t, 12, heap.Desencolar())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, -9, heap.VerMax())
	require.EqualValues(t, -9, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncoloVacioEncolo(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	heap.Encolar(-109)
	heap.Encolar(67)
	heap.Encolar(3)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 67, heap.VerMax())
	require.EqualValues(t, 67, heap.Desencolar())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, -109, heap.VerMax())
	require.EqualValues(t, -109, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	heap.Encolar(321)
	heap.Encolar(12)
	require.EqualValues(t, 321, heap.VerMax())
	require.EqualValues(t, 321, heap.Desencolar())
	require.EqualValues(t, 12, heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 12, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestCantidad(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	require.EqualValues(t, heap.Cantidad(), 0)
	heap.Encolar(5)
	require.EqualValues(t, heap.Cantidad(), 1)
	heap.Encolar(-20)
	require.EqualValues(t, heap.Cantidad(), 2)
	heap.Desencolar()
	require.EqualValues(t, heap.Cantidad(), 1)
	heap.Desencolar()
	require.EqualValues(t, heap.Cantidad(), 0)
}

func TestColaVolumen(t *testing.T) {
	heap := TDAheap.CrearHeap[int](func(V1, V2 int) int { return funcionCmpint(V1, V2) })
	for i := 0; i <= VOLUMEN; i++ {
		heap.Encolar(i)
	}
	for i := VOLUMEN; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func funcionCmpint(elemento1, elemento2 int) int {
	if elemento1 < elemento2 {
		return -1
	}
	if elemento1 > elemento2 {
		return 1
	}
	return 0
}

func funcionCmpfloat(elemento1, elemento2 float32) int {
	if elemento1 < elemento2 {
		return -1
	}
	if elemento1 > elemento2 {
		return 1
	}
	return 0
}

func funcionCmpstring(a, b string) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	return 0
}
