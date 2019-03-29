package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/produto"
)

//AlteraProduto altera produto pelo id do produto
func AlteraProduto(p *produto.Produtos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		produto := p.Lista(params["id_produto"])
		_ = json.NewDecoder(r.Body).Decode(&produto)
		resultado := p.Altera(produto, params["id_produto"])
		json.NewEncoder(w).Encode(resultado)
	}
}
