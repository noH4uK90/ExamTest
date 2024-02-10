package models

type TestTypes struct {
	TestID int64 `db:"test_id" json:"testId"`
	TypeID int64 `db:"type_id" json:"typeId"`
}
