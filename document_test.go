package domino

import "fmt"

func ExampleDocument() {
	// Create a document with the title Domino
	d := NewDocument("Domino")

	// Add an H1 to the body.
	d.Body.H1("Hello World!")

	// Add a script to the header.
	d.Head.Script(Attr{"type": "text/javascript", "src": "//google.com/jquery.js"})

	fmt.Println(d.String())
	// Output:
	// <!DOCTYPE html>
	// <html><head><title>Domino</title><script src="//google.com/jquery.js" type="text/javascript"></script></head><body><h1>Hello World!</h1></body></html>
}
