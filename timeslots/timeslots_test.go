package timeslots

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsTimePeriodValid(t *testing.T) {
	t.Run("time period is valid", func(t *testing.T) {
		testTime := time.Date(2020, 11, 9, 7, 0, 0, 0, time.UTC)
		valid := isTimePeriodValid(TimePeriod{From: testTime, To: testTime.Add(time.Hour * 22 * 7)})

		assert.True(t, valid)
	})

	t.Run("time period is not valid", func(t *testing.T) {
		testTime := time.Date(2020, 11, 9, 7, 0, 0, 0, time.UTC)
		valid := isTimePeriodValid(TimePeriod{From: testTime.Add(time.Hour * 22 * 7), To: testTime})

		assert.False(t, valid)
	})
}
