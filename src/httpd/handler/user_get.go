package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// ListaTodos os usuarios
func ListaTodos(u usuario.ListaTodos) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resultado := u.ListaTodos()
		json.NewEncoder(w).Encode(resultado)
	}
}

//Lista um usuario pelo id
func Lista(u usuario.Lista) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		resultado := u.Lista(params["id"])
		json.NewEncoder(w).Encode(resultado)
	}
}
