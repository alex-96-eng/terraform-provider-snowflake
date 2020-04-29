package snowflake

import (
	"fmt"
	"unicode"
)

// ViewSelectStatementExtractor is a simplistic parser that only exists to extract the select statement from a
// create view statement
type ViewSelectStatementExtractor struct {
	input []rune
	pos   int
}

func NewViewSelectStatementExtractor(input string) *ViewSelectStatementExtractor {
	return &ViewSelectStatementExtractor{
		input: []rune(input),
	}
}

func (e *ViewSelectStatementExtractor) Extract() (string, error) {
	e.consumeSpace()
	e.consumeToken("create")
	e.consumeSpace()
	e.consumeToken("view")
	e.consumeSpace()
	e.consumeIdentifier()
	e.consumeSpace()
	e.consumeToken("as")
	e.consumeSpace()

	return string(e.input[e.pos:]), nil
}

func (e *ViewSelectStatementExtractor) consumeToken(t string) {
	fmt.Printf("consume token %s\n", t)
	found := 0
	for i, r := range t {
		fmt.Printf("e.pos %d r %s\n", e.pos, string(r))
		if e.pos+i > len(e.input) || r != e.input[e.pos+i] {
			break
		}
		found += 1
	}
	fmt.Printf("found %d\n", found)

	if found == len(t) {
		e.pos += len(t)
	}
}

func (e *ViewSelectStatementExtractor) consumeSpace() {
	found := 0
	for {
		fmt.Printf("e.pos %d found %d r %s\n", e.pos, found, string(e.input[e.pos+found]))
		if e.pos+found > len(e.input)-1 || !unicode.IsSpace(e.input[e.pos+found]) {
			break
		}
		found += 1
	}
	e.pos += found
}

func (e *ViewSelectStatementExtractor) consumeIdentifier() {
	// TODO quoted identifiers
	e.consumeNonSpace()
}

func (e *ViewSelectStatementExtractor) consumeNonSpace() {
	found := 0
	for {
		if e.pos+found > len(e.input)-1 || unicode.IsSpace(e.input[e.pos+found]) {
			break
		}
		found += 1
	}
	e.pos += found
}
