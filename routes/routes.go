package routes

import (
	"net/http"

	"github.com/jonathantx/loja/controllers"
)

func LoadingRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/cadastrar-produto", controllers.RegisterProduct)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
