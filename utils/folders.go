package utils

import "strings"

// SplitPath returns an array that contains the hierarchy to access to a folder
func SplitPath(path string) []string {
	return strings.Split(path, "/")
}

// StripHeadSlash removes the first slash
func StripHeadSlash(path string) string {
	r := []rune(path)

	if len(r) > 0 {
		if r[0] == '/' {
			return string(r[1:len(r)])
		}
	}

	return path
}
