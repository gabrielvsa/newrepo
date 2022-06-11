package usuarios

type usuarios struct {
	Nome  string
	email string
}

func NewUsuario(nome, email string) *usuarios {
	return &usuarios{
		Nome:  nome,
		email: email,
	}
}
