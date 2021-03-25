package dal

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
)

type Collections struct {
	DB *sqlx.DB
}

func (c *Collections) Create(collection *models.Collection) error {
	if collection.User == nil {
		return errors.New("the user is not set")
	}

	row := c.DB.QueryRowx(
		`INSERT INTO COLLECTIONS (TITLE, USER_ID, VISIBILITY, DEFAULT_COLLECTION) VALUES ($1, $2, $3, $4) RETURNING (CLC_ID)`,
		collection.Title,
		collection.User.UserID,
		collection.Visibility,
		collection.IsDefault,
	)

	if row.Err() != nil {
		if pqErr, ok := row.Err().(*pq.Error); ok {
			if pqErr.Code == "23505" { // Duplicate key
				return serrors.CollectionAlreadyExists
			}
		}
		return row.Err()
	}

	var collectionId string
	if err := row.Scan(&collectionId); err != nil {
		return err
	}

	collection.CollectionID = collectionId
	return nil
}

func (c *Collections) CreateDefault(user *models.User) (*models.Collection, error) {
	collection := models.Collection{
		Title:      "Default collection", // @Todo: localize this given the registration locale (?)
		IsDefault:  true,
		User:       user,
		Visibility: models.VisibilityPublic,
		Medias:     []models.Media{},
	}

	err := c.Create(&collection)

	return &collection, err
}

func (c *Collections) FetchRoot(user *models.User) (*models.Collection, error) {
	row := c.DB.QueryRowx(`SELECT CLC_ID, TITLE, VISIBILITY, DEFAULT_COLLECTION FROM COLLECTIONS WHERE USER_ID = $1 AND DEFAULT_COLLECTION = true`, user.UserID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)

	if err == sql.ErrNoRows {
		err = serrors.CollectionNotFound
	}

	return &collection, err
}

func (c *Collections) Fetch(user *models.User, name string) (*models.Collection, error) {
	row := c.DB.QueryRowx(`SELECT CLC_ID, TITLE, VISIBILITY, DEFAULT_COLLECTION FROM COLLECTIONS WHERE USER_ID = $1 AND TITLE = $2`, user.UserID, name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)

	return &collection, err
}

// clcid is there when i'll have the faith to implement nested collections
// By procrastinating on the fact of doing it maybe I can prepare it fine enough that it won't be a pain to do
// when time will come
func (c *Collections) FetchSubCollection(user *models.User, clcid string) ([]models.Collection, error) {
	rows, err := c.DB.Queryx(`SELECT CLC_ID, TITLE, VISIBILITY FROM COLLECTIONS WHERE USER_ID = $1 AND DEFAULT_COLLECTION = false`, user.UserID)
	if err != nil {
		return nil, err
	}

	var collections []models.Collection
	for rows.Next() {
		collection := models.Collection{
			IsDefault: false,
		}
		err = rows.StructScan(&collection)
		if err != nil {
			return nil, err
		}

		collections = append(collections, collection)
	}

	return collections, err
}

func (c *Collections) FetchWithMedias(dalMedias Medias, user *models.User, name string, showHidden bool) (*models.Collection, error) {
	row := c.DB.QueryRowx(`
		SELECT 
			CLC_ID,
			TITLE,
			VISIBILITY,
			DEFAULT_COLLECTION
		FROM
			COLLECTIONS
		WHERE
			USER_ID = $1
		  AND (
				(LENGTH($2) > 0 AND TITLE = $2)
			   OR
				(LENGTH($2) = 0 AND DEFAULT_COLLECTION = TRUE)
			)
`,
		user.UserID,
		name,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)
	if err != nil {
		return nil, err
	}

	if collection.IsDefault {
		cols, err := c.FetchSubCollection(user, collection.CollectionID)
		if err != nil {
			return nil, err
		}

		collection.Collections = cols
	} else {
		collection.Collections = []models.Collection{}
	}

	medias, err := dalMedias.FindFromCollection(collection.CollectionID, showHidden, false)
	if err != nil {
		return nil, err
	}

	collection.Medias = medias

	return &collection, err
}

func (c *Collections) FetchFromUsernameWithMedias(dalMedias Medias, user string, name string, showHidden bool) (*models.Collection, error) {
	row := c.DB.QueryRowx(`
		SELECT 
			CLC_ID,
			TITLE,
			VISIBILITY,
			DEFAULT_COLLECTION
		FROM
			COLLECTIONS c
			INNER JOIN SCINNA_USER su ON su.USER_ID = c.USER_ID
		WHERE
			su.USER_NAME = $1
		  AND (
				(LENGTH($2) > 0 AND TITLE = $2)
			   OR
				(LENGTH($2) = 0 AND DEFAULT_COLLECTION = TRUE)
			)`,
		user,
		name,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, serrors.CollectionNotFound
		}

		return nil, err
	}

	medias, err := dalMedias.FindFromCollection(collection.CollectionID, showHidden, false)
	if err != nil {
		return nil, err
	}

	collection.Medias = medias

	return &collection, err
}

func (c *Collections) UpdateIfOwned(user *models.User, title, newTitle string, newVisibility models.Visibility) (*models.Collection, error) {
	res, err := c.DB.Exec(`
		UPDATE
			COLLECTIONS
		SET
			TITLE = $1,
			VISIBILITY = $2
		WHERE
			TITLE = $3
   		  AND
			USER_ID = $4
		  AND
		    DEFAULT_COLLECTION = false
`, newTitle, newVisibility, title, user.UserID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // Duplicate key
				return nil, serrors.CollectionAlreadyExists
			}
		}

		return nil, err
	}

	amtRowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if amtRowsAffected == 0 {
		return nil, serrors.CollectionNotFound
	}

	return c.Fetch(user, newTitle)
}

func (c *Collections) DeleteIfOwned(user *models.User, title string) error {
	res, err := c.DB.Exec(`
		DELETE FROM collections
		WHERE 
			USER_ID = $1
		  AND
		    TITLE = $2
		  AND
		      DEFAULT_COLLECTION = false
`, user.UserID, title)

	if err != nil {
		return err
	}

	amtRowsDeleted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if amtRowsDeleted == 0 {
		return serrors.CollectionNotFound
	}

	return nil
}
