package domino

import (
	"bytes"
	"fmt"
	"html"
	"sort"
)

type Attr map[string]interface{}

type Node interface {
	String() string
	StringBuild(*bytes.Buffer)
}

// node is a DOM tree node.
type DomNode struct {
	NodeName string
	Attrs    Attr
	Children []Node
}

type TextNode struct {
	Value string
}

// Node constructs a node with a name.
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

// NodeArgs constructs a text node.
func NewTextNode(value string) *TextNode {
	return &TextNode{Value: value}
}

// add a new node to the list of children and return the added node
func (n *DomNode) Add(child Node) Node {
	n.Children = append(n.Children, child)
	return child
}

func (n *DomNode) setAttribute(k string, v interface{}) *DomNode {
	n.Attrs[k] = v
	return n
}

// returns the HTML for this node and all its ancestors
func (n *TextNode) String() string {
	return n.Value
}

func (n *TextNode) StringBuild(b *bytes.Buffer) {
	b.WriteString(n.Value)
}

// returns the HTML for this node and all its ancestors
func (n *DomNode) String() string {
	b := &bytes.Buffer{}
	n.StringBuild(b)
	return b.String()
}

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

func (n *DomNode) Text(text string) *DomNode {
	n.Add(NewTextNode(text))
	return n
}

func (n *DomNode) Clear() *DomNode {
	n.Children = make([]Node, 0)
	return n
}