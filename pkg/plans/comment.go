package plans

import (
	"github.com/google/uuid"
)

type Comment struct {
	commentID         uuid.UUID
	item              uint
	commenter         Person
	commentType       CT
	text              string
	commentRationale  string
	decision          CD
	decisionRationale string
}
