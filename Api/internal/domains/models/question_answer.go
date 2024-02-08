package models

type QuestionAnswer struct {
	QuestionId int64 `db:"question_id" json:"questionId"`
	AnswerId   int64 `db:"answer_id" json:"answerId"`
}
