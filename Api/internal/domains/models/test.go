package models

type Test struct {
	ID   int64  `db:"test_id" json:"id"`
	Name string `json:"name"`
}

type TestResponse struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Answers []Answer   `json:"answers"`
	Types   []TestType `json:"types"`
}

type TestRequest struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Answers []int64 `json:"answers"`
	Types   []int64 `json:"types"`
}
