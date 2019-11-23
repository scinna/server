package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/services"
)

// RawPictureRoute is the route that render the picture: /{picture id}
func RawPictureRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["pict"]

		if len(id) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := dal.GetPicture(prv, id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if p.Visibility == 2 {
			/** @TODO Should verify the JWT, if it's given and the picture is private, it should be displayed nonetheless **/
			/** Usage: In the app or clients, they should be able to retreive the picture **/
			/** Need to think of the correct way to do it since it will prevent it from being directly put in an img tag, this
			will require a lot more work on the client to display.
			Yet we can't put it in the GET request since the user could send it to someone else **/
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// @TODO Replace with UUID
		pictFile, err := os.Open(prv.PicturePath + "/" + strconv.FormatInt(*p.Creator.ID, 10) + "/" + id + ".png")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			pictFile, err = os.Open("not_found.png")
			if err != nil {
				return
			}
		}
		defer pictFile.Close()

		w.Header().Set("Content-Type", "image/png")
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
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		user, err := auth.ValidateRequest(prv, w, r)
		if p.Visibility == 2 {
			if auth.RespondError(w, err) {
				return
			}

			if user.ID != p.Creator.ID {
				w.WriteHeader(http.StatusForbidden)
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
		w.Write([]byte("UploadPictureRoute - To be implemented"))
	}
}

// DeletePictureRoute is the route that let the user delete one OR MULTIPLE of his picture
func DeletePictureRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeletePictureRoute - To be implemented"))
	}
}
