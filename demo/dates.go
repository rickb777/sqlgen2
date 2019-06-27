package demo

import "github.com/rickb777/date"

// This example demonstrates
//   * dates and times stored in a variety of formats

//go:generate sqlgen -json -type demo.Dates -o dates_sql.go -all -v .

type Dates struct {
	Id      uint64          `sql:"pk: true, auto: true"`
	Integer date.Date       `sql:"type: integer"`
	String  date.DateString `sql:"type: text"`
	//Date      date.Date       `sql:"type: date"`
	//Timestamp date.Date       `sql:"type: timestamp"`
}

func NewDates(d date.Date) *Dates {
	return &Dates{
		Integer: d,
		String:  d.DateString(),
		//Date:      d,
		//Timestamp: d,
	}
}
