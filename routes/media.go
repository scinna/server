package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/log"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Medias(prv *services.Provider, r *mux.Router) {

	// Oh dear we are in trouble
	sr := r.PathPrefix("/").Subrouter()
	sr.Use(middlewares.LoggedInMiddleware(prv))
	sr.Use(middlewares.ContentTypeMiddleware)
	sr.HandleFunc("/upload", uploadMedia(prv))

	r.HandleFunc("/{media_id}", getMedia(prv))
}

func getMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mediaID := mux.Vars(r)["media_id"]
		if len(mediaID) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		media, err := dal.FindMedia(prv, mediaID)
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

			if dal.MediaBelongsToToken(prv, media, token) {
				serrors.NotOwner.Write(w)
				return
			}
		}

		http.ServeFile(w, r, file)
	}
}

func uploadMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		if user == nil {
			serrors.NoToken.Write(w)
			return
		}

		r.ParseMultipartForm(10 << 20) // @TODO Max upload size customizable

		title := r.FormValue("title")
		desc := r.FormValue("description")
		visib := r.FormValue("visibility")
		collection := r.FormValue("collection")
		visibInt, err := strconv.Atoi(visib)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if visibInt < 0 || visibInt > 2 {
			serrors.InvalidVisibility.Write(w)
			return
		}

		file, _, err := r.FormFile("picture")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Warn(fmt.Sprintf("An upload failed for user %v: %v", "user", err))
			return
		}
		defer file.Close()

		mime, _ := mimetype.DetectReader(file)
		if !mimetype.EqualsAny(mime.String(), utils.AllowedMimetypes...) {
			serrors.InvalidType.Write(w)
			return
		}

		parentFolder := prv.Config.MediaPath + "/" + user.UserID + "/"
		os.MkdirAll(parentFolder, os.ModePerm)

		pict := models.Media{
			Title:       title,
			Description: desc,
			Visibility:  visibInt,
			User:        user,
			Mimetype:    mime.String(),
		}

		err = dal.CreatePicture(prv, &pict, collection)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		outputFile, err := os.Create(parentFolder + pict.MediaID)
		if err != nil {
			serrors.WriteError(w, err)
			dal.DeleteMedia(prv, &pict)
			return
		}
		defer outputFile.Close()

		file.Seek(0, io.SeekStart)

		_, err = io.Copy(outputFile, file)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		str, _ := json.Marshal(pict)
		w.WriteHeader(http.StatusCreated)
		w.Write(str)
	}
}
