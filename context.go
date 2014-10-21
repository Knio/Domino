package domino

/*
Context is a stack of dom nodes, the last-in node becomes the focused node
and from there you can add elements, text etc. to it before popping to get
back up the stack. See examples.
*/
type Context struct {
	*DomNode
	stack []*DomNode
}

// NewContext creates a context focused on the passed in dom node.
func NewContext(n *DomNode) *Context {
	c := &Context{
		DomNode: n,
		stack:   make([]*DomNode, 0),
	}
	c.Push(n)
	return c
}

// Push a node on to the stack and focus it.
func (c *Context) Push(n *DomNode) {
	c.stack = append(c.stack, n)
	c.DomNode = n
}

// Pop a node off of the stack, focus moves to the last-in node.
func (c *Context) Pop() {
	if len(c.stack) == 0 {
		panic("Nothing to pop")
	}
	c.stack = c.stack[0 : len(c.stack)-1]
	c.DomNode = c.stack[len(c.stack)-1]
}
