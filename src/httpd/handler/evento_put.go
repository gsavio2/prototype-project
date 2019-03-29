package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrohmachado/prototype-project/src/platform/evento"
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
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

// AlteraEvento altera evento
func AlteraEvento(e *evento.Eventos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		evento := e.Lista(params["id_evento"])
		_ = json.NewDecoder(r.Body).Decode(&evento)
		eventoAtualizado := e.Altera(evento, params["id_evento"])
		resultado := e.Atualiza(eventoAtualizado)
		json.NewEncoder(w).Encode(resultado)
	}
}

// RemoveParticipante remove participante do evento
func RemoveParticipante(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		participante := u.Lista(params["id"])
		evento := e.Lista(params["id_evento"])
		eventoAtualizado := e.RemoveParticipante(participante, evento)
		resultado := e.Atualiza(eventoAtualizado)
		json.NewEncoder(w).Encode(resultado)
	}
}
