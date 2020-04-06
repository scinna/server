package routes

import (
	"encoding/json"
	"net/http"

	"github.com/scinna/server/auth"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

type content struct {
	IsFolder bool   `json:"IsFolder"`
	Name     string `json:"Name"`
	ID       *int64 `json:"ID"`
	Icon     string `json:"Icon"` // Only for media
}

type getContentResponse struct {
	Current *model.Folder `json:"CurrentFolder"`
	Content []content     `json:"Content"`
}

// GetFolderContentRoute lists all folders and files in a folder
func GetFolderContentRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		urlPath := utils.StripHeadSlash(r.URL.Path)

		f, err := dal.GetFolderByPath(prv, &user, urlPath)
		if serrors.WriteError(w, err) {
			return
		}

		contentArr := make([]content, 0)

		folders, err := dal.GetSubfolders(prv, &user, urlPath)
		for i := range folders {
			contentArr = append(contentArr, content{
				IsFolder: true,
				Name:     folders[i].FolderName,
				ID:       &folders[i].ID,
			})
		}

		medias, err := dal.GetMediasInFolder(prv, &user, urlPath)
		for i := range medias {
			contentArr = append(contentArr, content{
				IsFolder: false,
				Name:     medias[i].Title,
				ID:       medias[i].ID,
				Icon:     medias[i].Thumbnail,
			})
		}

		cr := getContentResponse{
			Current: f,
			Content: contentArr,
		}

		resp, err := json.Marshal(cr)
		if serrors.WriteError(w, err) {
			return
		}

		/**
			@TODO before opening it to see other users' folders:
				Clean up the user (Remove the email, createdat, role)
		**/

		w.Write(resp)
	}
}

// CreateFolderRoute lets a user create a folder
func CreateFolderRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		// @TODO: Not implemented yet
		f, err := dal.CreateFolder(prv, &user, r.URL.Path)
		if serrors.WriteError(w, err) {
			return
		}

		resp, err := json.Marshal(f)
		if serrors.WriteError(w, err) {
			return
		}

		w.Write(resp)
	}
}

// RenameFolderRoute lets a user rename a folder
func RenameFolderRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Not implemented yet
	}
}

// DeleteFolderRoute lets a user delete a folder and its content
func DeleteFolderRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Not implemented yet
	}
}

// MoveFolderRoute lets a user move a folder to another one
func MoveFolderRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Not implemented yet
	}
}
