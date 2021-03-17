package main

import (
	"embed"
	"github.com/scinna/server/services"
	"io/fs"
)

//go:embed frontend/build
var frontend embed.FS

//go:embed assets/templates
var templates embed.FS

//go:embed assets/logo.png
var logoWide []byte

//go:embed assets/logo_small.png
var logoSmall []byte

//go:embed assets/not_found.png
var notFoundPicture []byte

func GetEmbeddedAssets() (*services.Assets, error){
	correctedFS, err := fs.Sub(frontend, "frontend/build")
	if err != nil {
		return nil, err
	}

	return &services.Assets{
		Frontend:        &correctedFS,
		Templates:       &templates,
		LogoWide:        logoWide,
		LogoSmall:       logoSmall,
		NotFoundPicture: notFoundPicture,
	}, nil
}