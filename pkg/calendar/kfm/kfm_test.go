package kfm

import (
	"testing"

	"github.com/ag7if/go-files"
	"github.com/fxtlabs/date"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar/plancal"
)

func TestGenerateBlankKFM(t *testing.T) {
	bd, err := date.ParseISO(config.GetString(config.Birthday))
	assert.NoError(t, err)
	cal := plancal.NewCalendar(2026, bd)

	kfm := GenerateBlankKFM(cal)

	doc, err := yaml.Marshal(kfm)
	assert.NoError(t, err)

	f, err := files.NewFile("kfm.yaml")
	assert.NoError(t, err)
	err = f.WriteBytes(doc)
	assert.NoError(t, err)
}
