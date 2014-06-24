package domino

import (
	"testing"
	// . "github.com/Knio/Domino"
)

func TestDocument(t *testing.T) {
	d := NewDocument("Domino")

	d.Body.H1("Hello World!")

	if d.String() != "<!DOCTYPE html>\n<html><head><title>Domino</title></head><body><h1>Hello World!</h1></body></html>" {
		t.Error("got: ", d.String())
	}
}
