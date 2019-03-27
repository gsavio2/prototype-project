package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// NovoUsuario adiciona usuario
func NovoUsuario(u usuario.Adiciona) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var usuario usuario.Usuario
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		u.Adiciona(usuario)
		json.NewEncoder(w).Encode(usuario)
	}
}
