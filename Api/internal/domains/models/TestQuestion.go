package models

type TestQuestion struct {
	TestId     int `db:"test_id" json:"testId"`
	QuestionId int `db:"question_id" json:"questionId"`
}
