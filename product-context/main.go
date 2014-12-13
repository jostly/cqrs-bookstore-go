package main

import (
	"github.com/jostly/cqrs-bookstore-go/cqrslib"
	"github.com/jostly/cqrs-bookstore-go/product-context/server"
	"net/http"
	"os"
)

func main() {
	cqrslib.Greet()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server.RegisterHandlers()

	http.ListenAndServe("127.0.0.1:"+port, nil)
}
