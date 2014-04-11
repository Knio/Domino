package domino

import (
	"github.com/knio/domino/tags"
)

func HTML(args ...interface{}) *node {
	return NodeArgs(tags.HTML, args...)
}

func (n *node) HTML(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.HTML, args...))
}

func Body(args ...interface{}) *node {
	return NodeArgs(tags.Body, args...)
}

func (n *node) Body(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.Body, args...))
}

func Head(args ...interface{}) *node {
	return NodeArgs(tags.Head, args...)
}

func (n *node) Head(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.Head, args...))
}

func Title(args ...interface{}) *node {
	return NodeArgs(tags.Title, args...)
}

func (n *node) Title(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.Title, args...))
}

func Div(args ...interface{}) *node {
	return NodeArgs(tags.Div, args...)
}

func (n *node) Div(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.Div, args...))
}

func Span(args ...interface{}) *node {
	return NodeArgs(tags.Span, args...)
}

func (n *node) Span(args ...interface{}) *node {
	return n.setSibling(NodeArgs(tags.Span, args...))
}
