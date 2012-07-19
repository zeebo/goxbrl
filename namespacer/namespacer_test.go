package namespacer

import (
	"bytes"
	"encoding/xml"
	"log"
	"testing"
)

func TestEncode(t *testing.T) {
	var buf bytes.Buffer
	enc := NewEncoder(&buf)

	type Address struct {
		XMLName     xml.Name
		City, State string
	}
	type Person struct {
		XMLName   xml.Name
		Id        int     `xml:"id,attr"`
		FirstName string  `xml:"name>first"`
		LastName  string  `xml:"name>last"`
		Age       int     `xml:"age"`
		Height    float32 `xml:"height,omitempty"`
		Married   bool
		Address   Address
		Comment   string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{xml.Name{"bar", "address"}, "Hanga Roa", "Easter Island"}

	v.XMLName.Space = "url"
	v.XMLName.Local = "person"

	nsmap := map[string]string{
		"url": "foo",
		"bar": "baz",
	}
	if err := enc.Encode(v, nsmap); err != nil {
		t.Fatal(err)
	}

	log.Println(buf.String())
}
