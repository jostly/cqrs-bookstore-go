package server

import (
	"github.com/jostly/cqrs-bookstore-go/product-context/domain"
)

type BookDTO struct {
	BookId      string `json:"bookId" validate:"nonzero,regexp=^(([0-9a-fA-F]){8}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){12})$"`
	ISBN        string `json:"isbn" validate:"nonzero"`
	Title       string `json:"title" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
}

type ProductDTO struct {
	ProductId           string  `json:"productId" validate:"nonzero,regexp=^(([0-9a-fA-F]){8}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){12})$"`
	Book                BookDTO `json:"book" validate:"nonzero"`
	Price               uint64  `json:"price" validate:"min=0"`
	PublisherContractId string  `json:"publisherContractId" validate:"nonzero,regexp=^(([0-9a-fA-F]){8}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){12})$"`
}

func (dto *BookDTO) ToDomain() domain.Book {
	return domain.Book{
		BookId:      domain.BookId(dto.BookId),
		ISBN:        dto.ISBN,
		Title:       dto.Title,
		Description: dto.Description,
	}
}

func (dto *ProductDTO) ToDomain() domain.Product {
	return domain.Product{
		ProductId:           domain.ProductId(dto.ProductId),
		Book:                dto.Book.ToDomain(),
		Price:               dto.Price,
		PublisherContractId: dto.PublisherContractId,
	}
}

func newBookDTO(book domain.Book) BookDTO {
	return BookDTO{
		BookId:      string(book.BookId),
		ISBN:        book.ISBN,
		Title:       book.Title,
		Description: book.Description,
	}
}

func NewProductDTO(product domain.Product) ProductDTO {
	return ProductDTO{
		ProductId:           string(product.ProductId),
		Book:                newBookDTO(product.Book),
		Price:               product.Price,
		PublisherContractId: product.PublisherContractId,
	}
}
