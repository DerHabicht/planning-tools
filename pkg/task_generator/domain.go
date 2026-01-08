package task_generator

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type Domain int

const (
	Home Domain = iota
	Career
	Volunteering
)

func ParseDomain(s string) (Domain, error) {
	switch strings.ToUpper(s) {
	case "A", "HOME":
		return Home, nil
	case "B", "CAREER":
		return Career, nil
	case "C", "VOLUNTERING":
		return Volunteering, nil
	default:
		return 0, errors.Errorf("invalid domain: %s", s)

	}
}

func (d Domain) Code() string {
	switch d {
	case Home:
		return "A"
	case Career:
		return "B"
	case Volunteering:
		return "C"
	default:
		panic(errors.Errorf("unknown domain: %d", d))
	}
}

func (d Domain) String() string {
	switch d {
	case Home:
		return "HOME"
	case Career:
		return "CAREER"
	case Volunteering:
		return "VOLUNTEERING"
	default:
		panic(errors.Errorf("unknown domain: %d", d))
	}
}

func (d Domain) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Domain) UnmarshalJSON(raw []byte) error {
	val, err := ParseDomain(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*d = val
	return nil
}

func (d Domain) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *Domain) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*d, err = ParseDomain(val)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d Domain) Value() (driver.Value, error) {
	return []byte(d.String()), nil
}

func (d *Domain) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseDomain(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = val
	return nil
}
