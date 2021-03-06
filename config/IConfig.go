package config

import "errors"

type Configuration interface {
	Validate() error
}

func combineErrors(err ...error) error {
	str := ""
	for _, e := range err {
		if e != nil {
			str += e.Error() + "\n"
		}
	}

	if len(str) == 0 {
		return nil
	}

	return errors.New(str)
}