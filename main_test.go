package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/alextanhongpin/go-ginkgo-test"
)

var _ = Describe("Main", func() {
	var book Book

	BeforeEach(func() {
		book = Book{
			Title:  "A Dream",
			Author: "John Doe",
			Pages:  1000,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 1000 pages", func() {
			It("Should be a long book", func() {
				Expect(book.CategoryByLength()).To(Equal("long"))
			})

			It("Should have the correct fields", func() {
				Expect(book.Title).To(Equal("A Dream"))
				Expect(book.Author).To(Equal("John Doe"))
				Expect(book.Pages).To(Equal(1000))
			})
		})
	})
})
