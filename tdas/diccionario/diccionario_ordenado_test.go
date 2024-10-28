package diccionario_test

import (
	TDADiccionario "tdas/diccionario"
	"testing"

	"fmt"

	"github.com/stretchr/testify/require"
)

var TAM_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpstring)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDicOrdenadoUnicaClave(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](funcionCmpintMaximos)
	dic.Guardar(1, "A")
	require.True(t, dic.Pertenece(1))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "A", dic.Obtener(1))
	require.EqualValues(t, "A", dic.Borrar(1))
}

func TestDicOrdenadoVariosElems(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpstring)
	dic.Guardar("Aw", 1)
	dic.Guardar("Bqq", 2)
	dic.Guardar("C", 3)
	dic.Guardar("Dewww", 4)
	dic.Guardar("Eqweqweq", 5)
	dic.Guardar("Aw", 6)
	require.EqualValues(t, 5, dic.Cantidad())
	require.True(t, dic.Pertenece("Aw"))
	require.True(t, dic.Pertenece("Eqweqweq"))
	require.True(t, dic.Pertenece("Dewww"))
	require.True(t, dic.Pertenece("C"))
	require.True(t, dic.Pertenece("Bqq"))
	require.EqualValues(t, 4, dic.Obtener("Dewww"))
	require.EqualValues(t, 2, dic.Obtener("Bqq"))
	require.EqualValues(t, 6, dic.Obtener("Aw"))
	require.EqualValues(t, 4, dic.Obtener("Dewww"))
	require.EqualValues(t, 3, dic.Obtener("C"))
	require.EqualValues(t, 5, dic.Obtener("Eqweqweq"))
	require.EqualValues(t, 5, dic.Cantidad())
}

func TestDicOrdenadoBorrarUnico(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMinimos)
	dic.Guardar(10, 10000)
	require.EqualValues(t, 10000, dic.Borrar(10))
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(0) })
}

func TestDicOrdenadoBorrarVarios(t *testing.T) {
	dic := TDADiccionario.CrearABB[float32, string](funcionCmpfloat)
	dic.Guardar(3.14, "A")
	dic.Guardar(1.41, "C")
	dic.Guardar(2.42, "D")
	dic.Guardar(0.53, "B")
	dic.Guardar(0.86, "E")
	require.EqualValues(t, "A", dic.Borrar(3.14))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3.14) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(3.14) })
	require.EqualValues(t, "B", dic.Borrar(0.53))
	require.EqualValues(t, "C", dic.Borrar(1.41))
	require.EqualValues(t, "D", dic.Borrar(2.42))
	require.EqualValues(t, "E", dic.Borrar(0.86))
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(0.86) })
}

func TestBorrarRaizDosHijos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMinimos)
	dic.Guardar(5, 5)
	dic.Guardar(2, 10)
	dic.Guardar(7, 9)
	require.EqualValues(t, 5, dic.Borrar(5))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(5) })
	dic.Guardar(5, 5)
	dic.Guardar(1, 1)
	dic.Guardar(2, 6)
	require.EqualValues(t, 9, dic.Borrar(7))
	require.EqualValues(t, 5, dic.Borrar(5))
	require.EqualValues(t, 6, dic.Borrar(2))
}

func TestBorroRaizConReemplazoProfundo(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	dic.Guardar(100, 1)
	dic.Guardar(200, 2)
	dic.Guardar(-100, 3)
	for i := -99; i < 100; i++ {
		dic.Guardar(i, 0)
	}
	require.EqualValues(t, 1, dic.Borrar(100))
	require.EqualValues(t, 201, dic.Cantidad())
	for i := 99; i > -100; i-- {
		require.EqualValues(t, 0, dic.Borrar(i))
	}

}

