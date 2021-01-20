package models

type DatabaseModel interface {
	GetTableName() string
	GenerateTable() string
}
