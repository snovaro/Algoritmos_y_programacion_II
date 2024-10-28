package redsocial

import (
	"algogram/diseno/errores"
	"bufio"
	"fmt"
	"os"
	TDADiccionario "tdas/diccionario"
)

func LeerUsuarios(nombreArchivo string) Usuarios {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Fprintln(os.Stdout, new(errores.ErrorLeerArchivo).Error())
		return nil
	}
	hashUsuarios := TDADiccionario.CrearHash[string, *Usuario]()
	lineaArch := bufio.NewScanner(archivo)
	posUsuario := 0
	for lineaArch.Scan() {
		nombre := lineaArch.Text()
		usuario := CrearUsuario(nombre, posUsuario)
		hashUsuarios.Guardar(nombre, &usuario)
		posUsuario++
	}
	usuarios := CrearUsuarios(hashUsuarios)
	return usuarios
}