func TestBorrarYGuardarEnElMedio(t *testing.T) {
	dic := TDADiccionario.CrearABB[float32, string](funcionCmpfloat)
	dic.Guardar(3.5, "A")
	dic.Guardar(123, "B")
	dic.Guardar(-5, "C")
	dic.Guardar(-12.3, "D")
	dic.Guardar(0, "E")
	dic.Guardar(4, "F")
	dic.Guardar(4.5, "G")
	dic.Guardar(3.99, "H")
	dic.Guardar(3.6, "I")
	dic.Guardar(3.4, "J")
	dic.Guardar(100.3, "K")
	require.EqualValues(t, "F", dic.Borrar(4))
	require.EqualValues(t, "I", dic.Borrar(3.6))
	require.True(t, dic.Pertenece(3.4))
	require.False(t, dic.Pertenece(3.6))
	require.False(t, dic.Pertenece(4))
	require.EqualValues(t, "A", dic.Borrar(3.5))
	require.EqualValues(t, "B", dic.Borrar(123))
	require.EqualValues(t, "C", dic.Borrar(-5))
	require.False(t, dic.Pertenece(3.5))
	require.False(t, dic.Pertenece(123))
	require.False(t, dic.Pertenece(-5))
	dic.Guardar(3.5, "Y")
	dic.Guardar(123, "Z")
	dic.Guardar(-5, "X")
	require.EqualValues(t, "K", dic.Borrar(100.3))
	require.EqualValues(t, "H", dic.Borrar(3.99))
	require.EqualValues(t, "J", dic.Borrar(3.4))
	require.EqualValues(t, "Y", dic.Borrar(3.5))
	require.EqualValues(t, "G", dic.Borrar(4.5))
	require.EqualValues(t, "X", dic.Borrar(-5))
	require.EqualValues(t, "Z", dic.Borrar(123))
	require.EqualValues(t, "D", dic.Borrar(-12.3))
	require.EqualValues(t, "E", dic.Borrar(0))
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestBorrarNodo2HijosSinNietos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMinimos)
	dic.Guardar(3, 11)
	dic.Guardar(1, 1)
	dic.Guardar(4, 100)
	require.EqualValues(t, 11, dic.Borrar(3))
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })
	require.EqualValues(t, 1, dic.Borrar(1))
	require.EqualValues(t, 100, dic.Borrar(4))
}

func TestBorrarNodoReemplazoConUnHijoYVaciar(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMinimos)
	dic.Guardar(3, 11)
	dic.Guardar(2, 10)
	dic.Guardar(1, 1)
	dic.Guardar(4, 100)
	dic.Guardar(5, 101)
	require.EqualValues(t, 11, dic.Borrar(3))
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })
	require.EqualValues(t, 1, dic.Borrar(1))
	require.EqualValues(t, 100, dic.Borrar(4))
	require.EqualValues(t, 101, dic.Borrar(5))
	require.EqualValues(t, 10, dic.Borrar(2))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(4) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(5) })
}

