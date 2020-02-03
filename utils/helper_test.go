package utils

import "testing"

func TestValidMimetype(t *testing.T) {
	if !IsValidMimetype("image/jpeg") {
		t.Errorf("IsValidMimetype should return true for \"image/jpeg\"")
	}
}

func TestInvalidMimetype(t *testing.T) {
	if IsValidMimetype("application/json") {
		t.Errorf("IsValidMimetype should return false for \"application/json\"")
	}
}

func TestInvalidMimetypeImage(t *testing.T) {
	if IsValidMimetype("image/json") {
		t.Errorf("IsValidMimetype should check on the whole mimetype and return false for \"image/json\"")
	}
}
