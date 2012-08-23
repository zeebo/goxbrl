package xbrl

import "github.com/zeebo/goxbrl/marshal"

type DocType int

const (
	Instance DocType = iota
)

type serializer struct {
	f Filing
}

func Serialize(f Filing, doc DocType) *marshal.Node {
	s := &serializer{f}
	switch doc {
	case Instance:
		return s.instance()
	}
	panic("Unknown document type")
}
