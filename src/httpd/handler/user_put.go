package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

//AlteraUsuario altera usuario
func AlteraUsuario(u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		resultado := u.Altera(usuario, params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}
