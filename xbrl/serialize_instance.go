package xbrl

import (
	"encoding/xml"
	"github.com/zeebo/goxbrl/marshal"
)

func (s *serializer) instance() *marshal.Node {
	return &marshal.Node{
		Name:  xml.Name{"", ""},
		Attrs: []xml.Attr{},
		Nodes: []*marshal.Node{},
		Value: "",
	}
}
