package evento

import (
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Evento struct modelo
type Evento struct {
	ID            string            `json: "id"`
	Descricao     string            `json: "descricao"`
	Participantes []usuario.Usuario `json: "participantes"`
}

// Eventos struct modelo
type Eventos struct {
	Eventos []Evento
}
