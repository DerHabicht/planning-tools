package plans

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type PT int

const (
	PtStratplan PT = iota
	PtCamplan
	PtQplan
	PtBplan
	PtOplan
	PtOpord
	PtFragord
	PtSplan
	PtIap
)

func ParsePT(s string) (PT, error) {
	switch strings.ToUpper(s) {
	case "STRATPLAN":
		return PtStratplan, nil
	case "CAMPLAN":
		return PtCamplan, nil
	case "QPLAN":
		return PtQplan, nil
	case "BPLAN":
		return PtBplan, nil
	case "OPLAN":
		return PtOplan, nil
	case "OPORD":
		return PtOpord, nil
	case "FRAGORD":
		return PtFragord, nil
	case "SPLAN":
		return PtSplan, nil
	case "IAP":
		return PtIap, nil
	default:
		return -1, errors.Errorf("invalid plan type: %s", s)
	}
}

func (p PT) String() string {
	switch p {
	case PtStratplan:
		return "STRATPLAN"
	case PtCamplan:
		return "CAMPLAN"
	case PtQplan:
		return "QPLAN"
	case PtBplan:
		return "BPLAN"
	case PtOplan:
		return "OPLAN"
	case PtOpord:
		return "OPORD"
	case PtFragord:
		return "FRAGORD"
	case PtSplan:
		return "SPLAN"
	case PtIap:
		return "IAP"
	default:
		panic(errors.Errorf("invalid PT value: %d", p))
	}
}

func (pt PT) MarshalJSON() ([]byte, error) {
	return []byte(pt.String()), nil
}

func (pt *PT) UnmarshalJSON(raw []byte) error {
	val, err := ParsePT(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*pt = val
	return nil
}

func (pt PT) MarshalYAML() (interface{}, error) {
	return pt.String(), nil
}

func (pt *PT) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*pt, err = ParsePT(val)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (pt PT) Value() (driver.Value, error) {
	return []byte(pt.String()), nil
}

func (pt *PT) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParsePT(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*pt = val
	return nil
}
