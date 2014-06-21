package domino

import (
	"github.com/knio/domino/tags"
)

func HTML(args ...interface{}) DomNode {
	return MakeDomNode(tags.HTML, args...)
}

func (n *domNode) HTML(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.HTML, args...))
	 return n
}

func Head(args ...interface{}) DomNode {
	return MakeDomNode(tags.Head, args...)
}

func (n *domNode) Head(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.Head, args...))
	 return n
}

func Title(args ...interface{}) DomNode {
	return MakeDomNode(tags.Title, args...)
}

func (n *domNode) Title(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.Title, args...))
	 return n
}

func Body(args ...interface{}) DomNode {
	return MakeDomNode(tags.Body, args...)
}

func (n *domNode) Body(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.Body, args...))
	 return n
}

func Div(args ...interface{}) DomNode {
	return MakeDomNode(tags.Div, args...)
}

func (n *domNode) Div(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.Div, args...))
	 return n
}

func Span(args ...interface{}) DomNode {
	return MakeDomNode(tags.Span, args...)
}

func (n *domNode) Span(args ...interface{}) DomNode {
	 n.appendChild(MakeDomNode(tags.Span, args...))
	 return n
}
