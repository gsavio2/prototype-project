package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/pedrohmachado/prototype-project/src/httpd/handler"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	u := usuario.Novo()

	c := cors.AllowAll().Handler(r)

	r.HandleFunc("/api/user", handler.ListaTodos(u)).Methods("GET")
	r.HandleFunc("/api/user/{id}", handler.Lista(u)).Methods("GET")
	r.HandleFunc("/api/user", handler.Novo(u)).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.Exclui(u)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", c))
}
