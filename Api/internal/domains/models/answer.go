package models

type Answer struct {
	Id      int    `db:"answer_id" json:"id"`
	Text    string `json:"text"`
	IsRight bool   `json:"isRight"`
}
