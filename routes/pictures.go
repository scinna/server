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

// RawPictureRoute is the route that render the picture: /{picture id}
func RawPictureRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["pict"]

		if len(id) == 0 {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		p, err := dal.GetPicture(prv, id)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		if p.Visibility == 2 {
			user, err := auth.ValidateRequest(prv, w, r)
			if err != nil || p.Creator.ID == user.ID {
				serrors.ErrorPrivatePicture.Write(w)
				return
			}
		}

		pictFile, err := os.Open(prv.Config.PicturePath + "/" + strconv.FormatInt(*p.Creator.ID, 10) + "/" + strconv.FormatInt(*p.ID, 10) + "." + p.Ext)
		if err != nil {
			serrors.ErrorPictureNotFound.Write(w)
			return
		}
		defer pictFile.Close()

		w.Header().Set("Content-Type", utils.GetMimetypeForExt(p.Ext))
		io.Copy(w, pictFile)
	}
}

// PictureInfoRoute returns the informations of the picture like author, date, visibility, ...
func PictureInfoRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["URL_ID"]

		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := dal.GetPicture(prv, id)
		if serrors.WriteError(w, err) {
			return
		}

		user, err := auth.ValidateRequest(prv, w, r)
		if p.Visibility == 2 {
			if err == serrors.ErrorNoToken && serrors.WriteError(w, serrors.ErrorPrivatePicture) {
				return
			}

			if *user.ID != *p.Creator.ID {
				serrors.ErrorPrivatePicture.Write(w)
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

// UploadPictureRoute is the route that let user upload a picture
func UploadPictureRoute(prv *services.Provider) http.HandlerFunc {
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

		file, _, err := r.FormFile("picture")
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

		parentFolder := prv.Config.PicturePath + "/" + strconv.FormatInt(*user.ID, 10) + "/"

		_, err = os.Stat(parentFolder)
		if os.IsNotExist(err) {
			err = os.MkdirAll(parentFolder, os.ModePerm)
			if err != nil {
				serrors.WriteLoggableError(w, err)
				return
			}
		}

		pict := model.Picture{
			Title:       title,
			Description: desc,
			Creator:     &user,
			Visibility:  visibility,
			Ext:         utils.GetExtForMimetype(mimeType),
		}

		pict, err = dal.CreatePicture(prv, pict)
		if serrors.WriteError(w, err) {
			return
		}

		outputFile, err := os.Create(parentFolder + strconv.FormatInt(*pict.ID, 10) + "." + pict.Ext)
		if err != nil {
			serrors.WriteLoggableError(w, err)
			dal.DeletePicture(prv, pict)
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
		pict.ID = nil
		pict.Ext = ""

		json, err := json.Marshal(pict)
		if err != nil {
			// The picture is uploaded but something went wrong while encoding the response
			w.WriteHeader(http.StatusAccepted)
			return
		}

		w.Write(json)
	}
}

// DeletePictureRoute is the route that let the user delete one OR MULTIPLE of his picture
func DeletePictureRoute(prv *services.Provider) http.HandlerFunc {
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

		p, err := dal.GetPicture(prv, id)
		if serrors.WriteError(w, err) {
			return
		}

		if *p.Creator.ID != *user.ID && user.Role != model.UserRoleAdmin {
			serrors.ErrorWrongOwner.Write(w)
			return
		}

		err = os.Remove(prv.Config.PicturePath + "/" + strconv.FormatInt(*p.Creator.ID, 10) + "/" + strconv.FormatInt(*p.ID, 10) + "." + p.Ext)
		if err != nil {
			// @TODO: Log in database
			fmt.Println(err)
		}

		err = dal.DeletePicture(prv, p)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusGone)

	}
}
