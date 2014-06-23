package main

import (
	"fmt"
	. "github.com/Knio/Domino"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		page := Html()

		body := Body()
		page.Add(body)

		nums := Div()
		body.Add(nums)
		nums.Text("Here are the first 10 integers:")
		for i := 0; i < 10; i++ {
			nums.Span(fmt.Sprintf("%d", i))
		}

		fmt.Fprintln(w, page)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
