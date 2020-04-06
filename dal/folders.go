package dal

import (
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

// GetFolderByPath returns the the folder
func GetFolderByPath(prv *services.Provider, user *model.AppUser, folderPath string) (*model.Folder, error) {
	rq := `
		SELECT f.id, f.created_at, f.folder_name, f.parent_path,
		au.ID AS "creator.id", au.CREATED_AT AS "creator.created_at", au.ROLE as "creator.role", au.EMAIL as "creator.email", au.USERNAME AS "creator.username",
		p.ID as "parent_folder"
		FROM FOLDERS f
			INNER JOIN APPUSER au ON au.ID = f.CREATOR
			LEFT JOIN FOLDERS p ON p.ID = f.PARENT_FOLDER
			WHERE ((CASE WHEN f.parent_folder IS NOT NULL THEN f.parent_path || '/' ELSE '' END) || f.folder_name) = $2
			  AND au.id = $1
	`

	var folder model.Folder
	err := prv.Db.QueryRowx(rq, user.ID, utils.StripHeadSlash(folderPath)).StructScan(&folder)
	if err != nil {
		fmt.Println(err)
		return &model.Folder{}, serrors.ErrorBadInviteCode
	}

	return &folder, nil
}

// GetSubfolders returns a list of folders contained by the given one
func GetSubfolders(prv *services.Provider, user *model.AppUser, folderPath string) ([]model.Folder, error) {
	folders := make([]model.Folder, 0)
	rq := `SELECT ID, FOLDER_NAME FROM FOLDERS WHERE PARENT_PATH = $1 ORDER BY FOLDER_NAME`

	res, err := prv.Db.Queryx(rq, folderPath)
	if err != nil {
		return []model.Folder{}, err
	}

	for res.Next() {
		currFolder := model.Folder{}
		res.StructScan(&currFolder)
		folders = append(folders, currFolder)
	}

	return folders, nil
}

// GetMediasInFolder gives a list of medias in a folder
func GetMediasInFolder(prv *services.Provider, user *model.AppUser, folderPath string) ([]model.Media, error) {
	medias := make([]model.Media, 0)
	rq := `
		SELECT ID, TITLE, THUMBNAIL
		FROM MEDIAS
		WHERE (
			PARENT = NULL
			AND ($1 IS NULL OR LENGTH($1) = 0)
		) 
		OR (
			PARENT = (
				SELECT ID
				FROM FOLDER
				WHERE (PARENT_PATH  || '/' || FOLDER_NAME) = $1 ORDER BY FOLDER_NAME
			)
		)`

	res, err := prv.Db.Queryx(rq, folderPath)
	if err != nil {
		return []model.Media{}, err
	}

	for res.Next() {
		currMedia := model.Media{}
		res.StructScan(&currMedia)
		medias = append(medias, currMedia)
	}

	return medias, nil
}

// CreateFolder creates a folder
func CreateFolder(prv *services.Provider, user *model.AppUser, folderPath string) (*model.Folder, error) {
	folderPath = utils.StripHeadSlash(folderPath)
	path := utils.SplitPath(folderPath)

	var parentPath string
	var parent *model.Folder
	if len(path) > 1 {
		parentPath = folderPath[0:strings.LastIndex(folderPath, "/")]

		var err error
		parent, err = GetFolderByPath(prv, user, parentPath)
		if err != nil {
			return &model.Folder{}, err
		}
	}

	f := model.Folder{
		Creator:    *user,
		FolderName: path[len(path)-1],
		CreatedAt:  time.Now(),
	}

	// If there is a parent, we set it accordingly
	var parentID *int64 = nil
	if parent != nil {
		f.ParentPath = parentPath
		parentID = &parent.ID
	}

	// We relies on postgres' unique constraint on PARENT_FOLDER+FOLDERNAME to prevent creating folder with the same name
	// To keep in mind when I'll port it to GORM
	rq := ` 
		INSERT INTO FOLDERS (CREATOR, FOLDER_NAME, PARENT_FOLDER, PARENT_PATH)
		VALUES ($1, $2, $3, $4)	
		RETURNING ID`

	// We populate the folder that is returned with the ID
	var id int64
	rows, err := prv.Db.Query(rq, f.Creator.ID, f.FolderName, parentID, f.ParentPath)
	if err != nil {
		errPost, ok := err.(*pq.Error)
		if ok && errPost.Code.Name() == serrors.PostgresError["AlreadyExisting"] {
			return GetFolderByPath(prv, user, folderPath)
		}

		return &model.Folder{}, err
	}

	for rows.Next() {
		rows.Scan(&id)
	}

	f.ID = id

	return &f, nil

}
