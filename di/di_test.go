package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Shivam")

	got := buffer.String()
	want := "Hello, Shivam"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
