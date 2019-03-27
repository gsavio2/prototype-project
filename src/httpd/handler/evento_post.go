package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/evento"
)

// NovoEvento cria evento associado ao usuario criador
func NovoEvento(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		var evento evento.Evento
		_ = json.NewDecoder(r.Body).Decode(&evento)
		resultado := e.NovoEvento(usuario, evento)
		json.NewEncoder(w).Encode(resultado)
	}
}
