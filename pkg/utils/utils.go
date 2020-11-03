package utils

import (
	"time"
)

// GetCurrentMilliTime ...
func GetCurrentMilliTime() int64 {
	// 毫秒
	return time.Now().UnixNano() / 1e6
}
