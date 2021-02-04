package fixtures

import (
	"fmt"
	"github.com/scinna/server/log"
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"syreclabs.com/go/faker"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func InsertFakeData(prv *services.Provider, user *models.User, defaultCollection *models.Collection) {
	log.InfoAlwaysShown("\t- Inserting default pictures in collection")
	for i := 0; i < r.Intn(10); i++ {
		err := generatePicture(prv, user, defaultCollection)
		if err != nil {
			fmt.Printf("Error generating picture %v (Default collection): %v\n", i, err)
		}
	}

	for i := 0; i < r.Intn(3); i++ {
		col, err := generateCollection(prv, user)
		if err != nil {
			fmt.Printf("Failed to generate collection %v: %v\n", i, err)
			continue
		}

		for k := 0; k < r.Intn(10); k++ {
			err := generatePicture(prv, user, col)
			if err != nil {
				fmt.Printf("Error generating picture %v (collection #%v): %v\n", i, k, err)
			}
		}
	}

}

func generateCollection(prv *services.Provider, user *models.User) (*models.Collection, error) {
	visib := r.Intn(2)

	col := models.Collection{
		Title:       fmt.Sprintf("[%v] %v]", visib, faker.Lorem().Sentence(2)),
		User:         user,
		Visibility:   visib,
		IsDefault:    false,
		Medias:       nil,
	}

	err := prv.Dal.Collections.Create(&col)

	return &col, err
}


func generatePicture(prv *services.Provider, user *models.User, collection *models.Collection) error {
	parentFolder := prv.Config.MediaPath + "/" + user.UserID + "/"
	os.MkdirAll(parentFolder, os.ModePerm)

	uid, _ := prv.GenerateUID()

	visib := r.Intn(2)

	pict := models.Media{
		MediaID:     uid,
		Title:       fmt.Sprintf("[%v] %v]", visib, faker.Lorem().Sentence(3)),
		Description: faker.Lorem().Sentence(15),
		Visibility:  visib,
		User:        user,
		Mimetype:    "image/jpeg",
	}

	err := prv.Dal.Medias.CreatePicture(&pict, collection.Title)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(parentFolder + pict.MediaID)
	if err != nil {
		prv.Dal.Medias.DeleteMedia(&pict)
		return err
	}
	defer outputFile.Close()

	file, err := findRandomPicture()
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, file)
	return err
}

func findRandomPicture() (*os.File, error) {
	var files = []string{}

	filepath.Walk("assets/fake_data", func(path string, info os.FileInfo, err error) error {
		if path != "assets/fake_data" && path != "assets/fake_data/FakeDataSource.txt" {
			files = append(files, path)
		}
		return nil
	})

	rndFile := r.Intn(len(files))
	return os.Open(files[rndFile])
}