package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Browser(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.Json)

	// Implemented as such because in the future we'll have nested collections
	r.PathPrefix("/{user}").Handler(http.StripPrefix("/api/browse/", list(prv))).Methods(http.MethodGet)
	r.PathPrefix("/{user}").Handler(http.StripPrefix("/api/browse/", middlewares.LoggedInMiddleware(prv)(create(prv)))).Methods(http.MethodPost)
	r.PathPrefix("/{user}").Handler(http.StripPrefix("/api/browse/", middlewares.LoggedInMiddleware(prv)(edit(prv)))).Methods(http.MethodPut)
	r.PathPrefix("/{user}").Handler(http.StripPrefix("/api/browse/", middlewares.LoggedInMiddleware(prv)(delete(prv)))).Methods(http.MethodDelete)
}

// stripPrefix removes the username from a URI requested in order to have only the "path" for the given collection
func stripPrefix(uri, username string) string {
	uri = uri[len(username):]

	if len(uri) > 0 && uri[0] == '/' {
		uri = uri[1:]
	}

	if len(uri) > 0 && uri[len(uri)-1:] == "/" {
		uri = uri[:len(uri)-1]
	}

	return uri
}

/** @TODO: Idea for nested => add a parent key to collection & primary key it along with the name & the username) **/

func list(prv *services.Provider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		path, err := url.QueryUnescape(r.URL.RequestURI())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		path = stripPrefix(path, username)

		token, err := middlewares.GetTokenFromRequest(r)
		if err != nil && err != serrors.NoToken {
			serrors.WriteError(w, err)
			return
		}

		var user *models.User
		if err == nil {
			user, err = prv.Dal.User.FetchUserFromToken(token)
			if serrors.WriteError(w, err) {
				return
			}
		}

		var collection *models.Collection

		// If the requested user is the logged in user, we can fetch more infos and private pictures
		if user != nil && user.Name == username {
			collection, err = prv.Dal.Collections.FetchWithMedias(prv.Dal.Medias, user, path, true)
		} else {
			collection, err = prv.Dal.Collections.FetchFromUsernameWithMedias(prv.Dal.Medias, username, path, false)
		}

		if serrors.WriteError(w, err) {
			return
		}

		collectionJSON, _ := json.Marshal(collection)
		_, _ = w.Write(collectionJSON)
	})
}

func create(prv *services.Provider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username := mux.Vars(r)["user"]
		if user.Name != username {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		uriParsed, err := url.QueryUnescape(r.URL.RequestURI())
		uriParsed = stripPrefix(uriParsed, username)

		body, err := ioutil.ReadAll(r.Body)
		if serrors.WriteError(w, err) {
			return
		}

		var newCollectionRequest struct {
			Visibility models.Visibility
		}

		err = json.Unmarshal(body, &newCollectionRequest)
		if err != nil || !newCollectionRequest.Visibility.IsValid() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		collection := models.Collection{
			Title:      uriParsed,
			User:       user,
			Visibility: newCollectionRequest.Visibility,
			IsDefault:  false,
		}

		err = prv.Dal.Collections.Create(&collection)
		if serrors.WriteError(w, err) {
			return
		}

		collection.User = nil

		collectionJSON, _ := json.Marshal(collection)
		_, _ = w.Write(collectionJSON)
	})
}

func edit(prv *services.Provider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username := mux.Vars(r)["user"]
		if user.Name != username {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		uriParsed, err := url.QueryUnescape(r.URL.RequestURI())
		uriParsed = stripPrefix(uriParsed, username)

		if len(uriParsed) == 0 {
			serrors.CollectionNotFound.Write(w)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if serrors.WriteError(w, err) {
			return
		}

		var query struct {
			Title      string
			Visibility models.Visibility
		}

		err = json.Unmarshal(body, &query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(query.Title) == 0 || !query.Visibility.IsValid() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		col, err := prv.Dal.Collections.UpdateIfOwned(user, uriParsed, query.Title, query.Visibility)
		if serrors.WriteError(w, err) {
			return
		}

		collectionJson, _ := json.Marshal(&col)
		_, _ = w.Write(collectionJson)
	})
}

func delete(prv *services.Provider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		username := mux.Vars(r)["user"]
		if user.Name != username {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		path, err := url.QueryUnescape(r.URL.RequestURI())
		path = stripPrefix(path, username)

		if len(path) == 0 {
			serrors.CollectionNotFound.Write(w)
			return
		}

		err = prv.Dal.Collections.DeleteIfOwned(user, path)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusGone)
	})
}
