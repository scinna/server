package services

import (
	"github.com/scinna/server/config"
)

type Provider struct {

}

func NewProvider(cfg *config.Config) *Provider {
	return &Provider{}
}
