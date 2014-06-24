package domino

type Context struct {
	*DomNode
	Stack []*DomNode
}

func MakeContext(n *DomNode) *Context {
	c := &Context{
		DomNode: n,
		Stack:   make([]*DomNode, 0),
	}
	c.Push(n)
	return c
}

func (c *Context) Push(n *DomNode) {
	c.Stack = append(c.Stack, n)
	c.DomNode = n
}

func (c *Context) Pop() {
	if len(c.Stack) == 0 {
		panic("Nothing to pop")
	}
	c.Stack = c.Stack[0 : len(c.Stack)-1]
	c.DomNode = c.Stack[len(c.Stack)-1]
}
