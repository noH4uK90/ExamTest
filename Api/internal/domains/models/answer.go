package models

type Answer struct {
	ID      int64  `db:"answer_id" json:"id"`
	Text    string `json:"text" validate:"required"`
	IsRight bool   `db:"is_right" json:"isRight"`
}
