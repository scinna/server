package dto

import (
	"github.com/scinna/server/models"
	"time"
)

type MediaInfo struct {
	MediaID string

	Title       string
	Description string

	Visibility  models.Visibility
	PublishedAt time.Time

	Collection string
	Author     string
}

func GetMediasInfos(media *models.Media) MediaInfo {
	return MediaInfo{
		MediaID:     media.MediaID,
		Title:       media.Title,
		Description: media.Description,
		Visibility:  media.Visibility,
		PublishedAt: media.PublishedAt,
		Collection:  media.Collection.Title,
		Author:      media.User.Name,
	}
}
