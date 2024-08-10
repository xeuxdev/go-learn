package main

import (
	"net/http"

	"github.com/go-learn/api-calc/internals/handlers"
)

func main() {

	server := &http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("POST /api/calc/add", handlers.Add)
	http.HandleFunc("POST /api/calc/sub", handlers.Subtract)
	http.HandleFunc("POST /api/calc/mul", handlers.Multiply)
	http.HandleFunc("POST /api/calc/div", handlers.Divide)

	server.ListenAndServe()

}
