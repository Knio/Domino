package domino

import (
	"github.com/knio/domino/tags"
)

func HTML(args ...interface{}) *DomNode {
	return MakeDomNode(tags.HTML, args...)
}

func Head(args ...interface{}) *DomNode {
	return MakeDomNode(tags.Head, args...)
}

func Title(args ...interface{}) *DomNode {
	return MakeDomNode(tags.Title, args...)
}

func Body(args ...interface{}) *DomNode {
	return MakeDomNode(tags.Body, args...)
}

func Div(args ...interface{}) *DomNode {
	return MakeDomNode(tags.Div, args...)
}

func Span(args ...interface{}) *DomNode {
	return MakeDomNode(tags.Span, args...)
}


func (n *DomNode) HTML(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.HTML, args...))
	 return n
}
func (n *DomNode) Head(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.Head, args...))
	 return n
}

func (n *DomNode) Title(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.Title, args...))
	 return n
}

func (n *DomNode) Body(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.Body, args...))
	 return n
}
func (n *DomNode) Div(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.Div, args...))
	 return n
}

func (n *DomNode) Span(args ...interface{}) *DomNode {
	 n.appendChild(MakeDomNode(tags.Span, args...))
	 return n
}
