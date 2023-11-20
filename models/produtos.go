package models

import "main/db"

type Produto struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func BuscaProdutos() []Produto {
	db := db.ConectandoComBanco()
	selectDeProdutos, err := db.Query("select * from produtos")
	produtos := []Produto{}
	p := Produto{}
	if err != nil {
		panic(err.Error())
	}
	for selectDeProdutos.Next() {
		// Aqui estou criando variavel locals que irão servir de referente
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		// Verifico se tem erro e caso tenho falo o erro
		if err != nil {
			panic(err.Error())
		}
		// atribui para variavel as informações que vão vim do banco
		p.Id = id
		p.Nome = nome
		p.Preco = preco
		p.Quantidade = quantidade
		p.Descricao = descricao
		// e adiciono na minha lista de produtos.
		produtos = append(produtos, p)

	}
	defer db.Close() // fechou a conexão caso termine tudo.
	return produtos  // Retorno os produtos da lista

}
