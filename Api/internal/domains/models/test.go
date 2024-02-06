package models

type Test struct {
	Id     int    `db:"test_id" json:"id"`
	Name   string `json:"name"`
	TypeId int    `db:"type_id" json:"typeId"`
}
