package autenticacao

import (
	"API/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour *6).Unix()
	permissoes["usuarioID"] = usuarioID
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

//ValidarToken valida o token recebido da requisicao
func ValidarToken(r *http.Request) error {
	tokenString := extrairTokens(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
		if erro != nil {
			return erro
		}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

			return errors.New("Ixi Token invalido")
}

func extrairTokens(r *http.Request) string {
	token := r.Header.Get("Authorization")
	//Bearer andiandnqklladsa

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return " "

}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}