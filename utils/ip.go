package utils

import (
	"github.com/scinna/server/log"
	"net"
	"net/http"
	"strings"
)

func IPForRequest(r *http.Request) string {
	// @TODO: Write explicitely in the wiki that scinna MUST be used behind a reverse-proxy
	// @TODO: and that it MUST set the X-Forwarded-For field ITSELF AND REMOVE ANY EXISTING ONE COMING FROM THE USER
	// Otherwise the client could craft a custom header to spoof as someone else
	// This should NOT be handled in any other way
	if len(r.Header.Get("X-Forwarded-For")) > 0 {
		return strings.Split(r.Header.Get("X-Forwarded-For"), ", ")[0]
	} else {
		log.Fatal("SECURITY ISSUE! PLEASE VERIFY YOUR REVERSE PROXY SETTINGS. IT MUST SET AND REPLACE THE X-Forwarded-For HEADER !")
		log.Fatal("If you are not using a reverse proxy please consider so. Scinna was NOT designed to work without one")
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
