package plans

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type CA int

// Action: APP = Approve  COORD = Coordinate  INFO = Information  SIGN = Signature  REV = Review  POC = Point of contact  LOG = Log
const (
	ApproveAction CA = iota
	CoordAction
	InfoAction
	SignAction
	RevAction
	PocAction
	LogAction
)

func ParseCA(s string) (CA, error) {
	switch strings.ToLower(s) {
	case "appr":
		return ApproveAction, nil
	case "coord":
		return CoordAction, nil
	case "info":
		return InfoAction, nil
	case "sign":
		return SignAction, nil
	case "rev":
		return RevAction, nil
	case "poc":
		return PocAction, nil
	case "log":
		return LogAction, nil
	default:
		return -1, fmt.Errorf("invalid CA action: %s", s)
	}
}

func (ca CA) String() string {
	switch ca {
	case ApproveAction:
		return "Appr"
	case CoordAction:
		return "Coord"
	case InfoAction:
		return "Info"
	case SignAction:
		return "Sign"
	case RevAction:
		return "Rev"
	case PocAction:
		return "POC"
	case LogAction:
		return "Log"
	default:
		panic(errors.Errorf("invalid CA value: %d", ca))
	}
}
