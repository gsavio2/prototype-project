package evento

import (
	"time"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Evento struct modelo
type Evento struct {
	ID            string            `json: "id"`
	Descricao     string            `json: "descricao"`
	Criador       usuario.Usuario   `json: "criador"`
	Participantes []usuario.Usuario `json: "participantes"`
	DataCriacao   string            `json: "data_criacao"`
	DataEvento    string            `json: "data_evento"`
}

// Eventos struct modelo
type Eventos struct {
	Eventos []Evento
}

// Novo inicia eventos
func Novo() *Eventos {
	return &Eventos{
		Eventos: []Evento{},
	}
}

// NovoEvento cria novo evento
func (e *Eventos) NovoEvento(usuario usuario.Usuario, evento Evento) []Evento {
	evento.Criador = usuario
	layout := "01/02/2006"
	evento.DataCriacao = time.Now().Format(layout)
	e.Eventos = append(e.Eventos, evento)
	return e.Eventos
}

// ListaEventoTodos lista todos os eventos
func (e *Eventos) ListaEventoTodos() []Evento {
	return e.Eventos
}

// ListaEvento lista evento pelo id do evento
func (e *Eventos) ListaEvento() {

}

// ListaEvento lista evento pelo id do evento
func (e *Eventos) ListaEventoTodosUsuario(usuario usuario.Usuario) []Evento {
	var eventos []Evento
	for _, evento := range e.Eventos {
		if evento.Criador.ID == usuario.ID {
			eventos = append(eventos, evento)
		}
	}
	return eventos
}

// ExcluiEvento pelo id do evento
func (e *Eventos) ExcluiEvento() {

}

// AlteraEvento pelo id do evento
func (e *Eventos) AlteraEvento() {

}

// EncerraEvento pelo id do evento
func (e *Eventos) EncerraEvento() {

}
