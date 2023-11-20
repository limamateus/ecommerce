package controller

import (
	"main/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // variavel que irá receber dados de templates

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaProdutos() // Aqui estou pegando os produtos da função que estão na model de produtos que retorna uma lista de produtos

	temp.ExecuteTemplate(w, "index", todosOsProdutos) // aqui eu passo os dados para index

}
