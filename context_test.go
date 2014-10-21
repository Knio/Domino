package domino

import "fmt"

func ExampleContext() {
	// Establish a new context on a <div>
	c := NewContext(Div())

	// Push a <ul> on to the stack.
	c.Push(Ul(c))
	for i := 0; i < 2; i++ {
		// Add to the current context's focus <ul>
		c.Li(fmt.Sprintf("hello %d", i+1))
	}
	// Pop the <ul> off the stack, context focuses on the <div> again.
	c.Pop()

	// Add a span to the div from before
	Span(c, "end")
	fmt.Println(c)
	// Output:
	// <div><ul><li>hello 1</li><li>hello 2</li></ul><span>end</span></div>
}
