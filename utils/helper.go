package utils

import (
	"net/http"

	"github.com/oxodao/scinna/services"
)

var mimetypes = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/gif":  "gif",
}

// IsValidMimetype returns true if the mimetype is allowed to be uploaded to server (Image files)
func IsValidMimetype(mt string) bool {
	for k := range mimetypes {
		if k == mt {
			return true
		}
	}
	return false
}

// GetExtForMimetype returns the extension that a file usually have for the given mimetype
func GetExtForMimetype(mt string) string {
	return mimetypes[mt]
}

var validVisibility = [...]int8{
	// Public
	0,
	// Unlisted
	1,
	// Private
	2,
}

// IsValidVisibility returns whether the visibility exists or not
func IsValidVisibility(vis int8) bool {
	for v := 0; v < len(validVisibility); v++ {
		if validVisibility[v] == vis {
			return true
		}
	}
	return false
}

// ReadUserIP retreive the IP from the client, whether its coming directly or through a properly configured reverse-proxy
func ReadUserIP(prv *services.Provider, r *http.Request) string {
	if len(prv.HeaderIPField) > 0 {
		return r.Header.Get(prv.HeaderIPField)
	}

	return r.RemoteAddr
}
