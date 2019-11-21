package routes

import (
	"fmt"
	"io/ioutil"
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

type UpdateInfoRequest struct {
	Username string /** Temporary, won't be needed as soon as JWT is implemented, maybe repurposed to change username later **/
	Email string
	Password string
}

func UpdateMyInfosRoute (prv *services.Provider) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var rc UpdateInfoRequest 

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

		if u.Email != rc.Email && lenMail > 0 {
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

		if result.RowsAffected() == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}
}