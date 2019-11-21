package routes

import (
	"net/http"
	"fmt"
	"os"
	"io"

	"github.com/gorilla/mux"

	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/dal"
)

func RawPictureRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["pict"]

		p, err := dal.GetPicture(prv, id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if p.Visibility == 2 {
			/** @TODO Should verify the JWT, if it's given and the picture is private, it should be displayed nonetheless **/
			/** Usage: In the app or clients, they should be able to retreive the picture **/
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		pictFile, err := os.Open(prv.PicturePath + "/" + id + ".png")
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

func PictureInfoRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PictureInfoRoute - To be implemented"))
	}
}

func UploadPictureRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UploadPictureRoute - To be implemented"))
	}
}

func DeletePictureRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeletePictureRoute - To be implemented"))
	}
}