package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// AlteraProduto altera produto
func AlteraProduto(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// ###### melhorar o bloco abaixo #######
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		produtos := u.ListaTodosProduto(params["id"])
		produto := u.ListaProduto(produtos, params["id_produto"])
		// ######################################
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		resultado := u.AlteraProduto(usuario, produto, params["id"], params["id_produto"])
		json.NewEncoder(w).Encode(resultado)
	}
}
