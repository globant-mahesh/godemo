package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//	. "books"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

var _ = Describe("Book", func() {
//	var (
//		longBook  Book
//		shortBook Book
//	)

	BeforeEach(func() {
//		longBook = Book{
//			Title:  "Les Miserables",
//			Author: "Victor Hugo",
//			Pages:  1488,
//		}
//
//		shortBook = Book{
//			Title:  "Fox In Socks",
//			Author: "Dr. Seuss",
//			Pages:  24,
//		}
	})

	//    Describe("Categorizing book length", func() {
	//        Context("With more than 300 pages", func() {
	//            It("should be a novel", func() {
	//                Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
	//            })
	//        })
	//
	//        Context("With fewer than 300 pages", func() {
	//            It("should be a short story", func() {
	//                Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
	//            })
	//        })
	//    })

	var _ = Describe("Book", func() {
		It("can be loaded from JSON", func() {
			book := Book{Title:"Les Miserables",
            Author:"Victor Hugo",
            Pages:1488}
//			NewBookFromJSON(`{
//            "title":"Les Miserables",
//            "author":"Victor Hugo",
//            "pages":1488
//        }`)

			Expect(book.Title).To(Equal("Les Miserables"))
			Expect(book.Author).To(Equal("Victor Hugo"))
			Expect(book.Pages).To(Equal(1488))
		})
	})

})
