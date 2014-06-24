package domino

import (
	"bytes"
)

type Document struct {
	Html  *DomNode
	Head  *DomNode
	Body  *DomNode
	title *DomNode
}

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

func (d *Document) SetTitle(title string) {
	d.title.Clear()
	d.title.Text(title)
}

func (d *Document) String() string {
	b := &bytes.Buffer{}
	b.WriteString("<!DOCTYPE html>\n")
	d.Html.StringBuild(b)
	return b.String()
}
