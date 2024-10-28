package main

import (
	"algogram/diseno/errores"
	TDARedSocial "algogram/diseno/redsocial"
	"bufio"
	"fmt"
	"os"
)

const (
	SELECCION1 = "login"
	SELECCION2 = "logout"
	SELECCION3 = "publicar"
	SELECCION4 = "ver_siguiente_feed"
	SELECCION5 = "likear_post"
	SELECCION6 = "mostrar_likes"
)

func main() {
	publicaciones := TDARedSocial.CrearPublicaciones() // base de datos de todas las publicaciones (hash + ID_Nuevo)
	nombreArchivo := os.Args[1]
	usuarios := TDARedSocial.LeerUsuarios(nombreArchivo) // Hash con todos los usuarios + bool + logueado actual

	if usuarios == nil {
		fmt.Println(new(errores.ErrorLeerArchivo).Error())
	}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() && scanner.Text() != "" {
		input := scanner.Text()
		vectorIngreso := TDARedSocial.ConvertirIngreso(input)

		if vectorIngreso[0] == SELECCION1 {
			TDARedSocial.Login(usuarios, vectorIngreso[1])
		} else if vectorIngreso[0] == SELECCION2 {
			TDARedSocial.Logout(usuarios)
		} else if vectorIngreso[0] == SELECCION3 {
			TDARedSocial.Publicar(usuarios, vectorIngreso[1], publicaciones)
		} else if vectorIngreso[0] == SELECCION4 {
			TDARedSocial.VerSiguienteFeed(usuarios)
		} else if vectorIngreso[0] == SELECCION5 {
			TDARedSocial.LikearPost(usuarios, publicaciones, vectorIngreso[1])
		} else if vectorIngreso[0] == SELECCION6 {
			TDARedSocial.MostrarLikes(publicaciones, vectorIngreso[1])
		}
	}
}
