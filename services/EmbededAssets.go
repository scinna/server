package services

import (
	"embed"
	"io/fs"
)

type Assets struct {
	Frontend        *fs.FS
	Templates       *embed.FS
	LogoWide        []byte
	LogoSmall       []byte
	NotFoundPicture []byte
}
