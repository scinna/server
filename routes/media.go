package routes

import (
	"github.com/gorilla/mux"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"net/http"
	"os"
)

func Medias(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/{media_id}", getMedia(prv))
}

func getMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := prv.Dal.Medias.FindMedia(mediaID)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		file := prv.Config.MediaPath + media.Path
		if _, err := os.Stat(file); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if media.Visibility == 2 {
			token, err := middlewares.GetTokenFromRequest(r)
			if err != nil {
				serrors.WriteError(w, err)
				return
			}

			if prv.Dal.Medias.MediaBelongsToToken(media, token) {
				serrors.NotOwner.Write(w)
				return
			}
		}

		http.ServeFile(w, r, file)
	}
}
