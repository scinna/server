package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

type userInfoResponse struct {
	model.AppUser
	Pictures []model.Picture
}

// UserInfoRoute is the route that gives the user infos & his pictures
func UserInfoRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		if len(username) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var usr model.AppUser
		var err error
		if username == "me" {
			usr, err = auth.ValidateRequest(prv, w, r)
		} else if username != "me" {
			usr, err = dal.GetUser(prv, username)
		}

		if serrors.WriteError(w, err) {
			return
		}

		uir := userInfoResponse{}
		uir.CreatedAt = usr.CreatedAt
		uir.Username = usr.Username
		uir.Role = usr.Role

		if username == "me" {
			uir.Email = usr.Email
		}

		picts, err := dal.GetPicturesFromUser(prv, usr.ID, username != "me")
		if err != nil {
			serrors.WriteError(w, err)
			return
		}
		uir.Pictures = picts

		json, err := json.Marshal(uir)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

type updateInfoRequest struct {
	Email    string
	Password string
}

// UpdateMyInfosRoute lets the user change it's infos
func UpdateMyInfosRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

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
			if _, ok := err.(*pq.Error); ok {
				serrors.ErrorRegExistingMail.Write(w)
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ra, err := result.RowsAffected()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if ra == 0 {
			serrors.WriteError(w, serrors.ErrorUserNotFound)
			return
		}

	}
}
