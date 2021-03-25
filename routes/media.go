package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"net/http"
	"os"
)

func Medias(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/{media_id}", getMedia(prv))
	r.HandleFunc("/{media_id}/infos", getMediaInfos(prv))
}

func getMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := prv.Dal.Medias.Find(mediaID)
		if err != nil {
			serrors.WriteError(w, r, err)
			return
		}

		switch media.MediaType {
		case models.MEDIA_SHORTURL:
			http.Redirect(w, r, media.CustomData["url"].(string), 301)
			return

		case models.MEDIA_PICTURE:
			file := prv.Config.MediaPath + media.Path
			if _, err := os.Stat(file); os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			if media.Visibility.IsPrivate() {
				token, err := middlewares.GetTokenFromRequest(r)
				if err != nil {
					if err == serrors.NoToken {
						serrors.WriteError(w, r, serrors.NotOwner)
						return
					}
					serrors.WriteError(w, r, err)
					return
				}

				if !prv.Dal.Medias.MediaBelongsToToken(media, token) {
					serrors.NotOwner.Write(w, r)
					return
				}
			}

			http.ServeFile(w, r, file)
			return
		}
	}
}

func getMediaInfos(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := prv.Dal.Medias.Find(mediaID)
		if serrors.WriteError(w, r, err) {
			return
		}

		switch media.MediaType {
		case models.MEDIA_SHORTURL:
			jsonBytes, _ := json.Marshal(dto.GetShortenLinkInfo(media))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(jsonBytes)
			return
		case models.MEDIA_PICTURE:
			if media.Visibility.IsPrivate() {
				token, err := middlewares.GetTokenFromRequest(r)
				if err != nil {
					if err == serrors.NoToken {
						serrors.WriteError(w, r, serrors.NotOwner)
						return
					}

					serrors.WriteError(w, r, err)
					return
				}

				if !prv.Dal.Medias.MediaBelongsToToken(media, token) {
					serrors.NotOwner.Write(w, r)
					return
				}
			}

			mediaDto := dto.GetMediasInfos(media)

			jsonBytes, _ := json.Marshal(mediaDto)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(jsonBytes)

			return
		}
	}
}
