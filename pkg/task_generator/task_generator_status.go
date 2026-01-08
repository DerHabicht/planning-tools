package task_generator

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type TGS int

const (
	StatusBlue TGS = iota
	StatusGreen
	StatusYellow
	StatusRed
	StatusPurple
)

func ParseTGS(s string) (TGS, error) {
	switch strings.ToUpper(s) {
	case "BLUE":
		return StatusBlue, nil
	case "GREEN":
		return StatusGreen, nil
	case "YELLOW":
		return StatusYellow, nil
	case "RED":
		return StatusRed, nil
	case "PURPLE":
		return StatusPurple, nil
	default:
		return StatusBlue, errors.Errorf("invalid status: %s", s)
	}
}

func (tgs TGS) String() string {
	switch tgs {
	case StatusBlue:
		return "BLUE"
	case StatusGreen:
		return "GREEN"
	case StatusYellow:
		return "YELLOW"
	case StatusRed:
		return "RED"
	case StatusPurple:
		return "PURPLE"
	default:
		panic(errors.Errorf("invalid status value: %d", tgs))
	}
}

func (tgs TGS) MarshalJSON() ([]byte, error) {
	return []byte(tgs.String()), nil
}

func (tgs *TGS) UnmarshalJSON(raw []byte) error {
	val, err := ParseTGS(string(raw))
	if err != nil {
		return errors.WithStack(err)
	}

	*tgs = val
	return nil
}

func (tgs TGS) MarshalYAML() (interface{}, error) {
	return tgs.String(), nil
}

func (tgs *TGS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*tgs, err = ParseTGS(val)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (tgs TGS) Value() (driver.Value, error) {
	return []byte(tgs.String()), nil
}

func (tgs *TGS) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseTGS(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*tgs = val
	return nil
}
