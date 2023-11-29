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

func TestAreTimeslotsOverlapping(t *testing.T) {
	t.Run("time periods are not overlapping", func(t *testing.T) {
		now := time.Now()
		twoHoursLater := time.Now().Add(time.Hour * 2)
		threeHoursLater := time.Now().Add(time.Hour * 3)

		t1 := TimePeriod{From: now, To: twoHoursLater}
		t2 := TimePeriod{From: threeHoursLater, To: threeHoursLater}
		testTime := []TimePeriod{t1, t2}
		valid := areTimePeriodsOverlapping(testTime)

		assert.False(t, valid)
	})
	t.Run("time periods are overlapping", func(t *testing.T) {
		now := time.Now()
		hourLater := time.Now().Add(time.Hour * 1)
		twoHoursLater := time.Now().Add(time.Hour * 2)

		t1 := TimePeriod{From: now, To: hourLater}
		t2 := TimePeriod{From: now, To: twoHoursLater}
		testTime := []TimePeriod{t1, t2}
		valid := areTimePeriodsOverlapping(testTime)

		assert.True(t, valid)
	})
	t.Run("one time period starts when previous one ends", func(t *testing.T) {
		now := time.Now()
		twoHoursLater := time.Now().Add(time.Hour * 2)
		threeHoursLater := time.Now().Add(time.Hour * 3)

		t1 := TimePeriod{From: now, To: twoHoursLater}
		t2 := TimePeriod{From: twoHoursLater, To: threeHoursLater}
		testTime := []TimePeriod{t1, t2}
		valid := areTimePeriodsOverlapping(testTime)

		assert.True(t, valid)
	})
}

func TestTimeperiodsBelongToTheDay(t *testing.T) {
	t.Run("time periods belong to today", func(t *testing.T) {
		now := time.Now()

		t1 := TimePeriod{From: time.Now(), To: time.Now().Add(time.Hour * 1)}
		t2 := TimePeriod{From: time.Now().Add(time.Hour * 1), To: time.Now().Add(time.Hour * 2)}

		testTime := []TimePeriod{t1, t2}
		valid := timeperiodsBelongToTheDay(testTime, now)

		assert.True(t, valid)
	})

	t.Run("time periods don't belong to today", func(t *testing.T) {
		now := time.Now()
		nextYear := time.Now().AddDate(1, 1, 1)

		t1 := TimePeriod{From: nextYear, To: nextYear}

		testTime := []TimePeriod{t1}
		valid := timeperiodsBelongToTheDay(testTime, now)

		assert.False(t, valid)
	})

}
