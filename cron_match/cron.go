package cron_match

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type CronFilter struct {
	Start           cron.Schedule
	StartNthOfMonth int
	StartDayOfMonth int
	End             *cron.Schedule
	Duration        *time.Duration
}

func NewCronFilter(data map[string]any) (*CronFilter, error) {
	start, ok := data["start"].(string)
	if !ok {
		return nil, errors.New("start is required")
	}
	startParts := strings.Split(start, "#")
	schedule, err := cron.ParseStandard(startParts[0])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	startN := 0
	startDow := -1
	if len(startParts) > 1 {
		nth, err := strconv.Atoi(startParts[1])
		if err != nil {
			return nil, errors.WithMessage(err, "start nth of month is not a number")
		}
		if nth < 1 || nth > 5 {
			return nil, errors.New("start nth of month should be between 1 and 5")
		}

		cronParts := strings.Split(startParts[0], " ")
		dow := strings.ToLower(cronParts[4])
		switch dow {
		case "sun", "0":
			startDow = 0
		case "mon", "1":
			startDow = 1
		case "tue", "2":
			startDow = 2
		case "wed", "3":
			startDow = 3
		case "thu", "4":
			startDow = 4
		case "fri", "5":
			startDow = 5
		case "sat", "6":
			startDow = 6
		}
		startN = nth
	}

	var end *cron.Schedule
	endStr, ok := data["end"].(string)
	if ok {
		schedule, err := cron.ParseStandard(endStr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		end = &schedule

		if len(startParts) > 1 {
			return nil, errors.New("start end does not support nth of month, in case of reverse cron.")
		}
	}
	var duration *time.Duration
	durationStr, ok := data["duration"].(string)
	if ok {
		d, err := time.ParseDuration(durationStr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		duration = &d
	}

	if startN > 0 && startDow == -1 {
		return nil, errors.New("start nth of month requires a day of week")
	}
	return &CronFilter{
		Start:           schedule,
		StartNthOfMonth: startN,
		StartDayOfMonth: startDow,
		End:             end,
		Duration:        duration,
	}, nil
}

func (f CronFilter) Matches(tm time.Time) bool {
	if f.End != nil {
		return f.StartEndMatches(tm)
	}

	return f.StartDurationMatches(tm)

}

func (f CronFilter) StartEndMatches(tm time.Time) bool {
	st := f.Start.Next(tm)
	end := (*f.End).Next(tm)
	return st.After(end)
}

func (f CronFilter) StartDurationMatches(tm time.Time) bool {
	check := f.Start.Next(tm)
	minBefore := tm
	pSt := f.Start.Next(minBefore)
	for pSt.Equal(check) {
		minBefore = minBefore.Add(-1 * time.Minute)
		pSt = f.Start.Next(minBefore)
	}
	if pSt.Add(*f.Duration).Before(tm) || pSt.After(tm) {
		return false
	}
	return f.isNthOfMonth(pSt, f.StartNthOfMonth, f.StartDayOfMonth)
}

func (f CronFilter) isNthOfMonth(tm time.Time, n int, dow int) bool {
	count := 0
	for i := 1; i <= tm.Day(); i++ {
		if int(time.Date(tm.Year(), tm.Month(), i, 0, 0, 0, 0, time.UTC).Weekday()) == dow {
			count++
		}
	}
	return count == n
}
