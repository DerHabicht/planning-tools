package plancal

import (
	"testing"
	"time"

	"github.com/ag7if/calendar/location"
	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestCalendar_FetchWeek(t *testing.T) {
	loc, err := location.FromMGRS("KSLC", "12TVL1750415720", location.AcpT)
	assert.NoError(t, err)

	cal := NewCalendar(2025, loc, date.New(1988, time.September, 27))

	wk, err := cal.FetchWeek(2025, 27)
	assert.NoError(t, err)
	assert.Equal(t, "+2025-06-30", wk.StartDay().ISODate())

	wk, err = cal.FetchWeek(2024, 1)
	assert.NoError(t, err)
	assert.Equal(t, "+2024-01-01", wk.StartDay().ISODate())
}
