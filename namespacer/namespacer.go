package namespacer

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}, nsmap map[string]string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	r, w := io.Pipe()
	enc := xml.NewEncoder(w)
	dec := xml.NewDecoder(r)

	//run the encoder
	go func() { w.CloseWithError(enc.Encode(v)) }()

	//loop the tokens from the decoder rewriting the xml nodes
	var tok xml.Token
	for {
		tok, err = dec.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}

		switch v := tok.(type) {
		case xml.StartElement:
			e.writeString("<")
			e.writeString(nameOf(v.Name, nsmap))
			for _, attr := range v.Attr {
				if attr.Name.Local == "xmlns" {
					continue
				}
				e.writeString(" ")
				e.writeString(nameOf(attr.Name, nsmap))
				e.writeString("=\"")
				e.writeString(attr.Value)
				e.writeString("\"")
			}
			e.writeString(">")
		case xml.EndElement:
			e.writeString("</")
			e.writeString(nameOf(v.Name, nsmap))
			e.writeString(">")
		case xml.CharData:
			xml.Escape(e.w, v)
		case xml.Comment:
		case xml.ProcInst:
			e.writeString("<?")
			e.writeString(v.Target + " ")
			e.writeBytes(v.Inst)
			e.writeString("?>")
		case xml.Directive:
			e.writeString("<!")
			e.writeBytes(v)
			e.writeString(">")
		default:
			panic("unknown token")
		}
	}

	return
}

func nameOf(name xml.Name, nsmap map[string]string) string {
	if name.Space == "" {
		return name.Local
	}

	local, ok := nsmap[name.Space]
	if !ok {
		panic("unknown namespace: " + name.Space)
	}
	return fmt.Sprintf("%s:%s", local, name.Local)
}

func (e *Encoder) writeString(s string) {
	n, err := io.WriteString(e.w, s)
	if n != len(s) || err != nil {
		panic(fmt.Sprintf("short write or %v", err))
	}
}

func (e *Encoder) writeBytes(p []byte) {
	n, err := e.w.Write(p)
	if n != len(p) || err != nil {
		panic(fmt.Sprintf("short write or %v", err))
	}
}

func (e *Encoder) indent(n int) {
	e.writeString("\n")
	for i := 0; i < n; i++ {
		e.writeString("\t")
	}
}
