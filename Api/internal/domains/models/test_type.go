package models

type TestType struct {
	Id   int64  `db:"type_id" json:"id"`
	Name string `json:"name"`
}
