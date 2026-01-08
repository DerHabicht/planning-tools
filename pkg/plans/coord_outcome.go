package plans

type CO int

const (
	NoOutcome CO = iota
	ConcurOutcome
	ConcurWCommentOutcome
	NonConcurOutcome
)
