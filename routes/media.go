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
	"strconv"
)

func Medias(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/{media_id}", getMedia(prv)).Methods(http.MethodGet)
	r.HandleFunc("/{media_id}", editMedia(prv)).Methods(http.MethodPut)
	r.HandleFunc("/{media_id}", deleteMedia(prv)).Methods(http.MethodDelete)
	r.HandleFunc("/{media_id}/thumbnail", getThumbnail(prv)).Methods(http.MethodGet)
	r.HandleFunc("/{media_id}/infos", getMediaInfos(prv)).Methods(http.MethodGet)
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
			_ = prv.Dal.Medias.IncrementViewCount(media)
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

			_ = prv.Dal.Medias.IncrementViewCount(media)
			http.ServeFile(w, r, file)
			return
		}
	}
}

func getThumbnail(prv *services.Provider) http.HandlerFunc {
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

		if media.MediaType != models.MEDIA_PICTURE && media.MediaType != models.MEDIA_VIDEO {
			serrors.MediaNoThumbnail.Write(w, r)
			return
		}

		file := prv.Config.MediaPath + media.Path + "_thumb"
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

func editMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, err := middlewares.GetUserFromRequest(prv, r)
		if serrors.WriteError(w, r, err) {
			return
		}

		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := prv.Dal.Medias.Find(mediaID)
		if serrors.WriteError(w, r, err) {
			return
		}

		if media.User.UserID != user.UserID {
			serrors.NotOwner.Write(w, r)
			return
		}

		switch media.MediaType {
		case models.MEDIA_SHORTURL:
			// Shortened URL should not be modifiable
			// We don't want users to edit where they're pointing to after creation
			// And they do not have title nor description
			return
		case models.MEDIA_PICTURE:
			fallthrough
		case models.MEDIA_VIDEO:
			title := r.FormValue("title")
			description := r.FormValue("description")

			visibilityStr := r.FormValue("visibility")
			visibilityInt, err := strconv.Atoi(visibilityStr)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if len(title) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			visibility := models.VisibilityFromInt(visibilityInt)
			if !visibility.IsValid() {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			media.Title = title
			media.Description = description
			media.Visibility = visibility

			err = prv.Dal.Medias.UpdateMedia(media)
			if serrors.WriteError(w, r, err) {
				return
			}

			bytes, _ := json.Marshal(dto.GetMediasInfos(media))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(bytes)
		}
	}
}

func deleteMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, err := middlewares.GetUserFromRequest(prv, r)
		if serrors.WriteError(w, r, err) {
			return
		}

		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := prv.Dal.Medias.Find(mediaID)
		if serrors.WriteError(w, r, err) {
			return
		}

		if media.User.UserID != user.UserID {
			serrors.NotOwner.Write(w, r)
			return
		}

		err = prv.Dal.Medias.DeleteMedia(media)
		if serrors.WriteError(w, r, err) {
			return
		}

		w.WriteHeader(http.StatusGone)
	}
}
