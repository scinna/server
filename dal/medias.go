package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type Medias struct {
	DB *sqlx.DB
}

func (m *Medias) Find(mediaID string) (*models.Media, error) {
	rq := `
	SELECT m.MEDIA_ID, m.TITLE, m.DESCRIPTION, m.PATH, m.VISIBILITY, su.USER_ID as "User.user_id", su.user_name as "User.user_name", '' as "User.user_email", '' as "User.user_password", true as "User.validated", '' as "User.validation_code" 
	FROM MEDIA m
	INNER JOIN SCINNA_USER su ON su.USER_ID = m.USER_ID
	WHERE m.MEDIA_ID = $1
`

	var media models.Media
	row := m.DB.QueryRowx(rq, mediaID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.StructScan(&media)
	return &media, err
}

func (m *Medias) FindFromCollection(id string, withHidden bool) ([]models.Media, error) {
	rows, err := m.DB.Queryx(`
		SELECT MEDIA_ID, TITLE, DESCRIPTION, PATH, VISIBILITY
		FROM MEDIA
		WHERE 
			CLC_ID = $1
		  AND
			(
				(VISIBILITY = 0)
				OR
				$2
			)
	`, id, withHidden)

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
		INSERT INTO MEDIA (MEDIA_ID, USER_ID, TITLE, DESCRIPTION, PATH, VISIBILITY, MIMETYPE, CLC_ID)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 
		        CASE WHEN LENGTH($8) > 0 THEN (SELECT CLC_ID FROM COLLECTIONS WHERE user_id = $2 AND TITLE = $8)
		        ELSE (SELECT CLC_ID FROM COLLECTIONS WHERE user_id = $2 AND DEFAULT_COLLECTION = true)
		END)`, pict.MediaID, pict.User.UserID, pict.Title, pict.Description, pict.Path, pict.Visibility, pict.Mimetype, collection)

	return err
}

func (m *Medias) DeleteMedia(pict *models.Media) error {
	_, err := m.DB.Exec("DELETE FROM MEDIA WHERE media_id = $1", pict.MediaID)
	return err
}
