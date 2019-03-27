package evento

import (
	"time"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Evento struct modelo
type Evento struct {
	ID            string            `json: "id"`
	Descricao     string            `json: "descricao"`
	Criador       usuario.Usuario   `json: "criador"`
	Participantes []usuario.Usuario `json: "participantes"`
	DataCriacao   string            `json: "dataCriacao"`
	DataEvento    string            `json: "dataEvento"`
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

// Adiciona cria novo evento
func (e *Eventos) Adiciona(usuario usuario.Usuario, evento Evento) []Evento {
	evento.Criador = usuario
	//layout do formato de data/hora
	layout := "02/01/2006 15:04"
	evento.DataCriacao = time.Now().Format(layout)
	e.Eventos = append(e.Eventos, evento)
	return e.Eventos
}

// ListaTodos lista todos os eventos
func (e *Eventos) ListaTodos() []Evento {
	return e.Eventos
}

// Lista lista evento pelo id do evento
func (e *Eventos) Lista(id string) Evento {
	for _, evento := range e.Eventos {
		if evento.ID == id {
			return evento
		}
	}
	return Evento{}
}

// ListaTodosUsuario lista todos os eventos de um usuario pelo id do usuario
func (e *Eventos) ListaTodosUsuario(usuario usuario.Usuario) []Evento {
	var eventos []Evento
	for _, evento := range e.Eventos {
		if evento.Criador.ID == usuario.ID {
			eventos = append(eventos, evento)
		}
	}
	return eventos
}

// ListaUsuarios lista todos os usuarios de um evento pelo id do evento
func (e *Eventos) ListaUsuarios(evento Evento) []usuario.Usuario {
	participantes := append(evento.Participantes, evento.Criador)
	return participantes
}

// ListaProdutos lista todos os produtos de um evento pelo id do evento
func (e *Eventos) ListaProdutos(participantes []usuario.Usuario, produtos []produto.Produto) []produto.Produto {
	var lista []produto.Produto
	for _, participante := range participantes {
		for _, produto := range produtos {
			if participante.ID == produto.Dono.ID {
				lista = append(lista, produto)
			}
		}
	}
	return lista
}

// Exclui pelo id do evento
func (e *Eventos) Exclui() {

}

// Altera pelo id do evento
func (e *Eventos) Altera() {

}

// EncerraEvento pelo id do evento
func (e *Eventos) EncerraEvento() {

}
