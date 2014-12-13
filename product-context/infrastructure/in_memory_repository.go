package infrastructure

import (
	"errors"
	"github.com/jostly/cqrs-bookstore-go/product-context/domain"
)

type inMemoryRepository struct {
	productMap map[string]domain.Product
}

func NewInMemoryRepository() *inMemoryRepository {
	r := new(inMemoryRepository)
	r.productMap = make(map[string]domain.Product)
	return r
}

func (r *inMemoryRepository) GetProduct(productId string) (product domain.Product, err error) {
	var ok bool

	product, ok = r.productMap[productId]
	if !ok {
		err = errors.New("No product with id '" + productId + "' exists")
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

	return v
}
