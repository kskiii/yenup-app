package rate

import (
	"time"
)

type Rate struct {
	Base   string
	Target string
	Value  float64
	Date   time.Time
}
