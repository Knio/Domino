package domino

import (
	"fmt"
	"testing"
	// . "github.com/Knio/Domino"
)

func TestContext(t *testing.T) {
	c := NewContext(Div())

	c.Push(Ul(c))
	for i := 0; i < 2; i++ {
		c.Li(fmt.Sprintf("hello %d", i+1))
	}
	c.Pop()

	Span(c, "end")

	if c.String() != "<div><ul><li>hello 1</li><li>hello 2</li></ul><span>end</span></div>" {
		t.Error("Expected a span in a div, got: ", c.String())
	}
}
