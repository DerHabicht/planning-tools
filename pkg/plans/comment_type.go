package plans

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type CT int

const (
	CriticalComment CT = iota
	MajorComment
	SubstantiveComment
	AdminComment
)

func ParseCT(s string) (CT, error) {
	switch strings.ToLower(s) {
	case "critical":
		return CriticalComment, nil
	case "major":
		return MajorComment, nil
	case "substantive":
		return SubstantiveComment, nil
	case "admin":
		return AdminComment, nil
	default:
		panic(errors.Errorf("invalid CT %s", s))
	}
}

func (ct CT) String() string {
	switch ct {
	case CriticalComment:
		return "Critical"
	case MajorComment:
		return "Major"
	case SubstantiveComment:
		return "Substantive"
	case AdminComment:
		return "Admin"
	default:
		panic(errors.Errorf("invalid CT %d", ct))
	}
}

func (ct CT) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *CT) UnmarshalJSON(raw []byte) error {
	val, err := ParseCT(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*ct = val
	return nil
}

func (ct CT) MarshalYAML() (interface{}, error) {
	return ct.String(), nil
}

func (ct *CT) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*ct, err = ParseCT(val)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ct CT) Value() (driver.Value, error) {
	return []byte(ct.String()), nil
}

func (ct *CT) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseCT(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*ct = val
	return nil
}
