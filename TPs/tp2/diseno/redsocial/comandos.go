package redsocial

import (
	"algogram/diseno/errores"
	"fmt"
	"strconv"
)

func Login(users Usuarios, nombre string) {
	if users.HayUsuario() {
		fmt.Println(new(errores.YaHayUsuario).Error())
		return
	}
	if users.Login(nombre) != nil {
		fmt.Println(new(errores.ErrorUsuarioInexistente).Error())
	} else {
		fmt.Println("Hola", nombre)
	}
}

func Logout(users Usuarios) {
	if !users.HayUsuario() {
		fmt.Println(new(errores.ErrorNoHayUsuario).Error())
		return
	}
	users.Logout()
	fmt.Println("Adios")
}

func Publicar(users Usuarios, texto string, publicaciones Publicaciones) {
	if !users.HayUsuario() {
		fmt.Println(new(errores.ErrorNoHayUsuario).Error())
		return
	}
	post := CrearPost(texto, users, publicaciones.ObtenerNuevoID())
	publicaciones.Publicar(&post)
	users.PublicarPost(&post)
	fmt.Println("Post publicado")
}

func VerSiguienteFeed(users Usuarios) {
	if !users.HayUsuario() || users.VerProximoPost() != nil {
		fmt.Println(new(errores.NoHayMasPostsONoHayUsuario).Error())
	}
}

func LikearPost(users Usuarios, publicaciones Publicaciones, input string) {
	if !users.HayUsuario() {
		fmt.Println(new(errores.UsuarioNoLoggeadoOPostInexistente).Error())
		return
	}
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(new(errores.ErrorParametros).Error())
		return
	}
	if users.LikearPost(id, publicaciones) != nil {
		fmt.Println(new(errores.UsuarioNoLoggeadoOPostInexistente).Error())
	} else {
		fmt.Println("Post likeado")
	}
}

func MostrarLikes(publicaciones Publicaciones, input string) {
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(new(errores.ErrorParametros).Error())
		return
	}
	if publicaciones.MostrarLikes(id_post(id)) != nil {
		fmt.Println(new(errores.PostInexistenteOSinLikes).Error())
	}
}
