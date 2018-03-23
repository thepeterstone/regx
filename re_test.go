package builder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/thepeterstone/regx"
)

var _ = Describe("CharacterExpression", func() {
	var e CharacterExpression

	BeforeEach(func() {
		e = CharacterExpression{}
	})

	It("matches a single entry", func() {
		e.AddRune('a')

		Expect(e.Match('a')).To(BeTrue())
	})

	It("matches each input", func() {
		e.AddRune('a')
		e.AddRune('b')
		e.AddRune('c')

		Expect(e.Match('a')).To(BeTrue())
		Expect(e.Match('b')).To(BeTrue())
		Expect(e.Match('c')).To(BeTrue())
	})

	It("condenses a series of digits", func() {
		for _, i := range []rune{'1', '2', '3', '4', '5'} {
			e.AddRune(i)
		}

		Expect(e.Match('6')).To(BeTrue())
		Expect(e.Match('7')).To(BeTrue())
		Expect(e.Match('a')).To(BeFalse())
	})

	It("matches hex digits", func() {
		for _, i := range []rune{'1', 'a', '4', 'f', 'e', '6'} {
			e.AddRune(i)
		}

		Expect(e.Match('1')).To(BeTrue())
		Expect(e.Match('B')).To(BeTrue())
		Expect(e.Match('g')).To(BeFalse())
	})

	It("matches hex digits added after decimal digits", func() {
		for _, i := range []rune{'1', '2', '3', '4', '5', 'e', 'A'} {
			e.AddRune(i)
		}

		Expect(e.Match('1')).To(BeTrue())
		Expect(e.Match('B')).To(BeTrue())
		Expect(e.Match('g')).To(BeFalse())
	})

	It("incrementally loosens its scope", func() {
		for _, i := range []rune{'1', '2', '3', '4', '5'} {
			e.AddRune(i)
		}

		Expect(e.Match('6')).To(BeTrue())
		Expect(e.Match('c')).To(BeFalse())
		Expect(e.Match('w')).To(BeFalse())
		Expect(e.Match('#')).To(BeFalse())

		e.AddRune('a') // upgrade to xdigit
		Expect(e.Match('6')).To(BeTrue())
		Expect(e.Match('c')).To(BeTrue())
		Expect(e.Match('w')).To(BeFalse())
		Expect(e.Match('#')).To(BeFalse())

		e.AddRune('M') // upgrade to \w
		Expect(e.Match('6')).To(BeTrue())
		Expect(e.Match('c')).To(BeTrue())
		Expect(e.Match('w')).To(BeTrue())
		Expect(e.Match('#')).To(BeFalse())

		e.AddRune('#') // upgrade to .
		Expect(e.Match('6')).To(BeTrue())
		Expect(e.Match('c')).To(BeTrue())
		Expect(e.Match('w')).To(BeTrue())
		Expect(e.Match('#')).To(BeTrue())
	})
})

var _ = Describe("Expression", func() {
	var e Expression

	BeforeEach(func() {
		e = Expression{}
	})

	It("matches a single entry", func() {
		e.AddLine("a")

		Expect(e.Match("a")).To(BeTrue())
	})

	It("matches each input", func() {
		e.AddLine("a")
		e.AddLine("b")
		e.AddLine("c")

		Expect(e.Match("a")).To(BeTrue())
		Expect(e.Match("b")).To(BeTrue())
		Expect(e.Match("c")).To(BeTrue())
	})

	It("condenses a series of digits", func() {
		for _, i := range []string{"1", "2", "3", "4", "5"} {
			e.AddLine(i)
		}

		Expect(e.Match("6")).To(BeTrue())
		Expect(e.Match("c")).To(BeFalse())
	})

	It("matches hex digits", func() {
		for _, i := range []string{"1", "a", "4", "f", "e", "6"} {
			e.AddLine(i)
		}

		Expect(e.Match("1")).To(BeTrue())
		Expect(e.Match("B")).To(BeTrue())
		Expect(e.Match("g")).To(BeFalse())
	})

	It("matches hex digits added after decimal digits", func() {
		e := Expression{}

		for _, i := range []string{"1", "2", "3", "4", "5", "e", "A"} {
			e.AddLine(i)
		}

		Expect(e.Match("1")).To(BeTrue())
		Expect(e.Match("B")).To(BeTrue())
		Expect(e.Match("g")).To(BeFalse())
	})

	It("matches two-digit numbers", func() {
		e.AddLine("10")
		e.AddLine("11")
		e.AddLine("12")

		Expect(e.Match("10")).To(BeTrue())
		Expect(e.Match("11")).To(BeTrue())
		Expect(e.Match("12")).To(BeTrue())
	})

	It("condenses a series of two-digit numbers", func() {
		for _, i := range []string{"11", "12", "13", "14"} {
			e.AddLine(i)
		}

		Expect(e.Match("15")).To(BeTrue())
		Expect(e.Match("6")).To(BeFalse())
	})

	It("matches dates", func() {
		e := Expression{}

		for _, i := range []string{"2011-01-01", "2014-09-17", "2019-04-23", "2010-12-31", "2009-06-24"} {
			e.AddLine(i)
			Expect(e.Match(i)).To(BeTrue())
		}

		Expect(e.Match("2018-03-12")).To(BeTrue())
		Expect(e.Match("hello bob!")).To(BeFalse())
		Expect(e.Match("2018-95-42")).To(BeFalse())
	})

})
