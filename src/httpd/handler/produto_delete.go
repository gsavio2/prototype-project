package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// ExcluiProduto exclui produto relacionado a um usuario
func ExcluiProduto(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		resultado := u.ExcluiProduto(usuario, params["id"], params["id_produto"])
		json.NewEncoder(w).Encode(resultado)
	}
}
