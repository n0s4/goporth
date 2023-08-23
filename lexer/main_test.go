package lexer

import (
	"golang.org/x/exp/slices"
	"testing"
)

func TestLex(t *testing.T) {
	expected := map[string]([]Token){
		"":                     {},
		" \t \n \r\n ":         {},
		"+":                    {{Op: PLUS}},
		"print":                {{Op: PRINT}},
		"101":                  {{PUSH, 101}},
		"35  34    + \n print": {{PUSH, 35}, {PUSH, 34}, {Op: PLUS}, {Op: PRINT}},
	}

	for input, expected := range expected {
		t.Logf("testing tokenizing '%v'\n", input)
		got, _ := Lex(input)
		if !slices.Equal(got, expected) {
			t.Errorf("expected: %v\ngot: %v\n", expected, got)
		}
	}
}
