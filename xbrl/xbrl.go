package xbrl

import (
	"encoding/xml"
	"time"
)

type Company struct {
	CIK    string
	URL    string
	Ticker string
}

func (c *Company) Namespace(t time.Time) string {
	return c.URL + t.Format("20060102")
}

type ContextPeriod int

const (
	ContextPeriodDuration ContextPeriod = iota
	ContextPeriodInstant
)

type Context struct {
	Start time.Time
	End   time.Time
}

func (c *Context) PeriodType() ContextPeriod {
	if c.End.IsZero() {
		return ContextPeriodInstant
	}
	return ContextPeriodDuration
}

type Unit struct {
	Name    string
	Measure string
}

type Calc struct {
	Mul  int
	Fact *Fact
}

type Fact struct {
	Name        xml.Name
	LabelText   string
	Href        string
	Title       string
	Calculation []*Calc
}

type Datum struct {
	Fact  *Fact
	Unit  *Unit
	Value string
}

type Item struct {
	Datum    *Datum
	Children Items
}

type Items []*Item

type Data struct {
	Context *Context
	Items   Items
}

type Chart struct {
	Title string
	Data  []*Data
}

type Filing struct {
	Charts  []*Chart
	Date    time.Time
	Company *Company
}
