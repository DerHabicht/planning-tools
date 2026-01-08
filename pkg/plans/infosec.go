package plans

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type IL int

const (
	NotSensitive IL = iota
	SomewhatSensitive
	Sensitive
	VerySensitive
)

func ParseIL(s string) (IL, error) {
	switch strings.ToUpper(s) {
	case "NS", "NOT SENSITIVE":
		return NotSensitive, nil
	case "SS", "SOMEWHAT SENSITIVE":
		return SomewhatSensitive, nil
	case "SE", "SENSITIVE":
		return Sensitive, nil
	case "VS", "VERY SENSITIVE":
		return VerySensitive, nil
	default:
		return -1, fmt.Errorf("unknown IL value: %s", s)
	}
}

func (i IL) String() string {
	switch i {
	case NotSensitive:
		return "NOT SENSITIVE"
	case SomewhatSensitive:
		return "SOMEWHAT SENSITIVE"
	case Sensitive:
		return "SENSITIVE"
	case VerySensitive:
		return "VERY SENSITIVE"
	default:
		panic(fmt.Sprintf("unknown IL value: %d", i))
	}
}

func (i IL) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *IL) UnmarshalJSON(raw []byte) error {
	val, err := ParseIL(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*i = val
	return nil
}

func (i IL) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *IL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*i, err = ParseIL(val)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (i IL) Value() (driver.Value, error) {
	return []byte(i.String()), nil
}

func (i *IL) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseIL(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*i = val
	return nil
}
