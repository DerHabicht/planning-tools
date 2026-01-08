package plans

import (
	"github.com/fxtlabs/date"
)

type Revision struct {
	draft        bool
	version      string
	revDate      date.Date
	authors      []Person
	notes        string
	staffSummary StaffSummary
}
