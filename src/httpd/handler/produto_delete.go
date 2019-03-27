package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/gorilla/mux"
)

//ExcluiProduto exclui produto pelo id
func ExcluiProduto(p *produto.Produtos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := p.Exclui(params["id_produto"])
		json.NewEncoder(w).Encode(resultado)
	}
}
