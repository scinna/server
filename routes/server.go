package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
	"net/http"
)

func Server(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/infos", configRoute(prv))

	r.HandleFunc("/logo", logoWide(prv))
	r.HandleFunc("/logo/wide", logoWide(prv))
	r.HandleFunc("/logo/small", logoSmall(prv))

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.Json)
	admin.Use(middlewares.LoggedInMiddleware(prv))

	admin.HandleFunc("/invite", listInviteCode(prv)).Methods(http.MethodGet)
	admin.HandleFunc("/invite", generateInviteCode(prv)).Methods(http.MethodPost)
	admin.HandleFunc("/invite/{code}", deleteInviteCode(prv)).Methods(http.MethodDelete)

	admin.HandleFunc("/users", listUsers(prv)).Methods(http.MethodGet)
}


func newServerConfig(prv *services.Provider, isAdmin bool) dto.ServerConfig {
	cfg := dto.ServerConfig{
		RegistrationAllowed: prv.Config.Registration.Allowed,
		Validation:          prv.Config.Registration.Validation,
		WebURL:              prv.Config.WebURL,
		CustomBranding:      prv.Config.CustomBranding,
	}

	if isAdmin {
		cfg.ScinnaVersion = fmt.Sprintf("%v.%v", utils.SCINNA_VERSION, utils.SCINNA_PATCH)
	}

	return cfg
}

func configRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, _ := middlewares.GetUserFromRequest(prv, r)
		isAdmin := false
		if user != nil {
			isAdmin = user.IsAdmin
		}

		w.Header().Set("Content-Type", "application/json")

		bytes, err := json.Marshal(newServerConfig(prv, isAdmin))
		if serrors.WriteError(w, r, err) {
			return
		}

		_, _ = w.Write(bytes)
	}
}

func logoWide(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(prv.Config.CustomLogoWide) > 0 {
			http.ServeFile(w, r, prv.Config.CustomLogoWide)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(prv.EmbeddedAssets.LogoWide)
	}
}

func logoSmall(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(prv.Config.CustomLogoSmall) > 0 {
			http.ServeFile(w, r, prv.Config.CustomLogoSmall)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(prv.EmbeddedAssets.LogoSmall)
	}
}

func listInviteCode(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		inviteCodes, err := prv.Dal.Server.ListInviteCode()
		if serrors.WriteError(w, r, err){
			return
		}

		bytes, _ := json.Marshal(inviteCodes)
		w.Write(bytes)
	}
}

func generateInviteCode(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		/** @TODO: Replace with ROLE_GENERATE_INVITE, by default in the ROLE_ADMIN group **/
		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		inviteCode, err := prv.GenerateUID()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = prv.Dal.Server.GenerateInviteCode(user, inviteCode)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, _ := json.Marshal(struct {
			Code string
		} {
			Code: inviteCode,
		})

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(bytes)
	}
}

func deleteInviteCode(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		/** @TODO: Replace with ROLE_GENERATE_INVITE, by default in the ROLE_ADMIN group **/
		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		code := mux.Vars(r)["code"]

		err := prv.Dal.Server.Delete(code)
		if serrors.WriteError(w, r, err) {
			return
		}

		w.WriteHeader(http.StatusGone)
	}
}

func listUsers(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		/** @TODO: Replace with ROLE_USER_MANAGER, by default in the ROLE_ADMIN group **/
		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		users, err := prv.Dal.Admin.ListUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, _ := json.Marshal(users)
		w.Write(bytes)
	}
}