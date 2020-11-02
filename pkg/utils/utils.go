package utils

import (
	"time"
)

// GetCurrentMilliTime ...
func GetCurrentMilliTime() int64 {

	return time.Now().UnixNano() / 1e6
}
