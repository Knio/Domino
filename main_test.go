package domino

import (
	"github.com/knio/domino/tags"
	. "testing"
)

func TestDocument(t *T) {
	t.Parallel()
	doc := Document("", nil, nil)

	if doc.Title != nil {
		t.Error("Expected no title node.")
	}

	if doc.Body.Name != tags.Body {
		t.Error("Expected a body node.")
	}

	if doc.Head.Name != tags.Head {
		t.Error("Expected a head node.")
	}
}

func TestDocumentTitle(t *T) {
	t.Parallel()
	title := "testTitle"
	doc := Document(title, nil, nil)

	if doc.Title.Name != tags.Title {
		t.Error("Expected a title node.")
	}

	if len(doc.Title.Children) == 0 {
		t.Fatal("Title has no child nodes.")
	}

	if check := doc.Title.Children[0].Name; title != check {
		t.Errorf("Expected a title with text: %v, got: %v", title, check)
	}
}

func TestDocumentChildren(t *T) {
	t.Parallel()
	doc := Document("", Div(), Span())

	if ln := len(doc.Head.Children); ln != 1 {
		t.Fatal("Head must have exactly 1 child, has: ", ln)
	}

	if check := doc.Head.Children[0].Name; tags.Div != check {
		t.Errorf("Expected a title with text: %v, got: %v", tags.Div, check)
	}

	if ln := len(doc.Body.Children); ln != 1 {
		t.Fatal("Body must have exactly 1 child, has: ", ln)
	}

	if check := doc.Body.Children[0].Name; tags.Span != check {
		t.Errorf("Expected a title with text: %v, got: %v", tags.Span, check)
	}
}

func TestNode(t *T) {
	t.Parallel()
	if node := Node(tags.HTML); node == nil {
		t.Error("Node should not be nil.")
	} else if node.Name != tags.HTML {
		t.Error("Name should be:", tags.HTML)
	} else if node.Children == nil {
		t.Error("It should pre-allocate the array of children.")
	}
}

func TestDiv(t *T) {
	t.Parallel()
	if div := Div().String(); div != "<div></div>" {
		t.Error("Expected a div, got: ", div)
	}

	if div := Div("hey").String(); div != "<div>hey</div>" {
		t.Error("Expected a div with hey, got: ", div)
	}
}

func TestAttrs(t *T) {
	div := Div(Attr{"style": "height: 57px;"}).String()
	if div != `<div style="height: 57px;"></div>` {
		t.Error("Expected a div with a style attribute, got: ", div)
	}
}

func TestChildren(t *T) {
	if div := Div(Span("hey")).String(); div != "<div><span>hey</span></div>" {
		t.Error("Expected a div with a span with hey, got: ", div)
	}
}

func TestSiblings(t *T) {
	if div := Div().Span("hey").String(); div != "<div></div><span>hey</span>" {
		t.Error("Expected a div next to a span with hey, got: ", div)
	}
}

func TestEscapes(t *T) {
	t.Parallel()
	_ = escapeHTMLArg
	t.SkipNow()
}
