package config

import (
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
	"time"
)

type RateLimit struct {
	Amount   uint64
	Interval time.Duration
}

func (rl *RateLimit) NewStore() (*limiter.Store, *httplimit.Middleware, error) {
	currentLimiter, err := memorystore.New(&memorystore.Config{
		Tokens:   rl.Amount,
		Interval: rl.Interval * time.Minute,
	})
	if err != nil {
		return nil, nil, err
	}

	middleware, err := httplimit.NewMiddleware(currentLimiter, httplimit.IPKeyFunc("X-Forwarded-For"))
	return &currentLimiter, middleware, err
}

type RateLimiter struct {
	Registration      RateLimit
	Login             RateLimit
	ForgottenPassword RateLimit
}

func (rl *RateLimit) validOrDefaults(amount uint64, interval time.Duration) {
	if rl.Amount <= 0 {
		rl.Amount = amount
	}

	if rl.Interval <= 0 {
		rl.Interval = interval
	}
}

func (rl *RateLimiter) Validate() error {
	// Setting defaults
	rl.Registration.validOrDefaults(10, 30)
	rl.Login.validOrDefaults(10, 5)
	rl.ForgottenPassword.validOrDefaults(3, 15)

	return nil
}
