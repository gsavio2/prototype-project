package produto

import (
	"github.com/pedrohmachado/prototype-project/src/platform/usuario"
)

// Produto struct modelo
type Produto struct {
	ID        string          `json: "id"`
	Descricao string          `json: "descricao"`
	Dono      usuario.Usuario `json: "dono"`
	Status    string          `json: "status"`
}

// Produtos struct modelo
type Produtos struct {
	Produtos []Produto
}

// Novo inicia produtos
func Novo() *Produtos {
	return &Produtos{
		Produtos: []Produto{},
	}
}

// Adiciona produto a produtos
func (p *Produtos) Adiciona(produto Produto, usuario usuario.Usuario) []Produto {
	produto.Dono = usuario
	produto.Status = "ativo"
	p.Produtos = append(p.Produtos, produto)
	return p.Produtos
}

// ListaTodos lista todos os produtos
func (p *Produtos) ListaTodos() []Produto {
	return p.Produtos
}

// Lista lista produto pelo id
func (p *Produtos) Lista(id string) Produto {
	for _, produto := range p.Produtos {
		if produto.ID == id {
			return produto
		}
	}
	return Produto{}
}

// ListaTodosUsuario lista todos os produtos de um usuario pelo id do usuario
func (p *Produtos) ListaTodosUsuario(usuario usuario.Usuario) []Produto {
	var produtos []Produto
	for _, produto := range p.Produtos {
		if produto.Dono.ID == usuario.ID {
			produtos = append(produtos, produto)
		}
	}
	return produtos
}

// Exclui exclui produto pelo id (e retorna o produto excluído)
func (p *Produtos) Exclui(id string) []Produto {
	for index, produto := range p.Produtos {
		if produto.ID == id {
			p.Produtos = append(p.Produtos[:index], p.Produtos[index+1:]...)
			return p.Produtos
		}
	}
	return p.Produtos
}

// Altera altera produto pelo id (e retorna o produto alterado)
func (p *Produtos) Altera(produto Produto, id string) Produto {
	p.Produtos = p.Exclui(id)
	produto.ID = id
	p.Produtos = append(p.Produtos, produto)
	return produto
}

// AtualizaUsuario atualiza os dados do usuario (dono do produto)
func (p *Produtos) AtualizaUsuario(usuario usuario.Usuario, produto Produto) Produto {
	produto.Dono = usuario
	return produto
}

// Atualiza atualiza a lista dos produtos
func (p *Produtos) Atualiza(produto Produto) []Produto {
	id := produto.ID
	p.Produtos = p.Exclui(id)
	p.Produtos = append(p.Produtos, produto)
	return p.Produtos
}
