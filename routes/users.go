package routes

import (
	"encoding/json"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
	"io/ioutil"
	"net/http"
)

// UserPicturesRoute is the route that list all the given user's picture, and their infos
func UserPicturesRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

// MyPicturesRoute is pretty much the same as UserPicturesRoute but for the current user
func MyPicturesRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		w.Write(json)

	}
}

// MyInfosRoute returns the user's infos (Username, Mail, Qty of public pictures, ...)
func MyInfosRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: JWT / Get user
		user, err := auth.ValidateRequest(prv, w, r)
		if err != nil {
			if err == serrors.ErrorNoToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if err == serrors.ErrorBadToken {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(json)

	}
}

type updateInfoRequest struct {
	Username string /** Temporary, won't be needed as soon as JWT is implemented, maybe repurposed to change username later **/
	Email    string
	Password string
}

// UpdateMyInfosRoute lets the user change it's infos
func UpdateMyInfosRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var rc updateInfoRequest

		err = json.Unmarshal(body, &rc)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := dal.GetUser(prv, rc.Username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lenMail, lenPass := len(rc.Email), len(rc.Password)

		if lenMail == 0 && lenPass == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rq := `UPDATE APPUSER
			   SET EMAIL = ($1::VARCHAR),
			       PASSWORD = ($2::VARCHAR)
			   WHERE ID = $3::INTEGER
		`

		if lenMail > 0 {
			/** @TODO Regex the hell out of it **/
			u.Email = rc.Email
		}

		if lenPass > 0 {
			u.Password, err = prv.HashPassword(rc.Password)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		result, err := prv.Db.Exec(rq, u.Email, u.Password, u.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ra, err := result.RowsAffected()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if ra == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}
}
