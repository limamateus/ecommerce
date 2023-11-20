package routes

import (
	"main/controller"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controller.Index)  // Rota da pagina principal
	http.HandleFunc("/New", controller.New) // Reta para formulario de novo produto
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
}
