package minicss

import (
	"testing"
)

func TestMinify(t *testing.T) {
	data := []byte(`
	// This is a line comment
	p {
		color: red; /* This is a comment */
		content: "Hello, world!";
	}`)

	expect := `p{color:red;content:"Hello, world!";}`
	result := Minify(data)
	if string(result) != expect {
		t.Fatalf(`Expected "%s", error`, expect)
	}
}
