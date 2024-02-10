package models

type TestQuestion struct {
	TestID     int64 `db:"test_id" json:"testId"`
	QuestionId int64 `db:"question_id" json:"questionId"`
}
