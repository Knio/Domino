package domino

import (
	// "github.com/knio/domino/tags"
	"testing"
)

func TestDiv(t *testing.T) {
	t.Parallel()
	if div := Div().String(); div != "<div></div>" {
		t.Error("Expected a div, got: ", div)
	}

	if div := Div("hey").String(); div != "<div>hey</div>" {
		t.Error("Expected a div with hey, got: ", div)
	}
}

func TestAttrs(t *testing.T) {
	div := Div(Attr{"html": nil, "no": "quotes", "attr": "value"}).String()
	if div != `<div html no="quotes" attr="value"></div>` {
		t.Error("Expected a div with several attributes, got: ", div)
	}
}

func TestChildren(t *testing.T) {
	if div := Div(Span("hey")).String(); div != "<div><span>hey</span></div>" {
		t.Error("Expected a div with a span with hey, got: ", div)
	}
}

func TestSiblings(t *testing.T) {
	if div := Div().Span("hey").String(); div != "<div><span>hey</span></div>" {
		t.Error("Expected a span in a div, got: ", div)
	}
}

func TestEscapes(t *testing.T) {
	t.Parallel()
	_ = HTMLEscape
	t.SkipNow()
}
