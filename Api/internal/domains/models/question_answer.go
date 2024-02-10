package models

type QuestionAnswer struct {
	QuestionID int64 `db:"question_id" json:"questionId"`
	AnswerID   int64 `db:"answer_id" json:"answerId"`
}
