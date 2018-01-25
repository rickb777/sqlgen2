package schema

import (
	"io"
	"bytes"
)

type Identifiers []string

func (ids Identifiers) Quoted(w io.Writer, quoter func(string) string) {
	comma := ""
	for _, id := range ids {
		io.WriteString(w, comma)
		io.WriteString(w, quoter(id))
		comma = ","
	}
}

func (ids Identifiers) MkString(sep string) string {
	return ids.MkString3("", sep, "")
}

func (ids Identifiers) MkString3(before, separator, after string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(ids)*12))
	ids.MkString3W(w, before, separator, after)
	return w.String()
}

func (ids Identifiers) MkString3W(w io.Writer, before, separator, after string) {
	if len(ids) > 0 {
		comma := before
		for _, id := range ids {
			io.WriteString(w, comma)
			io.WriteString(w, string(id))
			comma = separator
		}
		io.WriteString(w, after)
	}
}
