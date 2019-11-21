package routes

import (
	"net/http"
	"encoding/json"
	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/dal"
)

func UserPicturesRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		picts, err := dal.GetPicturesFromUser(prv, "admin", true)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := json.Marshal(picts)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func MyPicturesRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		picts, err := dal.GetPicturesFromUser(prv, "admin", false)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := json.Marshal(picts)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

	}
}

func MyInfosRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// @TODO: JWT / Get user
		user, err := dal.GetUser(prv, "admin")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json, err := json.Marshal(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

	}
}

func UpdateMyInfosRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("UpdateMyInfosRoute - To be implemented"))
	}
}