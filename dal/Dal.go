package dal

import "github.com/jmoiron/sqlx"

type Dal struct {
	User         User
	Registration Registration
	Medias       Medias
}

func NewDal(db *sqlx.DB) Dal {
	return Dal{
		User:         User{DB: db},
		Registration: Registration{DB: db},
		Medias:       Medias{DB: db},
	}
}
