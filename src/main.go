package main

import (
	"log"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/evento"

	"github.com/rs/cors"

	"github.com/pedrohmachado/prototype-project/src/httpd/handler"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	u := usuario.Novo()
	e := evento.Novo()

	c := cors.AllowAll().Handler(r)

	// crud usuarios
	r.HandleFunc("/api/user", handler.ListaTodos(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}", handler.Lista(u)).Methods("GET")
	r.HandleFunc("/api/user", handler.Novo(u)).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.Exclui(u)).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", handler.Altera(u)).Methods("PUT")

	// crud produtos/usuario
	r.HandleFunc("/api/user/{id}/produto", handler.ListaProdutoTodos(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}/produto/{id_produto}", handler.ListaProduto(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}/produto", handler.AdicionaProduto(u)).Methods("POST")
	r.HandleFunc("/api/user/{id}/produto/{id_produto}", handler.ExcluiProduto(u)).Methods("DELETE")
	r.HandleFunc("/api/user/{id}/produto/{id_produto}", handler.AlteraProduto(u)).Methods("PUT")

	// crud eventos
	r.HandleFunc("/api/evento", handler.ListaEventoTodos(e)).Methods("GET")
	// r.HandleFunc("/api/evento/{id_evento}", handler.ListaEvento(e)).Methods("GET")

	r.HandleFunc("/api/user/{id}/evento", handler.NovoEvento(e, u)).Methods("POST")
	r.HandleFunc("/api/user/{id}/evento", handler.ListaEventoTodosUsuario(e, u)).Methods("GET")
	// r.HandleFunc("/api/user/{id}/evento/{id_evento}", handler.ExcluiEvento(e)).Methods("DELETE")
	// r.HandleFunc("/api/user/{id}/evento/{id_evento}", handler.AlteraEvento(e)).Methods("PUT")
	// r.HandleFunc("/api/evento/{id_evento}", handler.EncerraEvento(e)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", c))
}
