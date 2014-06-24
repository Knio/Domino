package domino

type Context struct {
	*DomNode
	stack []*DomNode
}

func NewContext(n *DomNode) *Context {
	c := &Context{
		DomNode: n,
		stack:   make([]*DomNode, 0),
	}
	c.Push(n)
	return c
}

func (c *Context) Push(n *DomNode) {
	c.stack = append(c.stack, n)
	c.DomNode = n
}

func (c *Context) Pop() {
	if len(c.stack) == 0 {
		panic("Nothing to pop")
	}
	c.stack = c.stack[0 : len(c.stack)-1]
	c.DomNode = c.stack[len(c.stack)-1]
}
