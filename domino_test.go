package domino

import "testing"

func TestDiv(t *testing.T) {
	t.Parallel()
	if div := Div().String(); div != "<div></div>" {
		t.Errorf("Expected a div, got: %q", div)
	}

	if div := Div("hey").String(); div != "<div>hey</div>" {
		t.Errorf("Expected a div with hey, got: %q", div)
	}
}

func TestAttrs(t *testing.T) {
	div := Div(Attr{"html": nil, "no": "qu\"otes", "attr": "value"}).String()
	if div != `<div attr="value" html no="qu&#34;otes"></div>` {
		t.Errorf("Expected a div with several attributes, got: %q", div)
	}
}

func TestChildren(t *testing.T) {
	if div := Div(Span("hey")).String(); div != "<div><span>hey</span></div>" {
		t.Errorf("Expected a div with a span with hey, got: %q", div)
	}
}

func TestSiblings(t *testing.T) {
	if div := Div().Span("hey").String(); div != "<div><span>hey</span></div>" {
		t.Errorf("Expected a span in a div, got: %q", div)
	}
}

func TestIndent(t *testing.T) {
	if div := Div().Span("hey").IndentString(); div != "<div>\n    <span>hey</span>\n</div>\n" {
		t.Errorf("Expected a span in a div, got: %q", div)
	}
}
