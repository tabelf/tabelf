package utils

import "strings"

// ValidWeekPeriod 拼接周一至周日的可用日期，传参为周一到周日的是否可用的数组.
func ValidWeekPeriod(period [7]bool) string {
	periodStr := [7]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	var l, r int
	left, right := -1, -1
	var weekdays []string
	for i := 0; i < 7; i++ {
		insert := false
		if period[i] {
			if left == -1 {
				left = i
			}
			right = i
			if i == 6 {
				l, r, insert = left, right, true
			}
		} else {
			if left == -1 {
				continue
			}
			l, r, insert = left, right, true
			left = -1
		}
		if insert {
			if l == r {
				weekdays = append(weekdays, periodStr[l])
			} else {
				weekdays = append(weekdays, periodStr[l]+"至"+periodStr[r])
			}
		}
	}
	if len(weekdays) == 0 {
		return ""
	}
	return strings.Join(weekdays, "、")
}
