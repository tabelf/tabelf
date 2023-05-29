package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/zaihui/go-hutils"
)

var ErrValidHour = errors.New("可用时段错误")

func ParseStartAndEndTime(start, end string) (startTime time.Time, endTime time.Time, err error) {
	now := time.Now()
	if start == "" {
		startTime = hutils.Time(now.Year(), now.Month(), now.Day(), 0, 0)
	} else {
		startTime, err = time.ParseInLocation(hutils.DateLayout, start, time.Local)
		if err != nil {
			return
		}
	}
	if end == "" {
		endTime = hutils.Time(now.Year(), now.Month(), now.Day(), 0, 0)
	} else {
		endTime, err = time.ParseInLocation(hutils.DateLayout, end, time.Local)
		if err != nil {
			return
		}
	}
	// 结束时间应该到这一天的结束
	endTime = endTime.AddDate(0, 0, 1)

	// 24h会出现计算了下一天的，再减1us
	oneSecondDuration, err := time.ParseDuration("-1us")
	endTime = endTime.Add(oneSecondDuration)
	return startTime, endTime, err
}

type Period struct {
	Start      time.Time
	End        time.Time
	IsCrossDay bool
}

func IsTimeOverlap(a, b Period) error {
	if a.IsCrossDay && (b.Start.After(a.Start) || b.Start.Before(a.End) || b.End.After(a.Start) || b.End.Before(a.End)) {
		return fmt.Errorf("%w: %s", ErrValidHour, "时间段重叠")
	} else if b.Start.After(a.Start) && b.Start.Before(a.End) || b.End.After(a.Start) && b.End.Before(a.End) {
		return fmt.Errorf("%w: %s", ErrValidHour, "时间段重叠")
	}
	return nil
}
