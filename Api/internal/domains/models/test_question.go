package models

type TestQuestion struct {
	TestId     int64 `db:"test_id" json:"testId"`
	QuestionId int64 `db:"question_id" json:"questionId"`
}
