package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/evento"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"
)

// ListaTodosProduto lista todos os produtos
func ListaTodosProduto(p *produto.Produtos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resultado := p.ListaTodos()
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaProduto lista produto pelo id do produto
func ListaProduto(p *produto.Produtos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := p.Lista(params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaTodosProdutoUsuario lista todos os produtos de um usuario pelo id do usuario
func ListaTodosProdutoUsuario(p *produto.Produtos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		resultado := p.ListaTodosUsuario(usuario)
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaTodosProdutoEvento lista todos os produtos de um evento pelo id do evento
func ListaTodosProdutoEvento(p *produto.Produtos, e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		evento := e.Lista(params["id_evento"])
		participantes := e.ListaUsuarios(evento)
		produtos := p.ListaTodos()
		resultado := e.ListaProdutos(participantes, produtos)
		json.NewEncoder(w).Encode(resultado)
	}
}
