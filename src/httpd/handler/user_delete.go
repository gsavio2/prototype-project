package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Exclui usuario
func Exclui(u usuario.Exclui) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := u.Exclui(params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}
