package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/evento"
)

// EncerraEvento encerra um evento (muda o status para inativo)
func EncerraEvento(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		evento := e.Lista(params["id_evento"])
		eventoAtualizado := e.Encerra(evento)
		resultado := e.Atualiza(eventoAtualizado)
		json.NewEncoder(w).Encode(resultado)
	}
}
