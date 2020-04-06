package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"

	"github.com/scinna/server/auth"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

// RawMediaRoute is the route that render the media: /{media id}
func RawMediaRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["media"]

		if len(id) == 0 {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		p, err := dal.GetMedia(prv, id)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		if p.Visibility == 2 {
			user, err := auth.ValidateRequest(prv, w, r)
			if err != nil || p.Creator.ID == user.ID {
				serrors.ErrorPrivateMedia.Write(w)
				return
			}
		}

		pictFile, err := os.Open(prv.Config.MediaPath + "/" + strconv.FormatInt(*p.Creator.ID, 10) + "/" + strconv.FormatInt(*p.ID, 10) + "." + p.Ext)
		if err != nil {
			serrors.ErrorMediaNotFound.Write(w)
			return
		}
		defer pictFile.Close()

		w.Header().Set("Content-Type", utils.GetMimetypeForExt(p.Ext))
		io.Copy(w, pictFile)
	}
}

// MediaInfoRoute returns the informations of the media like author, date, visibility, ...
func MediaInfoRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["URL_ID"]

		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := dal.GetMedia(prv, id)
		if serrors.WriteError(w, err) {
			return
		}

		user, err := auth.ValidateRequest(prv, w, r)
		if p.Visibility == 2 {
			if err == serrors.ErrorNoToken && serrors.WriteError(w, serrors.ErrorPrivateMedia) {
				return
			}

			if *user.ID != *p.Creator.ID {
				serrors.ErrorPrivateMedia.Write(w)
				return
			}
		}

		if err != nil || *user.ID != *p.Creator.ID {
			p.Creator = &model.AppUser{
				Username: p.Creator.Username,
			}
		}

		json, err := json.Marshal(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(json)

	}
}

// MoveMediaRoute lets the user move the media to another folder
func MoveMediaRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Not implemented yet
	}
}

// EditMediaRoute lets the user change the name, description and visibility of a media
func EditMediaRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Not implemented yet
	}
}

// UploadMediaRoute is the route that let user upload a media
func UploadMediaRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		r.ParseMultipartForm(10 << 20) // 10 meg max

		title := r.FormValue("title")
		desc := r.FormValue("description")
		visib := r.FormValue("visibility")
		visibInt, err := strconv.Atoi(visib)
		visibility := int8(visibInt)

		if len(title) == 0 || len(title) > 128 || err != nil || !utils.IsValidVisibility(visibility) {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		file, _, err := r.FormFile("media")
		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}
		defer file.Close()

		mimeType, _, err := mimetype.DetectReader(file)
		if err != nil || !utils.IsValidMimetype(mimeType) {
			serrors.ErrorInvalidMimetype.Write(w)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		parentFolder := prv.Config.MediaPath + "/" + strconv.FormatInt(*user.ID, 10) + "/"

		_, err = os.Stat(parentFolder)
		if os.IsNotExist(err) {
			err = os.MkdirAll(parentFolder, os.ModePerm)
			if err != nil {
				serrors.WriteLoggableError(w, err)
				return
			}
		}

		media := model.Media{
			Title:       title,
			Description: desc,
			Creator:     &user,
			Visibility:  visibility,
			Ext:         utils.GetExtForMimetype(mimeType),
		}

		media, err = dal.CreateMedia(prv, media)
		if serrors.WriteError(w, err) {
			return
		}

		outputFile, err := os.Create(parentFolder + strconv.FormatInt(*media.ID, 10) + "." + media.Ext)
		if err != nil {
			serrors.WriteLoggableError(w, err)
			dal.DeleteMedia(prv, media)
			return
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		// Clearing out the fields we don't want to send
		media.ID = nil
		media.Ext = ""

		json, err := json.Marshal(media)
		if err != nil {
			// The media is uploaded but something went wrong while encoding the response
			w.WriteHeader(http.StatusAccepted)
			return
		}

		w.Write(json)
	}
}

// DeleteMediaRoute is the route that let the user delete one OR MULTIPLE of his media
func DeleteMediaRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		params := mux.Vars(r)
		id := params["URL_ID"]

		if len(id) == 0 {
			serrors.ErrorMissingURLID.Write(w)
			return
		}

		p, err := dal.GetMedia(prv, id)
		if serrors.WriteError(w, err) {
			return
		}

		if *p.Creator.ID != *user.ID && user.Role != model.UserRoleAdmin {
			serrors.ErrorWrongOwner.Write(w)
			return
		}

		err = os.Remove(prv.Config.MediaPath + "/" + strconv.FormatInt(*p.Creator.ID, 10) + "/" + strconv.FormatInt(*p.ID, 10) + "." + p.Ext)
		if err != nil {
			// @TODO: Log in database
			fmt.Println(err)
		}

		err = dal.DeleteMedia(prv, p)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusGone)

	}
}
