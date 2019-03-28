package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/evento"
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

//AlteraUsuario altera usuario
func AlteraUsuario(u *usuario.Usuarios, e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuario := u.Lista(params["id"])
		_ = json.NewDecoder(r.Body).Decode(&usuario)
		usuarioAlterado := u.Altera(usuario, params["id"])
		// ^^^^^^^^^ registro de usuario ja foi alterado ^^^^^^^^

		//vvvvvvvvvvvvv atualizar eventos vvvvvvvvvvvvvvvvvvvvvv
		eventos := e.ListaTodosParticipante(usuario)

		// caso o usuario esteja em algum evento
		for _, evento := range eventos {
			evento = e.AtualizaUsuario(usuarioAlterado, evento)
			e.Atualiza(evento)
		}

		resultado := usuarioAlterado

		json.NewEncoder(w).Encode(resultado)
	}
}
