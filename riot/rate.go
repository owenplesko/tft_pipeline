package riot

import (
	"time"

	"golang.org/x/time/rate"
)

var regionalRateLimit *rate.Limiter
var continentalRateLimit *rate.Limiter

func InitRateLimit(requests int, interval time.Duration) {
	regionalRateLimit = rate.NewLimiter(rate.Every(interval/time.Duration(requests)), 1)
	continentalRateLimit = rate.NewLimiter(rate.Every(interval/time.Duration(requests)), 1)
}
