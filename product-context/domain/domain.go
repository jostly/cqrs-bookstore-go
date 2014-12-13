package domain

type ProductId string
type BookId string

type Book struct {
	BookId      BookId
	ISBN        string
	Title       string
	Description string
}

type Product struct {
	ProductId           ProductId
	Book                Book
	Price               uint64
	PublisherContractId string
}

type Repository interface {
	GetProducts() []Product
	GetProduct(productId ProductId) (Product, error)
	StoreProduct(product Product)
}
