package routes

import (
	"net/http"
	"fmt"
	"os"
	"io"

	"github.com/oxodao/scinna/services"
)

func RawPictureRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		img, err := os.Open(prv.PicturePath + "/photo.jpg")
		if err != nil {
			fmt.Println(err)
		}
		defer img.Close()

		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, img)
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