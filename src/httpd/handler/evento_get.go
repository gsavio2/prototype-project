package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/pedrohmachado/prototype-project/src/platform/evento"
)

// ListaEventoTodos lista todos os eventos
func ListaEventoTodos(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resultado := e.ListaEventoTodos()
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaEventoTodosUsuario lista os eventos de um usuario pelo id do usuaio
func ListaEventoTodosUsuario(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		resultado := e.ListaEventoTodosUsuario(usuario)
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaEvento lista um evento pelo id do evento
func ListaEvento(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

	}
}
