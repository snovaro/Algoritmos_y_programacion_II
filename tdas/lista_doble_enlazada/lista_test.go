package lista2_test

import (
	"fmt"
	TDAlista "tdas/lista_doble_enlazada"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 1000
const VOLUMEN_LISTAS_DE_LISTAS = 100

func TestVacio(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestListaIntUnico(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(5)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, lista.BorrarPrimero(), 5)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestListaStrUnico(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[string]()
	lista.InsertarPrimero("Argentina Campeon")
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, "Argentina Campeon", lista.VerPrimero())
	require.EqualValues(t, "Argentina Campeon", lista.VerUltimo())
	require.EqualValues(t, lista.BorrarPrimero(), "Argentina Campeon")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestListaIntInsertarPrimeros(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(-10)
	lista.InsertarPrimero(7)
	require.EqualValues(t, 7, lista.VerPrimero())
	require.EqualValues(t, -10, lista.VerUltimo())
	require.EqualValues(t, lista.BorrarPrimero(), 7)
	lista.InsertarPrimero(66)
	require.EqualValues(t, lista.BorrarPrimero(), 66)
	require.EqualValues(t, lista.BorrarPrimero(), -10)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestListaStrInsertarPrimeros(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[string]()
	lista.InsertarPrimero("Argentina campeon")
	lista.InsertarPrimero("Francia segundo")
	require.EqualValues(t, "Francia segundo", lista.VerPrimero())
	require.EqualValues(t, "Argentina campeon", lista.VerUltimo())
	require.EqualValues(t, lista.BorrarPrimero(), "Francia segundo")
	lista.InsertarPrimero("Brasil segundo")
	require.EqualValues(t, lista.BorrarPrimero(), "Brasil segundo")
	require.EqualValues(t, lista.BorrarPrimero(), "Argentina campeon")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestListaIntInsertarMixtos(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarUltimo(-10)
	lista.InsertarPrimero(7)
	lista.InsertarUltimo(66)
	require.EqualValues(t, 7, lista.VerPrimero())
	require.EqualValues(t, 66, lista.VerUltimo())
	require.EqualValues(t, lista.BorrarPrimero(), 7)
	require.EqualValues(t, lista.BorrarPrimero(), -10)
	require.EqualValues(t, lista.BorrarPrimero(), 66)
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestListaIntInsertarUltimos(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarUltimo(-10)
	lista.InsertarUltimo(7)
	require.EqualValues(t, -10, lista.VerPrimero())
	require.EqualValues(t, 7, lista.VerUltimo())
	require.EqualValues(t, -10, lista.BorrarPrimero())
	lista.InsertarUltimo(66)
	require.EqualValues(t, 7, lista.BorrarPrimero())
	require.EqualValues(t, 66, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestVolumenInsertarPrimeros(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	for i := 0; i <= VOLUMEN; i++ {
		lista.InsertarPrimero(i)
	}
	for i := VOLUMEN; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
}

func TestVolumenInsertarUltimos(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	for i := 0; i <= VOLUMEN; i++ {
		lista.InsertarUltimo(i)
	}
	for i := 0; i <= VOLUMEN; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
}

func TestListaDeListas(t *testing.T) {
	listaMayor := TDAlista.CrearListaDobleEnlazada[TDAlista.Lista[int]]()
	for j := 0; j <= VOLUMEN_LISTAS_DE_LISTAS; j++ {
		lista := TDAlista.CrearListaDobleEnlazada[int]()
		for i := 0; i <= VOLUMEN; i++ {
			lista.InsertarUltimo(i + j)
		}
		listaMayor.InsertarUltimo(lista)
	}
	for j := 0; j <= VOLUMEN_LISTAS_DE_LISTAS; j++ {
		lista := listaMayor.BorrarPrimero()
		for i := 0; i <= VOLUMEN; i++ {
			require.EqualValues(t, i+j, lista.VerPrimero())
			require.EqualValues(t, i+j, lista.BorrarPrimero())
		}
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaMayor.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaMayor.BorrarPrimero() })
	require.True(t, listaMayor.EstaVacia())
}

func TestIterarSiempreFalse(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.Iterar(func(v int) bool {
		fmt.Println(v)
		return true
	})
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 4, lista.VerUltimo())
}

func TestIterarPrimeroVerdadero(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(10)
	dato := true
	lista.Iterar(func(v int) bool {
		if v == 10 {
			dato = false
		}
		return dato
	})
	require.False(t, dato)
}

func TestIterarTexto(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[string]()
	palabra := "Hola Mundo"
	palabraLista := ""
	for _, letra := range palabra {
		lista.InsertarUltimo(string(letra))
	}
	lista.Iterar(func(v string) bool {
		palabraLista += v
		return !(palabra == palabraLista)
	})
	require.EqualValues(t, palabra, palabraLista)
}

func TestIterador(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(0)
	iter := lista.Iterador()
	for i := 0; iter.HaySiguiente(); i++ {
		require.EqualValues(t, i, iter.VerActual())
		iter.Siguiente()
	}
}

func TestIteradorInsertarPrimero(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(7)
	iter := lista.Iterador()
	iter.Siguiente()
	for i := 0; i < 5; i++ {
		iter.Insertar(i)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.EqualValues(t, 7, lista.BorrarPrimero())
	for i := 0; i < 5; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
}

func TestIteradorInsertarUltimo(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(100)
	iter := lista.Iterador()
	iter.Insertar(200)
	require.EqualValues(t, 200, lista.VerPrimero())
	require.EqualValues(t, 200, iter.VerActual())
}

func TestIteradorInsertarMedio(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(30)
	iter := lista.Iterador()
	require.EqualValues(t, 10, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 30, iter.VerActual())
	iter.Insertar(20)
	require.EqualValues(t, 20, iter.VerActual())
	iter.Siguiente()
	require.EqualValues(t, 30, iter.VerActual())
	require.EqualValues(t, 10, lista.BorrarPrimero())
	require.EqualValues(t, 20, lista.VerPrimero())
}
func TestIteradorBorraPrimero(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(13)
	lista.InsertarUltimo(20)
	iter := lista.Iterador()
	require.EqualValues(t, 10, iter.Borrar())
	require.EqualValues(t, lista.VerPrimero(), iter.VerActual())
}
func TestIteradorBorrarMedio(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(13)
	lista.InsertarUltimo(20)
	iter := lista.Iterador()
	iter.Siguiente()
	require.EqualValues(t, 13, iter.Borrar())
	require.EqualValues(t, 20, iter.VerActual())
	require.EqualValues(t, 10, lista.BorrarPrimero())
	require.EqualValues(t, 20, lista.VerPrimero())
}

func TestLongitudes(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(15)
	lista.InsertarUltimo(-5)
	require.EqualValues(t, 2, lista.Largo())
	iter := lista.Iterador()
	iter.Insertar(10)
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 10, iter.Borrar())
	require.EqualValues(t, 15, iter.VerActual())
	require.EqualValues(t, 2, lista.Largo())
}

func TestIteradorBorraAlFinal(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.EqualValues(t, 5, iter.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.EqualValues(t, 4, lista.VerUltimo())
}

func TestIteradorInsertarYBorrarEnListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(4)
	iter.Insertar(3)
	iter.Insertar(2)
	iter.Insertar(1)
	iter.Siguiente()
	require.EqualValues(t, 2, iter.Borrar())
	iter.Siguiente()
	require.EqualValues(t, 4, iter.Borrar())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 3, lista.VerPrimero())
}

func TestIteradorInsertoListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(10)
	require.EqualValues(t, 10, lista.VerUltimo())
	require.EqualValues(t, 10, lista.VerPrimero())
}

func TestIterInsertarAlFinalListaNoVacia(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(0)
	iter.Siguiente()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(1300)
	lista.InsertarUltimo(13200)
	lista.InsertarUltimo(100)
	lista.InsertarUltimo(101)
	lista.InsertarUltimo(1010)
	lista.InsertarUltimo(10000)
	iter2 := lista.Iterador()
	require.EqualValues(t, 0, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 1, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 10, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 1300, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 13200, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 100, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 101, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 1010, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
	require.EqualValues(t, 10000, iter2.VerActual())
	require.EqualValues(t, iter2.VerActual(), lista.BorrarPrimero())
}

func TestLlenarListaVaciaConInsercionesIteararConOtro(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	iter.Insertar(100)
	iter.Insertar(95)
	iter.Siguiente()
	iter.Insertar(30)
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(12)
	iter.Insertar(32)
	iter.Insertar(22)
	iter.Siguiente()
	iter.Insertar(10)
	iter2 := lista.Iterador()
	require.EqualValues(t, 95, iter2.VerActual())
	require.EqualValues(t, 95, iter2.Borrar())
	require.EqualValues(t, 30, iter2.Borrar())
	require.EqualValues(t, 100, iter2.VerActual())
	require.EqualValues(t, 100, iter2.Borrar())
	require.EqualValues(t, 22, iter2.VerActual())
	require.EqualValues(t, 22, iter2.Borrar())
	require.EqualValues(t, 10, iter2.VerActual())
	require.EqualValues(t, 10, iter2.Borrar())
	require.EqualValues(t, 32, iter2.VerActual())
	require.EqualValues(t, 32, iter2.Borrar())
	require.EqualValues(t, 12, iter2.VerActual())
	require.EqualValues(t, 12, iter2.Borrar())
}

func TestIteradoresInternoYExterno(t *testing.T) {
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(10)
	iter.Insertar(3)
	iter.Insertar(5)
	iter.Insertar(1)
	iter.Siguiente() //1,2,5,3,0,10
	iter.Insertar(2)
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(0)
	iter.Siguiente()
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	suma := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
		}
		return true
	})
	require.EqualValues(t, 12, suma)
	iter2 := lista.Iterador()
	require.EqualValues(t, 1, iter2.Borrar())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 5, iter2.Borrar())

}

func TestInvertirCola(t *testing.T){
	lista := TDAlista.CrearListaDobleEnlazada[int]()
	array := []int{1,2,3,4,5,6,7,8}
	for _,elem:=range(array){
		lista.InsertarUltimo(elem)
	}
	iter := lista.Iterador()
	for i:=0;i<lista.Largo();i++{
		require.EqualValues(t, array[i], iter.VerActual())
		iter.Siguiente()
	}
	lista.InvertirLista()
	iter2 := lista.Iterador()
	for i:=len(array)-1; i>=0; i--{
		require.EqualValues(t, array[i], iter2.VerActual())
		iter2.Siguiente()
	} 
}