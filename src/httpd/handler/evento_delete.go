package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/evento"
)

// ExcluiEvento exclui evento pelo id
func ExcluiEvento(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := e.Exclui(params["id_evento"])
		json.NewEncoder(w).Encode(resultado)
	}
}
