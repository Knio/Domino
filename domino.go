package domino

import (
	"bytes"
	"fmt"
	"html"
	"sort"
	"strings"
)

// NSpaces can be adjusted to determine how many spaces are used in each
// indent.
var NSpaces = 4

// Attr is a shorthand for map[string]interface{}, used when declaring
// attributes for DomNodes.
type Attr map[string]interface{}

// Node is basically a stringer interface but with the ability to append
// to a buffer for performance reasons.
type Node interface {
	String() string
	IndentString() string
	StringBuild(*bytes.Buffer, bool, int)
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

// IndentString appends it's value to the provided buffer.
func (n *TextNode) IndentString() string {
	return n.Value
}

// StringBuild appends it's value to the provided buffer.
func (n *TextNode) StringBuild(b *bytes.Buffer, indent bool, depth int) {
	if indent && depth > 0 {
		indent := strings.Repeat(" ", NSpaces*depth)
		b.WriteString(indent)
		b.WriteString(strings.Replace(n.Value, "\n", "\n"+indent, -1))
		b.WriteByte('\n')
	} else {
		b.WriteString(n.Value)
	}
}

// String returns HTML for this node and all it's ancestors.
func (n *DomNode) String() string {
	b := &bytes.Buffer{}
	n.StringBuild(b, false, 0)
	return b.String()
}

// IndentString returns HTML for this node and all it's ancestors with
// indentation.
func (n *DomNode) IndentString() string {
	b := &bytes.Buffer{}
	n.StringBuild(b, true, 0)
	return b.String()
}

// StringBuild appends the html for this node and all it's ancestors to the
// provided buffer.
func (n *DomNode) StringBuild(b *bytes.Buffer, indent bool, depth int) {
	if indent && depth > 0 {
		b.WriteString(strings.Repeat(" ", NSpaces*depth))
	}

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

	if _, ws := whiteSpaced[n.NodeName]; indent && !ws && len(n.Children) == 1 {
		if _, ok := n.Children[0].(*TextNode); ok {
			fmt.Fprintf(b, "%s</%s>\n", n.Children[0].String(), n.NodeName)
			return
		}
	}

	_, isVoid := voidElems[n.NodeName]

	if indent && (len(n.Children) > 0 || isVoid) {
		b.WriteByte('\n')
	}

	if isVoid {
		return
	}

	for _, child := range n.Children {
		child.StringBuild(b, indent, depth+1)
	}

	if indent && depth > 0 && len(n.Children) > 0 {
		b.WriteString(strings.Repeat(" ", depth*NSpaces))
	}

	b.WriteString("</")
	b.WriteString(n.NodeName)
	b.WriteByte('>')
	if indent {
		b.WriteByte('\n')
	}
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

var voidElems = map[string]struct{}{
	"area":    struct{}{},
	"base":    struct{}{},
	"br":      struct{}{},
	"col":     struct{}{},
	"command": struct{}{},
	"embed":   struct{}{},
	"hr":      struct{}{},
	"img":     struct{}{},
	"input":   struct{}{},
	"keygen":  struct{}{},
	"link":    struct{}{},
	"meta":    struct{}{},
	"param":   struct{}{},
	"source":  struct{}{},
	"track":   struct{}{},
	"wbr":     struct{}{},
}

var whiteSpaced = map[string]struct{}{
	"style":    struct{}{},
	"script":   struct{}{},
	"textarea": struct{}{},
}
