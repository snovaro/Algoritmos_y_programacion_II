package redsocial

type id_post int

type Post interface {
	ImprimirLikes() error

	PosicionUsuario() int

	MostrarPost()

	DarLike(string)

	VerIDPost() id_post
}

type Publicaciones interface {
	Publicar(*Post)

	LikearPost(id_post, string) error

	MostrarLikes(id_post) error

	ObtenerNuevoID() id_post
}
