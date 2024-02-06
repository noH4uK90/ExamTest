package models

type QuestionAnswer struct {
	QuestionId int `db:"question_id" json:"questionId"`
	AnswerId   int `db:"answer_id" json:"answerId"`
}
