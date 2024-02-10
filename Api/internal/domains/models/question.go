package models

type Question struct {
	ID      int64  `db:"question_id" json:"id"`
	Text    string `json:"text,omitempty"`
	Image   []byte `json:"image,omitempty"`
	ScoreId int64  `db:"score_id" json:"scoreId"`
	TypeId  int64  `db:"type_id" json:"typeId"`
}
