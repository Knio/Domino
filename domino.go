package domino

import (
	"bytes"
	"fmt"
	"html"
	"sort"
)

// Attr is a shorthand for map[string]interface{}, used when declaring
// attributes for DomNodes.
type Attr map[string]interface{}

// Node is basically a stringer interface but with the ability to append
// to a buffer for performance reasons.
type Node interface {
	String() string
	StringBuild(*bytes.Buffer)
}

// DomNode is a node in a dom tree. It has children nodes and attributes.
type DomNode struct {
	NodeName string
	Attrs    Attr
	Children []Node
}

// TextNode represents text in the dom tree. Text can not have child nodes or
// attributes.
type TextNode struct {
	Value string
}

// NewDomNode creates a new dom node with it's name (div, ul, html) and any
// number of arguments, these can be: children dom or text nodes, or Attr lists.
func NewDomNode(name string, args ...interface{}) *DomNode {
	n := &DomNode{
		NodeName: name,
		Attrs:    make(Attr, 0),
		Children: make([]Node, 0),
	}
	for _, arg := range args {
		switch a := arg.(type) {
		case *DomNode:
			n.Add(a)
		case *TextNode:
			n.Add(a)
		case string:
			n.Add(NewTextNode(a))
		case *Context:
			a.Add(n)
		case Attr:
			for k, v := range a {
				n.Attrs[k] = v
			}
		default:
			panic("wrong argument type")
		}
	}
	return n
}

// NewTextNode creates a new text node with the provided text.
func NewTextNode(value string) *TextNode {
	return &TextNode{Value: value}
}

// Add a new child node.
func (n *DomNode) Add(child Node) Node {
	n.Children = append(n.Children, child)
	return child
}

func (n *DomNode) setAttribute(k string, v interface{}) *DomNode {
	n.Attrs[k] = v
	return n
}

// String returns HTML for this node and all its ancestors
func (n *TextNode) String() string {
	return n.Value
}

// StringBuild appends it's value to the provided buffer.
func (n *TextNode) StringBuild(b *bytes.Buffer) {
	b.WriteString(n.Value)
}

// String returns HTML for this node and all it's ancestors.
func (n *DomNode) String() string {
	b := &bytes.Buffer{}
	n.StringBuild(b)
	return b.String()
}

// StringBuild appends the html for this node and all it's ancestors to the
// provided buffer.
func (n *DomNode) StringBuild(b *bytes.Buffer) {
	b.WriteByte('<')
	b.WriteString(n.NodeName)

	keys := make([]string, 0, len(n.Attrs))
	for k, _ := range n.Attrs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := n.Attrs[k]

		b.WriteByte(' ')
		if v == nil {
			fmt.Fprint(b, k)
			continue
		}

		vstr := html.EscapeString(fmt.Sprint(v))
		fmt.Fprintf(b, `%s="%s"`, k, vstr)
	}

	b.WriteByte('>')

	for _, child := range n.Children {
		child.StringBuild(b)
	}

	b.WriteByte('<')
	b.WriteByte('/')
	b.WriteString(n.NodeName)
	b.WriteByte('>')
}

// Text adds a text node with the provided text.
func (n *DomNode) Text(text string) *DomNode {
	n.Add(NewTextNode(text))
	return n
}

// Clear removes all children.
func (n *DomNode) Clear() *DomNode {
	n.Children = make([]Node, 0)
	return n
}
