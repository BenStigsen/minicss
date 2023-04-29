package minicss

func Minify(data []byte) []byte {
	out := make([]byte, 0, len(data))

	for i := 0; i < len(data); i += 1 {
		a := data[i]

		start := i
		if i < len(data)-1 && a == '/' {
			if data[i+1] == '/' {
				for i < len(data) && data[i] != '\n' {
					i += 1
				}
				continue
			} else if data[i+1] == '*' {
				for i < len(data)-1 {
					i += 1

					if data[i] == '*' && data[i+1] == '/' {
						break
					}
				}

				if data[i+1] != '/' {
					out = append(out, data[start:i+1]...)
				} else {
					i += 1
				}

				continue
			}
			// Block Comment
		}

		// https://pkg.go.dev/unicode#IsSpace
		if a == '\t' || a == '\n' || a == '\v' || a == '\f' || a == '\r' || a == ' ' || a == '\u0085' || a == '\u00A0' {
			continue
		}

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

		out = append(out, a)
	}

	return out[:len(out)]
}
