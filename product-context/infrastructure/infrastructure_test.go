package infrastructure_test

import (
	"github.com/jostly/cqrs-bookstore-go/product-context/domain"
	"github.com/jostly/cqrs-bookstore-go/product-context/infrastructure"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("In Memory Repository", func() {

	var repo domain.Repository

	BeforeEach(func() {
		repo = infrastructure.NewInMemoryRepository()
	})

	Describe("Getting a product", func() {
		Context("With no matching product", func() {
			It("should return an error", func() {
				_, err := repo.GetProduct("foo")
				Ω(err).Should(HaveOccurred())
			})
		})
		Context("With a matching product", func() {
			It("should return the product", func() {
				expected := domain.Product{ProductId: "foo"}
				repo.StoreProduct(expected)

				Ω(repo.GetProduct("foo")).Should(Equal(expected))
			})
		})
	})

	Describe("Getting all products", func() {
		Context("With no products", func() {
			It("should return an empty array of products", func() {
				Ω(repo.GetProducts()).Should(Equal(make([]domain.Product, 0)))
			})
		})
		Context("With one product", func() {
			It("should return an array with one entry", func() {
				expected := domain.Product{ProductId: "foo"}
				repo.StoreProduct(expected)

				Ω(repo.GetProducts()).Should(ConsistOf(expected))
			})
		})
		Context("With multiple products", func() {
			It("should return an array of all entries", func() {
				expected1 := domain.Product{ProductId: "foo"}
				expected2 := domain.Product{ProductId: "bar"}
				repo.StoreProduct(expected1)
				repo.StoreProduct(expected2)

				Ω(repo.GetProducts()).Should(ConsistOf(expected1, expected2))
			})
		})
	})

})
