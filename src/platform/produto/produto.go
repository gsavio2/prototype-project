package produto

// Produto struct modelo
type Produto struct {
	ID        string `json: "id"`
	Descricao string `json: "descricao"`
}

// Produtos struct modelo
type Produtos struct {
	Produtos []Produto
}
