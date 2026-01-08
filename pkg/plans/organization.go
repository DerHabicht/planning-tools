package plans

type Organization struct {
	orgName   string
	unitName  string
	address   string
	orgSeal   []byte
	unitPatch []byte
}

func NewOrganization(
	orgName string,
	unitName string,
	address string,
	orgSeal []byte,
	unitPatch []byte,
) Organization {
	return Organization{
		orgName:   orgName,
		unitName:  unitName,
		address:   address,
		orgSeal:   orgSeal,
		unitPatch: unitPatch,
	}
}

func (o Organization) OrgName() string {
	return o.orgName
}

func (o Organization) UnitName() string {
	return o.unitName
}

func (o Organization) Address() string {
	return o.address
}

func (o Organization) OrgSeal() []byte {
	return o.orgSeal
}

func (o Organization) UnitPatch() []byte {
	return o.unitPatch
}
