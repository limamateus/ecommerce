package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectandoComBanco() *sql.DB {
	stringDeConexao := "user=postgres dbname=Alura_Loja password=root host=localhost sslmode=disable" // Varivel que irá armazena a string de conecção com banco

	db, err := sql.Open("postgres", stringDeConexao) // aqui estou abrindo minha conexão, caso de erro  a mesangem vai para variavel err caso de certo vai para varival db

	if err != nil { // aqui estou verificando se deu erro
		panic(err.Error()) // retornando a mensagem do erro.
	}

	return db // caso de tudo certo eu retorno

}

type Produto struct {
	id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectandoComBanco()
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
		p.id = id
		p.Nome = nome
		p.Preco = preco
		p.Quantidade = quantidade
		p.Descricao = descricao
		// e adiciono na minha lista de produtos.
		produtos = append(produtos, p)

	}

	temp.ExecuteTemplate(w, "index", produtos) // aqui eu passo os dados para index

	defer db.Close() // fechou a conexão caso termine tudo.
}
