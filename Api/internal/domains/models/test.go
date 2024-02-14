package models

type Test struct {
	ID   int64  `db:"test_id" json:"id"`
	Name string `json:"name"`
}

type TestResponse struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
	Types     []TestType `json:"types"`
}

type TestRequest struct {
	Name        string `json:"name"`
	QuestionIDs []int  `json:"questionIDs"`
	TypeIDs     []int  `json:"typeIDs"`
}
