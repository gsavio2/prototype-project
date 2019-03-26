package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// ListaProdutoTodos lista todos os produtos de um usuario pelo id do usuario
func ListaProdutoTodos(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := u.ListaTodosProduto(params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaProduto lista um produto pelo seu id através do id de um usuario
func ListaProduto(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		produtos := u.ListaTodosProduto(params["id"])
		resultado := u.ListaProduto(produtos, params["id_produto"])
		json.NewEncoder(w).Encode(resultado)
	}
}
