package plans

type Person interface {
	Name() string
	Initials() string
	OfficeSymbol() string
	Contact() string
}

type BasePerson struct {
	name         string
	initials     string
	officeSymbol string
	contact      string
}

func NewBasePerson(name, initials, officeSymbol, contact string) BasePerson {
	return BasePerson{
		name:         name,
		initials:     initials,
		officeSymbol: officeSymbol,
		contact:      contact,
	}
}

func (b BasePerson) Name() string {
	return b.name
}

func (b BasePerson) Initials() string {
	return b.initials
}

func (b BasePerson) OfficeSymbol() string {
	return b.officeSymbol
}

func (b BasePerson) Contact() string {
	return b.contact
}
