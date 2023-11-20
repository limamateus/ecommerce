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

func CriarNovoProduto(nome, descricao string, quantidade int, preco float64) {
	db := db.ConectandoComBanco()

	insertProduto, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade)values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectandoComBanco()

	deletarProtudo, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deletarProtudo.Exec(id)

	defer db.Close()
}

func BuscarProduto(id string) Produto {
	db := db.ConectandoComBanco() // Abro a conexão

	selectProduto, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}
	p := Produto{}

	for selectProduto.Next() {
		// Aqui estou criando variavel locals que irão servir de referente
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

	}
	defer db.Close()
	return p

}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {

	db := db.ConectandoComBanco()

	atualizarProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizarProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()

}
