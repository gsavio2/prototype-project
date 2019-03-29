package evento

import (
	"time"

	"github.com/pedrohmachado/prototype-project/src/platform/produto"

	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Evento struct modelo
type Evento struct {
	// definir quais serão os status do evento
	ID            string            `json: "id"`
	Descricao     string            `json: "descricao"`
	Criador       usuario.Usuario   `json: "criador"`
	Participantes []usuario.Usuario `json: "participantes"`
	DataCriacao   string            `json: "dataCriacao"`
	DataEvento    string            `json: "dataEvento"`
	Status        string            `json: "status"`
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
	evento.Participantes = append(evento.Participantes, evento.Criador)
	evento.Status = "ativo"
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

// ListaTodosDono lista todos os eventos que o usuario é dono pelo id do usuario
func (e *Eventos) ListaTodosDono(usuario usuario.Usuario) []Evento {
	var eventos []Evento
	for _, evento := range e.Eventos {
		if evento.Criador.ID == usuario.ID {
			eventos = append(eventos, evento)
		}
	}
	return eventos
}

// ListaTodosParticipante lista todos os eventos que um usuario participa pelo id do usuario
func (e *Eventos) ListaTodosParticipante(usuario usuario.Usuario) []Evento {
	var eventos []Evento
	for _, evento := range e.Eventos {
		for _, participante := range evento.Participantes {
			if participante.ID == usuario.ID {
				eventos = append(eventos, evento)
			}
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
func (e *Eventos) Exclui(id string) []Evento {
	for index, evento := range e.Eventos {
		if evento.ID == id {
			e.Eventos = append(e.Eventos[:index], e.Eventos[index+1:]...)
			return e.Eventos
		}
	}
	return e.Eventos
}

// AdicionaParticipante adiciona usuario participante no evento
func (e *Eventos) AdicionaParticipante(participante usuario.Usuario, evento Evento) Evento {
	evento.Participantes = append(evento.Participantes, participante)
	return evento
}

// RemoveParticipante remove usuario participante do evento
func (e *Eventos) RemoveParticipante(usuario usuario.Usuario, evento Evento) Evento {
	for index, participante := range evento.Participantes {
		if participante == usuario {
			evento.Participantes = append(evento.Participantes[:index], evento.Participantes[index+1:]...)
		}
	}
	return evento
}

// AtualizaUsuario atualiza a lista de eventos após a alteração de um usuario
func (e *Eventos) AtualizaUsuario(usuario usuario.Usuario, evento Evento) Evento {

	// altera criador (pelo id) do evento
	if usuario.ID == evento.Criador.ID {
		evento.Criador = usuario
	}

	// altera participante (pelo id) do evento
	for index, participante := range evento.Participantes {
		if participante.ID == usuario.ID {
			//exclui antigo registro de participante
			evento.Participantes = append(evento.Participantes[:index], evento.Participantes[index+1:]...)
		}
	}
	//adiciona novo registro de participante com o mesmo id do antigo
	evento.Participantes = append(evento.Participantes, usuario)
	return evento
}

// Atualiza atualiza a lista de eventos (passando o evento atualizado por parametro)
func (e *Eventos) Atualiza(evento Evento) []Evento {
	id := evento.ID
	e.Eventos = e.Exclui(id)
	e.Eventos = append(e.Eventos, evento)
	return e.Eventos
}

// Altera pelo id do evento
func (e *Eventos) Altera(evento Evento, id string) Evento {
	// resgata evento que será alterado
	auxEvento := e.Lista(evento.ID)
	// atributos imutáveis nessa alteração
	evento.ID = id
	evento.DataCriacao = auxEvento.DataCriacao
	evento.Criador = auxEvento.Criador
	evento.Participantes = auxEvento.Participantes
	return evento
}

// Encerra evento pelo id do evento
func (e *Eventos) Encerra(evento Evento) Evento {
	evento.Status = "encerrado"
	return evento
}
