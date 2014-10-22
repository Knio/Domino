Domino
======

`Domino` is a Go library for creating and manipulating HTML documents
using an elegant DOM API. It allows you to write HTML pages in pure Go
reasonably concisely, which eliminates the need to learn another template
language, and lets you take advantage of the more powerful features of Go.

Go:

```go
package main

import "fmt"
import . "github.com/Knio/Domino"

func main() {
	d := NewDocument("Dominate your HTML")

	d.Head.Add(
		Link(Attr{"rel": "stylesheet", "href": "style.css"}),
		Script(Attr{"type": "text/javascript", "src": "script.js"}),
	)

	c := NewContext(d.Body.Div(Attr{"id": "header"}))

	c.Push(Ol())
	for _, i := range []string{"home", "about", "contact"} {
		Li(c, A(i, Attr{"href": fmt.Sprintf("/%s", i)}))
	}
	c.Pop()

	c.Div(Attr{"class": "body"}, P("Lorem ipsum.."))

	fmt.Print(d.IndentString())
}
```

HTML:

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Dominate your HTML</title>
        <link href="style.css" rel="stylesheet">
        <script src="script.js" type="text/javascript"></script>
    </head>
    <body>
        <div id="header"></div>
        <ol>
            <li>
                <a href="/home">home</a>
            </li>
            <li>
                <a href="/about">about</a>
            </li>
            <li>
                <a href="/contact">contact</a>
            </li>
        </ol>
        <div class="body">
            <p>Lorem ipsum..</p>
        </div>
    </body>
</html>
```

Compatibility
-------------

`Domino` is compatible with Go 1.1 and above.

Installation
------------

The recommended way to install `domino` is with

```bash
go get github.com/Knio/Domino
```

Developed By
------------

* Tom Flanagan - <tom@zkpq.ca>

Git repository located at
[github.com/Knio/Domino](//github.com/Knio/Domino)
