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
		resultado := e.Adiciona(usuario, evento)
		json.NewEncoder(w).Encode(resultado)
	}
}

// AdicionaParticipante adiciona um usuario participante aos participantes do evento (e atualiza a lista dos eventos)
func AdicionaParticipante(e *evento.Eventos, u *usuario.Usuarios) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		participante := u.Lista(params["id"])
		evento := e.Lista(params["id_evento"])
		eventoAtualizado := e.AdicionaParticipante(participante, evento)
		resultado := e.Atualiza(eventoAtualizado)
		json.NewEncoder(w).Encode(resultado)
	}
}
