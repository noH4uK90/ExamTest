package models

type Test struct {
	ID   int64  `db:"test_id" json:"id"`
	Name string `json:"name"`
}