func TestBorrarConReemplazoProfundoUnicaRama(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, []int](funcionCmpintMaximos)
	dic.Guardar(21, []int{01, 06, 2022})
	dic.Guardar(-5, []int{12, 12, 12})
	dic.Guardar(0, []int{0, 0, 0})
	dic.Guardar(1, []int{1, 2, 3})
	dic.Guardar(2, []int{4, 5, 6})
	dic.Guardar(3, []int{18, 12, 2022})
	dic.Guardar(5, []int{7})
	dic.Guardar(10, []int{24, 06, 1987})
	dic.Guardar(11, []int{14, 02, 1988})
	dic.Guardar(15, []int{10, 07, 2021})
	dic.Guardar(16, []int{02, 06, 1988})
	require.EqualValues(t, []int{01, 06, 2022}, dic.Borrar(21))
	require.True(t, dic.Pertenece(0))
	require.True(t, dic.Pertenece(1))
	require.True(t, dic.Pertenece(2))
	require.True(t, dic.Pertenece(3))
	require.True(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(-5))
	require.True(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(11))
	require.True(t, dic.Pertenece(16))
	require.False(t, dic.Pertenece(21))
	require.EqualValues(t, []int{0, 0, 0}, dic.Borrar(0))
	require.EqualValues(t, []int{1, 2, 3}, dic.Borrar(1))
	require.EqualValues(t, []int{4, 5, 6}, dic.Borrar(2))
	require.EqualValues(t, []int{7}, dic.Borrar(5))
	require.EqualValues(t, []int{12, 12, 12}, dic.Borrar(-5))
	require.EqualValues(t, []int{18, 12, 2022}, dic.Borrar(3))
	require.EqualValues(t, []int{24, 06, 1987}, dic.Borrar(10))
	require.EqualValues(t, []int{10, 07, 2021}, dic.Borrar(15))
	require.EqualValues(t, []int{14, 02, 1988}, dic.Borrar(11))
	require.EqualValues(t, []int{02, 06, 1988}, dic.Borrar(16))
	require.False(t, dic.Pertenece(0))
	require.False(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(2))
	require.False(t, dic.Pertenece(3))
	require.False(t, dic.Pertenece(5))
	require.False(t, dic.Pertenece(-5))
	require.False(t, dic.Pertenece(10))
	require.False(t, dic.Pertenece(15))
	require.False(t, dic.Pertenece(11))
	require.False(t, dic.Pertenece(16))
}
func TestIteradorInternoVarios(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, []int](funcionCmpintMinimos)
	slices := [][]int{{01, 06, 2022}, {02, 06, 1988}, {14, 02, 1988}, {10, 07, 2021}, {24, 06, 1987}, {7}, {18, 12, 2022}, {12, 12, 12}, {4, 5, 6}, {1, 2, 3}, {0, 0, 0}}
	claves := []int{21, 16, 15, 11, 10, 5, 3, 2, 1, 0, -5}
	recorridos := 0
	dic.Guardar(claves[0], slices[0])
	dic.Guardar(claves[9], slices[9])
	dic.Guardar(claves[8], slices[8])
	dic.Guardar(claves[7], slices[7])
	dic.Guardar(claves[10], slices[10])
	dic.Guardar(claves[6], slices[6])
	dic.Guardar(claves[5], slices[5])
	dic.Guardar(claves[4], slices[4])
	dic.Guardar(claves[2], slices[2])
	dic.Guardar(claves[3], slices[3])
	dic.Guardar(claves[1], slices[1])
	iterado := 0
	require.EqualValues(t, dic.Cantidad(), len(slices))
	dic.Iterar(func(K int, V []int) bool {
		for i := range V {
			require.EqualValues(t, slices[iterado][i], V[i])
		}
		require.EqualValues(t, claves[iterado], K)
		iterado++
		recorridos++
		return true
	})
	require.EqualValues(t, recorridos, dic.Cantidad())
}

var INTERVALOSDERAMAS = [][]int{{100, 200}, {50, 100}, {-50, 0}, {-200, -100}}

func TestBorrarABBConGranVolumen(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	dic.Guardar(0, 0)
	for _, intervalo := range INTERVALOSDERAMAS {
		inicio := intervalo[0]
		final := intervalo[1]
		for i := inicio; i < final; i++ {
			dic.Guardar(i*5, 5)
			dic.Guardar(i*5-3, 2)
			dic.Guardar(i*5-1, 4)
			dic.Guardar(i*5-2, 3)
			dic.Guardar(i*5-4, 1)
		}
	}
	claves := []int{}
	valores := []int{}
	dic.Iterar(func(K, V int) bool {
		claves = append(claves, K)
		valores = append(valores, V)
		return true
	})
	for i, clave := range claves {
		require.EqualValues(t, valores[i], dic.Borrar(clave))
	}
}

func TestIterarFueraDeRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpstring)
	clave1, clave2 := "Clave que no existe", "clave que tampoco existe"
	dic.IterarRango(&clave1, &clave2, func(clave string, dato int) bool {
		panic("No deberia iterar nada")
	})
}

func TestIterarHastaFueraDeRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpstring)
	clave1, clave2 := "Clave que existe", "clave que tampoco existe"
	dic.Guardar(clave1, 10)
	dic.IterarRango(&clave1, &clave2, func(clave string, dato int) bool {
		if clave != "Clave que existe" {
			panic("No deberia iterar nada")
		}
		require.EqualValues(t, 10, dic.Obtener(clave))
		require.EqualValues(t, 10, dic.Borrar(clave))
		require.False(t, dic.Pertenece(clave1))
		return true
	})
}

func TestIterarDesdeFueraDeRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpstring)
	clave1, clave2 := "Clave que existe", "clave que tampoco existe"
	dic.Guardar(clave1, 10)
	dic.IterarRango(&clave2, &clave1, func(clave string, dato int) bool {
		if clave != "Clave que existe" {
			panic("No deberia iterar nada")
		}
		require.EqualValues(t, 10, dic.Obtener(clave))
		require.EqualValues(t, 10, dic.Borrar(clave))
		require.False(t, dic.Pertenece(clave1))
		return true
	})
}

func TestIterarABBInternoCombinaciones(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	for i := 0; i <= 7; i++ {
		if i != 6 {
			dic.Guardar(i, i)
		}
	}
	dic.Guardar(6, 6)
	dic.Guardar(-1, 6)
	i := -1
	dic.Iterar(func(clave, dato int) bool {
		require.EqualValues(t, clave, i)
		i++
		return true
	})
}

func TestIteradorABBVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorABBUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](funcionCmpintMaximos)
	dic.Guardar(10, "Messi")
	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())
	clave, dato := iter.VerActual()
	require.EqualValues(t, 10, clave)
	require.EqualValues(t, "Messi", dato)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorABBVarios(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](funcionCmpintMinimos)
	hash := TDADiccionario.CrearHash[int, string]()
	dic.Guardar(10, "Messi")
	dic.Guardar(23, "Dibu")
	dic.Guardar(13, "Cuti")
	dic.Guardar(5, "Enzo")
	dic.Guardar(11, "Angelito")
	dic.Guardar(22, "Lautaro")
	dic.Guardar(9, "Julian")
	dic.Guardar(3, "Taglia")
	dic.Guardar(21, "Dybala")
	dic.Iterar(func(K int, V string) bool {
		hash.Guardar(K, V)
		return true
	})
	recorridos := 0
	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		recorridos++
		clave, actual := iter.VerActual()
		expected := hash.Obtener(clave)
		require.EqualValues(t, expected, actual)
	}
	require.EqualValues(t, dic.Cantidad(), recorridos)
}

func TestIteradorOrdenadoNoLlegaAlFinal(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](funcionCmpstring)
	hash := TDADiccionario.CrearHash[string, string]()
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")
	dic.Iterar(func(K, V string) bool {
		hash.Guardar(K, V)
		return true
	})
	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, valorPrimero := iter3.VerActual()
	iter3.Siguiente()
	segundo, valorSegundo := iter3.VerActual()
	iter3.Siguiente()
	tercero, valorTercero := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.True(t, hash.Pertenece(primero))
	require.True(t, hash.Pertenece(segundo))
	require.True(t, hash.Pertenece(tercero))
	require.EqualValues(t, hash.Obtener(primero), valorPrimero)
	require.EqualValues(t, hash.Obtener(segundo), valorSegundo)
	require.EqualValues(t, hash.Obtener(tercero), valorTercero)
}

func funcionCmpintMaximos(elemento1, elemento2 int) int {
	if elemento1 < elemento2 {
		return -1
	}
	if elemento1 > elemento2 {
		return 1
	}
	return 0
}

func funcionCmpintMinimos(elemento1, elemento2 int) int {
	if elemento1 < elemento2 {
		return 1
	}
	if elemento1 > elemento2 {
		return -1
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
	if a == b {
		return 0
	}
	if len(a) > len(b) {
		return 1
	}
	return -1
}

func TestIterarRangoNiveles(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	abb.Guardar(10, 10)
	abb.Guardar(5, 5)
	abb.Guardar(15, 15)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)
	abb.Guardar(7, 7)
	abb.Guardar(12, 12)
	abb.Guardar(20, 20)
	abb.Guardar(14, 14)
	clave1 := 5
	clave2 := 15
	lista := abb.IteradorRangoNiveles(&clave1, &clave2, 3)
	lista.Iterar(func(valor int) bool {
		fmt.Println(valor)
		return true
	})
}

func TestIterarInverso(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](funcionCmpintMaximos)
	abb.Guardar(10, 10)
	abb.Guardar(5, 5)
	abb.Guardar(15, 15)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)
	abb.Guardar(7, 7)
	abb.Guardar(12, 12)
	abb.Guardar(20, 20)
	abb.Guardar(14, 14)
	abb.IterNivelesInversos(func(clave int, valor int) bool {
		fmt.Println(clave, valor)
		return true
	})
}
