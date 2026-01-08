package plans

type Plan interface {
	PlanType() PT
	InfosecLevel() IL
	VersionHistory() []Revision
	Sponsor() Organization
	PlanNumber() string
	SupplementID() string
	Operation() Operation
}
