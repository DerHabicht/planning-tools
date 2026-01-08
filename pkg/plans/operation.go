package plans

import (
	"time"
)

type Operation interface {
	OperationNumber() string
	OperationType() string
	OperationName() string
	OperationLocation() string
	OperationStart() time.Time
	OperationEnd() time.Time
	OperationLogos() [][]byte
	OperationCommander() Person
	OperationCommanderTitle() string
}
