package models

const (
	VisibilityPublic   Visibility = 0
	VisibilityUnlisted Visibility = 1
	VisibilityPrivate  Visibility = 2
)

type Visibility int

// weird flex but ok...
func VisibilityFromInt(v int) Visibility {
	switch v {
	case 0:
		return VisibilityPublic
	case 1:
		return VisibilityUnlisted
	case 2:
		return VisibilityPrivate
	}

	return 0
}

func (v Visibility) IsValid() bool {
	return v >= VisibilityPublic && v <= VisibilityPrivate
}

func (v Visibility) IsPublic() bool {
	return v == VisibilityPublic
}

func (v Visibility) IsUnlisted() bool {
	return v == VisibilityUnlisted
}

func (v Visibility) IsPrivate() bool {
	return v == VisibilityPrivate
}
