package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// UserPicturesRoute is the route that list all the given user's picture, and their infos
func UserPicturesRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		if len(username) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		picts, err := dal.GetPicturesFromUser(prv, username, true)

		if err != nil {
			serrors.WriteError(w, err)
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

// MyPicturesRoute is pretty much the same as UserPicturesRoute but for the current user - I'll may be deprecating this route in order to have only one that CAN use the auth token
func MyPicturesRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		picts, err := dal.GetPicturesFromUser(prv, user.Username, false)

		if serrors.WriteLoggableError(w, err) {
			return
		}

		json, err := json.Marshal(picts)

		if serrors.WriteLoggableError(w, err) {
			return
		}

		w.Write(json)

	}
}

// MyInfosRoute returns the user's infos (Username, Mail, Qty of public pictures, ...)
func MyInfosRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
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
			serrors.ErrorBadRequest.Write(w)
			return
		}

		var rc updateInfoRequest

		err = json.Unmarshal(body, &rc)

		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		u, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		lenMail, lenPass := len(rc.Email), len(rc.Password)

		if lenMail == 0 && lenPass == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !prv.Mail.IsEmail.MatchString(rc.Email) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rq := `UPDATE APPUSER
			   SET EMAIL = ($1::VARCHAR),
			       PASSWORD = ($2::VARCHAR)
			   WHERE ID = $3::INTEGER
		`

		if lenMail > 0 {
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
