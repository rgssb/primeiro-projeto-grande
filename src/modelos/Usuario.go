package modelos

import (
	"API/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Usuario representa um usuario ativo da rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

//Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == ""{
		return errors.New("O nome é obrigatorio, não será possivel deixa-lo em branco")
	}

	if usuario.Nick == ""{
		return errors.New("O nick é obrigatorio, não será possivel deixa-lo em branco")
	}

	if usuario.Email == ""{
		return errors.New("O email é obrigatorio, não será possivel deixa-lo em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é invalido")
	}

	if etapa == "cadastro" && usuario.Senha == ""{
		return errors.New("A senha é obrigatoria, não será possivel deixa-la em branco")
	}

		return nil

}

func(usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email  = strings.TrimSpace(usuario.Email)


	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		
		usuario.Senha = string(senhaComHash)
	
	}
		return nil
}