package xbrl

import "github.com/zeebo/goxbrl/marshal"

type DocType int

const (
	Instance DocType = iota
)

func Serialize(f Filing, doc DocType) *marshal.Node {
	switch doc {
	case Instance:
		return serializeInstance(f)
	}
	panic("Unknown document type")
}
