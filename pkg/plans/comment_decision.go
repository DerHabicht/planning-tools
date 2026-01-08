package plans

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type CD int

const (
	NoDecision CD = iota
	AcceptDecision
	RejectDecision
	ModifyDecision
)

func ParseCD(s string) CD {
	switch strings.ToLower(s) {
	case "accept":
		return AcceptDecision
	case "reject":
		return RejectDecision
	case "modify":
		return ModifyDecision
	default:
		return NoDecision
	}
}

func (c CD) String() string {
	switch c {
	case AcceptDecision:
		return "Accept"
	case RejectDecision:
		return "Reject"
	case ModifyDecision:
		return "Modify"
	default:
		return ""
	}
}

func (cd CD) MarshalJSON() ([]byte, error) {
	return []byte(cd.String()), nil
}

func (cd *CD) UnmarshalJSON(raw []byte) error {
	val := ParseCD(string(raw))

	*cd = val
	return nil
}

func (cd CD) MarshalYAML() (interface{}, error) {
	return cd.String(), nil
}

func (cd *CD) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*cd = ParseCD(val)

	return nil
}

func (cd CD) Value() (driver.Value, error) {
	return []byte(cd.String()), nil
}

func (cd *CD) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val := ParseCD(s)

	*cd = val
	return nil
}
