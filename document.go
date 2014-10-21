package domino

import (
	"bytes"
)

// Document is a convenience type to help build a standard HTML5 document. It
// provides a doctype, and header/title/body structure.
type Document struct {
	Html  *DomNode
	Head  *DomNode
	Body  *DomNode
	title *DomNode
}

// NewDocument creates a new HTML5 document with a title using the text provided
func NewDocument(title string) *Document {
	d := &Document{
		Html:  Html(),
		Head:  Head(),
		Body:  Body(),
		title: Title(),
	}
	d.Html.Add(d.Head)
	d.Html.Add(d.Body)
	d.Head.Add(d.title)
	d.SetTitle(title)
	return d
}

// SetTitle sets the title of the document.
func (d *Document) SetTitle(title string) {
	d.title.Clear()
	d.title.Text(title)
}

// String returns the HTML for the document.
func (d *Document) String() string {
	b := &bytes.Buffer{}
	b.WriteString("<!DOCTYPE html>\n")
	d.Html.StringBuild(b)
	return b.String()
}
