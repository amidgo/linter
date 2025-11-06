package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Правило для обнаружения fallthrough
func banFallthrough(m dsl.Matcher) {
	m.Match(`fallthrough`).
		Report("fallthrough is considered an antipattern")
}

// Правило для обнаружения goto
func banGoto(m dsl.Matcher) {
	m.Match(`goto $label`).
		Report("goto is considered an antipattern")
}
