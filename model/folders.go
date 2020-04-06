package model

import "time"

// Folder represents a virtual folder inside of a scinna account
type Folder struct {
	ID           int64     `db:"id" json:"ID"`
	CreatedAt    time.Time `db:"created_at" json:"CreatedAt"`
	Creator      AppUser   `db:"creator" json:"Creator"`
	FolderName   string    `db:"folder_name" json:"FolderName"`
	ParentFolder *int64    `db:"parent_folder" json:"ParentFolder"`
	ParentPath   string    `db:"parent_path" json:"ParentPath"`
}
