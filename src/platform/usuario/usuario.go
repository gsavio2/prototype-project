package usuario

import (
	"github.com/pedrohmachado/prototype-project/src/platform/produto"
)

// ListaTodos lista todos os usuarios
type ListaTodos interface {
	ListaTodos() []Usuario
}

// Lista um usuario
type Lista interface {
	Lista(id string) Usuario
}

// Altera usuario
type Altera interface {
	Altera(usuario Usuario, id string) Usuario
}

// Adiciona cria novo usuario
type Adiciona interface {
	Adiciona(usuario Usuario)
}

// Exclui usuario
type Exclui interface {
	Exclui(id string) []Usuario
}

// Usuario struct modelo
type Usuario struct {
	ID       string            `json: "id"`
	Nome     string            `json: "nome"`
	Email    string            `json: "email"`
	Produtos []produto.Produto `json: "produtos"`
}

// Usuarios struct modelo
type Usuarios struct {
	Usuarios []Usuario
}

// Novo inicia usuarios
func Novo() *Usuarios {
	return &Usuarios{
		Usuarios: []Usuario{},
	}
}

// ListaTodos retorna slice de todos os usuarios
func (u *Usuarios) ListaTodos() []Usuario {
	return u.Usuarios
}

// Lista retorna um usuario de todos os usuarios
func (u *Usuarios) Lista(id string) Usuario {
	for _, usuario := range u.Usuarios {
		if usuario.ID == id {
			return usuario
		}
	}
	return Usuario{}
}

// Adiciona adiciona usuario a usuarios
func (u *Usuarios) Adiciona(usuario Usuario) {
	u.Usuarios = append(u.Usuarios, usuario)
}

// Exclui usuario de usuarios
func (u *Usuarios) Exclui(id string) []Usuario {
	for index, usuario := range u.Usuarios {
		if usuario.ID == id {
			u.Usuarios = append(u.Usuarios[:index], u.Usuarios[index+1:]...)
			return u.Usuarios
		}
	}
	return u.Usuarios
}

// Altera usuario
func (u *Usuarios) Altera(usuario Usuario, id string) Usuario {
	usuario.ID = id
	u.Usuarios = u.Exclui(id)
	u.Usuarios = append(u.Usuarios, usuario)
	return usuario
}

// ##########################################################################
// -----------------------------PRODUTOS-------------------------------------
// ------------- melhorar ---------------------------------------------------
// ##########################################################################

// AdicionaProduto a usuario pelo id do usuario
func (u *Usuarios) AdicionaProduto(produto produto.Produto, usuario Usuario, id string) Usuario {
	usuario.ID = id
	u.Usuarios = u.Exclui(id)
	usuario.Produtos = append(usuario.Produtos, produto)
	u.Usuarios = append(u.Usuarios, usuario)
	return usuario
}

// ListaTodosProduto lista todos os produtos do usuario pelo id
func (u *Usuarios) ListaTodosProduto(id string) []produto.Produto {
	usuario := u.Lista(id)
	return usuario.Produtos
}

// ExcluiProduto de usuario pelo id do usuario e id do produto
func (u *Usuarios) ExcluiProduto(usuario Usuario, idUsuario, idProduto string) Usuario {
	usuario.ID = idUsuario
	u.Usuarios = u.Exclui(idUsuario)
	for index, produto := range usuario.Produtos {
		if produto.ID == idProduto {
			usuario.Produtos = append(usuario.Produtos[:index], usuario.Produtos[index+1:]...)
		}
	}
	u.Usuarios = append(u.Usuarios, usuario)
	return usuario
}

// ListaProduto pelo id do produto
func (u *Usuarios) ListaProduto(produtos []produto.Produto, id string) produto.Produto {
	for _, produto := range produtos {
		if produto.ID == id {
			return produto
		}
	}
	return produto.Produto{}
}

// AlteraProduto pelo id do usuario e do produto
func (u *Usuarios) AlteraProduto(usuario Usuario, produto produto.Produto, idUsuario, idProduto string) Usuario {
	usuario.ID = idUsuario
	produto.ID = idProduto
	u.Usuarios = u.Exclui(idUsuario)
	for index, produtoItem := range usuario.Produtos {
		if produtoItem.ID == idProduto {
			usuario.Produtos = append(usuario.Produtos[:index], usuario.Produtos[index+1:]...)
		}
	}
	usuario.Produtos = append(usuario.Produtos, produto)
	u.Usuarios = append(u.Usuarios, usuario)
	return usuario
}

// ##########################################################################
// -----------------------------EVENTOS--------------------------------------
// ##########################################################################
