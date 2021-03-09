package forms

import (
	"encoding/json"
	"fmt"
	"github.com/scinna/server/services"
	"net/http"
)

type Constraint struct {
	Field      string
	Message    string
	IsViolated func() bool
}

func ConstraintUniqueString(prv *services.Provider, field, table, column, value, message string) Constraint {
	return Constraint{
		Field: field,
		Message: message,
		IsViolated: func() bool {
			row := prv.DB.QueryRow(fmt.Sprintf("SELECT %v FROM %v WHERE %v = $1", column, table, column), value)
			if row.Err() != nil {
				return true
			}

			valFromDB := ""
			_ = row.Scan(&valFromDB)

			return len(valFromDB) > 0
		},
	}
}

func HasViolations(constraints []Constraint, w http.ResponseWriter) bool {
	errors := map[string]string{}

	for _, c := range constraints {
		if c.IsViolated() {
			errors[c.Field] = c.Message
		}
	}

	if len(errors) == 0 {
		return false
	}

	message, _ := json.Marshal(struct {
		Violations map[string]string
	} { errors })

	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(message)

	return true
}
