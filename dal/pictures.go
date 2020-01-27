package dal

import (
	"database/sql"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

// GetPicture retreive a picture from the database and its user given an URL ID
func GetPicture(p *services.Provider, urlID string) (model.Picture, error) {
	rq := ` SELECT p.ID, p.CREATED_AT, p.TITLE, p.URL_ID, p.DESCRIPT, p.VISIBILITY, p.EXT,
				   au.ID AS "creator.id", au.CREATED_AT AS "creator.created_at", au.ROLE as "creator.role", au.EMAIL as "creator.email", au.USERNAME AS "creator.username"
			FROM PICTURES p
			INNER JOIN APPUSER au ON au.ID = p.CREATOR
			WHERE p.URL_ID = $1`

	var pict model.Picture
	err := p.Db.QueryRowx(rq, urlID).StructScan(&pict)
	if err == sql.ErrNoRows {
		return pict, serrors.ErrorPictureNotFound
	}

	return pict, err
}

// GetPicturesFromUser returns all the pictures from a user given its username and whether it report only public pictures
func GetPicturesFromUser(p *services.Provider, userID *int64, visibility bool) ([]model.Picture, error) {
	rq := ` SELECT ID, CREATED_AT, TITLE, URL_ID, DESCRIPT, VISIBILITY, EXT
			FROM PICTURES
			WHERE CREATOR = $1`

	if visibility {
		rq += " AND VISIBILITY = 0"
	}

	// This assignement is needed so that the JSON Marshal does not return nil instead of empty array when the user has no pictures
	var pictures []model.Picture = []model.Picture{}
	rows, err := p.Db.Queryx(rq, &userID)
	if err != nil {
		// Should never happen
		return []model.Picture{}, err
	}

	for rows.Next() {
		var p model.Picture
		rows.StructScan(&p)

		pictures = append(pictures, p)
	}

	return pictures, nil
}

// CreatePicture inserts the picture in the database, and returs the URL ID generated
func CreatePicture(prv *services.Provider, pict model.Picture) (model.Picture, error) {
	rq := `INSERT INTO PICTURES(TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR, EXT) 
		   VALUES ($1, $2, $3, $4, $5, $6)
		   RETURNING ID`

	URLID, err := prv.GenerateUID()
	if err != nil {
		return pict, err
	}
	pict.URLID = URLID

	var lastInsertedID int64 = 0
	err = prv.Db.QueryRow(rq, pict.Title, pict.URLID, pict.Description, pict.Visibility, pict.Creator.ID, pict.Ext).Scan(&lastInsertedID)
	pict.ID = &lastInsertedID

	return pict, err
}

// DeletePicture removes a picture from the database
func DeletePicture(prv *services.Provider, pict model.Picture) error {
	rq := ` DELETE
			FROM PICTURES
			WHERE ID = $1`
	result, err := prv.Db.Exec(rq, pict.ID)

	if err != nil {
		count, err := result.RowsAffected()
		if err != nil && count == 0 {
			return serrors.ErrorPictureNotFound
		}
	}

	return err
}
