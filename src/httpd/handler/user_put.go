package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Altera usuario
func Altera(uA usuario.Altera, uL usuario.Lista) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := uL.Lista(params["id"])
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		resultado := uA.Altera(params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}
