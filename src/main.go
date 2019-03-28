package main

import (
	"log"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/pedrohmachado/prototype-project/src/platform/evento"

	"github.com/rs/cors"

	"github.com/pedrohmachado/prototype-project/src/httpd/handler"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/gorilla/mux"
)

func main() {
	// roteador
	r := mux.NewRouter()

	// registros
	u := usuario.Novo()
	e := evento.Novo()
	p := produto.Novo()

	c := cors.AllowAll().Handler(r)

	// crud usuarios
	r.HandleFunc("/api/user", handler.NovoUsuario(u)).Methods("POST")
	r.HandleFunc("/api/user", handler.ListaTodosUsuario(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}", handler.ListaUsuario(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}", handler.AlteraUsuario(u, e)).Methods("PUT")

	// vvvvvvvvvvvvvvvvv deletar usuario não pode ser utilizado vvvvvvvvvvvvvvvv
	// r.HandleFunc("/api/user/{id}", handler.ExcluiUsuario(u)).Methods("DELETE")

	// crud produtos (agora)
	r.HandleFunc("/api/user/{id}/produto", handler.AdicionaProduto(p, u)).Methods("POST")
	r.HandleFunc("/api/produto", handler.ListaTodosProduto(p)).Methods("GET")
	r.HandleFunc("/api/produto/{id}", handler.ListaProduto(p)).Methods("GET")
	r.HandleFunc("/api/user/{id}/produto", handler.ListaTodosProdutoUsuario(p, u)).Methods("GET")
	r.HandleFunc("/api/evento/{id_evento}/produto", handler.ListaTodosProdutoEvento(p, e)).Methods("GET")
	// r.HandleFunc("/api/user/{id}/produto{id_produto}", handler.AlteraProduto(p, u)).Methods("PUT")

	// r.HandleFunc("/api/produto/{id_produto}", handler.ExcluiProduto(p, u)).Methods("DELETE")

	// crud eventos
	r.HandleFunc("/api/user/{id}/evento", handler.NovoEvento(e, u)).Methods("POST")
	r.HandleFunc("/api/evento", handler.ListaTodosEvento(e)).Methods("GET")
	r.HandleFunc("/api/user/{id}/evento", handler.ListaTodosEventoUsuario(e, u)).Methods("GET")
	r.HandleFunc("/api/evento/{id_evento}", handler.ListaEvento(e)).Methods("GET")
	r.HandleFunc("/api/evento/{id_evento}/user/{id}", handler.AdicionaParticipante(e, u)).Methods("POST")
	// r.HandleFunc("/api/user/{id}/evento/{id_evento}", handler.AlteraEvento(e)).Methods("PUT")
	r.HandleFunc("/api/evento/{id_evento}", handler.ExcluiEvento(e)).Methods("DELETE")
	r.HandleFunc("/api/evento/{id_evento}", handler.EncerraEvento(e)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", c))
}
