package controller

import (
	"log"
	"main/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // variavel que irá receber dados de templates

func Index(w http.ResponseWriter, r *http.Request) { // EndPoint do Formulario da Pagina principal

	todosOsProdutos := models.BuscaProdutos() // Aqui estou pegando os produtos da função que estão na model de produtos que retorna uma lista de produtos

	temp.ExecuteTemplate(w, "index", todosOsProdutos) // aqui eu passo os dados para index

}

func New(w http.ResponseWriter, r *http.Request) { // Endpoint do formulario de novo produto
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // Verifico se o metodo é Post
		// Aqui estou armazenando os valores dos inputs em uma variavel
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		// Realizo a conversão
		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro ao converter preco !")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter quantidade !")
		}
		// passo os dados para meu modelo
		models.CriarNovoProduto(nome, descricao, quantidadeConvertida, precoConvertido)

	}

	http.Redirect(w, r, "/", 301) // redireciono para tela da pagina inicial
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id") // Aqui estou pegando o id do produto que está contido na url

	models.DeletarProduto(idDoProduto) // depois passo para metodo DeletarProduto o id que eu peguei

	http.Redirect(w, r, "/", 301) // redirenciono para tela inicial

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")

	p := models.BuscarProduto(idDoProduto)

	temp.ExecuteTemplate(w, "Edit", p)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			panic(err.Error())
		}

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			panic(err.Error())
		}

		models.AtualizarProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}
