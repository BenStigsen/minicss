# minicss

Tiny, fast and simple CSS minification.
- Removes whitespaces
- Removes single-line comments
- Removes multi-line comments

## Usage

```go
import (
    "fmt"
    "github.com/benstigsen/minicss"
)

func main() {
    data := []byte(`
	// This is a line comment
	p {
		color: red; /* This is a comment */
		content: "Hello, world!";
	}`)

    result := minicss.Minify(data)
    fmt.Println(string(result)) // p{color:red;content:"Hello, world!";}
}
```