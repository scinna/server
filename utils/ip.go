package utils

import (
	"github.com/scinna/server/config"
	"net/http"
	"strings"
)

func IPForRequest(cfg *config.Config, r *http.Request) string {
	ip := r.RemoteAddr
	if strings.Index(ip, ":") > 0 {
		ip = strings.Split(ip, ":")[0]
	}

	if len(cfg.RealIpHeader) > 0 {
		ip = r.Header.Get(cfg.RealIpHeader)
	}

	return ip
}
