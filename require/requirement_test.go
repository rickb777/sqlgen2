package require

import (
	. "github.com/onsi/gomega"
	"testing"
	"fmt"
)

var e0 = fmt.Errorf("foo")

func TestChainErrorIfQueryNotSatisfiedBy_and_ChainErrorIfExecNotSatisfiedBy(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		input    error
		req      Requirement
		actual   int64
		expected bool
		qMessage string
		eMessage string
	}{
		{nil, nil, 0, false, "", ""},
		{e0, nil, 0, true, "foo", "foo"},

		{nil, None, 0, false, "", ""},
		{nil, None, 5, true, "expected to fetch none but got 5", "expected to change none but changed 5"},
		{nil, One, 5, true, "expected to fetch one but got 5", "expected to change one but changed 5"},
		{nil, Many, 5, false, "", ""},
		{nil, Many, 1, true, "expected to fetch many but got 1", "expected to change many but changed 1"},

		{nil, Exactly(3), 3, false, "", ""},
		{nil, Exactly(3), 2, true, "expected to fetch 3 but got 2", "expected to change 3 but changed 2"},

		{nil, AtLeast(3), 5, false, "", ""},
		{nil, AtLeast(3), 2, true, "expected to fetch at least 3 but got 2", "expected to change at least 3 but changed 2"},

		{nil, NoMoreThan(3), 2, false, "", ""},
		{nil, NoMoreThan(3), 5, true, "expected to fetch no more than 3 but got 5", "expected to change no more than 3 but changed 5"},
	}

	for _, c := range cases {
		e1 := ChainErrorIfQueryNotSatisfiedBy(c.input, c.req, c.actual)
		if c.expected {
			Ω(e1.Error()).Should(Equal(c.qMessage))
		} else {
			Ω(e1).Should(BeNil())
		}

		e2 := ChainErrorIfExecNotSatisfiedBy(c.input, c.req, c.actual)
		if c.expected {
			Ω(e2.Error()).Should(Equal(c.eMessage))
		} else {
			Ω(e2).Should(BeNil())
		}
	}
}
