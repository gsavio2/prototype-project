package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// AdicionaProduto relaciona produto a um usuario
func AdicionaProduto(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var produto produto.Produto
		_ = json.NewDecoder(r.Body).Decode(&produto)
		usuario := u.Lista(params["id"])
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		resultado := u.AdicionaProduto(produto, usuario, params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}
