package domain

type Book struct {
	BookId      string
	ISBN        string
	Title       string
	Description string
}

type Product struct {
	ProductId           string
	Book                Book
	Price               uint64
	PublisherContractId string
}

type Repository interface {
	GetProducts() []Product
	GetProduct(productId string) (Product, error)
	StoreProduct(product Product)
}
