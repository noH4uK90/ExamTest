package models

type QuestionType struct {
	Id   int    `db:"type_id" json:"id"`
	Name string `json:"name"`
}
