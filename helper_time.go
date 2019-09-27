package helper

import (
	"time"
)

func GetJakartaNowTime() time.Time {
	return time.Now().In(JakartaTime)
}