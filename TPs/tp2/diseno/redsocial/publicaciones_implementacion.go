package redsocial

import (
	"algogram/diseno/errores"
	TDADiccionario "tdas/diccionario"
)

type publicaciones struct {
	hash     TDADiccionario.Diccionario[id_post, *Post]
	id_nuevo id_post
}

func (publicaciones *publicaciones) LikearPost(id id_post, nombre string) error {
	if !publicaciones.hash.Pertenece(id) {
		return new(errores.UsuarioNoLoggeadoOPostInexistente)
	}
	publicacion := publicaciones.hash.Obtener(id)
	(*publicacion).DarLike(nombre)
	return nil
}

func (posts *publicaciones) MostrarLikes(id id_post) error {
	if !posts.hash.Pertenece(id) {
		return new(errores.PostInexistenteOSinLikes)
	}
	post := posts.hash.Obtener(id)
	return (*post).ImprimirLikes()
}

func CrearPublicaciones() Publicaciones {
	posts := new(publicaciones)
	hash := TDADiccionario.CrearHash[id_post, *Post]()
	posts.hash = hash
	return posts
}

func (posts *publicaciones) Publicar(post *Post) {
	posts.hash.Guardar(posts.id_nuevo, post)
	posts.id_nuevo++
}

func (posts *publicaciones) ObtenerNuevoID() id_post {
	return posts.id_nuevo
}
