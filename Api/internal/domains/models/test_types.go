package models

type TestTypes struct {
	TestId int64 `db:"test_id" json:"testId"`
	TypeId int64 `db:"type_id" json:"typeId"`
}
