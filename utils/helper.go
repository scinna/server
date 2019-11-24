package utils

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
