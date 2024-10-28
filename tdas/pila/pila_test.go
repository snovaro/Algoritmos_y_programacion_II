package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 10000

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaEnteroUnico(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(5)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 5, pila.VerTope())
	require.EqualValues(t, 5, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaFloatUnico(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float32]()
	pila.Apilar(3.1)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 3.1, pila.VerTope())
	require.EqualValues(t, 3.1, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestPilaFloatVarios(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float32]()
	pila.Apilar(3.1)
	pila.Apilar(0.0)
	pila.Apilar(9.9)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 9.9, pila.VerTope())
	require.EqualValues(t, 9.9, pila.Desapilar())
	require.EqualValues(t, 0.0, pila.VerTope())
	require.EqualValues(t, 0.0, pila.Desapilar())
	require.EqualValues(t, 3.1, pila.VerTope())
	require.EqualValues(t, 3.1, pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestEnterosVarios(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(30)
	pila.Apilar(7)
	pila.Apilar(-9)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, -9, pila.VerTope())
	require.EqualValues(t, -9, pila.Desapilar())
	require.EqualValues(t, 7, pila.VerTope())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 30, pila.VerTope())
	require.EqualValues(t, 30, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaStringUnico(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola Mundo")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Hola Mundo", pila.VerTope())
	require.EqualValues(t, "Hola Mundo", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestPilaStringVarios(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Buenos")
	pila.Apilar("Dias")
	pila.Apilar("Mundo")
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, "Mundo", pila.VerTope())
	require.EqualValues(t, "Mundo", pila.Desapilar())
	require.EqualValues(t, "Dias", pila.VerTope())
	require.EqualValues(t, "Dias", pila.Desapilar())
	require.EqualValues(t, "Buenos", pila.VerTope())
	require.EqualValues(t, "Buenos", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}

func TestApiloVacioApilo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	pila.Apilar(67)
	pila.Apilar(-109)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, -109, pila.VerTope())
	require.EqualValues(t, -109, pila.Desapilar())
	require.EqualValues(t, 67, pila.VerTope())
	require.EqualValues(t, 67, pila.Desapilar())
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	pila.Apilar(12)
	pila.Apilar(321)
	require.EqualValues(t, 321, pila.VerTope())
	require.EqualValues(t, 321, pila.Desapilar())
	require.EqualValues(t, 12, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 12, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= VOLUMEN; i++ {
		pila.Apilar(i)
	}
	for i := VOLUMEN; i >= 0; i-- {
		require.False(t, pila.EstaVacia())
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.True(t, pila.EstaVacia())
}
