package dal

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type Collections struct {
	DB *sqlx.DB
}

func (c *Collections) InsertCollection(collection *models.Collection) error {
	if collection.User == nil {
		return errors.New("the user is not set")
	}

	row := c.DB.QueryRow(
		`INSERT INTO COLLECTIONS (TITLE, USER_ID, VISIBILITY, DEFAULT_COLLECTION) VALUES ($1, $2, $3, $4) RETURNING (CLC_ID)`,
		collection.Title,
		collection.User.UserID,
		collection.Visibility,
		collection.IsDefault,
	)

	if row.Err() != nil {
		return row.Err()
	}

	var collectionId string
	if err := row.Scan(&collection); err != nil {
		return err
	}

	collection.CollectionID = collectionId
	return nil
}

func (c *Collections) CreateDefaultCollection(user *models.User) (*models.Collection, error) {
	collection := models.Collection{
		Title:      "Default collection", // @Todo: localize this given the registration locale (?)
		IsDefault:  true,
		User:       user,
		Visibility: 0,
		Medias:     []models.Media{},
	}

	err := c.InsertCollection(&collection)

	return &collection, err
}
