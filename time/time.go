package time

import (
	"fmt"
	"time"
)

var (
	// TimeIntervalWeekly period of time  - week
	TimeIntervalWeekly = 1

	// TimeIntervalDaily  period of time  - day
	TimeIntervalDaily = 2
)

// TimeIntervalBlock block of time period
type TimeIntervalBlock struct {
	Begin time.Time
	End   time.Time
}

// TimeIntervalBlocks blocks of time period
type TimeIntervalBlocks struct {
	Model  string
	Blocks []*TimeIntervalBlock
}

// GetTimeNowRFC3339 time RFC3339 string of now
func GetTimeNowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

// GetTimeIntervalBlocks returns the datetime blocks of special model(weekly,daily) between begin and end params
// model {1:weekly , 2: daily}
func GetTimeIntervalBlocks(model int, begin, end time.Time) (blocks *TimeIntervalBlocks, err error) {

	if !intIsIn(model, TimeIntervalDaily) {
		return nil, fmt.Errorf("only support model, daily:%d, unsupported model:%d", TimeIntervalDaily, model)
	}

	if end.After(begin) {

		blocks = &TimeIntervalBlocks{}
		theDayAfterLastDayBegin := GetOneDayBeginOfTime(end).Add(24 * time.Hour)
		for temp := begin; temp.Before(theDayAfterLastDayBegin); temp = temp.Add(24 * time.Hour) {
			blocks.Blocks = append(blocks.Blocks, &TimeIntervalBlock{GetOneDayBeginOfTime(temp), GetOneDayEndOfTime(temp)})
		}
		return blocks, nil

	}

	return nil, fmt.Errorf("bad begin and end time, start must before end ")

}

// GetOneDayBeginOfTime returns the begin of the t
func GetOneDayBeginOfTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetOneDayEndOfTime returns the end of the t
func GetOneDayEndOfTime(t time.Time) time.Time {
	return GetOneDayBeginOfTime(t).Add(24 * time.Hour).Add(-1 * time.Nanosecond)
}

// TimeWeekOffset  to offset weeks of v
// eg. if v = 1, it returns t of next week, if v = -1 ,it returns last week
func TimeWeekOffset(t time.Time, v int64) time.Time {
	d := time.Duration(v) * 7 * 24 * time.Hour
	t = t.Add(d)
	return t
}

// TimeBeginningOfWeek return the begin of the week of t
// bSundayFirst means that many country use the monday as the first day of week
func TimeBeginningOfWeek(t time.Time, bSundayFirst bool) time.Time {

	weekday := int(t.Weekday())
	if !bSundayFirst {
		if weekday == 0 {
			weekday = 7
		}
		weekday = weekday - 1
	}

	d := time.Duration(-weekday) * 24 * time.Hour
	t = t.Add(d)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// TimeEndOfWeek return the end of the week of t
// bSundayFirst means that many country use the monday as the first day of week
func TimeEndOfWeek(t time.Time, bSundayFirst bool) time.Time {
	return TimeBeginningOfWeek(t, bSundayFirst).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// TimeBeginningOfMonth return the begin of the month of t
func TimeBeginningOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// TimeEndOfMonth return the end of the month of t
func TimeEndOfMonth(t time.Time) time.Time {
	return TimeBeginningOfMonth(t).AddDate(0, 1, -1)
}

// TimeSubDaysOfTwoDays return the days bewteen d1 and d2
func TimeSubDaysOfTwoDays(d1 time.Time, d2 time.Time) int64 {
	ds1 := GetOneDayBeginOfTime(d1)
	ds2 := GetOneDayBeginOfTime(d2)
	return int64(ds1.Sub(ds2).Hours() / 24)
}

func intIsIn(in int, arr ...int) bool {

	for i, j := 0, len(arr); i < j; i++ {
		if in == arr[i] {
			return true
		}
	}

	return false
}
