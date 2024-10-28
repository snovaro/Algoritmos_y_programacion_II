package redsocial

import (
	"algogram/diseno/errores"
	TDADiccionario "tdas/diccionario"
)

type usuarios struct {
	hash       TDADiccionario.Diccionario[string, *Usuario]
	hayUsuario bool
	loggeado   *Usuario
}

func CrearUsuarios(hash TDADiccionario.Diccionario[string, *Usuario]) Usuarios {
	users := new(usuarios)
	users.hash = hash
	users.hayUsuario = false
	return users
}

func (users usuarios) HayUsuario() bool {
	return users.loggeado != nil
}

func (users usuarios) LikearPost(id int, publicaciones Publicaciones) error {
	usuario := users.loggeado
	nombre := (*usuario).ObtenerNombre()
	return publicaciones.LikearPost(id_post(id), nombre)
}

func (users *usuarios) Login(nombre string) error {
	if !users.hash.Pertenece(nombre) {
		return new(errores.ErrorUsuarioInexistente)
	}
	usuario := users.hash.Obtener(nombre)
	users.hayUsuario = true
	users.loggeado = usuario
	return nil
}

func (users *usuarios) Logout() {
	users.hayUsuario = false
	users.loggeado = nil
}

func (users usuarios) PublicarPost(post *Post) {
	users.hash.Iterar(func(nombre string, user *Usuario) bool {
		if user != users.loggeado {
			(*user).AgregarAlFeed(post)
		}
		return true
	})
}

func (users *usuarios) ObtenerLoggeado() *Usuario {
	return users.loggeado
}

func (users *usuarios) VerProximoPost() error {
	user := users.loggeado
	post, estaVacia := (*user).ObtenerPost()
	if estaVacia {
		return new(errores.NoHayMasPostsONoHayUsuario)
	}
	post.MostrarPost()
	return nil
}
