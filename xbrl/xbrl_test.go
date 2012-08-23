package xbrl

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

func MustParse(a, b string) time.Time {
	t, err := time.Parse(a, b)
	if err != nil {
		panic(err)
	}
	return t
}

func TestNothing(t *testing.T) {
	company := &Company{
		CIK:    "0000843006",
		URL:    "http://issuerdirect.com/",
		Ticker: "isdr",
	}
	cns := company.Namespace(time.Now())

	c1 := &Context{
		Start: MustParse("2006 01 02", "2009 12 31"),
	}
	u1 := &Unit{
		Name:    "USD",
		Measure: "iso4271:USD",
	}

	f1 := &Fact{Name: xml.Name{cns, "AssetsCurrent"}}
	f2 := &Fact{Name: xml.Name{cns, "OtherAssetsCurrent"}}
	f3 := &Fact{
		Name:        xml.Name{cns, "Assets"},
		Calculation: []*Calc{{1, f1}, {1, f2}},
	}

	//tree for first context
	data1 := &Data{
		Context: c1,
		Items: Items{
			{&Datum{f3, u1, "465005"}, Items{
				{&Datum{f1, u1, "323555"}, nil},
				{&Datum{f2, u1, "19201"}, nil},
			}},
		},
	}

	c := &Chart{
		Title: "DocumentSomethingOrOther",
		Data:  []*Data{data1},
	}

	f := Filing{
		Charts:  []*Chart{c},
		Date:    time.Now(),
		Company: company,
	}

	fmt.Printf("%#v\n", f)

	fmt.Printf("%+v\n", Serialize(f, Instance))
}
