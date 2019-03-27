package usuario

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
	ID    string `json: "id"`
	Nome  string `json: "nome"`
	Email string `json: "email"`
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
