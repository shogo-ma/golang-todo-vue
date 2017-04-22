package model

type Todo struct {
	TodoID string `json:"todo_id"`
	Text   string `json:"text"`
	Status bool   `json:"status"`
}
