package dal

import (
	"database/sql"

	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

// GetMedia retreive a media from the database and its user given an URL ID
func GetMedia(p *services.Provider, urlID string) (model.Media, error) {
	rq := ` SELECT p.ID, p.CREATED_AT, p.TITLE, p.URL_ID, p.DESCRIPT, p.VISIBILITY, p.EXT, p.THUMBNAIL,
				   au.ID AS "creator.id", au.CREATED_AT AS "creator.created_at", au.ROLE as "creator.role", au.EMAIL as "creator.email", au.USERNAME AS "creator.username"
			FROM MEDIAS p
			INNER JOIN APPUSER au ON au.ID = p.CREATOR
			WHERE p.URL_ID = $1`

	var pict model.Media
	err := p.Db.QueryRowx(rq, urlID).StructScan(&pict)
	if err == sql.ErrNoRows {
		return pict, serrors.ErrorMediaNotFound
	}

	return pict, err
}

// GetMediasFromUser returns all the medias from a user given its username and whether it report only public medias
func GetMediasFromUser(p *services.Provider, userID *int64, visibility bool) ([]model.Media, error) {
	rq := ` SELECT ID, CREATED_AT, TITLE, URL_ID, DESCRIPT, VISIBILITY, EXT, THUMBNAIL
			FROM MEDIAS
			WHERE CREATOR = $1`

	if visibility {
		rq += " AND VISIBILITY = 0"
	}

	// This assignement is needed so that the JSON Marshal does not return nil instead of empty array when the user has no medias
	var medias []model.Media = []model.Media{}
	rows, err := p.Db.Queryx(rq, &userID)
	if err != nil {
		// Should never happen
		return []model.Media{}, err
	}

	for rows.Next() {
		var p model.Media
		rows.StructScan(&p)

		medias = append(medias, p)
	}

	return medias, nil
}

// CreateMedia inserts the medias in the database, and returs the URL ID generated
func CreateMedia(prv *services.Provider, media model.Media) (model.Media, error) {
	rq := `INSERT INTO MEDIAS(TITLE, URL_ID, DESCRIPT, VISIBILITY, CREATOR, EXT, THUMBNAIL) 
		   VALUES ($1, $2, $3, $4, $5, $6, $7)
		   RETURNING ID`

	URLID, err := prv.GenerateUID()
	if err != nil {
		return media, err
	}
	media.URLID = URLID

	var lastInsertedID int64 = 0
	err = prv.Db.QueryRow(rq, media.Title, media.URLID, media.Description, media.Visibility, media.Creator.ID, media.Ext, media.Thumbnail).Scan(&lastInsertedID)
	media.ID = &lastInsertedID

	return media, err
}

// DeleteMedia removes a media from the database
func DeleteMedia(prv *services.Provider, media model.Media) error {
	rq := ` DELETE
			FROM MEDIAS
			WHERE ID = $1`
	result, err := prv.Db.Exec(rq, media.ID)

	if err != nil {
		count, err := result.RowsAffected()
		if err != nil && count == 0 {
			return serrors.ErrorMediaNotFound
		}
	}

	return err
}
