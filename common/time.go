package common

import "time"

func Now() time.Time {
	return time.Now()
}

func NowAdd(hour int, minute int, second int) time.Time {
	return time.Now().Local().Add(time.Hour*time.Duration(hour) +
		time.Minute*time.Duration(minute) +
		time.Second*time.Duration(second))
}
