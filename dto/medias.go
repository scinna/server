package dto

import (
	"github.com/scinna/server/models"
	"time"
)

type MediaInfo struct {
	MediaID   string
	MediaType int

	Title       string
	Description string

	Visibility  models.Visibility
	PublishedAt time.Time

	ViewCount  int
	Collection string
	Author     string
}

func GetMediasInfos(media *models.Media) MediaInfo {
	return MediaInfo{
		MediaID:     media.MediaID,
		MediaType:   media.MediaType,
		Title:       media.Title,
		Description: media.Description,
		Visibility:  media.Visibility,
		PublishedAt: media.PublishedAt,
		ViewCount:   media.ViewCount,
		Collection:  media.Collection.Title,
		Author:      media.User.Name,
	}
}

type ShortenLinkInfo struct {
	MediaID     string
	MediaType   int
	Url         string
	AccessCount int
	PublishedAt time.Time
	Author      string
}

func GetShortenLinkInfo(media *models.Media) ShortenLinkInfo {
	return ShortenLinkInfo{
		MediaID:     media.MediaID,
		MediaType:   media.MediaType,
		Url:         media.CustomData["url"].(string),
		PublishedAt: media.PublishedAt,
		AccessCount: media.ViewCount,
		Author:      media.User.Name,
	}
}
