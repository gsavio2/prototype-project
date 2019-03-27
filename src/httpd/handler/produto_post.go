package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// AdicionaProduto cria um produto relacionado a um usuario
func AdicionaProduto(p *produto.Produtos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		var produto produto.Produto
		_ = json.NewDecoder(r.Body).Decode(&produto)
		resultado := p.Adiciona(produto, usuario)
		json.NewEncoder(w).Encode(resultado)
	}
}
