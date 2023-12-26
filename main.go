package main

import (
	"net/http"

	"github.com/jonathantx/loja/routes"
)

func main() {

	routes.LoadingRoutes()
	http.ListenAndServe(":8000", nil)
}
