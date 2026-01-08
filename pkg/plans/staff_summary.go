package plans

import (
	"github.com/fxtlabs/date"
)

/*
% \actionOfficer{<NAME>}{<OFFICE_SYMBOL>}{<CONTACT>}
% \draftDate{<DRAFT_DATE>}
% \suspense{<SUSPENSE>}
% \staffCoord{<TO>,<ACTION>,<OUTCOME>,<DATE>}
*/
type StaffSummary struct {
	actionOfficer Person
	draftDate     date.Date
	suspense      date.Date
	coordChain    []CoordLink
	crm           []Comment
}
