package models

import "web-application/db"

type Produto struct {
	id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

// BuscarTodosOsProdutos retorna todos os produtos da tabela 'produtos'.
func BuscarTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	//exibir produtos do banco de dados
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		// Lê os dados da linha atual para o struct Produto
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
