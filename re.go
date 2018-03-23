package builder

import (
	"fmt"
	"regexp"
	"strings"
)

// Expression defines a regexp that evolves
type Expression struct {
	atoms   []CharacterExpression
	pattern *regexp.Regexp
}

// AddLine inserts candidates into an Expression
func (e *Expression) AddLine(line string) {
	if !e.Match(line) {
		for i, c := range []rune(line) {
			if i >= len(e.atoms) {
				e.atoms = append(e.atoms, newCharacterExpression(c))
			}
			e.atoms[i].AddRune(c)
		}
		var s []string
		for _, a := range e.atoms {
			s = append(s, a.Pattern())
		}
		e.pattern = regexp.MustCompile(strings.Join(s, ""))
	}
}

// String returns a printable form of the Expression
func (e *Expression) String() string {
	if e.pattern == nil {
		return fmt.Sprintf("%+v", e)
	}
	return e.pattern.String()
}

// Match checks whether the input satisfies the Expression
func (e *Expression) Match(input string) bool {
	if e.pattern != nil && e.pattern.MatchString(input) {
		return true
	}
	for i, a := range []rune(input) {
		if len(e.atoms) <= i || !e.atoms[i].Match(a) {
			return false
		}
	}
	return true
}

// CharacterExpression defines a regexp that matches against
// a single-character input
type CharacterExpression struct {
	pattern  *regexp.Regexp
	literals []rune
	atoms    []rune
}

func newCharacterExpression(c rune) CharacterExpression {
	return CharacterExpression{
		pattern:  nil,
		literals: nil,
		atoms:    []rune{c},
	}
}

// Match checks whether the input satisfies the CharacterExpression
func (e *CharacterExpression) Match(input rune) bool {
	if e.pattern != nil && allMatch(e.pattern, []rune{input}) {
		return true
	}
	for _, c := range e.literals {
		if c == input {
			return true
		}
	}
	return false
}

// AddRune adds a candidate to a CharacterExpression
func (e *CharacterExpression) AddRune(r rune) {
	if e.Match(r) {
		return
	}
	e.literals = append(e.literals, r)
	if len(append(e.atoms, e.literals...)) > 4 {
		var highest int
		for ex, w := range knownTypes {
			if allMatch(ex, append(e.atoms, e.literals...)) && w > highest {
				e.pattern = ex
				e.atoms = append(e.atoms, e.literals...)
				e.literals = nil
				highest = w
			}
		}
	}
}

// Pattern returns a regexp representing the CharacterExpression in string form
func (e *CharacterExpression) Pattern() string {
	var k []string
	if e.pattern != nil {
		k = append(k, e.pattern.String())
	}
	for _, a := range e.literals {
		k = append(k, string(a))
	}
	return "(" + strings.Join(k, "|") + ")"
}

func combine(a, b *regexp.Regexp) *regexp.Regexp {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return regexp.MustCompile(
		a.String() + "|" + b.String(),
	)
}

func allMatch(ex *regexp.Regexp, patterns []rune) bool {
	bounded := regexp.MustCompile("^" + ex.String() + "$")
	for _, p := range patterns {
		if !bounded.Match([]byte{byte(p)}) {
			return false
		}
	}
	return true
}

var knownTypes = map[*regexp.Regexp]int{
	regexp.MustCompile(`[A-Z]`):        200,
	regexp.MustCompile(`[a-z]`):        200,
	regexp.MustCompile(`[[:digit:]]`):  200,
	regexp.MustCompile(`[[:alpha:]]`):  100,
	regexp.MustCompile(`[[:xdigit:]]`): 100,
	regexp.MustCompile(`\w`):           50,
	regexp.MustCompile(`.`):            1,
}
