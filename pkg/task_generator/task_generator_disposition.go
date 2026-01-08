package task_generator

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type TGD int

const (
	Incomplete TGD = iota
	Done
	WontDo
	Merged
)

func ParseTGD(dt string) TGD {
	switch strings.ToUpper(dt) {
	case "DONE":
		return Done
	case "WONT DO":
		return WontDo
	case "MERGED":
		return Merged
	default:
		return Incomplete
	}
}

func (tgd TGD) String() string {
	switch tgd {
	case Done:
		return "DONE"
	case WontDo:
		return "WONT DO"
	case Merged:
		return "MERGED"
	default:
		return ""
	}
}

func (tgd TGD) MarshalJSON() ([]byte, error) {
	return []byte(tgd.String()), nil
}

func (tgd *TGD) UnmarshalJSON(raw []byte) error {
	val := ParseTGD(string(raw))

	*tgd = val
	return nil
}

func (tgd TGD) MarshalYAML() (interface{}, error) {
	return tgd.String(), nil
}

func (tgd *TGD) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	err := unmarshal(&val)
	if err != nil {
		return errors.WithStack(err)
	}

	*tgd = ParseTGD(val)

	return nil
}

func (tgd TGD) Value() (driver.Value, error) {
	return []byte(tgd.String()), nil
}

func (tgd *TGD) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val := ParseTGD(s)

	*tgd = val
	return nil
}
