package modelos


//DadosAutenticacao contem o token e o id do usuario autenticado
type DadosAutenticacao struct {
	Email    string `json:"email"`
	Senha    string `json:"senha"`
}