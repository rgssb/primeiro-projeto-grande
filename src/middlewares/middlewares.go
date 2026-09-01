package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Logger escreve as informacoes da requisicao no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}



//func (w http.ResponseWriter, r *http.Request)
//Autenticar verifica se o usuario que esta fazendo a autenticação, esta autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}
