package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
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

func Upload(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.LoggedInMiddleware(prv))
	r.Use(middlewares.Json)

	r.HandleFunc("", uploadMedia(prv)).Methods(http.MethodPost)
	r.HandleFunc("/shorten", shortenUrl(prv)).Methods(http.MethodPost)
}

func uploadMedia(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		err := r.ParseMultipartForm(10 << 20) // @TODO Max upload size customizable
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")
		collection := r.FormValue("collection")

		visibilityStr := r.FormValue("visibility")
		visibilityInt, err := strconv.Atoi(visibilityStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		visibility := models.VisibilityFromInt(visibilityInt)
		if !visibility.IsValid() {
			w.WriteHeader(http.StatusBadRequest)
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
			serrors.InvalidType.Write(w, r)
			return
		}

		parentFolder := prv.Config.MediaPath + "/" + user.UserID + "/"
		_ = os.MkdirAll(parentFolder, os.ModePerm)

		uid, err := prv.GenerateUID()
		if err != nil {
			serrors.WriteError(w, r, err)
			return
		}

		pict := models.Media{
			MediaID:     uid,
			MediaType:   models.MEDIA_PICTURE,
			Title:       title,
			Description: description,
			Visibility:  visibility,
			User:        user,
			Mimetype:    mime.String(),
		}

		err = prv.Dal.Medias.CreatePicture(&pict, collection)
		if err != nil {
			serrors.WriteError(w, r, err)
			return
		}

		outputFile, err := os.Create(parentFolder + pict.MediaID)
		if err != nil {
			serrors.WriteError(w, r, err)
			_ = prv.Dal.Medias.DeleteMedia(&pict)
			return
		}
		defer outputFile.Close()

		_, _ = file.Seek(0, io.SeekStart)

		_, err = io.Copy(outputFile, file)
		if err != nil {
			serrors.WriteError(w, r, err)
			return
		}

		err = pict.GenerateThumbnail(parentFolder + pict.MediaID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		str, _ := json.Marshal(pict)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(str)
	}
}

func shortenUrl(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		url := r.FormValue("url")
		if len(url) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		uid, err := prv.GenerateUID()
		if serrors.WriteError(w, r, err) {
			return
		}

		rootCollection, err := prv.Dal.Collections.FetchRoot(user)
		if serrors.WriteError(w, r, err) {
			return
		}

		shortenUrl := models.Media{
			MediaID:   uid,
			MediaType: models.MEDIA_SHORTURL,
			User:      user,
			CustomData: map[string]interface{}{
				"url": url,
			},
			Collection: rootCollection,
		}

		err = prv.Dal.Medias.CreateShortenUrl(&shortenUrl)
		if serrors.WriteError(w, r, err) {
			return
		}

		str, _ := json.Marshal(dto.GetShortenLinkInfo(&shortenUrl))
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(str)
	}
}
