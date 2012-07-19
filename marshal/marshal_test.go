package marshal

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	tree := &Node{
		Name: xml.Name{"", "bar"},
		Nodes: []*Node{
			{Name: xml.Name{"http://someurl", "baz"}},
			{
				Name:  xml.Name{"http://someurl", "bif"},
				Attrs: []xml.Attr{{xml.Name{"", "id"}, "2"}},
			},
			{Name: xml.Name{"", "doof"}, Value: "some data"},
		},
	}
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	if err := enc.Encode(tree, map[string]string{
		"http://someurl": "foo",
	}); err != nil {
		t.Fatal(err)
	}
	log.Print(buf.String())
}

func BenchmarkMarshal(b *testing.B) {
	tree := &Node{
		Name: xml.Name{"", "bar"},
		Nodes: []*Node{
			{Name: xml.Name{"http://someurl", "baz"}},
			{
				Name:  xml.Name{"http://someurl", "bif"},
				Attrs: []xml.Attr{{xml.Name{"", "id"}, "2"}},
			},
			{Name: xml.Name{"", "doof"}, Value: "some data"},
		},
	}
	enc := NewEncoder(ioutil.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(tree, map[string]string{"http://someurl": "foo"})
	}
}
