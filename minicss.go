package minicss

func Minify(data []byte) []byte {
	out := make([]byte, 0, len(data))

	prev_space := false
	for i := 0; i < len(data); i += 1 {
		a := data[i]

		start := i
		if i < len(data)-1 && a == '/' {
			// Single-line comment
			if data[i+1] == '/' {
				for i < len(data) && data[i] != '\n' {
					i += 1
				}
				continue
			// Block comment
			} else if data[i+1] == '*' {
				for i < len(data)-1 {
					if data[i] == '*' && data[i+1] == '/' {
						break
					}
					i += 1
				}

				if i < len(data) - 1 && data[i+1] == '/' {
					i += 1
				} else {
					out = append(out, data[start:i+1]...)
				}

				continue
			}
		}

		// String
		if a == '\'' || a == '"' {
			for _, b := range data[i+1:] {
				i += 1
				if b == a {
					break
				}
			}
			out = append(out, data[start:i+1]...)
			continue
		}

		// https://pkg.go.dev/unicode#IsSpace
		if a == '\n' {
			continue
		}

		if a == '\t' || a == '\v' || a == '\f' || a == '\r' || a == ' ' || a == '\u0085' || a == '\u00A0' {
			if !prev_space && len(out) > 0 {
				out = append(out, ' ')
				prev_space = true
			}
			continue
		}

		prev_space = false
		out = append(out, a)
	}

	return out[:len(out)]
}
