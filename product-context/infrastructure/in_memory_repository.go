package infrastructure

import (
	"errors"
	"github.com/jostly/cqrs-bookstore-go/product-context/domain"
)

type inMemoryRepository struct {
	productMap map[domain.ProductId]domain.Product
}

func NewInMemoryRepository() *inMemoryRepository {
	r := new(inMemoryRepository)
	r.productMap = make(map[domain.ProductId]domain.Product)
	return r
}

func (r *inMemoryRepository) GetProduct(productId domain.ProductId) (product domain.Product, err error) {
	var ok bool

	product, ok = r.productMap[productId]
	if !ok {
		err = errors.New("No product with id '" + string(productId) + "' exists")
	}
	return
}

func (r *inMemoryRepository) StoreProduct(product domain.Product) {
	r.productMap[product.ProductId] = product
}

func (r *inMemoryRepository) GetProducts() []domain.Product {
	v := make([]domain.Product, 0, len(r.productMap))
	for _, value := range r.productMap {
		v = append(v, value)
	}
	// TODO this should be sorted according to book title, but that was complex :)

	return v
}
