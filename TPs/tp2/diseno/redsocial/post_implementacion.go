package redsocial

import (
	"algogram/diseno/errores"
	"fmt"
	TDADiccionario "tdas/diccionario"
)

type post struct {
	publicacion string
	usuario     string
	posUsuario  int
	likes       TDADiccionario.DiccionarioOrdenado[string, int]
	id          id_post
}

func CrearPost(texto string, usuarios Usuarios, id id_post) Post {
	loggeado := usuarios.ObtenerLoggeado()
	post := new(post)
	post.publicacion = texto
	post.usuario = (*loggeado).ObtenerNombre()
	post.posUsuario = (*loggeado).ObtenerPosicion()
	post.likes = TDADiccionario.CrearABB[string, int](cmpstring)
	post.id = id
	return post
}

func (post *post) MostrarPost() {
	fmt.Println("Post ID", post.id)
	fmt.Println(post.usuario, "dijo:", post.publicacion)
	fmt.Println("Likes:", post.likes.Cantidad())
}

func (post *post) ImprimirLikes() error {
	cantidadLikes := post.likes.Cantidad()
	if cantidadLikes == 0 {
		return new(errores.PostInexistenteOSinLikes)
	}
	fmt.Printf("El post tiene %d likes:\n", cantidadLikes)
	iter := post.likes.Iterador()
	for iter.HaySiguiente() {
		actual, _ := iter.VerActual()
		fmt.Printf("\t%s\n", actual)
		iter.Siguiente()
	}
	return nil
}

func (post *post) PosicionUsuario() int {
	return post.posUsuario
}

func (post *post) VerIDPost() id_post {
	return post.id
}

func (publicacion *post) DarLike(nombre string) {
	publicacion.likes.Guardar(nombre, 0)
}
