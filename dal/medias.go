package dal

import (
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
)

func FindMedia(prv *services.Provider, mediaID string) (*models.Media, error) {
	rq := `
	SELECT m.MEDIA_ID, m.TITLE, m.DESCRIPTION, m.PATH, m.VISIBILITY, su.USER_ID as "User.user_id", su.user_name as "User.user_name", '' as "User.user_email", '' as "User.user_password", true as "User.validated", '' as "User.validation_code" 
	FROM MEDIA m
	INNER JOIN SCINNA_USER su ON su.USER_ID = m.USER_ID
	WHERE m.MEDIA_ID = $1
`

	var media models.Media
	row := prv.DB.QueryRowx(rq, mediaID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.StructScan(&media)
	return &media, err
}

func MediaBelongsToToken(prv *services.Provider, pict *models.Media, token string) bool {
	row := prv.DB.QueryRow(`
		SELECT TRUE
		FROM MEDIA m
		INNER JOIN SCINNA_USER u ON u.USER_ID = m.USER_ID
		INNER JOIN LOGIN_TOKENS lt ON lt.USER_ID = u.USER_ID
		WHERE lt.LOGIN_TOKEN = ?
		  AND m.MEDIA_ID = ?
`, token, pict.MediaID)

	if row.Err() != nil {
		return false
	}

	var isOwner bool
	row.Scan(&isOwner)

	return isOwner
}

func CreatePicture(prv *services.Provider, pict *models.Media) error {
	uid, err := prv.GenerateUID()
	if err != nil {
		return err
	}

	pict.MediaID = uid
	pict.Path = pict.User.UserID + "/" + uid

	_, err = prv.DB.Exec(`
		INSERT INTO MEDIA (MEDIA_ID, USER_ID, TITLE, DESCRIPTION, PATH, VISIBILITY, MIMETYPE)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, pict.MediaID, pict.User.UserID, pict.Title, pict.Description, pict.Path, pict.Visibility, pict.Mimetype)

	return err
}

func DeleteMedia(prv *services.Provider, pict *models.Media) error {
	_, err := prv.DB.Exec("DELETE FROM MEDIA WHERE media_id = ?", pict.MediaID)
	return err
}