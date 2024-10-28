package redsocial

import (
	TDAheap "tdas/cola_prioridad"
)

type usuario struct {
	nombre   string
	posicion int
	feed     TDAheap.ColaPrioridad[*Post]
}

func CrearUsuario(nombre string, posicion int) Usuario {
	heap := TDAheap.CrearHeap[*Post](func(a, b *Post) int {
		diferencia := Abs(posicion-(*b).PosicionUsuario()) - Abs(posicion-(*a).PosicionUsuario())
		if diferencia > 0 {
			return 1
		}
		if diferencia == 0 {
			if (*b).VerIDPost() > (*a).VerIDPost() {
				return 1
			}
			return -1
		}
		return -1
	})
	user := new(usuario)
	user.feed = heap
	user.posicion = posicion
	user.nombre = nombre
	return user
}

func (user *usuario) ObtenerPost() (Post, bool) {
	if user.feed.EstaVacia() {
		return nil, true
	}
	return *user.feed.Desencolar(), false
}

func (user usuario) AgregarAlFeed(post *Post) {
	user.feed.Encolar(post)
}

func (user usuario) ObtenerNombre() string {
	return user.nombre
}

func (user *usuario) ObtenerPosicion() int {
	return user.posicion
}
