package domino

import (
	"bytes"
	"fmt"
	"github.com/knio/domino/tags"
	// "strings"
)

var _ = tags.Title

type Attr map[string]interface{}


type Node interface {
	String() string
	StringBuild(*bytes.Buffer)
}

type DomNode interface {
	Node

	// DOM Core API
	appendChild(Node) Node
	// removeChild(Node) Node
	setAttribute(key string, value interface{}) DomNode

	// Fluent API
	Div(args ...interface{}) DomNode
	Span(args ...interface{}) DomNode
	Text(string) DomNode
}

// node is a DOM tree node.
type domNode struct {
	NodeName    string
	Attrs       Attr
	Children    []Node
}

type textNode struct {
	Value string
}

// Node constructs a node with a name.
func MakeDomNode(name string, args ...interface{}) DomNode {
	n := &domNode{
		NodeName:     name,
		Attrs:    make(Attr, 0),
		Children: make([]Node, 0),
	}
	for _, arg := range args {
		switch a := arg.(type) {
		case *domNode:
			n.appendChild(a)
		case *textNode:
			n.appendChild(a)
		case string:
			n.appendChild(TextNode(a))
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


// NodeArgs constructs a text node.
func TextNode(value string) Node {
	return &textNode{Value: value}
}

// add a new node to the list of children and return the added node
func (n *domNode) appendChild(child Node) Node {
	n.Children = append(n.Children, child)
	return child
}

func (n *domNode) setAttribute(k string, v interface{}) DomNode {
	n.Attrs[k] = v
	return n
}

// returns the HTML for this node and all its ancestors
func (n *textNode) String() string {
	return n.Value
}

func (n *textNode) StringBuild(b *bytes.Buffer) {
	b.WriteString(n.Value)
}

// returns the HTML for this node and all its ancestors
func (n *domNode) String() string {
	b := &bytes.Buffer{}
	n.StringBuild(b)
	return b.String()
}

func (n *domNode) StringBuild(b *bytes.Buffer) {
	b.WriteByte('<')
	b.WriteString(n.NodeName)

	for k, v := range n.Attrs {
		b.WriteByte(' ')

		if v == nil {
			fmt.Fprint(b, k)
			continue
		}

		vstr := HTMLEscape(fmt.Sprint(v))
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

func (n *domNode) Text(text string) DomNode {
	n.appendChild(TextNode(text))
	return n
}

func HTMLEscape(s string) string {
	return s
}
