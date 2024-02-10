package models

type TestType struct {
	ID   int64  `db:"type_id" json:"id"`
	Name string `json:"name"`
}
