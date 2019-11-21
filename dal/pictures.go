package dal

import (
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/services"
)

func GetPicture(p *services.Provider, urlId string) (model.Picture, error) {
	rq := ` SELECT p.ID, p.CREATED_AT, p.TITLE, p.URL_ID, p.DESCRIPT, p.VISIBILITY,
				   au.ID AS "creator.id", au.CREATED_AT AS "creator.created_at", au.EMAIL as "creator.email", au.USERNAME AS "creator.username"
			FROM PICTURES p
			INNER JOIN APPUSER au ON au.ID = p.CREATOR
			WHERE p.URL_ID = $1`

	var pict model.Picture
	err := p.Db.QueryRowx(rq, urlId).StructScan(&pict)
	return pict, err
}

func GetPicturesFromUser(p *services.Provider, user string, visibility bool) ([]model.Picture, error) {
	u, err := GetUser(p, user)
	if err != nil {
		return []model.Picture {}, err
	}

	rq := ` SELECT ID, CREATED_AT, TITLE, URL_ID, DESCRIPT, VISIBILITY
			FROM PICTURES
			WHERE CREATOR = $1`

	if visibility {
		rq += " AND VISIBILITY = 0"
	}

	var pictures []model.Picture
	rows, err := p.Db.Queryx(rq, u.ID)
	if err != nil {
		return []model.Picture {}, err
	}

	for rows.Next() {
		var p model.Picture
		rows.StructScan(&p)

		pictures = append(pictures, p)
	}

	return pictures, nil

}