package models

type Answer struct {
	Id      int64  `db:"answer_id" json:"id"`
	Text    string `json:"text"`
	IsRight bool   `json:"isRight"`
}
