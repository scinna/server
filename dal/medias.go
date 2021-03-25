package dal

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type Medias struct {
	DB *sqlx.DB
}

func (m *Medias) Find(mediaID string) (*models.Media, error) {
	rq := `
	SELECT 
		m.MEDIA_ID,
	    m.MEDIA_TYPE,
		m.TITLE,
		m.DESCRIPTION,
		m.PATH,
		m.THUMBNAIL,
		m.VISIBILITY,
	    m.CUSTOM_DATA,
		su.USER_ID as "User.user_id",
		su.user_name as "User.user_name",
		'' as "User.user_email",
		'' as "User.user_password",
		true as "User.validated",
		'' as "User.validation_code",
		c.clc_id as "collection.clc_id",
		CASE WHEN c.DEFAULT_COLLECTION = TRUE
			 THEN ''
			 ELSE c.title
		END as "collection.title",
		c.visibility as "collection.visibility",
		c.default_collection "collection.default_collection"
	FROM 
		MEDIA m
		INNER JOIN SCINNA_USER su ON su.USER_ID = m.USER_ID
		INNER JOIN COLLECTIONS c ON c.CLC_ID = m.CLC_ID
	WHERE
		m.MEDIA_ID = $1
`

	var media models.Media
	row := m.DB.QueryRowx(rq, mediaID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.StructScan(&media)
	return &media, err
}

func (m *Medias) FindFromCollection(id string, withHidden bool, showLinks bool) ([]models.Media, error) {
	rows, err := m.DB.Queryx(`
		SELECT MEDIA_ID, MEDIA_TYPE, TITLE, DESCRIPTION, PATH, THUMBNAIL, VISIBILITY, CUSTOM_DATA
		FROM MEDIA
		WHERE 
			CLC_ID = $1
		  AND
		      (
		          $3 = true
		          OR
		          MEDIA_TYPE <> 3
		      )
		  AND
			(
				(VISIBILITY = 0)
				OR
				$2
			)
	`, id, withHidden, showLinks)

	if err != nil {
		return nil, err
	}

	var medias []models.Media
	for rows.Next() {
		media := models.Media{}
		err = rows.StructScan(&media)
		if err != nil {
			continue
		}

		medias = append(medias, media)
	}

	return medias, err
}

func (m *Medias) MediaBelongsToToken(pict *models.Media, token string) bool {
	row := m.DB.QueryRow(`
		SELECT TRUE
		FROM MEDIA m
		INNER JOIN SCINNA_USER u ON u.USER_ID = m.USER_ID
		INNER JOIN LOGIN_TOKENS lt ON lt.USER_ID = u.USER_ID
		WHERE lt.LOGIN_TOKEN = $1
		  AND m.MEDIA_ID = $2
`, token, pict.MediaID)

	if row.Err() != nil {
		return false
	}

	var isOwner bool
	row.Scan(&isOwner)

	return isOwner
}

func (m *Medias) CreatePicture(pict *models.Media, collection string) error {
	pict.Path = pict.User.UserID + "/" + pict.MediaID

	_, err := m.DB.Exec(`
		INSERT INTO MEDIA (MEDIA_ID, MEDIA_TYPE, USER_ID, TITLE, DESCRIPTION, PATH, THUMBNAIL, VISIBILITY, MIMETYPE, CLC_ID, CUSTOM_DATA)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
		        CASE WHEN LENGTH($10) > 0 THEN (SELECT CLC_ID FROM COLLECTIONS WHERE user_id = $3 AND TITLE = $10)
		        ELSE (SELECT CLC_ID FROM COLLECTIONS WHERE user_id = $3 AND DEFAULT_COLLECTION = true)
		END, '{}')`, pict.MediaID, pict.MediaType, pict.User.UserID, pict.Title, pict.Description, pict.Path, pict.Thumbnail, pict.Visibility, pict.Mimetype, collection)

	return err
}

func (m *Medias) CreateShortenUrl(shortenUrl *models.Media) error {
	customData, err := json.Marshal(shortenUrl.CustomData)
	if err != nil {
		return err
	}

	_, err = m.DB.Exec(`
		INSERT INTO MEDIA (MEDIA_ID, MEDIA_TYPE, USER_ID, CUSTOM_DATA, CLC_ID, visibility)
		VALUES ($1, $2, $3, $4, $5, 0)
`, shortenUrl.MediaID, shortenUrl.MediaType, shortenUrl.User.UserID, customData, shortenUrl.Collection.CollectionID)

	return err
}

func (m *Medias) DeleteMedia(pict *models.Media) error {
	_, err := m.DB.Exec("DELETE FROM MEDIA WHERE media_id = $1", pict.MediaID)
	return err
}
