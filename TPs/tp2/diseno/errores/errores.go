package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type YaHayUsuario struct{}

func (e YaHayUsuario) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorNoHayUsuario struct{}

func (e ErrorNoHayUsuario) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type UsuarioNoLoggeadoOPostInexistente struct {
}

func (e UsuarioNoLoggeadoOPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type PostInexistenteOSinLikes struct{}

func (e PostInexistenteOSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

type NoHayMasPostsONoHayUsuario struct{}

func (e NoHayMasPostsONoHayUsuario) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}
