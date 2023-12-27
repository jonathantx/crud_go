package models

import (
	"fmt"

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

	selectProducts, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")

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

func Delete(id int) {

	db := db.ConnectionDataBase()

	deleteProduct, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Produto {

	db := db.ConnectionDataBase()

	produto, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	p := Produto{}

	for produto.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produto.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

	}

	defer db.Close()

	return p
}

func Update(id, nome, descricao string, preco float64, quantidade int) {

	db := db.ConnectionDataBase()

	update, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()

}
