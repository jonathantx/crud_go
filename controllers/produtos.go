package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jonathantx/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)

}

func RegisterProduct(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Cadastro", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		models.CreateProduct(nome, descricao, precoConvertido, quantidadeConvertida)

		http.Redirect(w, r, "/", 301)
	}
}
