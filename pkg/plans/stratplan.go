package plans

type Stratplan struct {
	infosecLevel   IL
	versionHistory []Revision
	sponsor        Organization
}

func (s *Stratplan) PlanType() PT {
	return PtSplan
}

func (s *Stratplan) InfosecLevel() IL {
	return s.infosecLevel
}

func (s *Stratplan) VersionHistory() []Revision {
	return s.versionHistory
}

func (s *Stratplan) Sponsor() Organization {
	return s.sponsor
}

func (s *Stratplan) PlanNumber() string {
	//TODO implement me
	panic("implement me")
}

func (s *Stratplan) SupplementID() string {
	//TODO implement me
	panic("implement me")
}

func (s *Stratplan) Operation() Operation {
	//TODO implement me
	panic("implement me")
}
