package utils

import (
	"net/http"
	"strings"
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

// GetMimetypeForExt returns the mimetype that correspond to the extension
func GetMimetypeForExt(ext string) string {
	for k, v := range mimetypes {
		if v == ext {
			return k
		}
	}

	return ""
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
func ReadUserIP(headerIPField string, r *http.Request) string {
	ip := r.RemoteAddr
	if len(headerIPField) > 0 {
		ip = r.Header.Get(headerIPField)
	}

	ip = strings.Split(ip, ":")[0]

	return ip
}
