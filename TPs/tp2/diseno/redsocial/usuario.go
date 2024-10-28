package redsocial

type Usuario interface {
	ObtenerNombre() string

	AgregarAlFeed(*Post)

	ObtenerPosicion() int

	ObtenerPost() (Post, bool)
}

type Usuarios interface {
	PublicarPost(*Post)

	ObtenerLoggeado() *Usuario

	LikearPost(int, Publicaciones) error

	VerProximoPost() error

	Login(string) error

	Logout()

	HayUsuario() bool
}
