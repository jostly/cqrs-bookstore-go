package server

type BookDTO struct {
	BookId      string `json:"bookId"`
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductDTO struct {
	ProductId           string `json:"productId"`
	Price               uint64 `json:"price"`
	PublisherContractId string `json:"publisherContractId"`
}
