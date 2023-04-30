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
		border: solid #5B6DCD 10px;
	}`)

	//expect := `p{color:red;content:"Hello, world!";border: solid #5B6DCD 10px;}`
	expect := `p { color: red; content: "Hello, world!"; border: solid #5B6DCD 10px; }`
	result := Minify(data)
	if string(result) != expect {
		t.Fatalf(`Expected "%s", error`, expect)
	}
}
