package domino

import (
	"bytes"
	"fmt"
	"github.com/knio/domino/tags"
	"strings"
)

var _ = tags.Title

type Attr map[string]interface{}

// nodeKind is the type of the node.
type nodeKind int

// The kinds of nodes
const (
	tagNode nodeKind = iota
	textNode
)

// node is a DOM tree node.
type node struct {
	Name        string
	Kind        nodeKind
	Attrs       Attr
	Children    []*node
	Left, Right *node
}

// Node constructs a node with a name.
func Node(name string) *node {
	return &node{
		Name:     name,
		Attrs:    make(Attr),
		Children: make([]*node, 0),
	}
}

// NodeArgs constructs a node with a name and arguments.
func NodeArgs(name string, args ...interface{}) *node {
	n := Node(name)
	n.inflateArgs(args...)
	return n
}

// NodeArgs constructs a text node.
func NodeText(content string) *node {
	n := Node(content)
	n.Kind = textNode
	return n
}

// pushChild adds a new node to the list of children, and returns the pushed
// child.
func (n *node) pushChild(child *node) *node {
	n.Children = append(n.Children, child)
	return child
}

// setSibling sets this node's sibling and returns the last node in the sibling
// list.
func (n *node) setSibling(sibling *node) *node {
	n.Right = sibling
	sibling.Left = n
	return sibling
}

// String returns the HTML for this DOM fragment.
func (n *node) String() string {
	b := &bytes.Buffer{}

	// ensure we find the first sibling to start printing
	for n.Left != nil {
		n = n.Left
	}

	n.stringBuild(b)
	return b.String()
}

// StringBuild uses a buffer to write out the HTML for this DOM fragment.
func (n *node) stringBuild(b *bytes.Buffer) {
	writeTag := n.Kind == tagNode

	if writeTag {
		b.WriteByte('<')
		b.WriteString(n.Name)

		for k, v := range n.Attrs {
			b.WriteByte(' ')

			if v == nil {
				fmt.Fprint(b, k)
				continue
			}

			vstr := fmt.Sprint(v)
			if !(strings.HasPrefix(vstr, `"`) && strings.HasSuffix(vstr, `"`)) {
				fmt.Fprintf(b, `%s="%s"`, k, escapeHTMLArg(vstr))
			} else {
				fmt.Fprintf(b, `%s="%s"`, k, escapeHTMLArg(vstr[1:len(vstr)-1]))
			}
		}

		b.WriteByte('>')
	} else {
		b.WriteString(n.Name)
	}

	for _, child := range n.Children {
		child.stringBuild(b)
	}

	if writeTag {
		fmt.Fprintf(b, "</%s>", n.Name)
	}

	for n = n.Right; n != nil; n = n.Right {
		n.stringBuild(b)
	}
}

type document struct {
	*node
	Body  *node
	Title *node
	Head  *node
}

// Document returns a standard HTML document type.
func Document(title string, head, body *node) document {
	d := document{}

	d.Body = Body()
	if len(title) == 0 {
		d.Head = Head()
	} else {
		d.Title = Title(title)
		d.Head = Head(d.Title)
	}

	if head != nil {
		d.Head.pushChild(head)
	}

	if body != nil {
		d.Body.pushChild(body)
	}

	d.node = NodeArgs(tags.DOCTYPE, Attr{"html": nil}).
		HTML(d.Head, d.Body)
	return d
}

func (n *node) inflateArgs(args ...interface{}) {
	for _, arg := range args {
		switch a := arg.(type) {
		case *node:
			n.pushChild(a)
		case string:
			n.pushChild(NodeText(a))
		case Attr:
			for k, v := range a {
				n.Attrs[k] = v
			}
		}
	}
}

// escapeHTMLArg makes the arg="arg" unbreakable.
func escapeHTMLArg(s string) string {
	return s
}
