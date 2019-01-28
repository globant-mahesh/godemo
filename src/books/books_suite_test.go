package books_test

import (
	"testing"
	//a command line interface with plenty of handy command line arguments for running your tests and generating test files.
	. "github.com/onsi/ginkgo"
	//	Ginkgo's Preferred Matcher Library
	. "github.com/onsi/gomega"
)

//testing test
func TestBooks(t *testing.T) {
	//	A Ginkgo test signals failure by calling Ginkgo’s Fail(description string) function. We pass this function to Gomega using RegisterFailHandler. This is the sole connection point between Ginkgo and Gomega.
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
