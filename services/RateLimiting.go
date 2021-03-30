package services

import (
	"github.com/scinna/server/config"
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/httplimit"
)

type RateLimiting struct {
	RegistrationLimiter      limiter.Store
	LoginLimiter             limiter.Store
	ForgottenPasswordLimiter limiter.Store

	RegistrationMiddleware      *httplimit.Middleware
	LoginMiddleware             *httplimit.Middleware
	ForgottenPasswordMiddleware *httplimit.Middleware
}

func NewRateLimiting(cfg *config.Config) (*RateLimiting, error) {
	storeRegister, middlewareRegister, err := cfg.RateLimiter.Registration.NewStore()
	if err != nil {
		return nil, err
	}

	storeLogin, middlewareLogin, err := cfg.RateLimiter.Login.NewStore()
	if err != nil {
		return nil, err
	}

	storeForgottenPassword, middlewareForgottenPassword, err := cfg.RateLimiter.ForgottenPassword.NewStore()
	if err != nil {
		return nil, err
	}

	return &RateLimiting{
		RegistrationLimiter:      *storeRegister,
		LoginLimiter:             *storeLogin,
		ForgottenPasswordLimiter: *storeForgottenPassword,

		RegistrationMiddleware:      middlewareRegister,
		LoginMiddleware:             middlewareLogin,
		ForgottenPasswordMiddleware: middlewareForgottenPassword,
	}, nil
}
