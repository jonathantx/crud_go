package models

import (
	"github.com/jonathantx/loja/db"
)

// Struct: Coleção de variaveis que formam novo tipo

// Estrutura de produtos

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProducts() []Produto {

	db := db.ConnectionDataBase()

	selectProducts, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos

}

func CreateProduct(nome, descricao string, preco float64, quantidade int) {

	db := db.ConnectionDataBase()

	insertProducts, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProducts.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
