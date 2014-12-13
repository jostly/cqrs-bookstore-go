package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hariharan-uno/cors"
	"gopkg.in/unrolled/render.v1"
)

var renderer = render.New(render.Options{IndentJSON: true})

func RegisterHandlers() {

	router := mux.NewRouter()

	router.HandleFunc("/products", GetProducts)

	n := negroni.Classic()

	options := cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}

	n.Use(negroni.HandlerFunc(options.Allow))
	n.UseHandler(router)

	http.Handle("/", n)
}

func GetProducts(w http.ResponseWriter, req *http.Request) {
	renderer.JSON(w, http.StatusOK, ProductDTO{"foo", 123, "vvd"})
}
