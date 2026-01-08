package task_generator

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTG(t *testing.T) {
	tg, err := NewTG(uuid.Nil, "C", "25-P-080", "GTL Mission Ready")
	assert.NoError(t, err)

	assert.NotEqual(t, uuid.Nil, tg.UUID())
	assert.Equal(t, Project, tg.TGT())
	assert.Equal(t, Volunteering, tg.Domain())
	assert.Equal(t, "25-P-080", tg.TGN())
	assert.Equal(t, "GTL Mission Ready", tg.Title())
	assert.Nil(t, tg.Initiated())
	assert.Nil(t, tg.DispositionDate())
	assert.Equal(t, Incomplete, tg.FinalDisposition())
}
