package server

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hariharan-uno/cors"
	"gopkg.in/unrolled/render.v1"
	"gopkg.in/validator.v1"

	"github.com/jostly/cqrs-bookstore-go/product-context/domain"
	"github.com/jostly/cqrs-bookstore-go/product-context/infrastructure"
)

var (
	renderer = render.New(render.Options{IndentJSON: true})
	repo     = infrastructure.NewInMemoryRepository()
)

func RegisterHandlers() {

	r := mux.NewRouter()

	r.HandleFunc("/products/{productId}", GetProduct)
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products", GetProducts)

	n := negroni.Classic()

	options := cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}

	n.Use(negroni.HandlerFunc(options.Allow))
	n.UseHandler(r)

	http.Handle("/", n)
}

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var dto ProductDTO
	err := decoder.Decode(&dto)

	if err != nil {
		panic(err)
	}

	valid, errs := validator.Validate(dto)
	if !valid {
		panic(errs)
	}

	product := dto.ToDomain()

	repo.StoreProduct(product)
}

func GetProducts(w http.ResponseWriter, req *http.Request) {
	products := repo.GetProducts()
	dtos := make([]ProductDTO, 0, len(products))
	for _, v := range products {
		dtos = append(dtos, NewProductDTO(v))
	}
	renderer.JSON(w, http.StatusOK, dtos)
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productId := vars["productId"]

	product, err := repo.GetProduct(domain.ProductId(productId))
	if err != nil {
		panic(err)
	}

	renderer.JSON(w, http.StatusOK, NewProductDTO(product))
}
