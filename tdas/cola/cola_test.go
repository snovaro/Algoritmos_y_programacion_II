package cola_test

import (
	TDAcola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 1000

func TestColaVacia(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaEnteroUnico(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 5, cola.VerPrimero())
	require.EqualValues(t, 5, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEnterosVarios(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(-9)
	cola.Encolar(7)
	cola.Encolar(30)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, -9, cola.VerPrimero())
	require.EqualValues(t, -9, cola.Desencolar())
	require.EqualValues(t, 7, cola.VerPrimero())
	require.EqualValues(t, 7, cola.Desencolar())
	require.EqualValues(t, 30, cola.VerPrimero())
	require.EqualValues(t, 30, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaFloatUnico(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[float32]()
	cola.Encolar(3.1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 3.1, cola.VerPrimero())
	require.EqualValues(t, 3.1, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestFloatVarios(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[float32]()
	cola.Encolar(9.9)
	cola.Encolar(0.0)
	cola.Encolar(3.1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 9.9, cola.VerPrimero())
	require.EqualValues(t, 9.9, cola.Desencolar())
	require.EqualValues(t, 0.0, cola.VerPrimero())
	require.EqualValues(t, 0.0, cola.Desencolar())
	require.EqualValues(t, 3.1, cola.VerPrimero())
	require.EqualValues(t, 3.1, cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestColaStringUnico(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[string]()
	cola.Encolar("Hola Mundo")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Hola Mundo", cola.VerPrimero())
	require.EqualValues(t, "Hola Mundo", cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestColaStringVarios(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[string]()
	cola.Encolar("Mundo")
	cola.Encolar("Dias")
	cola.Encolar("Buenos")
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, "Mundo", cola.VerPrimero())
	require.EqualValues(t, "Mundo", cola.Desencolar())
	require.EqualValues(t, "Dias", cola.VerPrimero())
	require.EqualValues(t, "Dias", cola.Desencolar())
	require.EqualValues(t, "Buenos", cola.VerPrimero())
	require.EqualValues(t, "Buenos", cola.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestVariosEncoloDesencoloDesordenados(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(30)
	cola.Encolar(7)
	cola.Encolar(-9)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 30, cola.VerPrimero())
	require.EqualValues(t, 30, cola.Desencolar())
	require.EqualValues(t, 7, cola.VerPrimero())
	require.EqualValues(t, 7, cola.Desencolar())
	cola.Encolar(12)
	require.EqualValues(t, -9, cola.VerPrimero())
	require.EqualValues(t, -9, cola.Desencolar())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 12, cola.VerPrimero())
	require.EqualValues(t, 12, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncoloVacioEncolo(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	cola.Encolar(-109)
	cola.Encolar(67)
	cola.Encolar(3)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, -109, cola.VerPrimero())
	require.EqualValues(t, -109, cola.Desencolar())
	require.EqualValues(t, 67, cola.VerPrimero())
	require.EqualValues(t, 67, cola.Desencolar())
	require.EqualValues(t, 3, cola.VerPrimero())
	require.EqualValues(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	cola.Encolar(321)
	cola.Encolar(12)
	require.EqualValues(t, 321, cola.VerPrimero())
	require.EqualValues(t, 321, cola.Desencolar())
	require.EqualValues(t, 12, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 12, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaVolumen(t *testing.T) {
	cola := TDAcola.CrearColaEnlazada[int]()
	for i := 0; i <= VOLUMEN; i++ {
		cola.Encolar(i)
	}
	for i := 0; i <= VOLUMEN; i++ {
		require.False(t, cola.EstaVacia())
		require.EqualValues(t, i, cola.VerPrimero())
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.True(t, cola.EstaVacia())
}

func TestColaDeColas(t *testing.T) {
	colaMayor := TDAcola.CrearColaEnlazada[TDAcola.Cola[int]]()
	for j := 0; j <= VOLUMEN; j++ {
		cola := TDAcola.CrearColaEnlazada[int]()
		for i := 0; i <= VOLUMEN; i++ {
			cola.Encolar(i)
		}
		colaMayor.Encolar(cola)
	}
	for j := 0; j <= VOLUMEN; j++ {
		cola := colaMayor.Desencolar()
		for i := 0; i <= VOLUMEN; i++ {
			require.False(t, cola.EstaVacia())
			require.EqualValues(t, i, cola.VerPrimero())
			require.EqualValues(t, i, cola.Desencolar())
		}
		require.True(t, cola.EstaVacia())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaMayor.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaMayor.Desencolar() })
	require.True(t, colaMayor.EstaVacia())
}
