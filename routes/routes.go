package routes

import (
	"main/controller"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controller.Index)
}
