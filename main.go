package main

import (
	"html/template"
	"main/models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaProdutos() // Aqui estou pegando os produtos da função que estão na model de produtos que retorna uma lista de produtos

	temp.ExecuteTemplate(w, "index", todosOsProdutos) // aqui eu passo os dados para index

}
