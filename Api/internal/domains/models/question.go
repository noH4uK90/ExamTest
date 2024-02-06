package models

type Question struct {
	Id      int    `db:"question_id" json:"id"`
	Text    string `json:"text,omitempty"`
	Image   []byte `json:"image,omitempty"`
	ScoreId int    `db:"score_id" json:"scoreId"`
	TypeId  int    `db:"type_id" json:"typeId"`
}
