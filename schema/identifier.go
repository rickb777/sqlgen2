package schema

import (
	"io"
	"bytes"
)

type Identifier string

func (id Identifier) Quoted(quoter func(Identifier) string) string {
	return quoter(id)
}

type Identifiers []Identifier

func (ids Identifiers) Quoted(w io.Writer, quoter func(Identifier) string) {
	comma := ""
	for _, id := range ids {
		io.WriteString(w, comma)
		io.WriteString(w, quoter(id))
		comma = ", "
	}
}

func (ids Identifiers) QuotedS(quoter func(Identifier) string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(ids)*10))
	ids.Quoted(w, quoter)
	return w.String()
}

func (ids Identifiers) MkString(sep string) string {
	w := bytes.NewBuffer(make([]byte, 0, len(ids)*10))
	comma := ""
	for _, id := range ids {
		io.WriteString(w, comma)
		io.WriteString(w, string(id))
		comma = sep
	}
	return w.String()
}
