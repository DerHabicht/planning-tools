package task_generator

import (
	"github.com/fxtlabs/date"
	"github.com/google/uuid"
)

type TaskGenerator interface {
	UUID() uuid.UUID
	Domain() Domain
	TGN() string
	Super() TaskGenerator
	Title() string
	Initiated() *date.Date
	DispositionDate() *date.Date
	FinalDisposition() TGD
}
