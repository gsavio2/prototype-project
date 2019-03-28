package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/pedrohmachado/prototype-project/src/platform/evento"
)

// ListaTodosEvento lista todos os eventos
func ListaTodosEvento(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resultado := e.ListaTodos()
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaTodosEventoUsuario lista os eventos de um DONO pelo id do usuaio
func ListaTodosEventoUsuario(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		resultado := e.ListaTodosDono(usuario)
		json.NewEncoder(w).Encode(resultado)
	}
}

// ListaEvento lista um evento pelo id do evento
func ListaEvento(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := e.Lista(params["id_evento"])
		json.NewEncoder(w).Encode(resultado)
	}
}
