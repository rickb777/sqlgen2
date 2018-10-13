package require

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestWrongSizeZero(t *testing.T) {
	RegisterTestingT(t)

	err := ErrWrongSize(0, "wanted %d", 5)
	Ω(err.Error()).Should(Equal("wanted 5"))
	Ω(err.(Sizer).Size()).Should(BeEquivalentTo(0))
	Ω(IsNotFound(err)).Should(BeTrue())
	Ω(IsNotUnique(err)).Should(BeFalse())
}

func TestWrongSizeMany(t *testing.T) {
	RegisterTestingT(t)

	err := ErrWrongSize(3, "wanted %d", 5)
	Ω(err.Error()).Should(Equal("wanted 5"))
	Ω(err.(Sizer).Size()).Should(BeEquivalentTo(3))
	Ω(IsNotFound(err)).Should(BeFalse())
	Ω(IsNotUnique(err)).Should(BeTrue())
}