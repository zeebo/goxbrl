package marshal

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Node struct {
	Name  xml.Name
	Attrs []xml.Attr
	Value string
	Nodes []*Node
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(node *Node, nsmap map[string]string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	return e.encodeNode(node, nsmap, 0, true)
}

func (e *Encoder) encodeNode(node *Node, nsmap map[string]string, indent int, first bool) error {
	e.indent(indent)
	e.writeString("<")
	e.writeString(nameOf(node.Name, nsmap))
	for _, attr := range node.Attrs {
		e.writeString(" ")
		e.writeString(nameOf(attr.Name, nsmap))
		e.writeString(`="`)
		e.writeString(attr.Value)
		e.writeString(`"`)
	}
	if first {
		for url, name := range nsmap {
			e.writeString(" xmlns:" + name + `="` + url + `"`)
		}
	}
	if node.Value == "" && len(node.Nodes) == 0 {
		e.writeString(" />")
		return nil
	}
	e.writeString(">")

	if node.Value != "" {
		e.writeString(node.Value)
	} else {
		for _, sub := range node.Nodes {
			e.writeString("\n")
			if err := e.encodeNode(sub, nsmap, indent+1, false); err != nil {
				return err
			}
		}
		e.writeString("\n")
		e.indent(indent)
	}
	e.writeString("</")
	e.writeString(nameOf(node.Name, nsmap))
	e.writeString(">")
	return nil
}

func nameOf(name xml.Name, nsmap map[string]string) string {
	if name.Space == "" {
		return name.Local
	}
	local, ok := nsmap[name.Space]
	if !ok {
		panic("unknown namespace: " + name.Space)
	}
	return local + ":" + name.Local
}

func (e *Encoder) indent(n int) {
	for i := 0; i < n; i++ {
		e.writeString("\t")
	}
}

func (e *Encoder) writeString(s string) {
	e.w.Write([]byte(s))
}
