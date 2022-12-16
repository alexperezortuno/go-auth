package common

import (
	"github.com/alexperezortuno/go-auth/common/environment"
	"time"
)

var env = environment.Server()

func Now() time.Time {
	return time.Now()
}

func NowAdd(hour int, minute int, second int) time.Time {
	return time.Now().Local().Add(time.Hour*time.Duration(hour) +
		time.Minute*time.Duration(minute) +
		time.Second*time.Duration(second))
}

func PrimaryCacheDuration() time.Duration {
	return time.Duration(env.TokenLifeTime) * time.Minute
}

func SecondaryCacheDuration() time.Duration {
	return time.Duration(env.RefreshTokenLifeTime) * time.Hour
}
