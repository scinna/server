package dal

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type Collections struct {
	DB *sqlx.DB
}

func (c *Collections) Create(collection *models.Collection) error {
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

func (c *Collections) CreateDefault(user *models.User) (*models.Collection, error) {
	collection := models.Collection{
		Title:      "Default collection", // @Todo: localize this given the registration locale (?)
		IsDefault:  true,
		User:       user,
		Visibility: 0,
		Medias:     []models.Media{},
	}

	err := c.Create(&collection)

	return &collection, err
}

func (c *Collections) FetchRoot(user *models.User) (*models.Collection, error) {
	row := c.DB.QueryRowx(`SELECT TITLE, USER_ID, VISIBILITY, DEFAULT_COLLECTION FROM COLLECTIONS WHERE USER_ID = $1 AND DEFAULT_COLLECTION = true`, user.UserID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)

	return &collection, err
}

func (c *Collections) Fetch(user *models.User, name string) (*models.Collection, error) {
	row := c.DB.QueryRowx(`SELECT TITLE, USER_ID, VISIBILITY, DEFAULT_COLLECTION FROM COLLECTIONS WHERE USER_ID = $1 AND TITLE = $2`, user.UserID, name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var collection models.Collection
	err := row.StructScan(&collection)

	return &collection, err
}

func (c *Collections) Delete(collection *models.Collection) error {

	return nil
}