package models

type Test struct {
	Id     int64  `db:"test_id" json:"id"`
	Name   string `json:"name"`
	TypeId int64  `db:"type_id" json:"typeId"`
}
