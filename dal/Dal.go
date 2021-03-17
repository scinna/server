package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Dal struct {
	User         User
	Registration Registration
	Medias       Medias
	Collections  Collections
	Server       Server
}

func (d Dal) IsPostgresError(err error, constraint string) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		if pgErr.Constraint == constraint {
			return true
		}
	}

	return false
}

func NewDal(db *sqlx.DB) Dal {
	return Dal{
		User:         User{DB: db},
		Registration: Registration{DB: db},
		Medias:       Medias{DB: db},
		Collections:  Collections{DB: db},
		Server:       Server{DB: db},
	}
}
