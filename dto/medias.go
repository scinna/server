package dto

import (
	"github.com/scinna/server/models"
	"time"
)

type MediaInfo struct {
	MediaID string `json:"mediaId"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Visibility  int       `json:"visibility"`
	PublishedAt time.Time `json:"publishedAt"`

	Author string `json:"author"`
}

func GetMediasInfos (media *models.Media) MediaInfo {
	return MediaInfo{
		MediaID:     media.MediaID,
		Title:       media.Title,
		Description: media.Description,
		Visibility:  media.Visibility,
		PublishedAt: media.PublishedAt,
		Author:      media.User.Name,
	}
}
